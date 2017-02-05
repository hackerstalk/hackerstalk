import React from 'react';
import {render} from 'react-dom';
import style from './style.less';

const Login = React.createClass({
  displayName: 'Login',
  render () {
    return (
      <div className={style.body}>
        <p>This is login page.</p>
        <p>Sign in with an existing service</p>
        <a href="/auth/github">Login via Github</a>
      </div>
    );
  }
});

exports.Login = Login;
