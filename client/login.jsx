import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import request from 'superagent';

const Login = React.createClass({
  displayName: 'Login',

  propTypes: {
    router: React.PropTypes.object
  },

  getInitialState() {
    return {
      err: null,
      githubId: ''
    };
  },

  handleGithubIdChange(evt) {
    this.setState({githubId: evt.target.value});
  },

  handleLoginClick(evt) {
    request.post('/api/login')
           .type('form')
           .send({githubId: this.state.githubId})
           .set('Accept', 'application/json')
           .end(function (err, res) {
             if (err || !res.ok) {
               this.setState({err: (res.body && res.body.msg) || err});
               return;
             }
             this.props.router.push('/');
           }.bind(this));
  },

  render () {
    return (
      <div className={style.body}>
        <p>This is login page.</p>
        <div style={{border: '1pt solid red'}}>
          <p style={{color: 'red'}}>It is for debugging</p>
          <input type='text' value={this.state.githubId} onChange={this.handleGithubIdChange} placeholder="Github ID" />
          <button onClick={this.handleLoginClick}>Login (Debug)</button>
          <div className='error'>
            {this.state.err ? this.state.err.toString() : ''}
          </div>
        </div>
        <p>Sign in with an existing service</p>
        <a href="/auth/github">Login via Github</a>
      </div>
    );
  }
});

exports.Login = Login;
