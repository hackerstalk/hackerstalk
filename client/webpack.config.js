var webpack = require('webpack');
var path = require('path');

var BUILD_DIR = path.resolve(__dirname, '..', 'static');
var APP_DIR = __dirname;

var config = {
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

var indexConfig = Object.assign({}, config, {
  entry: path.join(APP_DIR, 'index.jsx'),
  output: {
    path: BUILD_DIR,
    filename: 'bundleIndex.js'
  }
});

module.exports = [
  indexConfig,
];
