

### 简介
由于MaxMind会把部分污染的IP错误归类为中国地址，导致特定情况下分流出错，出现无法成功走代理的情况。而IPIP提供的分类更为准确。
因此[@ipip2mmdb](https://github.com/JMVoid/ipip2mmdb) 将ipip的china_ip_list生成MaxMind数据库中的CN条目，并用于clash分流。

而本项目是ipip2mmdb的全版本。与ipip2mmdb的区别在于, 本项目是将ipip中国IP列表替换掉MaxMind GeoLite mmdb的数据库文件中的中国项，保留其他国家的IP信息。

### 使用
对于Clash用户只需要直接下载Country.mmdb并覆盖原来的Country.mmdb

每月自动从 [ipip](https://raw.githubusercontent.com/17mon/china_ip_list/master/china_ip_list.txt) 抓取china_ip_list进行转换
并从 [@Country.mmdb](https://github.com/Dreamacro/maxmind-geoip/releases/latest/download/Country.mmdb)

### 代码及下载
源代码在master branch
生成mmdb库在release branch下载

### 链接
| 文件 | release分支 |
| ------ | ------ | 
|Country.mmdb | [链接](https://raw.githubusercontent.com/JMVoid/ipip2mmdb/release/Country.mmdb) |

### Credit
[@alecthw](https://github.com/alecthw/mmdb_china_ip_list)
[@Dreamacro](https://github.com/Dreamacro)
[@17mon](https://github.com/17mon)

### 扩展说明
verify代码用来验证在2020.9月发现某些IP会被MaxMind错误归类为中国IP. 