# uuid
<h5>使用</h5>
<p>
<pre>
var (
	v3_namespace = []byte("v3-namesapce")
	v3_name      = []byte("v3-name")
	v5_namespace = []byte("v5-namesapce")
	v5_name      = []byte("v5-name")
)

func print(t *testing.T, uuid *UUID) {
	var b1 [36]byte
	Hex(b1, uuid)
	t.Log(string(b1[:]))

	var b2 [32]byte
	HexWithoutHyphen(b2, uuid)
	t.Log(string(b2[:]))
}

func TestV1(t *testing.T) {
	t.Log(V1())

	uuid := UUID{}
	uuid.V1()
	print(t, &uuid)
}

func TestV2(t *testing.T) {
	t.Log(V2(1985))
	t.Log(V2GID())
	t.Log(V2UID())

	uuid := UUID{}
	uuid.V2(1985)
	print(t, &uuid)

	uuid.V2GID()
	print(t, &uuid)

	uuid.V2UID()
	print(t, &uuid)
}

func TestV3(t *testing.T) {
	t.Log(V3(v3_namespace, v3_name))

	uuid := UUID{}
	uuid.V3(v3_namespace, v3_name)
	print(t, &uuid)

	uuid.V3WithHash(v3_namespace, v3_name, md5.New())
	print(t, &uuid)
}

func TestV4(t *testing.T) {
	t.Log(V4())

	uuid := UUID{}
	uuid.V4()
	print(t, &uuid)
}

func TestV5(t *testing.T) {
	t.Log(V5(v5_namespace, v5_name))

	uuid := UUID{}
	uuid.V5(v5_namespace, v5_name)
	print(t, &uuid)

	uuid.V5WithHash(v5_namespace, v5_name, md5.New())
	print(t, &uuid)
}

</pre>
</p>

<h5>测试</h5>
<p>
<pre>
bogon:go ben$ go test -bench=. github.com/qq51529210/uuid
goos: darwin
goarch: amd64
pkg: github.com/qq51529210/uuid
BenchmarkV1_1-4         10000000               150 ns/op              48 B/op          1 allocs/op
BenchmarkV1_2-4         20000000               114 ns/op               0 B/op          0 allocs/op
BenchmarkV2_1-4         10000000               149 ns/op              48 B/op          1 allocs/op
BenchmarkV2_2-4         20000000               111 ns/op               0 B/op          0 allocs/op
BenchmarkV2_3-4         10000000               151 ns/op              48 B/op          1 allocs/op
BenchmarkV2_4-4         10000000               152 ns/op              48 B/op          1 allocs/op
BenchmarkV2_5-4         20000000               117 ns/op               0 B/op          0 allocs/op
BenchmarkV2_6-4         20000000               112 ns/op               0 B/op          0 allocs/op
BenchmarkV3_1-4          5000000               285 ns/op             160 B/op          3 allocs/op
BenchmarkV3_2-4          5000000               246 ns/op             112 B/op          2 allocs/op
BenchmarkV3_3-4         10000000               210 ns/op              16 B/op          1 allocs/op
BenchmarkV4_1-4         20000000                99.4 ns/op            48 B/op          1 allocs/op
BenchmarkV4_2-4         20000000                61.8 ns/op             0 B/op          0 allocs/op
BenchmarkV5_1-4          5000000               301 ns/op             192 B/op          3 allocs/op
BenchmarkV5_2-4          5000000               266 ns/op             144 B/op          2 allocs/op
BenchmarkV5_3-4          5000000               289 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/qq51529210/uuid      30.381s
</pre>
</p>
