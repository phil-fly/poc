# struts2
struts2 漏洞扫描POC/EXP

编号:
	CVE-2017-5638
漏洞说明:
	Apache Struts 2被曝存在远程命令执行漏洞，漏洞编号S2-045，CVE编号CVE-2017-5638，在使用基于Jakarta插件的文件上传功能时，有可能存在远程命令执行，导致系统被黑客入侵。

恶意用户可在上传文件时通过修改HTTP请求头中的Content-Type值来触发该漏洞，进而执行系统命令。

影响范围

Struts 2.3.5 – Struts 2.3.31 Struts 2.5 – Struts 2.5.10

### 编 译
```
cd poc/struts2/
go build
```
  
