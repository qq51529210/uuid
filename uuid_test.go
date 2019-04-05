package uuid

import (
	"testing"
	satori_uuid "github.com/satori/go.uuid"
)

func TestV1(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(V1())
	}
}

func BenchmarkV1_My(b *testing.B) {
	for i := 0; i < b.N; i++ {
		V1()
	}
}

func BenchmarkV1_Satori(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid, e := satori_uuid.NewV1()
		if nil != e {
			b.Fatal(e)
		}
		uuid.String()
	}
}
