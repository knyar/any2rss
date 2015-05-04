// vim: ts=2:sw=2:sts=2:expandtab
var ContentBox = React.createClass({
  getInitialState: function() {
    feed_info = new FeedInfo("", "", "", "", "", "")
    return {feed_info: feed_info}
  },

  handleChangeFeedInfo: function(params) {
    this.setState({feed_info: params});
  },

  render: function() {
    return (
      <div className="contentBox container">
        <div className="page-header"><h3>any2rss</h3></div>
        <p>Here you can turn any HTML page (or other piece of content available
          via HTTP) into an RSS feed.</p>
        <InputForm onChangeFeedInfo={this.handleChangeFeedInfo}/>
        <PreviewBox feed_info={this.state.feed_info}/>
      </div>
    );
  }
});
