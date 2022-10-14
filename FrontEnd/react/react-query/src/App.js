import "./App.css";
import fetchPosts from "./FetchAPI";
import { useQuery } from "react-query";

function App() {
  const { isLoading, isFetching, error, data, status } = useQuery(
    "users",
    fetchPosts
  );

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error! {error.message}</div>;
  }

  if (isFetching) {
    return <div>Fetching data...</div>;
  }

  return (
    <div className="App">
      <h1>{data}</h1>
      
      <h4>{status}</h4>
    </div>
  );
}

export default App;
