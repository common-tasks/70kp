import React from "react";
import "./Search.css";

class SearchBox extends React.Component {
  render() {
    return (
      <div class="main">
        <input
          type="text"
          placeholder="Search for a service..."
          onChange={() => {}}
        />
      </div>
    );
  }
}

export default SearchBox;
