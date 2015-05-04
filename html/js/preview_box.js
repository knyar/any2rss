// vim: ts=2:sw=2:sts=2:expandtab
var PreviewBox = React.createClass({
  getInitialState: function() {
    return {
      url: null,
      content: null,
      loading: false,
    }
  },

  update: function() {
    if (!this.props.feed_info) {
      return
    }
    url = window.location.protocol + '//' + window.location.host +
      '/go/' + this.props.feed_info.toBase64();
    this.setState({url: url})

    this.setState({loading: true})
    $.ajax({
      url: url,
      dataType: 'text',
      success: function(data) {
        this.setState({
          content: data,
          loading: false,
        })
      }.bind(this),
      error: function(xhr, status, err) {
        content = "Error (" + err.toString() + "): " + xhr.responseText
        this.setState({loading: false, content: content})
      }.bind(this)
    });
  },

  componentDidUpdate: function(prevProps, prevState) {
    if (prevProps != this.props) {
      this.update()
    }
  },

  componentDidMount: function() {
    this.update()
  },

  preview: function() {
    if (this.state.loading) {
      return (<div>Loading...</div>)
    }
    return (
      <div>{this.state.content}</div>
    );
  },

  render: function() {
    if (!this.state.url) {
      return <div/>
    }
    return (
      <div>
        <div>Your feed URL: {this.state.url}</div>
        {this.preview()}
      </div>
    );
  }
});
