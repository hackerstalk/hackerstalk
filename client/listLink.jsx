import React from 'react';
import {render} from 'react-dom';
import style from './style.less';
import request from 'superagent';

import { Table } from 'elemental';
import { LinkAdd } from './addLink.jsx';

const LinkList = React.createClass({
  propTypes: {
    router: React.PropTypes.object
  },

  getInitialState() {
    return {
      err: null,
      items: []
    };
  },

  componentWillMount() {
    this.reload();
  },

  reload() {
    request.get('/api/link')
           .set('Accept', 'application/json')
           .end(function (err, res) {
              if (err || !res.ok) {
                this.setState({err: (res.body && res.body.msg) || err});
                return;
              }

              console.log(res.body.items);

              this.setState({items: res.body.items});
            }.bind(this));
  },

  onAdded() {
    this.reload()
  },

  renderRow(item, i) {
    return (
      <tr key={i}>
        <td>
          {item.Id}
        </td>
        <td>
          <a href={item.Url} target="_blank">{item.Url}</a>
        </td>
        <td>{item.Comment}</td>
        <td>{item.Tags}</td>
      </tr>
    )
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
        <div className="lead">새로운 링크 추가</div>
        <LinkAdd onAdded={this.onAdded}/>
      </div>

    );
  }
});

exports.LinkList = LinkList;
