import React from 'react';
import {render} from 'react-dom';
import style from './style.less';

const NoMatch = React.createClass({
  displayName: 'NoMatch',
  render () {
    return (<div>잘못된 접근입니다.</div>);
  }
});

exports.NoMatch = NoMatch;
