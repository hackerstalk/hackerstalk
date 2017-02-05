import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import { Link } from 'react-router';

const App = React.createClass({
  displayName: 'App',
  render () {
    return (
      <div className={style.body}>
        <div className={style.header}>
          <Link to="/login" >Login</Link>
        </div>
        <div>
          {this.props.children}
        </div>
      </div>
    );
  }
});

exports.App = App;
