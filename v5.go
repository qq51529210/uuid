package uuid

import (
	"crypto/sha1"
	"hash"
	"sync"
)

var _sha1 sync.Pool

func init() {
	_sha1.New = func() interface{} {
		return sha1.New()
	}
}

// 版本3，计算名字和名字空间的sha1散列值
func (u *UUID) V5(namespace, name []byte) {
	hash := _sha1.Get().(hash.Hash)
	hash.Write(namespace)
	hash.Write(name)
	copy(u[0:], hash.Sum(nil))
	_sha1.Put(hash)
	// version & variant
	u.initVersionAndVariant(0x5f)
}

func V5(namespace, name []byte) string {
	var uuid UUID
	uuid.V5(namespace, name)
	return uuid.String()
}
