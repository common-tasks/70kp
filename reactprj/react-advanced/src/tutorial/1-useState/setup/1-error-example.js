import React, { useState } from "react";

const ErrorExample = () => {
  const [title, setTitle] = useState("random title");
  const handleClick = () => {
    if (title === "random title") {
      setTitle("hello world");
    } else {
      setTitle("random title");
    }
  };
  return (
    <React.Fragment>
      <h1>{title}</h1>
      <button type="button" className="btn" onClick={handleClick}>
        Change Title
      </button>
    </React.Fragment>
  );
};

export default ErrorExample;
