package uuid

import (
	"math/rand"
	"testing"
	satori_uuid "github.com/satori/go.uuid"
)

func TestV1(t *testing.T) {
	t.Log(V1())
}

func TestV2(t *testing.T) {
	t.Log(V2Gid())
	t.Log(V2Uid())
	t.Log(V2(111))
}

func TestV3(t *testing.T) {
	t.Log(V3(V1(), "md5"))
}

func TestV4(t *testing.T) {
	t.Log(V4())
}

func TestV5(t *testing.T) {
	t.Log(V5(V2(222), "sha1"))
}

func BenchmarkV1_My(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V1()
	}
}

func BenchmarkV1_Satori(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid, e := satori_uuid.NewV1()
		if nil != e {
			b.Fatal(e)
		}
		uuid.String()
	}
}

func BenchmarkV2_My(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V2(uint32(i))
	}
}

func BenchmarkV2_Satori(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid, e := satori_uuid.NewV2(byte(i % 2))
		if nil != e {
			b.Fatal(e)
		}
		uuid.String()
	}
}

func BenchmarkV3_My(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	id := V1()
	for i := 0; i < b.N; i++ {
		V3(id, "md5")
	}
}

func BenchmarkV3_Satori(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	id, e := satori_uuid.NewV1()
	if nil != e {
		b.Fatal(e)
	}
	for i := 0; i < b.N; i++ {
		uuid := satori_uuid.NewV3(id, "md5")
		uuid.String()
	}
}

func BenchmarkV4_My(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V4()
	}
}

func BenchmarkV4_Satori(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid, e := satori_uuid.NewV4()
		if nil != e {
			b.Fatal(e)
		}
		uuid.String()
	}
}

func BenchmarkV5_My(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	id := V2(rand.Uint32())
	for i := 0; i < b.N; i++ {
		V5(id, "namespace-sha1")
	}
}

func BenchmarkV5_Satori(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	id, e := satori_uuid.NewV2(0)
	if nil != e {
		b.Fatal(e)
	}
	for i := 0; i < b.N; i++ {
		uuid := satori_uuid.NewV5(id, "namespace-sha1")
		uuid.String()
	}
}
