var Photobox = React.createClass({
  getInitialState: function () {
    return { files: [] };
  },

  loadFileData: function (path) {
    var req = new XMLHttpRequest();

    // Set up request to server for files at `path`
    req.open("GET", "/files" + (path.length ? '?path=' + path : ''));
    req.onreadystatechange = function (evt) {
      if (req.readyState === 4) {
        this.setState({
          files: req.response.split(',')
        })
      }
    }.bind(this);

    req.send();
  },

  componentDidMount: function () {
    this.loadFileData('');
  },

  render: function () {
    return (
      <FileList
        files={this.state.files}
        notifyPathChange={this.loadFileData} />
    );
  }
});
