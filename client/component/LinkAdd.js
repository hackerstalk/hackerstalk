import React from 'react';
import {render} from 'react-dom';
import request from 'superagent';
import { Link } from 'react-router';
import { Button, Table, Form, FormField, FormInput, Glyph } from 'elemental';

import { loggedIn } from '../util.js';

import { LoginButton } from './LoginButton.js';
import { LinkForm } from './LinkForm.js';


const LinkAdd = React.createClass({
  propTypes: {
    router: React.PropTypes.object,
    onAdded: React.PropTypes.func,
  },

  getInitialState() {
    return {
      err: null,
      isOpen: false
    };
  },

  toggleModal(isOpen) {
    this.setState({isOpen: isOpen});
  },

  onSubmit(data) {
    request.post('/api/link')
           .type('form')
           .send({url: data.url})
           .send({comment: data.comment})
           .send({tags: data.tags})
           .set('Accept', 'application/json')
           .end(function (err, res) {
             if (err || !res.ok) {
               this.setState({err: (res.body && res.body.msg) || err});
               return;
             }


             this.setState(this.getInitialState());

             if(this.props.onAdded) {
              this.props.onAdded();
             }
             
           }.bind(this));
  },


  render () {
    if (!loggedIn()) {
      return null;
    }
    else {
      return (
        <div>
          <Button onClick={this.toggleModal.bind(this, true)}>
            <Glyph icon="repo-create" /> 새 링크 등록</Button>
          <LinkForm 
            title="링크 추가"
            onSubmit={this.onSubmit} 
            onCancel={this.toggleModal.bind(this, false)} 
            isOpen={this.state.isOpen} />
          <div className="error">
            {this.state.err ? this.state.err.toString() : ''}
          </div>
        </div>
      )
    }
  }
});

exports.LinkAdd = LinkAdd;
