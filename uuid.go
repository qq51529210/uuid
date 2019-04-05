package uuid

import (
	"math/rand"
	"net"
	"sync"
	"time"
)

const (
	uuidLength = 36
)

var (
	uuidV1Timestamp = uint64(time.Now().UnixNano())
	uuidV1Clock     = uint16(rand.Int())
	uuidV1Mux       sync.RWMutex
	uuidV1Node      [6]byte
	hexTable        = []byte("0123456789ABCDEF")
)

func init() {
	initV1Node()
}

// MAC address default
func initV1Node() {
	a, e := net.Interfaces()
	if nil != e {
		panic(e)
	}
	for i := 0; i < len(a); i++ {
		if len(a[i].HardwareAddr) >= 6 {
			copy(uuidV1Node[0:], a[i].HardwareAddr)
			break
		}
	}
}

// set node
func SetV1Node(node [6]byte) {
	copy(uuidV1Node[0:], node[0:])
}

func hex(b []byte, c byte) {
	b[0] = hexTable[c>>4]
	b[1] = hexTable[c&0x0f]
}

// rfc4122,timestamp
func V1() string {
	var uuid [uuidLength]byte
	var clock uint16
	// timestamp
	ts := uint64(time.Now().UnixNano())
	uuidV1Mux.Lock()
	if ts <= uuidV1Timestamp {
		uuidV1Clock ++
	} else {
		uuidV1Timestamp = ts
	}
	clock = uuidV1Clock
	uuidV1Mux.Unlock()
	// time low
	hex(uuid[0:], uint8(ts))
	hex(uuid[2:], uint8(ts>>8))
	hex(uuid[4:], uint8(ts>>16))
	hex(uuid[6:], uint8(ts>>24))
	uuid[8] = '-'
	// time mid
	hex(uuid[9:], uint8(ts>>32))
	hex(uuid[11:], uint8(ts>>40))
	uuid[13] = '-'
	// time high and version
	hex(uuid[14:], uint8(ts>>48))
	hex(uuid[16:], uint8((uint8(ts>>56)&0x0f)|0x1f))
	uuid[18] = '-'
	// clock and variant
	hex(uuid[19:], (uint8(clock>>8)&0x3f)|0x80)
	hex(uuid[21:], uint8(clock))
	uuid[23] = '-'
	// node
	i := 24
	for j := 0; j < len(uuidV1Node); j++ {
		hex(uuid[i:], uuidV1Node[j])
		i += 2
	}
	return string(uuid[0:])
}
