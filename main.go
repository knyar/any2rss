package any2rss

import (
	"fmt"
	"strings"
	"net/http"
	"bytes"
	"time"
	"io/ioutil"
	"regexp"
	"encoding/base64"
	"crypto/sha1"
	"compress/zlib"

	"feed_info"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/feeds"
	"appengine"
	"appengine/memcache"
	"appengine/urlfetch"
)

func init() {
	http.HandleFunc("/go/", process)
}

func process(w http.ResponseWriter, r *http.Request) {
	urlPath := strings.Split(r.URL.Path, "/")
	data, err := base64.StdEncoding.DecodeString(urlPath[len(urlPath)-1])
	if err != nil {
		http.Error(w, "Error decoding base64", 406)
		return
	}

	feedInfo := feed_info.FeedInfo{}
	err = proto.Unmarshal(data, &feedInfo)
	if err != nil {
		http.Error(w, "Error unmarshalling proto", 406)
		return
	}

	url := feedInfo.GetSourceUrl()
	if url == "" {
		http.Error(w, "URL cannot be empty, doh", 406)
		return
	}

	ctx := appengine.NewContext(r)
	content, err := getURL(ctx, url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting URL: %v", err), 406)
		return
	}

	rss, err := createFeed(&feedInfo, content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating RSS: %v", err), 406)
		return
	}

	w.Write(rss)
}


func createFeed(feedInfo *feed_info.FeedInfo, content []byte) ([]byte, error) {
	if bs := feedInfo.GetBlockSearch(); bs != "" {
		bsRE, err := regexp.Compile("(?s)" + feedInfo.GetBlockSearch())
		if err != nil {
			return nil, fmt.Errorf("Invalid block extraction regexp:", err)
		}

		idx := bsRE.FindSubmatchIndex(content)
		content = bsRE.Expand([]byte{}, []byte(feedInfo.GetBlockExtract()), content, idx)
	}

	feed := &feeds.Feed{
		Title: feedInfo.GetTitle(),
		Created: time.Now(),
		Link: &feeds.Link{Href: feedInfo.GetSourceUrl()},
	}

	itemRE, err := regexp.Compile("(?s)" + feedInfo.GetItemSearch())
	if err != nil {
		return nil, fmt.Errorf("Invalid item extraction regexp:", err)
	}

	titleTpl := []byte(feedInfo.GetItemTitle())
	urlTpl := []byte(feedInfo.GetItemUrl())
	textTpl := []byte(feedInfo.GetItemText())

	for _, idx := range(itemRE.FindAllSubmatchIndex(content, 0)) {
		item := &feeds.Item{
			Title: string(itemRE.Expand([]byte{}, titleTpl, content, idx)),
			Link: &feeds.Link{
				Href: string(itemRE.Expand([]byte{}, urlTpl, content, idx)),
			},
			Description: string(itemRE.Expand([]byte{}, textTpl, content, idx)),
		}
		feed.Items = append(feed.Items, item)
	}

	atom, err := feed.ToAtom()
	if err != nil {
		return nil, err
	}

	return []byte(atom), nil
}

func getURL(ctx appengine.Context, url string) ([]byte, error) {
	hash := sha1.Sum([]byte(url))
	key := string(hash[:])

	if content, err := getMC(ctx, key); err != nil {
		if err != memcache.ErrCacheMiss {
			return nil, err
		}
	} else if err == nil {
		ctx.Infof("Got %s from memcache", url)
		return content, nil
	}

	content, err := fetchURL(ctx, url)
	if err != nil {
		return nil, err
	}

	if err = setMC(ctx, key, content); err != nil {
		return nil, err
	}
	ctx.Infof("Got %s via HTTP", url)
	return content, nil
}

func getMC(ctx appengine.Context, key string) ([]byte, error) {
	item, err := memcache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(item.Value)
	r, err := zlib.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func setMC(ctx appengine.Context, key string, content []byte) error {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(content)
	w.Close()

	if b.Len() > 1e6 {
		return fmt.Errorf("Content length is too large")
	}

	item := &memcache.Item{
		Key: key,
		Value: b.Bytes(),
		Expiration: 5 * time.Minute,
	}
	if err := memcache.Set(ctx, item); err != nil {
		return err
	}
	return nil
}

func fetchURL(ctx appengine.Context, url string) ([]byte, error) {
	c := urlfetch.Client(ctx)
	response, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error: %v", response.Status)
	}
	if data, err := ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
