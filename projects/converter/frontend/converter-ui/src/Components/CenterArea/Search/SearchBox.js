import React from "react";
import "./Search.css";

function SearchBox({ onTextChanged }) {
  return (
    <div className="main">
      <input
        type="text"
        placeholder="Search for a service..."
        onChange={(e) => onTextChanged(e.target.value)}
      />
    </div>
  );
}

export default SearchBox;
