# docker
CVE-2020-15257 漏洞扫描POC  附带EXP

### 漏洞信息

```
Containerd 是一个控制 runC 的守护进程，提供命令行客户端和API，用于在一个机器上管理容器。当在docker使用–net=host参数启动，与宿主机共享net namespace时，容器中的攻击者可以绕过访问权限访问 containerd 的控制API 进而导致权限提升，从而实现Docker容器逃逸。

Containerd是行业标准的容器运行时，可作为Linux和Windows的守护程序使用。在版本1.3.9和1.4.3之前的容器中，容器填充的API不正确地暴露给主机网络容器。填充程序的API套接字的访问控制验证了连接过程的有效UID为0，但没有以其他方式限制对抽象Unix域套接字的访问。这将允许在与填充程序相同的网络名称空间中运行的恶意容器（有效UID为0，但特权降低）导致新进程以提升的特权运行。
```

### 影响版本

containerd < 1.4.3

containerd < 1.3.9

### 验证原理

```
探测到docker内部有containerd-shim启动的unix socket即可确认存在该逃逸漏洞。通过docker内部/proc/net/unix中匹配固定socket初步判断是否存在漏洞，进一步可以创建shim cliet通过grpc(ttrpc)协议调用API，这里调用shimClient.ShimInfo作为POC是因为这个接口简单，不需要传参，并通过返回值进一步确认该socket可用。
```





### docker组件版本查看

1、命令行
```
docker version
```
2、docker 接口

```
echo -e "GET /version HTTP/1.0\r\n" | sudo nc -U /var/run/docker.sock
```



