## 概述

## 隔离性-Linux namespace

> Linux Namespace是Linux Kernel提供的资源隔离方案

### 容器使用namespace进行资源隔离

- Pid namespace: 每个pid namespace可以有完整的一套进程编号，不同容器内可以有相同的pid
- net namespace： 每个net namespace有独立的network devices,IP addresses, IP routing tables,/proc/net目录，Docker 默认采用veth的方式将container中的虚拟网卡同host上的一个docker bridge:docker0连接起来
- ipc namespace:  每个ipc namespace共享System V IPC和POSIX消息队列，container的进程间交互实际上还是host上具有相同Pid namespace中的进程间交互，因此需要在IPC资源申请时加入namespace信息 - 每个IPC资源有一个唯一的32位ID
- mnt namespace: 允许不同的namespace的进程看到的文件结构不同，每个容器内都是看到一套独立的文件结构，互相修改不影响
- uts namespace:允许每个container拥有独立的hostname和domain name，使其在网络上可以被视作一个独立的节点而非host上的一个进程
- user namespace: 允许每个container可以有不同的user和group id，也就是说可以在container内部用container内部的用户执行程序而非host上的用户

## 可配额/可度量 - Control Groups

> Cgroups是Linux下用于对一个或一组进程进行资源控制和监控的机制

- 一个Cgroups可以绑定多个进程，进行整体的cpu使用时间、内存、磁盘I/O等资源的限制
- Cgroups本身是层级的，每个Cgroup可以包括其他子Cgroup，因此Cgroup能使用的资源除了受本Cgroup配置的资源参数限制，还受父Cgroup设置的资源限制

## 文件系统 - UnisonFs

> 将不同目录挂载到同一个虚拟文件系统下的文件系统

- 每个目录设定readonly、readwrite、whiteout-able权限，对挂载的文件系统进行增删改查动作，均不会影响到readonly权限目录下的文件

- 用途一：将多个disk挂到同一个目录下

- 用途二：将一个readonly的branch和一个writeable的branch联合起来

- 容器的存储驱动：

  ![1665334649924](H:\coding\go_study\cloudNativeStudy\notes\第三模块-docker核心技术.assets\1665334649924.png)

## OCI容器标准

> OCI定义了镜像标准、运行时标准、分发标准

- 镜像标准定义应用如何打包
- 运行时标准定义如何解压应用包并运行
- 分发标准定义如何分发容器镜像

## 网络

### 容器内部网络配置

docker给容器网络提供了以下几种配置策略

- Null(--net=None)

  > - 将容器放入独立的网络空间但不做任何网络配置
  > - 用户需要通过运行docker network命令完成网络配置

- Host

  > 使用主机网络名空间，复用主机网络，相当于占用主机网络的一个ip

- Container

  > 重用其他容器的网络

- Bridge(--net=bridge) -- 默认模式

  > 使用Linux网桥和iptables提供容器互联，Docker在每台主机创建一个名叫docker0的网桥，通过veth pair连连接改主机的每一个EndPoint
  >
  > 相当于每个主机上的容器，有一套网络ip,然后通过网桥来连接主机网络

### 不容主机间容器的网络

- Underlay: 使用现有底层网络，为每个容器配置可路由的网络IP
- Overlay: 通过网络封包实现