import React from "react";

class Toggle extends React.Component {
  constructor(props) {
    super(props);
    this.state = { toggleme: true };
    // this.HandleClick= this.HandleClick.bind(this);
  }
  HandleClick = (e) => {
    this.setState((prevState) => ({
      toggleme: !prevState.toggleme,
    }));
  };
  render() {
    return (
      <div>
        <button onClick={this.HandleClick}>
          {this.state.toggleme ? "ON" : "OFF"}
        </button>
      </div>
    );
  }
}
export default Toggle;
