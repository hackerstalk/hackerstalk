import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import { Link } from 'react-router';
import { getCookie } from './util.js';

const App = React.createClass({
  displayName: 'App',

  renderUserInfo() {
    const name = getCookie('name');
    if (name === "") {
      return (
        <Link to="/login" >Login</Link>
      );
    } else {
      return (
        <div>
          <span style={{marginRight: 10}}>Welcome, {name}</span>
          <Link style={{marginRight: 10}} to="/post">Post</Link>
          <Link to="/logout">Logout</Link>
        </div>
      );
    }
  },

  render () {
    return (
      <div className={style.body}>
        <div className={style.header}>
          {this.renderUserInfo()}
        </div>
        <div>
          {this.props.children}
        </div>
      </div>
    );
  }
});

exports.App = App;
