
module.exports = {
  name: 'app',
  version: '1.0.0',
  items: [
    {name:'视频上传', path: '/operation/upload-video'},
    {name:'其他上传', path: "/other"}
  ],
  menus: [{
    path:  "/home",
    label: "主页"
  },{
    path:  "/document",
    label: "文档"
  },{
    path:  "/about",
    label: "关于我"
  },{
    path:  "/operation",
    label: "运维管理"
  }]
}
