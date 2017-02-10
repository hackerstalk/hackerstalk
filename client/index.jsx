import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import { Router, Route, IndexRoute, Link, hashHistory } from 'react-router';
import { App } from './app.jsx';
import { LinkList } from './listLink.jsx';
import { NoMatch } from './noMatch.jsx';


render((
  <Router history={hashHistory}>
    <Route path="/" component={App}>
      <IndexRoute component={LinkList}/>
      <Route path="*" component={NoMatch}/>
    </Route>
  </Router>
), document.getElementById('app'))
