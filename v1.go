package uuid

import (
	"encoding/binary"
	"net"
	"time"
)

var (
	_clock uint16
	_node  [6]byte
)

func init() {
	// init node with MAC address
	addrs, err := net.Interfaces()
	if nil != err {
		panic(err)
	}
	for i := 0; i < len(addrs); i++ {
		if len(addrs[i].HardwareAddr) >= 6 {
			copy(_node[0:], addrs[i].HardwareAddr)
			break
		}
	}
}

// 设置node，默认是取第一个网卡的MAC地址
// 不同node，生成的uuid就不一样
// 虽然说不同的机器应该不会出现网卡MAC地址相同的情况
// 但是服务如果在容器中就难说了
// 所以这个函数可以用，但没必要
func SetNode(node [6]byte) {
	copy(_node[0:], node[0:])
}

// 设置随机的node
func SetRandomNode() {
	_rand.Read(_node[0:])
}

// 版本1，机器的时间戳和Node决定(默认是MAC地址)
func (this UUID) V1() {
	// timestamp
	ts := uint64(time.Now().UTC().UnixNano())
	// time low
	binary.BigEndian.PutUint32(this[0:], uint32(ts))
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
	this.initVersionAndVariant(0x1f)
}

func V1() string {
	var uuid UUID
	uuid.V1()
	return uuid.String()
}
