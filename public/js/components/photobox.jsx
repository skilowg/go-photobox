var Photobox = React.createClass({
  getInitialState: function () {
    return {
      files: [],
      fileRoot: ''
    };
  },

  loadFileData: function (path) {
    var req = new XMLHttpRequest(),
        filePath = '';

    // Set up request to server for files at `path`

    if (path.length) {
      filePath = '?path=';
      filePath += this.state.fileRoot.length ?
        // Prepend the fileroot to the path if we have one
        this.state.fileRoot + '/' + path :
        // No fileroot, so just send the path
        path;

    } else {
      filePath = '';
    }

    req.open("GET", "/files" + filePath);
    req.onreadystatechange = function (evt) {
      if (req.readyState === 4) {
        this.setState({
          fileRoot: path,
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
        notifyPathChange={this.loadFileData}
        fileRoot={this.state.fileRoot} />
    );
  }
});
