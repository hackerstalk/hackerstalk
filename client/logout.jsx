import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import request from 'superagent';
import { setCookie } from './util.js';

const Logout = React.createClass({
  displayName: 'Logout',

  propTypes: {
    router: React.PropTypes.object
  },

  componentWillMount() {
    setCookie('name', '', 0);
    this.props.router.push('/');
  },
  
  render () {
    return (
      <div className={style.body}>
        Bye
      </div>
    );
  }
});

exports.Logout = Logout;
