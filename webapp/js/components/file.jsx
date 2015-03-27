class File extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    let fileLink, fileURL, fileRoot = this.props.fileRoot;
    fileRoot = fileRoot.length ? fileRoot + '/' : '';
    fileURL = '/photos/' + fileRoot + encodeURIComponent(this.props.file.name);

    if (this.props.file.isDir) {
      fileLink = this.props.file.name;
    } else {
      fileLink = (
        <div className="img-wrap">
          <img src={fileURL} />
        </div>
      );
    }

    return (
      <li>
        <a href={fileURL} onClick={this.props.notifyClick}>
          {fileLink}
        </a>
      </li>
    );

  }
}

File.defaultProps = {
  file: { name: '', isDir: false },
  fileRoot: '',
  notifyClick: function () {}
};
