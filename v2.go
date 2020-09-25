package uuid

import (
	"encoding/binary"
	"os"
	"time"
)

var (
	_gid = os.Getgid()
	_uid = os.Getuid()
)

// 版本2，和版本1相同，但会把时间戳的前4位置换为POSIX的UID或GID
func (u *UUID) V2(id int) {
	// timestamp
	ts := uint64(time.Now().UTC().UnixNano())
	// id
	binary.BigEndian.PutUint32(u[0:], uint32(id))
	// time mid
	binary.BigEndian.PutUint16(u[4:], uint16(ts>>32))
	// time high and version
	binary.BigEndian.PutUint16(u[6:], uint16(ts>>48))
	// clock
	_clock++
	binary.BigEndian.PutUint16(u[8:], _clock)
	// node
	copy(u[10:], _node[0:])
	// version & variant
	u.initVersionAndVariant(0x2f)
}

// 版本2，用的是gid
func (u *UUID) V2_GID() {
	u.V2(_gid)
}

// 版本2，用的是uid
func (u *UUID) V2_UID() {
	u.V2(_uid)
}

func V2(id int, hyphen bool) string {
	var uuid UUID
	uuid.V2(id)
	if hyphen {
		return uuid.HexWithHyphen()
	}
	return uuid.Hex()
}

func V2gid(hyphen bool) string {
	var uuid UUID
	uuid.V2_GID()
	if hyphen {
		return uuid.HexWithHyphen()
	}
	return uuid.Hex()
}

func V2uid(hyphen bool) string {
	var uuid UUID
	uuid.V2_UID()
	if hyphen {
		return uuid.HexWithHyphen()
	}
	return uuid.Hex()
}
