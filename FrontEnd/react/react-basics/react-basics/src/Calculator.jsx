import React from "react";

class Calculator extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      counter: props.counter,
    };
  }

  componentDidMount() {
    console.log("Calculator component mounted");
  }
  componentDidUpdate() {
    console.log("component updated");
  }
  render() {
    const num1 = Number(this.props.num1);

    const num2 = Number(this.props.num2);
    return (
      <>
        <h1>result is {num1 + num2}</h1>
        <h2>{this.state.counter}</h2>
      </>
    );
  }
}
export default Calculator;
