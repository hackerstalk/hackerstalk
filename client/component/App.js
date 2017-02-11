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

          <div className="lead">해커스톡 메일링 리스트</div>
          <p>매주 개발 관련된 뉴스를 이메일로 보내드립니다.</p>
          <div id="mc_embed_signup">
          <form action="//dabangapp.us8.list-manage.com/subscribe/post?u=126e16ff4f89785ae63578adc&amp;id=020aedb0ef" method="post" id="mc-embedded-subscribe-form" name="mc-embedded-subscribe-form" className="validate" target="_blank" noValidate>
            <div id="mc_embed_signup_scroll">
            <InputGroup contiguous>
              <InputGroup.Section grow>
                <FormIconField width="one-half" iconPosition="left" iconColor="default" iconKey="mail">
                  <FormInput placeholder="이메일 주소" name="email" name="EMAIL"  id="mce-EMAIL" required />
                </FormIconField>
              </InputGroup.Section>
              <InputGroup.Section>
                <Button type="primary" submit={true} name="b_126e16ff4f89785ae63578adc_020aedb0ef">등록</Button>
              </InputGroup.Section>
            </InputGroup>
            </div>
          </form>
          </div>
        </div>
      </div>
    );
  }
});

exports.App = App;
