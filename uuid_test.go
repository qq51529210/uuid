package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"testing"
)

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

func BenchmarkV1_1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V1()
	}
}

func BenchmarkV1_2(b *testing.B) {
	uuid := &UUID{}
	buf := [36]byte{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V1()
		Hex(buf, uuid)
	}
}

func BenchmarkV2_1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V2(uint32(i))
	}
}

func BenchmarkV2_2(b *testing.B) {
	uuid := &UUID{}
	buf := [36]byte{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V2(uint32(i))
		Hex(buf, uuid)
	}
}

func BenchmarkV2_3(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V2GID()
	}
}

func BenchmarkV2_4(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V2UID()
	}
}

func BenchmarkV2_5(b *testing.B) {
	uuid := &UUID{}
	buf := [36]byte{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V2UID()
		Hex(buf, uuid)
	}
}

func BenchmarkV2_6(b *testing.B) {
	uuid := &UUID{}
	buf := [36]byte{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V2GID()
		Hex(buf, uuid)
	}
}

func BenchmarkV3_1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V3(v3_namespace, v3_name)
	}
}

func BenchmarkV3_2(b *testing.B) {
	uuid := &UUID{}
	buf := [36]byte{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V3(v3_namespace, v3_name)
		Hex(buf, uuid)
	}
}

func BenchmarkV3_3(b *testing.B) {
	uuid := &UUID{}
	buf := [36]byte{}
	hash := md5.New()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V3WithHash(v3_namespace, v3_name, hash)
		Hex(buf, uuid)
	}
}

func BenchmarkV4_1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V4()
	}
}

func BenchmarkV4_2(b *testing.B) {
	uuid := &UUID{}
	buf := [36]byte{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V4()
		Hex(buf, uuid)
	}
}

func BenchmarkV5_1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V5(v5_namespace, v5_name)
	}
}

func BenchmarkV5_2(b *testing.B) {
	uuid := &UUID{}
	buf := [36]byte{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V5(v5_namespace, v5_name)
		Hex(buf, uuid)
	}
}

func BenchmarkV5_3(b *testing.B) {
	uuid := &UUID{}
	buf := [36]byte{}
	hash := sha1.New()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.V5WithHash(v5_namespace, v5_name, hash)
		Hex(buf, uuid)
	}
}
