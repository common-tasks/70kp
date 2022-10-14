import React from "react";
import LeftAreaComponent from "./LeftArea/LeftAreaComponent";
import CenterAreaComponent from "./CenterArea/CenterAreaComponent";
import RightAreaComponent from "./RightArea/RightAreaComponent";
import "./PageDesign.css";

function PageLayout() {
  const handleSearchTextChanged = (searchText) => {
    console.log("search text in PageLayout " + searchText);
  };
  return (
    <div className="row">
      <div className="column left">
        <LeftAreaComponent />
      </div>
      <div className="column middle">
        <CenterAreaComponent OnSearchTextChanged={handleSearchTextChanged} />
      </div>
      <div className="column right">
        <RightAreaComponent />
      </div>
    </div>
  );
}

export default PageLayout;
