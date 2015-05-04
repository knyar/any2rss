// vim: ts=2:sw=2:sts=2:expandtab
var InputForm = React.createClass({

  getInitialState: function() {
    feed_info = new FeedInfo({
      title: "",
      source_url: "",
      item_search: "",
      item_title: "\\1",
      item_url: "\\2",
      item_text: "\\3",
    })
    return {feed_info: feed_info}
  },

  onChange: function(key, event) {
    feed_info = this.state.feed_info
    feed_info.set(key, event.target.value)
    this.updateParams(feed_info)
  },

  updateParams: function(feed_info) {
    this.setState({feed_info: feed_info})
    this.props.onChangeFeedInfo(feed_info)
  },

  render: function() {
    return (
      <form>
        <div className="form-group">
          <label for="feedTitle">Title of your feed:</label>
          <input className="form-control" id="feedTitle"
            value={this.state.feed_info.title} onChange={this.onChange.bind(this, "title")}/>
        </div>
        <div className="form-group">
          <label for="sourceURL">Source page URL:</label>
          <input className="form-control" id="sourceURL"
            value={this.state.feed_info.source_url} onChange={this.onChange.bind(this, "source_url")}/>
        </div>
        <div className="form-group">
          <label for="blockSearch">Content block extraction:</label>
          <textarea className="form-control" id="blockSearch" rows="2"
            placeholder="regular expression" value={this.state.feed_info.block_search}
            onChange={this.onChange.bind(this, "block_search")}/>
        </div>
        <div className="form-group">
          <input className="form-control" id="blockExtract"
            value={this.state.feed_info.block_extract}
            onChange={this.onChange.bind(this, "block_extract")}/>
        </div>
        <div className="form-group">
          <label for="itemSearch">Item extraction:</label>
          <textarea className="form-control" id="itemSearch" rows="2"
            placeholder="regular expression" value={this.state.feed_info.item_search}
            onChange={this.onChange.bind(this, "item_search")}/>
        </div>
        <div className="form-group">
          <div className="input-group input-group-sm">
            <span className="input-group-addon" id="title-addon">Title</span>
            <input type="text" className="form-control" id="itemTitle"
              aria-describedby="title-addon" value={this.state.feed_info.item_title}
              onChange={this.onChange.bind(this, "item_title")}/>
          </div>
        </div>
        <div className="form-group">
          <div className="input-group input-group-sm">
            <span className="input-group-addon" id="url-addon">URL</span>
            <input type="text" className="form-control" id="itemURL"
              aria-describedby="url-addon" value={this.state.feed_info.item_url}
              onChange={this.onChange.bind(this, "item_url")}/>
          </div>
        </div>
        <div className="form-group">
          <div className="input-group input-group-sm">
            <span className="input-group-addon" id="text-addon">Text</span>
            <textarea className="form-control" aria-describedby="text-addon"
              rows="2" id="itemText" value={this.state.feed_info.item_text}
              onChange={this.onChange.bind(this, "item_text")}/>
          </div>
        </div>
      </form>
    );
  }
});
