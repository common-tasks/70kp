import "./App.css";
import LeftAreaComponent from "./Components/LeftArea/LeftAreaComponent";
import CenterAreaComponent from "./Components/CenterArea/CenterAreaComponent";
import RightAreaComponent from "./Components/RightArea/RightAreaComponent";

function App() {
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

export default App;
