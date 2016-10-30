var path = require("path");

module.exports = {
  entry: './src/main.es6',
  output: {
    path: './dist',
    filename: 'build.js'
  },
  module: {
    loaders: [
      {
        test: /\.(js|es6)$/,
        loader: 'babel',
        exclude: /node_modules/
      },
      {
        test: /\.vue$/,
        loader: 'vue'
      }
    ]
  },
  resolve: {
    root: [
      path.resolve('./src'),
      path.resolve('./node_modules/vue/dist'),
      path.resolve('./node_modules')
    ]
  }
};


