package uuid

import (
	"testing"
)

func Test_V1(t *testing.T) {
	t.Log(UpperV1())
	t.Log(LowerV1())
	t.Log(UpperV1WithoutHyphen())
	t.Log(LowerV1WithoutHyphen())
}

func Test_V2GID(t *testing.T) {
	t.Log(UpperV2GID())
	t.Log(LowerV2GID())
	t.Log(UpperV2GIDWithoutHyphen())
	t.Log(LowerV2GIDWithoutHyphen())
}

func Test_V2UID(t *testing.T) {
	t.Log(UpperV2UID())
	t.Log(LowerV2UID())
	t.Log(UpperV2UIDWithoutHyphen())
	t.Log(LowerV2UIDWithoutHyphen())
}

func Test_V3(t *testing.T) {
	namespace := []byte("v3 namespace")
	data := []byte("v3 test data")
	t.Log(UpperV3(namespace, data))
	t.Log(LowerV3(namespace, data))
	t.Log(UpperV3WithoutHyphen(namespace, data))
	t.Log(LowerV3WithoutHyphen(namespace, data))
}

func Test_V4(t *testing.T) {
	t.Log(UpperV4())
	t.Log(LowerV4())
	t.Log(UpperV4WithoutHyphen())
	t.Log(LowerV4WithoutHyphen())
}

func Test_V5(t *testing.T) {
	namespace := []byte("v5 namespace")
	data := []byte("v5 test data")
	t.Log(UpperV5(namespace, data))
	t.Log(LowerV5(namespace, data))
	t.Log(UpperV5WithoutHyphen(namespace, data))
	t.Log(LowerV5WithoutHyphen(namespace, data))
}

func BenchmarkUUID_V1(b *testing.B) {
	var id UUID
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.V1()
	}
}

func BenchmarkUUID_V2GID(b *testing.B) {
	var id UUID
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.V2GID()
	}
}

func BenchmarkUUID_V2UID(b *testing.B) {
	var id UUID
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.V2UID()
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
