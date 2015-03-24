var FileList = React.createClass({
  handleClick: function (path, evt) {
    if (this.props.notifyPathChange) {
      this.props.notifyPathChange(path);
    }

    evt.preventDefault();
  },

  getDefaultProps: function () {
    return {
      files: []
    }
  },

  render: function () {
    var that = this;

    var files = this.props.files
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
