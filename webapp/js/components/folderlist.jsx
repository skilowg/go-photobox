class FolderList extends React.Component {
  constructor (props) {
    super(props);
  }

  handleClick (path, evt) {
    this.props.notifyPathChange(path);
    evt.preventDefault();
  }

  render () {
    let flist = this.props.files.filter(file => file.isDir);
    flist.sort((left, right) => left.name <= right.name ? -1 : 1);

    if (this.props.showBrowseUp) {
      flist.unshift({name: 'Go Back', isDir: true});
    }

    let folders = flist
      .map(folder => {
        return (
          <File file={folder} key={folder.name}
                fileRoot={this.props.fileRoot}
                notifyClick={this.handleClick.bind(this, folder)} />
        );
      });

    return (
      <ul className="folders">
        {folders}
      </ul>
    );
  }
}

FolderList.defaultProps = {
  files: [],
  fileRoot: ''
};
