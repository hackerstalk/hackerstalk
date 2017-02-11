import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import { Router, Route, IndexRoute, Link, hashHistory } from 'react-router';
import { App } from './app.js';
import { LinkList } from './listLink.js';
import { NoMatch } from './noMatch.js';


render((
  <Router history={hashHistory}>
    <Route path="/" component={App}>
      <IndexRoute component={LinkList}/>
      <Route path="*" component={NoMatch}/>
    </Route>
  </Router>
), document.getElementById('app'))
