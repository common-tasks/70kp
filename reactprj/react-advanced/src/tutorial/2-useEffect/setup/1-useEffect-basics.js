import React, { useState, useEffect } from 'react';
// by default runs after every re-render
// cleanup function
// second parameter
const UseEffectBasics = () => {
  const [value,setValue] = useState(0);
  useEffect(()=>{
    console.log('use effect is called');
    document.title=`New Message ${value}`;
  },[value]);
  console.log('rendered component');
  useEffect(()=>{
    console.log('hello world')
  },[]);
  return <>
  <h1>{value}</h1>
  <button className='btn' onClick={()=>setValue(value+1)}>Click</button>
  </>;
};

export default UseEffectBasics;
