import React, { Component, Fragment } from 'react';

export default class AppHeader extends Component{
    constructor(props){
        super(props);
        this.handleListChange=this.handleListChange.bind(this);

    }
    handleListChange(lists){
        console.log('handleListChange of AppHeader is called');
        this.props.handleListChange(lists);
    }
    render(){
        return(
            <Fragment>
            <h1>{this.props.greetingmsg}</h1>
            <h6>{this.props.msg1} { this.props.msg2} { this.props.msg3}</h6>
            <p>There are {this.props.lists.length} enries</p>
            </Fragment>
        );
    }
}