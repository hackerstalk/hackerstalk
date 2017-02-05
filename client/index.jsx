import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import { Router, Route, IndexRoute, Link, hashHistory } from 'react-router';
import { App } from './app.jsx';
import { Login } from './login.jsx';
import { Logout } from './logout.jsx';
import { NoMatch } from './noMatch.jsx';
import { NewLinkForm } from './newLinkForm.jsx';

render((
  <Router history={hashHistory}>
    <Route path="/" component={App}>
      <Route path="login" component={Login}/>
      <Route path="logout" component={Logout}/>
      <Route path="post" component={NewLinkForm}/>
      <IndexRoute component={NoMatch}/>
    </Route>
  </Router>
), document.getElementById('app'))
