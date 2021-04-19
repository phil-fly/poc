# docker
CVE-2020-15257 漏洞扫描POC/EXP


### 漏洞信息
```
containerd->containerd-shim->runc 的通信模型中，containerd-shim的接口作为abstract unix socket暴露，在docker使用net=host参数启动、与宿主机共享net namespace时，其中的unix socket可以被容器内部访问到，容器中攻击者通过该socket可以通过API控制下游runc进程启动新的恶意镜像，并通过该镜像逃逸。

探测到docker内部有containerd-shim启动的unix socket即可确认存在该逃逸漏洞。通过docker内部/proc/net/unix中匹配固定socket即可判断是否存在漏洞，进一步可以创建shim cliet通过grpc(ttrpc)协议调用API，这里调用shimClient.ShimInfo作为POC是因为这个接口简单，不需要传参，并通过返回值进一步确认该socket可用。
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
