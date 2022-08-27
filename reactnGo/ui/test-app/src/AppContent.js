import React, { Component } from "react";

export default class AppContent extends Component {
  constructor(props){
    super(props);
    this.handleListChange=this.handleListChange.bind(this);

  } 
  handleListChange(lists){
    console.log('handleListChange of AppContent is called');
    this.props.handleListChange(lists)
  } 
  state = { ulist: [] };
  clickedItem = (x) => {
    console.log("item clicked", x);
  };

  mouseMoved = () => {
    console.log("mouse moved");
  };
  fetchList = () => {
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then((response) => response.json())
      .then((json) => {
        this.setState({ ulist: json });
        this.handleListChange(json);
        // console.log(json);
      });
  };

  render() {
    return (
      <p>
        base content
        <br></br>
        <button href="#" className="btn btn-primary" onClick={this.fetchList}>
          Click!
        </button>
      </p>
    );
  }
}
