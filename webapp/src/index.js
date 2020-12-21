import React from 'react';
import ReactDOM from 'react-dom';

import { Provider } from 'react-redux';
import store from './app/store';

import "./styles/normalizer-v8.0.1.css";
import "./styles/daleri-mega-v1.2.css";

import App from './App';

ReactDOM.render(
  <React.StrictMode>
    <Provider store={store}>
      <App />
    </Provider>
  </React.StrictMode>,
  document.getElementById('root')
);
