import React,{Component} from 'react';
import "./App.css";


export default class AppContent extends Component{
  fetchList = ()=>{
    console.log('I am clicked');
  }
  render(){
  return (
    <div className="app">
      <div>
      <h1>{this.props.msg1} {this.props.msg2} {this.props.msg3}</h1>
      <br></br>
      <button className='btn btn-primary' onClick={this.fetchList}>Fetch Data</button>
    </div>
    </div>
  );
}
}


