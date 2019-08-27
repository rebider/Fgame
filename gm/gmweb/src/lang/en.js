export default {
  route: {
    dashboard: 'Dashboard',
    introduction: 'Introduction',
    documentation: 'Documentation',
    guide: 'Guide',
    permission: 'Permission',
    pagePermission: 'Page Permission',
    directivePermission: 'Directive Permission',
    icons: 'Icons',
    components: 'Components',
    componentIndex: 'Introduction',
    tinymce: 'Tinymce',
    markdown: 'Markdown',
    jsonEditor: 'JSON Editor',
    dndList: 'Dnd List',
    splitPane: 'SplitPane',
    avatarUpload: 'Avatar Upload',
    dropzone: 'Dropzone',
    sticky: 'Sticky',
    countTo: 'CountTo',
    componentMixin: 'Mixin',
    backToTop: 'BackToTop',
    dragDialog: 'Drag Dialog',
    dragKanban: 'Drag Kanban',
    charts: 'Charts',
    keyboardChart: 'Keyboard Chart',
    lineChart: 'Line Chart',
    mixChart: 'Mix Chart',
    example: 'Example',
    nested: 'Nested Routes',
    menu1: 'Menu 1',
    'menu1-1': 'Menu 1-1',
    'menu1-2': 'Menu 1-2',
    'menu1-2-1': 'Menu 1-2-1',
    'menu1-2-2': 'Menu 1-2-2',
    'menu1-3': 'Menu 1-3',
    menu2: 'Menu 2',
    Table: 'Table',
    dynamicTable: 'Dynamic Table',
    dragTable: 'Drag Table',
    inlineEditTable: 'Inline Edit',
    complexTable: 'Complex Table',
    treeTable: 'Tree Table',
    customTreeTable: 'Custom TreeTable',
    tab: 'Tab',
    form: 'Form',
    createArticle: 'Create Article',
    editArticle: 'Edit Article',
    articleList: 'Article List',
    errorPages: 'Error Pages',
    page401: '401',
    page404: '404',
    errorLog: 'Error Log',
    excel: 'Excel',
    exportExcel: 'Export Excel',
    selectExcel: 'Export Selected',
    uploadExcel: 'Upload Excel',
    zip: 'Zip',
    exportZip: 'Export Zip',
    theme: 'Theme',
    clipboardDemo: 'Clipboard',
    i18n: 'I18n',
    externalLink: 'External Link',
    basemanage:'基础管理',
    usermanagebk:'用户管理',
    channelmanage:'渠道管理',
    platformmanage:'平台管理',
    gameplayer:'玩家管理',
    playersearch:'玩家查询',
    chat:'聊天监控',
    chatset:'聊天配置',
    centerset:'运维管理',
    centerplat:'中心平台配置',
    server:'中心服务配置',
    chatlog:'日志查询',
    log:"日志查询",
    playerlog:"玩家日志",
    serverlog:"服务器日志",
    serverlogin:"登陆设置",
    alliance:"仙盟查询",
    applymail:"邮件提交",
    approvemail:"邮件审核",
    serversupportplayer:"玩家扶持",
    serversupportpool:"扶持池",
    orderlist:"订单查询",
    orderstatic:"订单统计",
    servergm:"gm命令",
    onlinereport:"在线汇总",
    ngonlinereport:"内挂汇总",
    notice:"公告发送",
    redeem:"兑换码",
    redeemcode:"兑换码查看",
    centeruser:"中心用户查看",
    centernewguauser:"中心内挂用户",
    usermanage:"中心用户管理",
    goldchange:"元宝增加汇总",
    goldreduce:"元宝消耗汇总",
    centernotice:"中心公告配置",
    newbindgold:"新绑元汇总增加",
    newbindgoldreduce:"新绑元汇总减少",
    newgoldchange:"新元宝变化汇总增加",
    newgoldchangereduce:"新元宝变化汇总减少",
    alliancelog:"仙盟日志",
    serverset:"战区服设置",
    playerstats:"玩家点击统计",
    retentionreport:"存留率统计",
    centerserverquery:"服务器查看",
    platformmarry:"平台结婚配置",
    clientversionset:"客户端版本设置",
    platformserverconfig:"交易服务器设置",
    serverplayerlevel:"玩家等级统计",
    tradequery:"交易行查询",
    platformsetting:"中心服配置",
    playerquest:"玩家任务统计",
    ordersignlequery:"订单单服查询",
    serverdaily:"分服详情",
    zhanquServer:"战区服查看",
    ipquery:"ip封禁查询",
    recycle:"交易行回购",
    tradelog:"交易日志",
    jiaoYiZhanQu:"交易战区",
    jiaoYiZhanQuServerSet:"交易战区服务器配置",
    quanPingTaiServerSet:"全平台服务器配置",
    zhanQuServerSet:"战区服务器配置",
    platformSupportPoolSet:"平台扶植池设置",
    platformChatSet:"平台聊天设置",
    chengZhanServerSet:"城战平台设置",
    feedBackFee:"提现查询"
  },
  navbar: {
    logOut: 'Log Out',
    dashboard: 'Dashboard',
    github: 'Github',
    screenfull: 'Screenfull',
    theme: 'Theme',
    size: 'Global Size'
  },
  login: {
    title: 'Login Form',
    logIn: 'Log in',
    username: 'Username',
    password: 'Password',
    any: 'any',
    thirdparty: 'Or connect with',
    thirdpartyTips: 'Can not be simulated on local, so please combine you own business simulation! ! !'
  },
  documentation: {
    documentation: 'Documentation',
    github: 'Github Repository'
  },
  permission: {
    roles: 'Your roles',
    switchRoles: 'Switch roles'
  },
  guide: {
    description: 'The guide page is useful for some people who entered the project for the first time. You can briefly introduce the features of the project. Demo is based on ',
    button: 'Show Guide'
  },
  components: {
    documentation: 'Documentation',
    tinymceTips: 'Rich text editor is a core part of management system, but at the same time is a place with lots of problems. In the process of selecting rich texts, I also walked a lot of detours. The common rich text editors in the market are basically used, and the finally chose Tinymce. See documentation for more detailed rich text editor comparisons and introductions.',
    dropzoneTips: 'Because my business has special needs, and has to upload images to qiniu, so instead of a third party, I chose encapsulate it by myself. It is very simple, you can see the detail code in @/components/Dropzone.',
    stickyTips: 'when the page is scrolled to the preset position will be sticky on the top.',
    backToTopTips1: 'When the page is scrolled to the specified position, the Back to Top button appears in the lower right corner',
    backToTopTips2: 'You can customize the style of the button, show / hide, height of appearance, height of the return. If you need a text prompt, you can use element-ui el-tooltip elements externally',
    imageUploadTips: 'Since I was using only the vue@1 version, and it is not compatible with mockjs at the moment, I modified it myself, and if you are going to use it, it is better to use official version.'
  },
  table: {
    dynamicTips1: 'Fixed header, sorted by header order',
    dynamicTips2: 'Not fixed header, sorted by click order',
    dragTips1: 'The default order',
    dragTips2: 'The after dragging order',
    title: 'Title',
    importance: 'Imp',
    type: 'Type',
    remark: 'Remark',
    search: 'Search',
    add: 'Add',
    export: 'Export',
    reviewer: 'reviewer',
    id: 'ID',
    date: 'Date',
    author: 'Author',
    readings: 'Readings',
    status: 'Status',
    actions: 'Actions',
    edit: 'Edit',
    publish: 'Publish',
    draft: 'Draft',
    delete: 'Delete',
    cancel: 'Cancel',
    confirm: 'Confirm'
  },
  errorLog: {
    tips: 'Please click the bug icon in the upper right corner',
    description: 'Now the management system are basically the form of the spa, it enhances the user experience, but it also increases the possibility of page problems, a small negligence may lead to the entire page deadlock. Fortunately Vue provides a way to catch handling exceptions, where you can handle errors or report exceptions.',
    documentation: 'Document introduction'
  },
  excel: {
    export: 'Export',
    selectedExport: 'Export Selected Items',
    placeholder: 'Please enter the file name(default excel-list)'
  },
  zip: {
    export: 'Export',
    placeholder: 'Please enter the file name(default file)'
  },
  theme: {
    change: 'Change Theme',
    documentation: 'Theme documentation',
    tips: 'Tips: It is different from the theme-pick on the navbar is two different skinning methods, each with different application scenarios. Refer to the documentation for details.'
  },
  tagsView: {
    refresh: 'Refresh',
    close: 'Close',
    closeOthers: 'Close Others',
    closeAll: 'Close All'
  }
}
