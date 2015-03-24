var File = React.createClass({
  getDefaultProps: function () {
    return {
      file: "",
      notifyClick: function () {}
    }
  },

  render: function () {
    return (
      <li>
        <a href={this.props.file} onClick={this.props.notifyClick}>
          {this.props.file}
        </a>
      </li>
    );
  }
});
