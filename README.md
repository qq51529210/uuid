# uuid

<h5>时间戳的版本，uuid.V1()</h5>
<p>
默认使用机器MAC地址，可以用SetNode()修改。<br>
MAC地址和和时间戳都决定生成的uuid是否可能出现重复。
</p>

<h5>posix id的版本，uuid.V2()，uuid.V2Uid()，uuid.V2Gid()</h5>
<p>
默认使用机器MAC地址系统gid和uid，可以用SetGid()和SetUid()修改。<br>
时间戳的前4位置换为id，其他与v1相同。
</p>

<h5>命名空间md5的版本，uuid.V3()</h5>
<p>使用的是md5哈希算法，namespace+name的做哈希。</p>

<h5>random版本，uuid.V4()</h5>
<p>使用随机数生成的，会有重复的。</p>

<h5>命名空间sha1的版本，uuid.V5()</h5>
<p>使用的是sha1哈希算法，其他与v3相同。</p>

<h5>使用注意</h5>
<p>在分布式下，最后使用V1/V2，通过设置node的值，不使用MAC，每个uuid服务的就不会一样（保证时间正常）。<br>
V4的话，随机数是有可能重复的！<br>
V3/V5，由于是哈希，只要namespace和name不一样，就是不一样。</p>

<h5>下面是测试</h5>
<pre>
<code>
bogon:go ben$ go test -bench=. github.com/qq51529210/uuid
goos: darwin
goarch: amd64 
pkg: github.com/qq51529210/uuid 
BenchmarkV1_My-4        10000000               143 ns/op 
BenchmarkV1_Satori-4    10000000               238 ns/op 
BenchmarkV2_My-4        10000000               143 ns/op 
BenchmarkV2_Satori-4     5000000               255 ns/op 
BenchmarkV3_My-4         5000000               291 ns/op 
BenchmarkV3_Satori-4     5000000               328 ns/op 
BenchmarkV4_My-4        20000000               103 ns/op 
BenchmarkV4_Satori-4     1000000              1217 ns/op 
BenchmarkV5_My-4         5000000               309 ns/op 
BenchmarkV5_Satori-4     5000000               344 ns/op 
PASS 
ok      github.com/qq51529210/uuid      18.457s 
</cdoe>
</pre>
