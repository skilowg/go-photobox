class Photobox extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      files: [],
      fileStack: []
    };
  }

  navDir(path) {
    let req = new XMLHttpRequest(),
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

    req.open("GET", "/files" + (filePath.length ? "?path=" + encodeURIComponent(filePath) : ''));
    req.onreadystatechange = (evt) => {
      let files = [];

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
    };

    req.send();

  }

  loadFileData(file) {
    if (file.isDir) {
      this.navDir(file.name);
    }
  }

  componentDidMount() {
    this.loadFileData({name: '', isDir: true});
  }

  render() {
    return (
      <div className="photobox">
        <FolderList
          files={this.state.files}
          fileRoot={this.state.fileStack.join('/')}
          notifyPathChange={this.loadFileData.bind(this)}
          showBrowseUp={this.state.fileStack.length > 0} />

        <FileList
          files={this.state.files}
          fileRoot={this.state.fileStack.join('/')} />
      </div>
    );
  }
}
