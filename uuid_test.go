package uuid

import (
	"testing"
)

func Test_V1(t *testing.T) {
	t.Log(V1(true))
	t.Log(V1(false))
}

func Test_V2(t *testing.T) {
	t.Log(V2gid(true))
	t.Log(V2gid(false))
	t.Log(V2uid(true))
	t.Log(V2uid(false))
}

func Test_V3(t *testing.T) {
	t.Log(V3([]byte("v3-namesapce"), []byte("v3-name"),true))
	t.Log(V3([]byte("v3-namesapce"), []byte("v3-name"),false))
}

func Test_V4(t *testing.T) {
	t.Log(V4(true))
	t.Log(V4(false))
}

func Test_V5(t *testing.T) {
	t.Log(V5([]byte("v5-namesapce"), []byte("v5-name"),true))
	t.Log(V5([]byte("v5-namesapce"), []byte("v5-name"),false))
}

func BenchmarkUUID_V1(b *testing.B) {
	var id UUID
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.V1()
	}
}

func BenchmarkUUID_V2(b *testing.B) {
	var id UUID
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.V2(i)
	}
}

func BenchmarkUUID_V3(b *testing.B) {
	var id UUID
	namespace := []byte("name space")
	name := []byte("name")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.V3(namespace, name)
	}
}

func BenchmarkUUID_V4(b *testing.B) {
	var id UUID
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.V4()
	}
}

func BenchmarkUUID_V5(b *testing.B) {
	var id UUID
	namespace := []byte("name space")
	name := []byte("name")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.V5(namespace, name)
	}
}
