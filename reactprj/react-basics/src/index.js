import React, { Fragment } from "react";
import ReactDOM from "react-dom";
import './index.css';

function Greeting() {
  return (
    <Fragment>
      <NameOfPerson />
      <Message />
    </Fragment>
  );
}

const NameOfPerson = () => {
  return <h1>James Mahto</h1>;
};
const Message = () => {
  return <h1>Sab ko parnam</h1>;
};
ReactDOM.render(<Greeting />, document.getElementById("root"));
