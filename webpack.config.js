module.exports = {
    entry: {
        regular: './api.jsx'
    },
    output: {
        path: 'build',
        filename: 'build.js',
    },
    resolve: {
        extensions: ['', '.js', '.jsx'],
    		alias: {
    			'react': 'preact-compat',
    			'react-dom': 'preact-compat'
    		}
    },
    module: {
      loaders: [
        {
          test: /\.jsx?$/,
          exclude: /(node_modules|bower_components)/,
          loader: 'babel', 
          query: {
            presets: ['es2015', 'react']
          }
        }]
    }
}