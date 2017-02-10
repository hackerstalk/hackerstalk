import React from 'react';
import {render} from 'react-dom';
import style from './style.less';

import { Table } from 'elemental';
import { LinkAdd } from './addLink.jsx';

const LinkList = React.createClass({
  propTypes: {
    router: React.PropTypes.object
  },

  componentWillMount() {
    
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
            <tr>
              <td>
                1
              </td>
              <td>
                <a href="https://hackerstalk.com/" target="_blank">https://hackerstalk.com/</a>
              </td>
              <td>GOOD</td>
              <td>#해피해킹</td>
            </tr>
          </tbody>
        </Table>
        <LinkAdd/>
      </div>

    );
  }
});

exports.LinkList = LinkList;
