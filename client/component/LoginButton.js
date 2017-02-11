import React from 'react';
import {render} from 'react-dom';
import { Button, Glyph } from 'elemental';

const LoginButton = React.createClass({
  render () {
    return (
      <a href="/auth/github">
        <Button><Glyph icon="mark-github" /> GitHub 계정으로 로그인</Button>
      </a>);
  }
});

exports.LoginButton = LoginButton;
