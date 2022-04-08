## 基于kratos的注册登录服务

- 使用kratos模板创建工程，实现注册登录功能

step：

1、kratos根据layout创建工程

2、api/.proto定义方法的路由，入参，出参；生成pb.go文件

3、configs/以及deploy/docker-compose.yaml配置DB的相关设置，conf.proto中的设置根据configs/进行修改

4、data/中实现数据库连接，docker启动数据库，并进行连接测试

5、server/注册grpc和http，实现service-biz-data，逐级引用

6、internal/中每个文件都要wire注入相应的组件

7、cmd/中进行wire依赖更新

8、根目录启动项目