import React from 'react';
import {render} from 'react-dom';
import { Link } from 'react-router';
import { Button, InputGroup, FormIconField, FormInput } from 'elemental';

import { getCookie, setCookie, loggedIn } from '../util.js';

import { LinkList } from './LinkList.js';
import { LoginButton } from './LoginButton.js';


const App = React.createClass({
  displayName: 'App',

  logout() {
    if (confirm("로그아웃 하시겠습니까?")) {
      setCookie('name', '', 0); 
      setCookie('userId', '', 0); 
      location.href = '/';
    }
  },  

  renderUserInfo() {
    if (!loggedIn()) {
      return (
        <span className="info">
          <LoginButton/>
        </span>
      );
    } else {
      const name = getCookie('name');
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
        <hr/>
        <div className="footer">
          © 2017 HackersTalk. <a href="https://github.com/hackerstalk/hackerstalk" target="_blank">GitHub</a>에서 함께 만들어가요.
        </div>
      </div>
    );
  }
});

exports.App = App;
