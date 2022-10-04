import "./App.css";
import "./AboutEvents";
import AboutEvents from "./AboutEvents";
import Toggle from "./Toggle";
import ConditionalRendereing from "./ConditionalRendering";

function App() {
  return (
    <div className="App">
      <AboutEvents />
      <Toggle />
      <ConditionalRendereing isLoggedIn = {false}/>
    </div>
  );
}

export default App;
