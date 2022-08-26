import react, {Component, Fragment} from 'react'
import './AppFooter.css'

export default class AppFooter extends Component{
    render(){
        var currentYear = new Date().getFullYear()
        return (
            <Fragment className='app'>
            <hr></hr>
            <p className='footer'>Copyright {currentYear} @Anurag Kumar</p>
            </Fragment>
        );
    }
}