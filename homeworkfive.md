# 作业1(结论)
## 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
### 使用命令(redis-benchmark -t get,set -n 300000 -d 字节数)
| 字节大小 | SET耗时 | SET每秒请求次数 | GET耗时 | GET每秒请求次数 |
| ------- | ------ | ------------- | ------ | ------------  |
| 10 | 4.99s | 60084.12 | 5.12s | 58570.87  |
| 20 | 4.85s | 61830.17 | 5.30s | 56646.53  |
| 50 | 4.94s | 60741.04 | 5.57s | 53908.36  |
| 100 | 4.82s | 62240.66 | 5.05s | 59464.82  |
| 200 | 4.62s | 64892.93 | 4.80s | 62526.05  |
| 1000 | 5.01s | 59916.12 | 5.18s | 57881.54  |
| 5000 | 5.28s | 56839.71 | 5.69s | 52724.08  |
# 作业2(结论)
## 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
### 使用命令(redis-benchmark -h 127.0.0.1 -p 6379 -t set -r 200000 -n 300000 -d 字节数生成20万数据，redis-cli flushdb 清空KEY，redis-cli dbsize获取KEY数量)
#### 设置钱内存1123904(byte)
| 字节大小 | 平均key内存占用(byte) |
| ------- | ------ |
| 10 | 109.48 |
| 20 | 125.50 |
| 50 | 157.50 | 
| 100 | 205.49 |
| 200 | 301.52 |
| 1000 | 1101.53 |
| 5000 | 5213.49 | 
## 总结redis保存的value值跟实际情况不符合，随着value值越来越大，存储值会越来越接近真实值，所以redis不仅存值，应该还会存储一些结构信息等数据