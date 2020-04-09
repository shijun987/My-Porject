<template>
  <div class="hello">
    <el-tabs :tab-position="tabPosition" style="height: tableHeight+'px';">
      <el-tab-pane label="用户配置">用户配置</el-tab-pane>
      <el-tab-pane label="实时数据">实时数据</el-tab-pane>
      <el-tab-pane label="设备设置">设备设置</el-tab-pane>
      <el-tab-pane label="数据调试">数据调试</el-tab-pane>
    </el-tabs>
    <p>{{ msg }}</p>
    <h2>
      <a>written by{{ auther }}</a>
    </h2>
    <input v-model="msg" />
    <el-button plain>点我点我</el-button>
    <p>段落<a href="https://baidu.com">点我链接跳转</a></p>
  </div>
  <!--窗口-->
</template>

<script>
export default {
  name: 'HelloWorld',
  data() {
    return {
      msg: 'msg123456789',
      auther: '石俊',
      tabPosition: 'right',
      tableHeight: 0,
    };
  },

  methods: {
    getTreeHeight (dom) {
    return window.innerHeight - dom.offsetTop;
},
  },
  mounted() {
    this.$nextTick(() => {
      this.tableHeight = this.getTreeHeight(this.$refs['singleTable'].$el) - 50;
    })
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
