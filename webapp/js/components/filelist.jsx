class FileList extends React.Component {
  constructor(props) {
    super(props);
  }

  handleClick (path, evt) {
    this.props.notifyPathChange(path);
    evt.preventDefault();
  }

  render () {
    let flist = this.props.files.filter(file => !file.isDir);
    flist.sort((left, right) => left.name <= right.name ? -1 : 1);

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
  fileRoot: ''
};
