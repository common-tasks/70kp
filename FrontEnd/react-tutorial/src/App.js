import React from "react";
import Table from "./Table";

class App extends React.Component {
  state = {
    people: [
      { name: "anurag", job: "maid" },
      { name: "kavi", job: "fulltime" },
    ],
  };

  render() {
    const { people } = this.state;
    return (
      <div className="container">
        <Table people={people} removepeople={this.removepeople} />
      </div>
    );
  }

  removepeople = (index) => {
    const { people } = this.state;

    this.setState({
      people: people.filter((person, i) => {
        return i !== index;
      }),
    });
  };
}

export default App;
