import React, { useState } from "react";

const UseStateCounter = () => {
  const [value, setValue] = useState(0);
  const btnClickHandler = (caller) => {
    console.log(caller);
  };
  return (
    <>
      <section style={{ margin: "4 rem 0" }}>
        <h2>regular counter</h2>
        <h3>{value}</h3>
        <button
          className="btn"
          onClick={() => {
            setValue(value + 1);
          }}
        >
          Increase
        </button>
        <button
          className="btn"
          onClick={() => {
            setValue(0);
          }}
        >
          Reset
        </button>
        <button
          className="btn"
          onClick={() => {
            setValue(value - 1);
          }}
        >
          Decrease
        </button>
      </section>
    </>
  );
};

export default UseStateCounter;
