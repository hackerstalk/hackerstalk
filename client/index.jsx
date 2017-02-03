import React from 'react';
import {render} from 'react-dom';
import style from './style.less';

class App extends React.Component {
  render () {
    return (<div className={style.body}>Hackers Talk!</div>);
  }
}

render(<App/>, document.getElementById('app'));
