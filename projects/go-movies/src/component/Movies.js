import { Component, Fragment } from "react";
import { Link } from "react-router-dom";

export default class Movies extends Component {
  state = { movies: [] };

  componentDidMount(){
    console.log('componentDidMount');
    this.setState({
        movies:[
            {id:1,title:"The Shawshank Redemption",runtime:142},
            {id:2,title:"The Godfather",runtime:180},
            {id:3,title:"The Dark Knight",runtime:120}
        ]
    });
  }

  componentDidUpdate(){
    console.log('component did update');
  }
  componentWillUnmount(){
    console.log('component will unmount');
  }
  componentDidCatch(){
    console.log('component did catch');
  }
  render() {
    return (
      <Fragment>
        <h1>Choose a Movie</h1>
        <ul>
          {this.state.movies.map((m) => (
            <li key={m.id}>
              <Link to={`/movies/${m.id}`}>  
              {m.title}
              </Link>
            </li>
          ))}
        </ul>
      </Fragment>
    );
  }
}
