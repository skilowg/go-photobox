var Photobox = React.createClass({
  getInitialState: function () {
    return {
      files: [],
      fileStack: []
    };
  },

  loadFileData: function (path) {
    var req = new XMLHttpRequest(),
        filePath = '',
        newFileStack = this.state.fileStack;

    // Set up request to server for files at `path`
    if (path.length) {
      filePath = '';

      if (path === '..') {
        this.state.fileStack.pop();
        path = '';
      } else {
        newFileStack.push(path);
      }

      filePath = this.state.fileStack.join('/');
    }

    req.open("GET", "/files" + (filePath.length ? "?path=" + filePath : ''));
    req.onreadystatechange = function (evt) {
      if (req.readyState === 4) {
        this.setState({
          fileStack: newFileStack,
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
        showBrowseUp={this.state.fileStack.length > 0} />
    );
  }
});
