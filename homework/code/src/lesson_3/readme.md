使用说明

> 本程序提供一个httpserver的docker镜像，运行起来达到以下效果：
>
> 1. 接收客户端 request，并将 request 中带的 header 写入 response header
> 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
> 3. 当访问 localhost/healthz 时，应返回 200

## 作业步骤

- 将本目录下的文件上传到一个ubuntu 20.04或者相近版本的操作系统中

- 执行下面命令进行镜像构建

  ```sh
  # xjxxxj/xieqx-http-server为自己在docker镜像仓库中创建的仓库
  docker build  -t xjxxxj/xieqx-http-server:v2 .
  ```

  ![1665593077604](https://github.com/xjxxxj/cloudNativeStudy/blob/master/homework/code/src/lesson_3/readme.assets/1665593077604.png?raw=true)

- 执行下面命令推送到镜像仓库

  ```sh
  # 需要先通过docker login登录自己的账户
  docker push xjxxxj/xieqx-http-server:v2
  ```

  ![1665593242744.png](https://github.com/xjxxxj/cloudNativeStudy/blob/master/homework/code/src/lesson_3/readme.assets/1665593242744.png?raw=true)

- 然后在官网上就可以看到了：https://hub.docker.com/repository/docker/xjxxxj/xieqx-http-server/tags?page=1&ordering=last_updated

  ![1665593291071](https://github.com/xjxxxj/cloudNativeStudy/blob/master/homework/code/src/lesson_3/readme.assets/1665593291071.png?raw=true)

- 本地运行镜像

  ```sh
  docker run -d -p 8080:80 xjxxxj/xieqx-http-server:v2
  ```

  ![1665593399770](https://github.com/xjxxxj/cloudNativeStudy/blob/master/homework/code/src/lesson_3/readme.assets/1665593399770.png?raw=true)

- 使用curl验证httpserver应用已经生效

  ```sh
  curl -v 127.0.0.1:8080/test
  curl -v 127.0.0.1:8080/healthz
  ```

  ![1665593549271](https://github.com/xjxxxj/cloudNativeStudy/blob/master/homework/code/src/lesson_3/readme.assets/1665593549271.png?raw=true)

- 获取容器pid并且通过nsenter查看容器网络配置

  ```sh
  # 获取容器Pid
  docker inspect $(docker ps | grep http-server | awk '{print $1}') | grep Pid
  # 根据获取到的pid查看容器的网络配置,其中17478为前面查询到的Pid
  nsenter -t 17478 -n ip addr
  ```

  ![1665594498727](https://github.com/xjxxxj/cloudNativeStudy/blob/master/homework/code/src/lesson_3/readme.assets/1665594498727.png?raw=true)

  ![1665594528359](https://github.com/xjxxxj/cloudNativeStudy/blob/master/homework/code/src/lesson_3/readme.assets/1665594528359.png?raw=true)