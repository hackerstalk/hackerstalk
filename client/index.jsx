import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import { Router, Route, Link, hashHistory } from 'react-router';
import { App } from './app.jsx';
import { Login } from './login.jsx';
import { NoMatch } from './noMatch.jsx';

render((
  <Router history={hashHistory}>
    <Route path="/" component={App}>
      <Route path="/login" component={Login}/>
      <Route path="*" component={NoMatch}/>
    </Route>
  </Router>
), document.getElementById('app'))
