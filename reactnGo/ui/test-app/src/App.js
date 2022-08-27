import React, { Component } from "react";
import "./App.css";

export default class AppContent extends Component {
  fetchList = () => {
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then((response) => response.json())
      .then((json) => {
        var ulist = document.getElementById('ulist');
        json.forEach(element => {
          let list = document.createElement("li");
          list.appendChild(document.createTextNode(element.title));
          ulist.appendChild(list);
        });
        console.log(json);
      });
  };
  render() {
    return (
      <div className="app">
        <div>
          <h1>
            {this.props.msg1} {this.props.msg2} {this.props.msg3}
          </h1>
          <br></br>
          <button className="btn btn-primary" onClick={this.fetchList}>
            Fetch Data
          </button>
        </div>
        <br></br>
        <ul id="ulist"></ul>
      </div>
    );
  }
}
