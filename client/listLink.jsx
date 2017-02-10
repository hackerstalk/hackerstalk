import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import request from 'superagent';

import { Table, Spinner, Pagination } from 'elemental';
import { LinkAdd } from './addLink.jsx';

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

  renderRow(item, i) {
    return (
      <tr key={i}>
        <td>
          {item.id}
        </td>
        <td>
          <a href={item.url} target="_blank">{item.url}</a>
        </td>
        <td>{item.comment}</td>
        <td>{item.tags}</td>
      </tr>
    )
  },

  handlePageSelect(page) {
    this.setState({page: page}, this.reload)
  },
  
  render () {
    return (
      <div>
        <Table>
          <colgroup>
            <col width="50" />
            <col width="30%" />
            <col width="" />
            <col width="30%" />
          </colgroup>
          <thead>
            <tr>
              <th>#</th>
              <th>링크</th>
              <th>메모</th>
              <th>태그</th>
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
