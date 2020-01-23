# markdown demo blog file

> 这是一个测试markdown文用于展示markdown显示效果

`mvn clean`：清理所有生成的class和jar；

`mvn clean compile`：先清理，再执行到compile；

`mvn clean test`：先清理，再执行到test，因为执行test前必须执行compile，所以这里不必指定compile；

`mvn clean package`：先清理，再执行到package。

- clean：清理
- compile：编译
- test：运行测试
- package：打包

scope | 说明 | 示例
-|-|-
compile | 编译是需要用到该jar包（默认） | commons-logging
test | 编译Test时需要用到该jar包 | junit
runtime | 编译时不需要，但运行时需要用到 | mysql
provided | 编译时需要用到，但运行时需要由JDK或者某个服务器提供 | servlet-api

**POM**（Project Object Model项目对象模型）


![img](./assets/lifecycle.png)
