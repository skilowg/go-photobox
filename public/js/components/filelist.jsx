var FileList = React.createClass({
  handleClick: function (path, evt) {
    this.props.notifyPathChange(path);
    evt.preventDefault();
  },

  getDefaultProps: function () {
    return {
      files: [],
      fileRoot: '',
      notifyPathChange: function () {}
    }
  },

  render: function () {
    var that = this,
        flist = this.props.files;

    flist.sort(function (left, right) {
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

    var files = flist
      .map(function (file) {
        return (
          <File file={file} key={file.name}
                fileRoot={that.props.fileRoot}
                notifyClick={that.handleClick.bind(that, file)} />
        );
      });

    return (
      <ul className="files">
        {files}
      </ul>
    );
  }
});
