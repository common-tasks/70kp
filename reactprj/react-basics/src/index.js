import React, { Fragment } from "react";
import ReactDOM from "react-dom";
import "./index.css";

const Book = () => {
  const authorDetails = {
    authorNamewa:'kunar',
    age:50
  }
  return(
  <Fragment>
     <h1>Book01</h1>
     <AuthorName uthorName={authorDetails.authorNamewa} age ={authorDetails.age}/>
     <AuthorAge age={authorDetails.age}/>
  </Fragment>
  );
};
const AuthorAge =(props) =>{
  return(
    <h1>{props.age}</h1>
  );
}
const AuthorName =(props) =>{
  return(
<Fragment>
  <h1>{props.uthorName}</h1>

</Fragment>
  );
}

ReactDOM.render(<Book />, document.getElementById("root"));
