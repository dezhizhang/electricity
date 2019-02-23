const path = require('path')
module.exports = {
    mode:'development',
    entry: { 
        'index':['./src/page/index/index.js'],
        'login':['./src/page/login/login.js']
    },
    output: {
        path:path.resolve(__dirname,'dist'),
        filename:'js/[name].js'
    }
}