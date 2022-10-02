import "./App.css";
const user = {
  name: "james khalkho",
  job: "full time",
};
const name = "murmu kumar";

function App() {
  return (
    <div className="App">
      <h1>
        hello {name} and {formatName(user)}
      </h1>
      <GetName name="kunar" />
      <GetName name="gammu" />
    </div>
  );
}
function formatName(user) {
  return user.name + "::" + user.job;
}
function GetName(props) {
  return <h1>{props.name}</h1>;
}
export default App;
