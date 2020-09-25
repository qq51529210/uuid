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
	h := _sha1.Get().(hash.Hash)
	h.Write(namespace)
	h.Write(name)
	copy(u[0:], h.Sum(nil))
	_sha1.Put(h)
	// version & variant
	u.initVersionAndVariant(0x5f)
}

func V5(namespace, name []byte, hyphen bool) string {
	var uuid UUID
	uuid.V5(namespace, name)
	if hyphen {
		return uuid.HexWithHyphen()
	}
	return uuid.Hex()
}
