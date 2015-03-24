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

    if (this.props.showBrowseUp) {
      flist.unshift('..');
    }

    var files = flist
      .map(function (file) {
        return file.trim();
      })
      .map(function (file) {
        return (
          <File file={file} key={file}
                notifyClick={that.handleClick.bind(that, file)} />
        );
      });

    return (
      <ul>
        {files}
      </ul>
    );
  }
});
