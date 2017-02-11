import React from 'react';
import {render} from 'react-dom';
import { Router, Route, IndexRoute, Link, hashHistory } from 'react-router';

import { App } from './component/App.js';
import { LinkList } from './component/LinkList.js';
import { NoMatch } from './component/NoMatch.js';

import style from './style.less';

render((
  <Router history={hashHistory}>
    <Route path="/" component={App}>
      <IndexRoute component={LinkList}/>
      <Route path="*" component={NoMatch}/>
    </Route>
  </Router>
), document.getElementById('app'))
