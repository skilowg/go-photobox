var File = React.createClass({
  getDefaultProps: function () {
    return {
      file: {
        name: '',
        isDir: false
      },
      notifyClick: function () {}
    }
  },

  render: function () {
    return (
      <li>
        <a href={this.props.file.name} onClick={this.props.notifyClick}>
          {this.props.file.name}
        </a>
      </li>
    );
  }
});
