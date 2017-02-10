import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import { Link } from 'react-router';
import { getCookie, setCookie } from './util.js';
import { LinkList } from './listLink.jsx';


import { Button } from 'elemental';

const App = React.createClass({
  displayName: 'App',

  logout() {
    if (confirm("로그아웃 하시겠습니까?")) {
      setCookie('name', '', 0); 
      location.href = '/';
    }
  },  

  renderUserInfo() {
    const name = getCookie('name');
    if (name === "") {
      return null;
    } else {
      return (
        <span className="info">
          <span style={{marginRight: 10}}>Welcome, {name}</span>
          <Button onClick={this.logout}>Logout</Button>
        </span>
      );
    }
  },

  render () {
    return (
      <div>
        <div className="nav">
          <span className="logo">해커스톡</span>
          {this.renderUserInfo()}
        </div>
        <div className="content">
          {this.props.children}
        </div>
      </div>
    );
  }
});

exports.App = App;
