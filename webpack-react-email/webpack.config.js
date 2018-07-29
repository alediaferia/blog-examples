const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const PrerenderSPAPlugin = require('prerender-spa-plugin');
const HTMLInlineCSSWebpackPlugin = require('html-inline-css-webpack-plugin').default;

/**
 * This Webpack Configuration File defines how an email template should be rendered
 * into a static HTML file.
 *
 * This configuration expects to be given the following env options (--env.<name> notation):
 *    --env.entry=<template name matching the folder name under the src/templates/ folder>
 *    --env.ASSETS_URL=<the assets URL that will be used to reference all the assets in the email>
 */
const config = (env) => {
  return ({
    mode: 'production',
    entry: {
      // this will be filled in dynamically
    },
    output: {
      filename: '[name].js',
      path: path.join(__dirname, 'output'),
    },
    module: {
      rules: [
        {
          test: /\.(png|jpg|gif|svg)$/,
          use: [
            {
              loader: 'file-loader',
              options: {
                name: '[name].[hash].[ext]',
                outputPath: 'assets/',
                publicPath: (url) => {
                  if (env.ASSETS_URL) {
                    return env.ASSETS_URL + 'assets/' + url;
                  }
                  return 'assets/' + url;
                },
              },
            },
          ],
        },
        {
          test: /\.(scss|css)$/,
          use: [
            'css-loader',
            'sass-loader',
          ],
        },
        {
          test: /\.js$/,
          exclude: /node_modules/,
          use: {
            loader: 'babel-loader',
          },
        },
      ],
    },
    plugins: [
      new PrerenderSPAPlugin({
        staticDir: path.join(__dirname, 'output'),
        indexPath: path.join(__dirname, 'output', `${env.entry}.html`),
        routes: ['/'],
        postProcess: context => Object.assign(context, { outputPath: path.join(__dirname, 'output', `${env.entry}.html`) }),
      }),
      new HtmlWebpackPlugin({
        filename: `${env.entry}.html`,
        chunks: [env.entry],
      }),
      new HTMLInlineCSSWebpackPlugin(),

      // the DefinePlugin helps us defining a const
      // variable that will be 'visible' by the JS code
      // we are rendering at runtime
      new webpack.DefinePlugin({
        "EMAIL_TEMPLATE": JSON.stringify(env.entry),
      }),
    ],
  });
};

module.exports = (env) => {
  const entry = {};
  entry[env.entry] = './src/index.js';
  const cfg = config(env);
  cfg.entry = entry;
  return (cfg);
};