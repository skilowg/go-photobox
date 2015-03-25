var Photobox = React.createClass({
  getInitialState: function () {
    return {
      files: [],
      fileStack: []
    };
  },

  navDir: function (path) {
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
      var files = [];

      if (req.readyState === 4) {
        try {
          files = JSON.parse(req.response);
          this.setState({
            fileStack: newFileStack,
            files: files
          });
        } catch(e) {
          console.log(e);
        }
      }
    }.bind(this);

    req.send();

  },

  loadFileData: function (file) {
    if (file.isDir) {
      this.navDir(file.name);
    }
  },

  componentDidMount: function () {
    this.loadFileData({name: '', isDir: true});
  },

  render: function () {
    return (
      <FileList
        files={this.state.files}
        fileRoot={this.state.fileStack.join('/')}
        notifyPathChange={this.loadFileData}
        showBrowseUp={this.state.fileStack.length > 0} />
    );
  }
});
