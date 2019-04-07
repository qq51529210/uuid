# uuid

<h5>时间戳的版本，V1</h5>
<p>
默认使用机器MAC地址，可以用SetNode()修改。<br>
MAC地址和和时间戳都决定生成的uuid是否可能出现重复。
</p>

<h5>posix id的版本，V2</h5>
<p>
默认使用机器MAC地址系统gid和uid，可以用SetGid()和SetUid()修改。<br>
时间戳的前4位置换为id，其他与v1相同。
</p>

<h5>命名空间md5的版本，V3</h5>
<p>使用的是md5哈希算法，namespace+name的做哈希。</p>

<h5>random版本，V4</h5>
<p>使用随机数生成的，会有重复的。</p>

<h5>命名空间sha1的版本，V5</h5>
<p>使用的是sha1哈希算法，其他与v3相同。</p>

<h5>其他不适标准的版本</h5>
<p>自己实现了，比如facebook</p>

<h5>使用注意</h5>
<p>在分布式下，最后使用V1/V2，通过设置node的值，不使用MAC，每个uuid服务的就不会一样（保证时间正常）。<br>
V4的话，随机数是有可能重复的！<br>
V3/V5，由于是哈希，只要namespace和name不一样，就是不一样。</p>

<h5>使用方法</h5>
<pre>
<code>
// 不使用MAC地址，针对v1,v2
var new_node [6]byte
uuid.SetNode(new_node)
v1id := uuid.V1()
// 使用指定的id
v2id1 := uuid.V2(random.uint32())
// 使用默认系统id
v2id2 := uuid.V2Gid()
v2id3 := uuid.V2Uid()
// md5，namespace和name不同就会不一样
v3id := uuid.V3("specify namespace", "specify name")
// 随机数，可能有重复哦
v4id := uuid.V4()
// sha1，同v3
v5id := uuid.V3("specify namespace", "specify name")
</cdoe>
</pre>

<h5>下面是测试</h5>
<pre>
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
</pre>
