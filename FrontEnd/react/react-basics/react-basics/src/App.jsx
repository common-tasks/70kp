import React from "react";
import "./App.css";
import CommonSearchBar from "./CommomSearchBar";
import Clock from "./Clock";

class App extends React.Component {
  constructor() {
    super();
    this.state = { searchText: "" };
  }

  updateSearchText = (text) => {
    this.setState({ searchText: text });
    console.log("search text in App " + this.state.searchText);
  };
  render() {
    return (
      <div className="App">
        <CommonSearchBar searchText={this.state.searchText} />
        <Clock />
      </div>
    );
  }
}

export default App;
