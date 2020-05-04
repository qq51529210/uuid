package uuid

import (
	"crypto/md5"
	"hash"
	"sync"
)

var _md5 sync.Pool

func init() {
	_md5.New = func() interface{} {
		return md5.New()
	}
}

// 版本3，计算名字和名字空间的md5散列值
func (this UUID) V3(namespace, name []byte) {
	hash := _md5.Get().(hash.Hash)
	hash.Reset()
	hash.Write(namespace)
	hash.Write(name)
	copy(this[0:], hash.Sum(nil))
	_md5.Put(hash)
	// version & variant
	this.initVersionAndVariant(0x3f)
}

func V3(namespace, name []byte) string {
	var uuid UUID
	uuid.V3(namespace, name)
	return uuid.String()
}