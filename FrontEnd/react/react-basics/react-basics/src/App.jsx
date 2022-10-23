import "./App.css";
import Calculator from "./Calculator";
import RequestAnalyzer from "./RequestsAnalyzer";

function App() {
  return (
    <div className="App">
      <DisplayMessage name="kunar" />
      <Calculator num1="2" num2="5" counter="-9" />
      <RequestAnalyzer counter={0} />
    </div>
  );
}
function DisplayMessage(props) {
  return (
    <>
      <h1>{props.name}</h1>
    </>
  );
}
export default App;
