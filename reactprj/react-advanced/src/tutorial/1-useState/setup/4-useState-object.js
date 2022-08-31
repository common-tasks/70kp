import React, { useState } from "react";

const UseStateObject = () => {
  const [person, setPerson] = useState({
    name: "james",
    title: "khalkho",
    motto: "murga daru dhaklo",
  });
  const changeMotto = () => {
    if (person.motto === "murga daru dhaklo") {
      setPerson({ ...person, motto: "study 24*7" });
    } else {
      setPerson({ ...person, motto: "murga daru dhaklo" });
    }
  };
  return (
    <>
      <h1>
        {person.name} {person.title}
      </h1>
      <h2>{person.motto}</h2>
      <button type="button" className="btn" onClick={changeMotto}>
        change motto
      </button>
    </>
  );
};

export default UseStateObject;
