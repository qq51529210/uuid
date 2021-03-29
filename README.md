# uuid
分别实现了标准的5个版本的uuid算法，推荐使用版本1。  
1. 版本1，机器的时间戳和node共同生成。node，默认是取第一个网卡的MAC地址，  
理论上不同的机器，生成的uuid就不一样，但是不排除相同的可能性。  
如果使用这个版本，应用程序通过设置不同的node，就可以生成不一样的uuid。  
2. 版本2，和版本1大致相同，只是把时间戳的前4位置换为POSIX的UID或GID，不推荐。  
3. 版本3，根据指定不同的名字空间+名字（组合唯一），生成md5散列值，不推荐。  
4. 版本4，用的是随机数（随机数可能会重复），不推荐。  
5. 版本5，和版本3一样，根据指定的名字空间和名字，生成sha1散列值，不推荐。  
## 使用  
使用非常简单，具体看uuid_test.go文件
## 测试
```
goos: darwin
goarch: amd64
pkg: github.com/qq51529210/uuid
BenchmarkUUID_V1-4              12868052                88.7 ns/op             0 B/op          0 allocs/op
BenchmarkUUID_V2GID-4           13302963                88.7 ns/op             0 B/op          0 allocs/op
BenchmarkUUID_V2UID-4           13163793                88.4 ns/op             0 B/op          0 allocs/op
BenchmarkUUID_V3-4               5967187               196 ns/op              48 B/op          1 allocs/op
BenchmarkUUID_V4-4              49302597                24.0 ns/op             0 B/op          0 allocs/op
BenchmarkUUID_V5-4               5272489               226 ns/op              48 B/op          1 allocs/op
PASS
ok      github.com/qq51529210/uuid      7.908s
linwenbin@linwenbins-mac uuid % 
PASS
```