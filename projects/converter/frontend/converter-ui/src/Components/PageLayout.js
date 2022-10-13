import React from "react";
import LeftAreaComponent from "./LeftArea/LeftAreaComponent";
import CenterAreaComponent from "./CenterArea/CenterAreaComponent";
import RightAreaComponent from "./RightArea/RightAreaComponent";
import './PageDesign.css';

class PageLayout extends React.Component {
    render() {
        return (
            <div className="row">
              <div className="column left">
               <LeftAreaComponent />
              </div>
              <div className="column middle" >
                <CenterAreaComponent />
              </div>
              <div className="column right">
                <RightAreaComponent />
              </div>
            </div>
          );
    }
}

export default PageLayout;