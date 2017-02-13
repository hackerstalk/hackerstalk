import React from 'react';
import {render} from 'react-dom';

import { Modal, ModalHeader, ModalBody, ModalFooter, Form, FormField, FormInput, Button } from 'elemental';

const LinkForm = React.createClass({

  propTypes: {
    title: React.PropTypes.string,
    onSubmit: React.PropTypes.func,
    onCancel: React.PropTypes.func,
    isOpen: React.PropTypes.bool,
    data: React.PropTypes.object,
  },

  getInitialState() {
    if(this.props.data) {
      return {
        url: this.props.data.url || '',
        comment: this.props.data.comment || '',
        tags: this.props.data.tags || []
      }
    }

    return {
      url: '',
      comment: '',
      tags: []
    };
  },

  handleFieldChange(field, evt) {
    this.setState({[field]: evt.target.value});
  },

  handleTag(field, evt) {
    console.log(evt.target.value.split(','));
    this.setState({[field]: evt.target.value.split(',')})
  },

  onSubmit() {
    this.props.onSubmit(this.state)
    this.setState(this.getInitialState());
  },

  onCancel() {
    this.props.onCancel();
    this.setState(this.getInitialState());
  },

  render() {
    return (
      <Modal isOpen={this.props.isOpen}>
        <ModalHeader text={this.props.title} />
        <ModalBody>
          <Form>
            <FormField label="링크 주소" htmlFor="link-input">
              <FormInput type="text" placeholder="https://hackerstalk.com" name="link-input"
                value={this.state.url}
                onChange={this.handleFieldChange.bind(this, 'url')} />
            </FormField>
            <FormField label="태그" htmlFor="link-tag">
              <FormInput type="text" placeholder="#해피해킹" name="link-tag"
                value={this.state.tags.join(',')}
                onChange={this.handleTag.bind(this, 'tags')} />
            </FormField>
            <FormField label="메모" htmlFor="link-memo">
              <FormInput type="text" placeholder="해커스톡 짱" name="link-memo"
                value={this.state.comment}
                onChange={this.handleFieldChange.bind(this, 'comment')} 
                multiline />
            </FormField>
          </Form>
        </ModalBody>
        <ModalFooter>
          <Button type="primary" onClick={this.onSubmit}>등록</Button>
          <Button type="link-cancel" onClick={this.onCancel}>취소</Button>
        </ModalFooter>
      </Modal>
    )
  }
})

exports.LinkForm = LinkForm;
