
##项目初始化步骤

* 安装nodejs环境,推荐使用v6.12.3
    下载地址 : https://nodejs.org/download/release/v6.12.3/

* 在项目根目录执行npm初始化
    命令: npm install --registry=https://registry.npm.taobao.org

* 启动项目
    开发模式: npm run dev (windows系统上为npm run dev_win)
    生产模式: npm run dist (windows系统上为npm run dist_win)

* 开发模式下预览项目
    访问：http://localhost:8088/dist/view/index.html


---
##mmall-fe v1.0.1改动：

* 升级nodejs版本到6.12.3（升级后需要把原项目中node_modules目录删除，重新执行初始化）

* hogan插件被废弃，替换为hogan.js

* 启动方式改动，不再需要安装全局的webpack和webpack-dev-server（需要nodejs版本在6.12.3）

* 移除webpack-dev-server里的client注入，webpack-dev-server可以自动注入client了

* 添加devServer的proxy选项，开发模式下可以直接访问接口，不再需要charles或fiddler做代理

* 将启动命令中的inline和port选项从命令移到了配置文件中