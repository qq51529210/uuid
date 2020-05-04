package uuid

import (
	"math/rand"
	"time"
)

var _rand = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

// 版本4，用的是随机数，最好不要用这个版本，可能会重复
func (this UUID) V4() {
	_rand.Read(this[0:])
	// version & variant
	this.initVersionAndVariant(0x4f)
}

func V4() string {
	var uuid UUID
	uuid.V4()
	return uuid.String()
}
