import React, { Component } from "react";
import "./App.css";
import AppHeader from "./AppHeader";
import AppFooter from "./AppFooter";
import AppContent from "./AppContent";
import AppFooterFunctionalComponent from "./AppFooterFunctionalComponent";

export default class App extends Component {
  myVars = {
    msg1: "hello",
    msg2: " world",
    msg3: " callme anytime and I will ignore the call ",
  };
  constructor(props) {
    super(props);
    this.handleListChange = this.handleListChange.bind(this);
    this.state = { lists: [] };
  }
  handleListChange(lists) {
    console.log("handleListChange called of App component called");
    this.setState({ lists: lists });
  }

  render() {
    return (
      <div className="app">
        <AppHeader
          greetingmsg="Hello World !"
          {...this.myVars}
          lists={this.state.lists}
          handleListChange={this.handleListChange}
        ></AppHeader>
        <AppContent handleListChange={this.handleListChange}></AppContent>
        <AppFooter></AppFooter>
        <AppFooterFunctionalComponent AuthorName={"Anurag Kumar"}></AppFooterFunctionalComponent>
      </div>
    );
  }
}
