# f5
F5 漏洞扫描POC/EXP


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
![rce](https://github.com/phil-fly/poc/blob/main/f5/img/rce.png)

