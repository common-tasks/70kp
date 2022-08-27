import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import AppFooter from './AppFooter';
import AppContent from './AppContent';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap/dist//js/bootstrap';
import AppHeader from './AppHeader';

const root = ReactDOM.createRoot(document.getElementById('root'));
const myVars= {
  msg1:"hello",
  msg2:"world",
  msg3:"callme anytime and I will ignore the call"
}
root.render(
  <React.StrictMode>
    <div className='app'>
    <AppHeader greetingmsg="Hello World !"></AppHeader>
    <App {...myVars}></App>
    <AppContent></AppContent>
    <AppFooter></AppFooter>
    </div>
  </React.StrictMode>
);


reportWebVitals();
