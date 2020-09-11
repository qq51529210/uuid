package uuid

import (
	"encoding/binary"
)

// 前16字节：数据，后36字节：格式化后的字符串
type UUID [16 + 36]byte

var (
	hexTable = []byte("0123456789ABCDEF")
)

func (u *UUID) initVersionAndVariant(n byte) {
	// version
	u[7] = (u[7] & 0x0f) | n
	// variant
	u[8] = u[8]&0x1f | 0x80
}

// uuid其实是一个128位，所以它可以是两个64位的整数，有时候使用两个64位的整数，可能查询性能高一点
func (u *UUID) Uint64() (uint64, uint64) {
	return binary.BigEndian.Uint64(u[0:]), binary.BigEndian.Uint64(u[8:])
}

// 返回uuid的字符串，xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// buf的长度必须大于32
func (u *UUID) HexWithHyphen() string {
	i, j := 0, 16
	for ; i < 4; i++ {
		u[j] = hexTable[u[i]>>4]
		j++
		u[j] = hexTable[u[i]&0x0f]
		j++
	}
	u[j] = '-'
	j++
	for ; i < 6; i++ {
		u[j] = hexTable[u[i]>>4]
		j++
		u[j] = hexTable[u[i]&0x0f]
		j++
	}
	u[j] = '-'
	j++
	for ; i < 8; i++ {
		u[j] = hexTable[u[i]>>4]
		j++
		u[j] = hexTable[u[i]&0x0f]
		j++
	}
	u[j] = '-'
	j++
	for ; i < 10; i++ {
		u[j] = hexTable[u[i]>>4]
		j++
		u[j] = hexTable[u[i]&0x0f]
		j++
	}
	u[j] = '-'
	j++
	for ; i < 16; i++ {
		u[j] = hexTable[u[i]>>4]
		j++
		u[j] = hexTable[u[i]&0x0f]
		j++
	}
	return string(u[16:])
}

// 返回uuid的字符串，没有'-'
// buf的长度必须大于28
func (u *UUID) Hex() string {
	i, j := 0, 16
	for ; i < 16; i++ {
		u[j] = hexTable[u[i]>>4]
		j++
		u[j] = hexTable[u[i]&0x0f]
		j++
	}
	// 16+36-4
	return string(u[16:48])
}

// 返回uuid的字符串，xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (u *UUID) String() string {
	return u.HexWithHyphen()
}
