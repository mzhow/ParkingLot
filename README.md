# 项目介绍

供需不对称的停车场资源容易造成停车秩序混乱，交通拥堵等现象。本项目实现了停车场车位预约功能，可根据车主需要定制下单请求，限制下单时间，每天22:00开启第二天车位预约，还实现了进入/离开停车场登记、取消订单、停车费用结算等功能，一定程度上缓解了停车场资源管理的问题。

项目前端使用到的技术有HTML、JavaScript、CSS和Bootstrap，后端使用Go语言和MySQL数据库。通过RabbitMQ消息队列实现性能优化，安全策略采用JWT用户认证、Bcrypt密码加密、验证码机制，还实现了用户行为日志记录以及后台管理系统。

|      |                           项目地址                           |  账号   |        密码         |
| ---- | :----------------------------------------------------------: | :-----: | :-----------------: |
| 用户 |        [47.97.82.144:9090](http://47.97.82.144:9090)         |   mzh   | parkingPwd_mzh.User |
| 后台 | [47.97.82.144:9090/adminLogin](http://47.97.82.144:9090/adminLogin) | admin01 | parkingPwd_01.Admin |

# 项目结构

```
│  go.mod
│  go.sum
│  LICENSE
│  main // linux编译后程序
│  main.exe // windows编译后程序
│  main.go // 项目入口
│  README.md // 项目文档
│
├─.idea
│      ...
├─controller // 项目主体
│      adminhandler.go // 管理员有关操作
│      controller.go // 初始化及监听客户端行为
│      rabbitmq.go // RabbitMQ实现
│      safety.go // 安全相关策略实现
│      statichandler.go // 响应前端静态页面
│      userhandler.go // 用户注册、登录、下单等功能实现
│
├─dao // 程序操作数据库相关代码
│      admindao.go // 管理员相关数据库操作
│      bookingdao.go // 和订单相关数据库操作
│      cardao.go // 车辆相关数据库操作
│      initdao.go // 初始化数据库
│      paginator.go // 分页生成程序
│      spotdao.go // 车位相关数据库操作
│      userdao.go // 用户相关数据库操作
│
├─memory
│      memory.go // 内存实现，用于session
│
├─model
│      admin.go  // admin结构体
│      booking.go // 订单结构体
│      car.go // 车辆结构体
│      spot.go // 车位结构体
│      user.go // 用户结构体
│
├─session
│      session.go // session实现
│
└─views // 前端
    ├─pages
    │  ├─admin
    │  │      admin_403.html
    │  │      admin_bookings.html
    │  │      admin_index.html
    │  │      admin_login.html
    │  │      admin_spots.html
    │  │
    │  ├─error
    │  │      403.html
    │  │      404.html
    │  │
    │  └─user
    │         index.html
    │         login.html
    │         pay.html
    │         register.html
    │
    └─static
        ├─admin
        │      … 
        ├─css
        │      403.css
        │      404.css
        │      index.css
        │      login.css
        │      time.css
        │
        ├─img
        │      …
        └─js
               err.js
               index.js
               login.js
               register.js
               time.js
```

# 部署说明

- 申请云服务器，操作系统CentOS 7.3（64位）

- 云服务器添加9090端口安全组

- 云服务器安装Apache服务

  ```
  yum install httpd -y
  ```

- 启动服务：

  ```
  systemctl start httpd
  ```

- 将项目下载到云服务器：

  ```
  git clone https://github.com/mzhow/ParkingLot.git
  ```

- 将main文件赋予可执行权限：

  ```
  chmod 777 main
  ```

- 不中断地运行main程序：

  ```
  nohup ./main &
  ```

- 运行成功后，在浏览器访问云服务器IP:9090即可访问成功

# 重新编译说明

如需重新编译该应用，则还需以下操作：

- 安装MySQL，见教程[Linux 安装MySQL](https://www.cnblogs.com/wangpeng00700/p/13539856.html)。

- 新建数据库并加入数据：由ParkingLotData.sql文件导入。

- 修改dao/initdao.go中有关初始化数据库的代码，包括IP地址和账号。

- 在controller/safety.go中添加私钥SECRETKEY。

- 安装RabbitMQ：见教程[Linux centos7安装RabbitMQ](https://www.jianshu.com/p/ee9f7594212b)。

- 修改controller/rabbitmq.go中RabbitMQ的代码，包括RabbitMQ初始化的地址和账号。

- 重新编译项目，在主目录下进行以下操作：

  ```
  SET CGO_ENABLED=0
  SET GOOS=linux
  SET GOARCH=amd64
  go build main.go
  ```

- 依照部署说明部署应用。