import "./App.css";
import { getHealth, getServices } from "./RestClient";
import { useQuery } from "react-query";

function App() {
  const healthStatus = useQuery("health", getHealth);

  const services = useQuery("services", getServices);

  if (services.data) {
    console.log('list of services ' + JSON.stringify(services.data));
  }

  if (healthStatus.isLoading) {
    return <div>Loading...</div>;
  }

  if (healthStatus.error) {
    return <div>Error! {healthStatus.error.message}</div>;
  }

  if (healthStatus.isFetching) {
    return <div>Fetching data...</div>;
  }

  return (
    <div className="App">
      <h1>{healthStatus.data}</h1>

      <h4>{healthStatus.status}</h4>
    </div>
  );
}

export default App;
