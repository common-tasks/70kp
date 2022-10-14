import React from "react";
import SearchBox from "./Search/SearchBox";

function CenterAreaComponent({ OnSearchTextChanged }) {
  const handleSearchText = (text) => {
    console.log("text is " + text);
    OnSearchTextChanged(text);
  };
  return (
    <>
      <SearchBox onTextChanged={handleSearchText} />
    </>
  );
}

export default CenterAreaComponent;
