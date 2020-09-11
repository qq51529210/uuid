package uuid

import (
	"encoding/binary"
)

type UUID [16]byte

var (
	hexTable = []byte("0123456789ABCDEF")
)

func (u UUID) initVersionAndVariant(n byte) {
	// version
	u[7] = (u[7] & 0x0f) | n
	// variant
	u[8] = u[8]&0x1f | 0x80
}

// uuid其实是一个128位，所以它可以是两个64位的整数，有时候使用两个64位的整数，可能查询性能高一点
func (u UUID) Uint64() (uint64, uint64) {
	return binary.BigEndian.Uint64(u[0:]), binary.BigEndian.Uint64(u[8:])
}

// 返回uuid的字符串，xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// buf的长度必须大于32
func (u UUID) HexWithHyphen(buf []byte) string {
	i, j := 0, 0
	for ; i < 4; i++ {
		buf[j] = hexTable[u[i]>>4]
		j++
		buf[j] = hexTable[u[i]&0x0f]
		j++
	}
	buf[8] = '-'
	j = 9
	for ; i < 6; i++ {
		buf[j] = hexTable[u[i]>>4]
		j++
		buf[j] = hexTable[u[i]&0x0f]
		j++
	}
	buf[13] = '-'
	j = 14
	for ; i < 8; i++ {
		buf[j] = hexTable[u[i]>>4]
		j++
		buf[j] = hexTable[u[i]&0x0f]
		j++
	}
	buf[18] = '-'
	j = 19
	for ; i < 10; i++ {
		buf[j] = hexTable[u[i]>>4]
		j++
		buf[j] = hexTable[u[i]&0x0f]
		j++
	}
	buf[23] = '-'
	j = 24
	for ; i < 16; i++ {
		buf[j] = hexTable[u[i]>>4]
		j++
		buf[j] = hexTable[u[i]&0x0f]
		j++
	}
	return string(buf[0:])
}

// 返回uuid的字符串，没有'-'
// buf的长度必须大于28
func (u UUID) Hex(buf []byte) {
	i, j := 0, 0
	for ; i < 16; i++ {
		buf[j] = hexTable[u[i]>>4]
		j++
		buf[j] = hexTable[u[i]&0x0f]
		j++
	}
}

// 返回uuid的字符串，xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (u UUID) String() string {
	var buf [36]byte
	u.HexWithHyphen(buf[:])
	return string(buf[:])
}
