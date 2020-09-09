package uuid

import (
	"encoding/binary"
	"os"
	"time"
)

var (
	_gid int
	_uid int
)

func init() {
	_gid = os.Getgid()
	_uid = os.Getuid()
}

// 版本2，和版本1相同，但会把时间戳的前4位置换为POSIX的UID或GID
func (this UUID) V2(id int) {
	// timestamp
	ts := uint64(time.Now().UTC().UnixNano())
	// id
	binary.BigEndian.PutUint32(this[0:], uint32(id))
	// time mid
	binary.BigEndian.PutUint16(this[4:], uint16(ts>>32))
	// time high and version
	binary.BigEndian.PutUint16(this[6:], uint16(ts>>48))
	// clock
	_clock++
	binary.BigEndian.PutUint16(this[8:], _clock)
	// node
	copy(this[10:], _node[0:])
	// version & variant
	this.initVersionAndVariant(0x2f)
}

// 版本2，用的是gid
func (this UUID) V2_GID() {
	this.V2(_gid)
}

// 版本2，用的是uid
func (this UUID) V2_UID() {
	this.V2(_uid)
}

func V2(id int) string {
	var uuid UUID
	uuid.V2(id)
	return uuid.String()
}

func V2_GID() string {
	var uuid UUID
	uuid.V2_GID()
	return uuid.String()
}

func V2_UID() string {
	var uuid UUID
	uuid.V2_UID()
	return uuid.String()
}