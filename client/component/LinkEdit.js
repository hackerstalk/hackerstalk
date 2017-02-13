import React from 'react';
import {render} from 'react-dom';
import request from 'superagent';
import { Link } from 'react-router';
import { Button, Table, Form, FormField, FormInput, Glyph } from 'elemental';

import { loggedIn } from '../util.js';

import { LoginButton } from './LoginButton.js';
import { LinkForm } from './LinkForm.js';


const LinkEdit = React.createClass({
  propTypes: {
    router: React.PropTypes.object,
    onEdited: React.PropTypes.func,
    linkId: React.PropTypes.number,
    data: React.PropTypes.object,
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
    request.put('/api/link/' + this.props.linkId)
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

             if(this.props.onEdited) {
              this.props.onEdited();
             }
             
           }.bind(this));
  },


  render () {
    var form;
    if(this.state.isOpen) {
      form = (
        <LinkForm 
          title="링크 수정"
          onSubmit={this.onSubmit} 
          onCancel={this.toggleModal.bind(this, false)} 
          isOpen={this.state.isOpen}
          data={this.props.data}
           />
      );
    }
    return (
      <div>
        <Button onClick={this.toggleModal.bind(this, true)}>수정</Button>
        {form}
        <div className="error">
          {this.state.err ? this.state.err.toString() : ''}
        </div>
      </div>
    )
  }
});

exports.LinkEdit = LinkEdit;
