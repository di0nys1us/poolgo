'use strict';

import path from "path";
import HtmlWebpackPlugin from "html-webpack-plugin";

export default {
    entry: path.resolve(__dirname, './index.js'),
    output: {
        path: path.resolve(__dirname, './public'),
        filename: 'bundle.js'
    },
    module: {
        loaders: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: 'babel-loader'
            },
        ]
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: path.resolve(__dirname, './index.ejs')
        })
    ]
};