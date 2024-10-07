# 代码风格与规范

## JSON命名

采用`Snake Case`（下划线）命名方式

```json
"code": 0,
"msg": "success",
"data": {
    "chat_infos": []
}
```

## 请求体

`GET`请求与`DELETE`请求不需要 Request Body

## Git Commit 规范

- 🚀init: 项目初始化
- 🎸feat: 添加新特性
- 🐞fix: 修复 bug
- 📚docs: 仅仅修改了文档
- 💄style: 仅仅修改了空格、格式缩进、逗号等等，不改变代码逻辑
- 🔧refactor: 代码重构，没有加新功能或者修复 bug
- 🚀perf: 优化相关，比如提升性能、体验
- 🧪test: 添加测试用例
- 🛠️build: 依赖相关的内容
- 🔄ci: ci 配置相关，例如对 k8s、docker 的配置文件的修改
- 🗃️chore: 改变构建流程、或者增加依赖库、工具等
- ⏪️revert: 回滚到上一个版本
