import React from "react";
import SearchBox from "./Search/SearchBox";

function CenterAreaComponent({handleSearchTextChanged}) {
  const handleSearchText = (text) => {
    console.log("text is " + text);
    ;
  };
  return (
    <>
      <SearchBox onTextChanged={handleSearchText} />
    </>
  );
}

export default CenterAreaComponent;
