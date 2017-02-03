var webpack = require('webpack');
var path = require('path');

var BUILD_DIR = path.resolve(__dirname, '..', 'static');
var APP_DIR = __dirname;

var config = {
  entry: path.join(APP_DIR, 'index.jsx'),
  output: {
    path: BUILD_DIR,
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {
        test : /\.jsx?/,
        include : APP_DIR,
        loader : 'babel-loader'
      },
      {
        test: /\.less$/,
        use: [
          'style-loader',
          { loader: 'css-loader', options: { importLoaders: 1 } },
          'less-loader'
        ]
      }
    ]
  }
};

module.exports = config;
