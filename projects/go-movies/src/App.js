import React, { Component, Fragment } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
  useParams,
  useLocation,
  useMatch,
} from "react-router-dom";
import Movies from "./component/Movies";
import Admin from "./component/Admin";
import Home from "./component/Home";

export default class App extends Component {
  render() {
    return (
      <Router>
        <div className="container">
          <div className="row">
            <h1 className="mt-3">Watch Movies!!!</h1>
            <hr className="mb-3"></hr>
          </div>
          <div className="row">
            <div className="col-md-2">
              <nav>
                <ul className="list-group">
                  <li className="list-group-item">
                    <Link to="/home">Home</Link>
                  </li>
                  <li className="list-group-item">
                    <Link to="/movies">Movies</Link>
                  </li>
                  <li className="list-group-item">
                    <Link to="/category">Catogories</Link>
                  </li>
                  <li className="list-group-item">
                    <Link to="/admin">Manage Catalog</Link>
                  </li>
                </ul>
              </nav>
            </div>
            <div className="col-md-10">
              <Routes>
                <Route path="/movies/:id" element={<Movie />}></Route>
                <Route path="/movies" element={<Movies />}></Route>
                <Route path="/admin" element={<Admin />}></Route>
                <Route path="/home" element={<Home />}></Route>
                <Route exact path="/category" element={<Category />}></Route>
                <Route path="/" element={<Home />}></Route>
              </Routes>
            </div>
          </div>
        </div>
      </Router>
    );
  }
}
function Movie() {
  let { id } = useParams();
  return <h1>Movie id {id}</h1>;
}

function Category() {
  const { pathname } = useLocation();
  console.log("pathname is ", pathname);

  //   let { id } = useParams();
  // console.log('path is '+id);
  return (
    <Fragment>
      <h2>Category</h2>
      <ul>
        <li>
          <Link to={Category / Comedy}>Comedy</Link>
        </li>
        <li>
          <Link to={Category / Action}>Action</Link>
        </li>
        <li>
          <Link to={Category / Drama}>Drama</Link>
        </li>
      </ul>
    </Fragment>
  );
}
function Comedy() {
  return <h1>Comedy</h1>;
}
function Action() {
  return <h1>Action</h1>;
}
function Drama() {
  return <h1>Drama</h1>;
}
