import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import request from 'superagent';
import { Link } from 'react-router';
import { getCookie } from './util.js';

import { Button, Table, Form, FormField, FormInput, Glyph } from 'elemental';


const LinkAdd = React.createClass({
  propTypes: {
    router: React.PropTypes.object
  },

  getInitialState() {
    return {
      err: null,
      url: '',
      comment: '',
      tags: ''
    };
  },

  componentWillMount() {
    
  },

  handleFieldChange(field, evt) {
    this.setState({[field]: evt.target.value});
  },

  handlePostClick(evt) {
    request.post('/api/link')
           .type('form')
           .send({url: this.state.url})
           .send({comment: this.state.comment})
           .send({tags: this.state.tags.split(',')})
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
    const name = getCookie('name');
    if (name === "") {
      return (
        <div>
          <p>새로운 링크를 등록하려면 로그인이 필요합니다.</p>
          <a href="/auth/github"><Button><Glyph icon="mark-github" /> GitHub 계정으로 로그인</Button></a>
        </div>
      )
    }
    else {
      return (
        <div>
          <Form type="inline">
            <FormField label="링크 주소" htmlFor="link-input">
              <FormInput type="text" placeholder="https://hackerstalk.com" name="link-input"
                onChange={this.handleFieldChange.bind(this, 'url')} />
            </FormField>
            <FormField label="메모" htmlFor="link-memo">
              <FormInput type="text" placeholder="해커스톡 짱" name="link-memo"
                onChange={this.handleFieldChange.bind(this, 'comment')} />
            </FormField>
            <FormField label="태그" htmlFor="link-tag">
              <FormInput type="text" placeholder="#해피해킹" name="link-tag"
                onChange={this.handleFieldChange.bind(this, 'tags')} />
            </FormField>
            <FormField>
              <Button onClick={this.handlePostClick}>등록</Button>
            </FormField>
          </Form>
          <div className="error">
            {this.state.err ? this.state.err.toString() : ''}
          </div>
        </div>
      )
    }
  }
});

exports.LinkAdd = LinkAdd;
