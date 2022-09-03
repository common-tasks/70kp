import React, { useState, useEffect } from "react";

// cleanup function
// second argument

const UseEffectCleanup = () => {
  const [sz,SetSz] = useState(window.innerWidth);
  console.log(sz);
  const checkSize = ()=>{
    SetSz((currentSize)=>{
      console.log(currentSize);
      return window.innerWidth;
    });
  }
  useEffect(()=>{
    window.addEventListener('resize',checkSize);
  })
  return <>
  <h1>{sz} window size </h1>
  </>;
};

export default UseEffectCleanup;
