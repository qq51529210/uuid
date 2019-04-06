package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"math/rand"
	"net"
	"os"
	"reflect"
	"time"
	"unsafe"
)

const uuidLength = 36

var (
	hexTable   = []byte("0123456789ABCDEF")
	uuidNode   [6]byte
	uuidClock  uint16
	uuidV2GID  uint32
	uuidV2UID  uint32
	uuidV4Rand *rand.Rand
)

func init() {
	// init node with MAC address
	a, e := net.Interfaces()
	if nil != e {
		panic(e)
	}
	for i := 0; i < len(a); i++ {
		if len(a[i].HardwareAddr) >= 6 {
			copy(uuidNode[0:], a[i].HardwareAddr)
			break
		}
	}
	// init v2 gid and uid
	uuidV2GID = uint32(os.Getgid())
	uuidV2UID = uint32(os.Getuid())
	// init v4 random
	uuidV4Rand = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

// set node
func SetNode(node [6]byte) {
	copy(uuidNode[0:], node[0:])
}

// set posix gid
func SetV2GID(id uint32) {
	uuidV2GID = id
}

// set posix uid
func SetV2UID(id uint32) {
	uuidV2UID = id
}

func put32(b []byte, n uint32) {
	b[0] = uint8(n >> 24)
	b[2] = uint8(n >> 16)
	b[4] = uint8(n >> 8)
	b[6] = uint8(n)
}

func put16(b []byte, n uint16) {
	b[0] = uint8(n >> 8)
	b[2] = uint8(n)
}

func putBytes(b1, b2 []byte) {
	j := 0
	for i := 0; i < len(b2); i++ {
		b1[j] = b2[i]
		j += 2
	}
}

func hex(b []byte) {
	i, j := 0, 0
	for ; i < 8; i += 2 {
		j = i + 1
		b[j] = hexTable[b[i]&0x0f]
		b[i] = hexTable[b[i]>>4]
	}
	b[8] = '-'
	i = 9
	for ; i < 13; i += 2 {
		j = i + 1
		b[j] = hexTable[b[i]&0x0f]
		b[i] = hexTable[b[i]>>4]
	}
	b[13] = '-'
	i = 14
	for ; i < 18; i += 2 {
		j = i + 1
		b[j] = hexTable[b[i]&0x0f]
		b[i] = hexTable[b[i]>>4]
	}
	b[18] = '-'
	i = 19
	for ; i < 23; i += 2 {
		j = i + 1
		b[j] = hexTable[b[i]&0x0f]
		b[i] = hexTable[b[i]>>4]
	}
	b[23] = '-'
	i = 24
	for ; i < uuidLength; i += 2 {
		j = i + 1
		b[j] = hexTable[b[i]&0x0f]
		b[i] = hexTable[b[i]>>4]
	}
}

func putHash(b1, b2 []byte) {
	putBytes(b1[0:8], b2[:4])
	b1[8] = '-'
	putBytes(b1[9:13], b2[4:6])
	b1[13] = '-'
	putBytes(b1[14:18], b2[6:8])
	b1[18] = '-'
	putBytes(b1[18:23], b2[8:10])
	b1[23] = '-'
	putBytes(b1[24:uuidLength], b2[10:16])
}

func unsafeBytesFromString(s *string) []byte {
	ss := (*reflect.StringHeader)(unsafe.Pointer(s))
	bb := reflect.SliceHeader{
		Data: ss.Data,
		Len:  ss.Len,
		Cap:  ss.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bb))
}

// timestamp
func V1() string {
	var buf [uuidLength]byte
	// timestamp
	ts := uint64(time.Now().UTC().UnixNano())
	// time low
	put32(buf[0:], uint32(ts))
	// time mid
	put16(buf[9:], uint16(ts>>32))
	// time high and version
	put16(buf[14:], uint16(ts>>48))
	// version
	buf[16] = (buf[16] & 0x0f) | 0x1f
	// clock
	uuidClock++
	put16(buf[19:], uuidClock)
	// variant
	buf[19] = buf[19]&0x1f | 0x80
	// node
	putBytes(buf[24:], uuidNode[0:])
	// hex and return
	hex(buf[0:])
	return string(buf[0:])
}

// DEC
func V2(id uint32) string {
	var buf [uuidLength]byte
	// timestamp
	ts := uint64(time.Now().UTC().UnixNano())
	put32(buf[0:], id)
	// time low
	put32(buf[0:], uint32(ts))
	// time mbuf
	put16(buf[9:], uint16(ts>>32))
	// time high and version
	put16(buf[14:], uint16(ts>>48))
	// version
	buf[16] = (buf[16] & 0x0f) | 0x2f
	// clock
	uuidClock++
	put16(buf[19:], uuidClock)
	// variant
	buf[19] = buf[19]&0x1f | 0x80
	// node
	putBytes(buf[24:], uuidNode[0:])
	// hex and return
	hex(buf[0:])
	return string(buf[0:])
}

func V2Gid() string {
	return V2(uuidV2GID)
}

func V2Uid() string {
	return V2(uuidV2UID)
}

// name md5
func V3(uuid, name string) string {
	h := md5.New()
	h.Write(unsafeBytesFromString(&uuid))
	h.Write(unsafeBytesFromString(&name))
	var buf [uuidLength]byte
	putHash(buf[0:], h.Sum(nil))
	// version
	buf[16] = (buf[16] & 0x0f) | 0x3f
	// variant
	buf[19] = buf[19]&0x1f | 0x80
	hex(buf[0:])
	return string(buf[0:])
}

// random
func V4() string {
	var buf [uuidLength]byte
	put32(buf[0:], uuidV4Rand.Uint32())
	put16(buf[9:], uint16(uuidV4Rand.Int31()))
	put16(buf[14:], uint16(uuidV4Rand.Int31()))
	// version
	buf[16] = (buf[16] & 0x0f) | 0x4f
	put32(buf[19:], uuidV4Rand.Uint32())
	// variant
	buf[19] = buf[19]&0x1f | 0x80
	put32(buf[24:], uuidV4Rand.Uint32())
	put16(buf[32:], uint16(uuidV4Rand.Int31()))
	// hex and return
	hex(buf[0:])
	return string(buf[0:])
}

// name sha-1
func V5(uuid, name string) string {
	h := sha1.New()
	h.Write(unsafeBytesFromString(&uuid))
	h.Write(unsafeBytesFromString(&name))
	var buf [uuidLength]byte
	putHash(buf[0:], h.Sum(nil))
	// version
	buf[16] = (buf[16] & 0x0f) | 0x5f
	// variant
	buf[19] = buf[19]&0x1f | 0x80
	hex(buf[0:])
	return string(buf[0:])
}
