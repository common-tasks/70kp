import React, { Component } from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Movies from "./component/Movies";
import Admin from "./component/Admin";
import Home from "./component/Home";

export default class App extends Component {
  render() {
    return (
      <Router>
        <div className="container">
          <div className="row">
            <h1 className="mt-3">Plaza Cinema Hall</h1>
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
                    <Link to="/admin">Manage Catalog</Link>
                  </li>
                </ul>
              </nav>
            </div>
            <div className="col-md-10">
              <Routes>
                <Route path='/movies' element={<Movies/>}>
                </Route>
                <Route path='/admin' element ={<Admin/>}>                
                </Route>
                <Route path='/home' element={<Home/>}>
                </Route>
              </Routes>
            </div>
          </div>
        </div>
      </Router>
    );
  }
}
