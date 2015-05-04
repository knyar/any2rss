// vim: ts=2:sw=2:sts=2:expandtab

var ProtoBuf = dcodeIO.ProtoBuf;
var FeedInfo = ProtoBuf.loadProtoFile("/feed_info.proto").build("FeedInfo");

React.render(
  <ContentBox />,
  document.getElementById('content')
);
