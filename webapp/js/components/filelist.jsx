class FileList extends React.Component {
  constructor(props) {
    super(props);
  }

  handleClick (path, evt) {
    this.props.notifyPathChange(path);
    evt.preventDefault();
  }

  render () {
    let flist = this.props.files;

    flist.sort((left, right) => {
      if (left.isDir && right.isDir) {
        return left.name <= right.name ? -1 : 1;
      } else if (left.isDir) {
        return -1;
      } else if (right.isDir) {
        return 1;
      } else {
        return left.name <= right.name ? -1 : 1;
      }
    });

    if (this.props.showBrowseUp) {
      flist.unshift({name: '..', isDir: true});
    }

    let files = flist
      .map(file => {
        return (
          <File file={file} key={file.name}
                fileRoot={this.props.fileRoot}
                notifyClick={this.handleClick.bind(this, file)} />
        );
      });

    return (
      <ul className="files">
        {files}
      </ul>
    );
  }
}

FileList.defaultProps = {
  files: [],
  fileRoot: '',
  notifyPathChange: function () {}
};
