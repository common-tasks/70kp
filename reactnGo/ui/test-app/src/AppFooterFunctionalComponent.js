import React,{ Fragment } from "react";
import './AppFooter.css'

export default function AppFooterFunctionalComponent(props){
    var currentYear = new Date().getFullYear()
    return (
        <Fragment>
        <hr></hr>
        <p className='footer'>Copyright {currentYear} @ {props.AuthorName}</p>
        </Fragment>
    );
} 