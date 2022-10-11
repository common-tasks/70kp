import React from "react";

const TableHeader = () => {
  return (
    <thead>
      <tr>
        <th>Name</th>
        <th>Job</th>
      </tr>
    </thead>
  );
};
const TableBody = (props) => {
  const rows = props.people.map((row, index) => {
    return (
      <tr key={index}>
        <td>{row.name}</td>
        <td>{row.job}</td>
        <td>
          <button onClick={() => props.removepeople(index)}>Delete</button>
        </td>
      </tr>
    );
  });
  return <tbody>{rows}</tbody>;
};
const Table = (props) => {
  {
    const { people, removepeople } = props;
    return (
      <>
        <table>
          <TableHeader />
          <TableBody people={people} removepeople={removepeople} />
        </table>
      </>
    );
  }
};

export default Table;
