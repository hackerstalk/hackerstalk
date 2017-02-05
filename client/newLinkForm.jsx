import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import request from 'superagent';

const NewLinkForm = React.createClass({
  displayName: 'NewLinkForm',

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

  handleFieldChange(field, evt) {
    this.setState({[field]: evt.target.value});
  },

  handlePostClick(evt) {
    request.post('/api/link/add')
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
    return (
      <div>
        <input type='text'
               placeholder='http://hackerstalk.com'
               onChange={this.handleFieldChange.bind(this, 'url')} />
        <br />
        <input type='text'
               placeholder='Awesome Site!'
               onChange={this.handleFieldChange.bind(this, 'comment')} />
        <br />
        <input type='text'
               placeholder='web,server'
               onChange={this.handleFieldChange.bind(this, 'tags')} />
        <br />
        <button onClick={this.handlePostClick}>Post</button>
        <div className="error">
          {this.state.err ? this.state.err.toString() : ''}
        </div>
      </div>
    );
  }
});

exports.NewLinkForm = NewLinkForm;
