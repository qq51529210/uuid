package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"hash"
	"io"
	"math/rand"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const (
	UUIDLength              = 36
	UUIDLengthWithoutHyphen = UUIDLength - 4
)

var (
	clockSequence int32
	hexLowerTable = []byte("0123456789abcdef")
	hexUpperTable = []byte("0123456789ABCDEF")
	v1Node        [6]byte
	v2GID         = os.Getgid()
	v2UID         = os.Getuid()
	v3Pool        sync.Pool
	v4Rand        = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	v5Pool        sync.Pool
)

func init() {
	// init node with MAC address
	addr, err := net.Interfaces()
	if nil != err {
		panic(err)
	}
	for i := 0; i < len(addr); i++ {
		if len(addr[i].HardwareAddr) >= 6 {
			copy(v1Node[0:], addr[i].HardwareAddr)
			break
		}
	}
	v3Pool.New = func() interface{} {
		return md5.New()
	}
	v5Pool.New = func() interface{} {
		return sha1.New()
	}
}

// Set version-1 node. This value was set with MAC address by default.
func SetV1Node(node [6]byte) {
	copy(v1Node[0:], node[0:])
}

type UUID struct {
	b [20]byte
}

func (d *UUID) V1() {
	timestamp := time.Now().UTC().UnixNano()
	clock := uint16(atomic.AddInt32(&clockSequence, 1))
	// d.b[0:3] time low
	d.b[0] = byte(timestamp >> 24)
	d.b[1] = byte(timestamp >> 16)
	d.b[2] = byte(timestamp >> 8)
	d.b[3] = byte(timestamp)
	// d.b[4:5] time mid
	d.b[4] = byte(timestamp >> 40)
	d.b[5] = byte(timestamp >> 32)
	// d.b[6] version and time high
	d.b[6] = (byte(timestamp>>56) & 0b00001111) | 0b00010000
	// d.b[7] time high
	d.b[7] = byte(timestamp >> 48)
	// d.b[8] variant and clock sequence high
	d.b[8] = (byte(clock>>8) & 0b00011111) | 0b10000000
	// d.b[9] clock sequence low
	d.b[7] = byte(clock)
	// d.b[10:15] node
	copy(d.b[12:15], v1Node[:])
}

func (d *UUID) V2GID() {
	timestamp := time.Now().UTC().UnixNano()
	clock := uint16(atomic.AddInt32(&clockSequence, 1))
	// d.b[0:3] gid
	d.b[0] = byte(v2GID >> 24)
	d.b[1] = byte(v2GID >> 16)
	d.b[2] = byte(v2GID >> 8)
	d.b[3] = byte(v2GID)
	// d.b[4:5] time mid
	d.b[4] = byte(timestamp >> 40)
	d.b[5] = byte(timestamp >> 32)
	// d.b[6] version and time high
	d.b[6] = (byte(timestamp>>56) & 0b00001111) | 0b00010000
	// d.b[7] time high
	d.b[7] = byte(timestamp >> 48)
	// d.b[8] variant and clock sequence high
	d.b[8] = (byte(clock>>8) & 0b00011111) | 0b10000000
	// d.b[9] clock sequence low
	d.b[7] = byte(clock)
	// d.b[10:15] node
	copy(d.b[12:15], v1Node[:])
}

func (d *UUID) V2UID() {
	timestamp := time.Now().UTC().UnixNano()
	clock := uint16(atomic.AddInt32(&clockSequence, 1))
	// d.b[0:3] uid
	d.b[0] = byte(v2UID >> 24)
	d.b[1] = byte(v2UID >> 16)
	d.b[2] = byte(v2UID >> 8)
	d.b[3] = byte(v2UID)
	// d.b[4:5] time mid
	d.b[4] = byte(timestamp >> 40)
	d.b[5] = byte(timestamp >> 32)
	// d.b[6] version and time high
	d.b[6] = (byte(timestamp>>56) & 0b00001111) | 0b00100000
	// d.b[7] time high
	d.b[7] = byte(timestamp >> 48)
	// d.b[8] variant and clock sequence high
	d.b[8] = (byte(clock>>8) & 0b00011111) | 0b10000000
	// d.b[9] clock sequence low
	d.b[7] = byte(clock)
	// d.b[10:15] node
	copy(d.b[12:15], v1Node[:])
}

func (d *UUID) V3(namespace, data []byte) {
	h := v3Pool.Get().(hash.Hash)
	h.Reset()
	h.Write(namespace)
	h.Write(data)
	h.Sum(d.b[:0])
	v3Pool.Put(h)
	// d.b[6] version
	d.b[6] &= 0b00001111
	d.b[6] |= 0b00110000
	// d.b[8] variant
	d.b[8] &= 0b00011111
	d.b[8] |= 0b10000000
}

func (d *UUID) V4() {
	v4Rand.Read(d.b[:])
	// d.b[6] version
	d.b[6] &= 0b00001111
	d.b[6] |= 0b01000000
	// d.b[8] variant
	d.b[8] &= 0b00011111
	d.b[8] |= 0b10000000
}

func (d *UUID) V5(namespace, data []byte) {
	h := v5Pool.Get().(hash.Hash)
	h.Reset()
	h.Write(namespace)
	h.Write(data)
	h.Sum(d.b[:0])
	v5Pool.Put(h)
	// d.b[6] version
	d.b[6] &= 0b00001111
	d.b[6] |= 0b01010000
	// d.b[8] variant
	d.b[8] &= 0b00011111
	d.b[8] |= 0b10000000
}

func (d *UUID) String() string {
	return d.UpperString()
}

func (d *UUID) hex(table, buff []byte) {
	i, j := 0, 0
	for i < 4 {
		buff[j] = table[d.b[i]>>4]
		j++
		buff[j] = table[d.b[i]&0x0f]
		j++
		i++
	}
	buff[j] = '-'
	j++
	for i < 10 {
		buff[j] = table[d.b[i]>>4]
		j++
		buff[j] = table[d.b[i]&0x0f]
		j++
		i++
		buff[j] = table[d.b[i]>>4]
		j++
		buff[j] = table[d.b[i]&0x0f]
		j++
		i++
		buff[j] = '-'
		j++
	}
	for i < 16 {
		buff[j] = table[d.b[i]>>4]
		j++
		buff[j] = table[d.b[i]&0x0f]
		j++
		i++
	}
}

func (d *UUID) hexWithoutHyphen(table, buff []byte) {
	i, j := 0, 0
	for i < 4 {
		buff[j] = table[d.b[i]>>4]
		j++
		buff[j] = table[d.b[i]&0x0f]
		j++
		i++
	}
	for i < 10 {
		buff[j] = table[d.b[i]>>4]
		j++
		buff[j] = table[d.b[i]&0x0f]
		j++
		i++
		buff[j] = table[d.b[i]>>4]
		j++
		buff[j] = table[d.b[i]&0x0f]
		j++
		i++
	}
	for i < 16 {
		buff[j] = table[d.b[i]>>4]
		j++
		buff[j] = table[d.b[i]&0x0f]
		j++
		i++
	}
}

func (d *UUID) UpperString() string {
	var str [36]byte
	d.hex(hexUpperTable, str[:])
	return string(str[:])
}

func (d *UUID) LowerString() string {
	var str [36]byte
	d.hex(hexLowerTable, str[:])
	return string(str[:])
}

func (d *UUID) UpperStringWithoutHyphen() string {
	var str [32]byte
	d.hexWithoutHyphen(hexLowerTable, str[:])
	return string(str[:])
}

func (d *UUID) LowerStringWithoutHyphen() string {
	var str [32]byte
	d.hexWithoutHyphen(hexLowerTable, str[:])
	return string(str[:])
}

func (d *UUID) Upper(w io.Writer) error {
	var str [36]byte
	d.hex(hexUpperTable, str[:])
	_, err := w.Write(str[:])
	return err
}

func (d *UUID) Lower(w io.Writer) error {
	var str [36]byte
	d.hex(hexLowerTable, str[:])
	_, err := w.Write(str[:])
	return err
}

func (d *UUID) UpperWithoutHyphen(w io.Writer) error {
	var str [32]byte
	d.hexWithoutHyphen(hexLowerTable, str[:])
	_, err := w.Write(str[:])
	return err
}

func (d *UUID) LowerWithoutHyphen(w io.Writer) error {
	var str [32]byte
	d.hexWithoutHyphen(hexLowerTable, str[:])
	_, err := w.Write(str[:])
	return err
}

func UpperV1() string {
	var id UUID
	id.V1()
	return id.UpperString()
}

func LowerV1() string {
	var id UUID
	id.V1()
	return id.LowerString()
}

func UpperV1WithoutHyphen() string {
	var id UUID
	id.V1()
	return id.UpperStringWithoutHyphen()
}

func LowerV1WithoutHyphen() string {
	var id UUID
	id.V1()
	return id.LowerStringWithoutHyphen()
}

func UpperV2GID() string {
	var id UUID
	id.V2GID()
	return id.UpperString()
}

func LowerV2GID() string {
	var id UUID
	id.V2GID()
	return id.LowerString()
}

func UpperV2GIDWithoutHyphen() string {
	var id UUID
	id.V2GID()
	return id.UpperStringWithoutHyphen()
}

func LowerV2GIDWithoutHyphen() string {
	var id UUID
	id.V2GID()
	return id.LowerStringWithoutHyphen()
}

func UpperV2UID() string {
	var id UUID
	id.V2UID()
	return id.UpperString()
}

func LowerV2UID() string {
	var id UUID
	id.V2UID()
	return id.LowerString()
}

func UpperV2UIDWithoutHyphen() string {
	var id UUID
	id.V2UID()
	return id.UpperStringWithoutHyphen()
}

func LowerV2UIDWithoutHyphen() string {
	var id UUID
	id.V2UID()
	return id.LowerStringWithoutHyphen()
}

func UpperV3(namespace, data []byte) string {
	var id UUID
	id.V3(namespace, data)
	return id.UpperString()
}

func LowerV3(namespace, data []byte) string {
	var id UUID
	id.V3(namespace, data)
	return id.LowerString()
}

func UpperV3WithoutHyphen(namespace, data []byte) string {
	var id UUID
	id.V3(namespace, data)
	return id.UpperStringWithoutHyphen()
}

func LowerV3WithoutHyphen(namespace, data []byte) string {
	var id UUID
	id.V3(namespace, data)
	return id.LowerStringWithoutHyphen()
}

func UpperV4() string {
	var id UUID
	id.V4()
	return id.UpperString()
}

func LowerV4() string {
	var id UUID
	id.V4()
	return id.LowerString()
}

func UpperV4WithoutHyphen() string {
	var id UUID
	id.V4()
	return id.UpperStringWithoutHyphen()
}

func LowerV4WithoutHyphen() string {
	var id UUID
	id.V4()
	return id.LowerStringWithoutHyphen()
}

func UpperV5(namespace, data []byte) string {
	var id UUID
	id.V5(namespace, data)
	return id.UpperString()
}

func LowerV5(namespace, data []byte) string {
	var id UUID
	id.V5(namespace, data)
	return id.LowerString()
}

func UpperV5WithoutHyphen(namespace, data []byte) string {
	var id UUID
	id.V5(namespace, data)
	return id.UpperStringWithoutHyphen()
}

func LowerV5WithoutHyphen(namespace, data []byte) string {
	var id UUID
	id.V5(namespace, data)
	return id.LowerStringWithoutHyphen()
}
