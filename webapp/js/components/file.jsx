class File extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    let fileURL, fileRoot = this.props.fileRoot, contents;

    fileRoot = fileRoot.length ? fileRoot + '/' : '';
    fileURL = '/photos/' + fileRoot + encodeURIComponent(this.props.file.name);

    if (this.props.file.isDir) {
      contents = (
        <a href={fileURL} onClick={this.props.notifyClick} className="photoboxItem__dir">
          <span>{this.props.file.name}</span>
        </a>
      );
    } else {
      contents = (
        <div className="img-wrap photoboxItem__photo">
          <img src={fileURL} />
        </div>
      );
    }

    return (
      <li className="photoboxItem">{contents}</li>
    );

  }
}

File.defaultProps = {
  file: { name: '', isDir: false },
  fileRoot: '',
  notifyClick: function () {}
};
