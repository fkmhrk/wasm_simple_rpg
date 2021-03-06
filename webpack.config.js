var path = require("path");

module.exports = {
  mode: "development",
  //mode: 'production',

  entry: "./src/main.ts",

  output: {
    path: path.resolve(__dirname, "docs"),
  },

  module: {
    rules: [
      {
        test: /\.ts$/,
        use: "ts-loader",
      },
      {
        test: /\.scss$/,
        use: [
          "style-loader",
          {
            loader: "css-loader",
            options: {
              sourceMap: false,
              importLoaders: 2,
            },
          },
          {
            loader: "sass-loader",
            options: {
              sassOptions: {
                sourceMap: false,
                includePaths: ["./node_modules"],
              },
            },
          },
        ],
      },
    ],
  },

  resolve: {
    extensions: [".ts", ".js"],
  },
};
