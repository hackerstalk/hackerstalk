import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import request from 'superagent';

import { Button, Table, Spinner, Pagination } from 'elemental';
import { LinkAdd } from './addLink.js';
import { getUserId } from './util.js';

const LinkList = React.createClass({
  propTypes: {
    router: React.PropTypes.object
  },

  getInitialState() {
    return {
      page: 1,
      total: 0,
      limit: 0,
      err: null,
      loading: false,
      items: []
    };
  },

  componentWillMount() {
    this.reload();
  },

  reload() {
    this.setState({loading: true});
    request.get('/api/link')
           .query({ page: this.state.page })
           .set('Accept', 'application/json')
           .end(function (err, res) {
              if (err || !res.ok) {
                this.setState({err: (res.body && res.body.msg) || err, loading: false});
                return;
              }

              this.setState({
                items: res.body.items, 
                total: res.body.total,
                limit: res.body.limit,
                loading: false});
            }.bind(this));
  },

  onAdded() {
    this.reload()
  },

  onEditClick(id) {
    console.log('edit', id);
  },

  onDeleteClick(id) {
    console.log('delete', id);

    request.delete('/api/link/' + id)
           .type('form')
           .set('Accept', 'application/json')
           .end(function (err, res) {
             if (err || !res.ok) {
               this.setState({err: (res.body && res.body.msg) || err});
               return;
             }

             alert('삭제 성공');
             this.reload();
             
           }.bind(this));
  },

  renderRow(item, i) {
    var userId = getUserId();
    var opt;
    if(userId && userId == item.user_id) {
      opt = (
        <div>
          <Button onClick={this.onEditClick.bind(this, item.id)}>수정</Button>
          <Button onClick={this.onDeleteClick.bind(this, item.id)}>삭제</Button>
        </div>
      );
    }

    return (
      <tr key={i}>
        <td>
          {item.id}
        </td>
        <td>
          <a href={item.url} target="_blank">{item.url}</a>
        </td>
        <td>
          <p>{item.comment}</p>
          <p>{item.tags}</p>
        </td>
        <td>
          {opt}
        </td>
      </tr>
    )
  },

  handlePageSelect(page) {
    this.setState({page: page}, this.reload)
  },
  
  render () {
    return (
      <div>
        <div className="error">{this.state.err}</div>
        <Table>
          <colgroup>
            <col width="50" />
            <col width="30%" />
            <col width="" />
            <col width="10%" />
          </colgroup>
          <thead>
            <tr>
              <th>#</th>
              <th>링크</th>
              <th>메모</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {this.state.items.map(this.renderRow)}
          </tbody>
        </Table>

        <Pagination
          currentPage={this.state.page}
          onPageSelect={this.handlePageSelect}
          pageSize={this.state.limit}
          total={this.state.total}
          />

        <div style={{'textAlign': 'center'}}>
          { this.state.loading ? (<Spinner size="md" />) : null }
        </div>

        <div className="lead">새로운 링크 추가</div>
        <LinkAdd onAdded={this.onAdded}/>
      </div>

    );
  }
});

exports.LinkList = LinkList;
