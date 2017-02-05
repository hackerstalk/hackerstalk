import React from 'react';
import {render} from 'react-dom';
import style from './style.less';

const NoMatch = React.createClass({
  displayName: 'NoMatch',
  render () {
    return (<div className={style.body}>Hackers Talk!</div>);
  }
});

exports.NoMatch = NoMatch;
