import React, { Component } from 'react';

export default class AppHeader extends Component{
    render(){
        return(
            <h1>{this.props.greetingmsg}</h1>
        );
    }
}