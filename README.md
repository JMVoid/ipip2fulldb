

### 简介
本项目是ipip2mmdb的全版本。
由于MaxMind会把部分污染的IP错误归类为中国地址，导致不定时出现无法代理的情况。而IPIP提供的分类更为准确
即将ipip中国地址替换掉MaxMind GeoLite mmdb的数据库文件。
因此这是一个较完善的解决方案，用于将ipip的china_ip_list替换MaxMind数据库中的CN条目，并用于clash分流。

### 使用
对于Clash用户只需要直接下载Country.mmdb并覆盖原来的Country.mmdb

每月自动从 [ipip](https://raw.githubusercontent.com/17mon/china_ip_list/master/china_ip_list.txt) 抓取china_ip_list进行转换

### 代码及下载
源代码在master branch
生成mmdb库在release branch下载

### 链接
| 文件 | release分支 |
| ------ | ------ | 
|Country.mmdb | [链接](https://raw.githubusercontent.com/JMVoid/ipip2mmdb/release/Country.mmdb) |

### Credit
[@alecthw](https://github.com/alecthw/mmdb_china_ip_list)

### 扩展说明
verify代码用来验证在2020.9月发现某些IP会被MaxMind错误归类为中国IP