import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import AppFooter from './AppFooter'

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <div className='app'>
    <App></App>
    <AppFooter></AppFooter>
    </div>
  </React.StrictMode>
);


reportWebVitals();
