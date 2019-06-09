package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"hash"
	"net"
	"os"
	"math/rand"
	"time"
)

type UUID [16]byte

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

// 设置node，默认是取第一个网卡的MAC地址
// 不同node，生成的uuid就不一样
// 虽然说不同的机器应该不会出现网卡MAC地址相同的情况
// 但是服务如果在容器中就难说了
// 所以这个函数可以用，但没必要
func SetNode(node [6]byte) {
	copy(uuidNode[0:], node[0:])
}

// 设置node，随机的
// 看SetNode
func InitRandomNode() {
	uuidV4Rand.Read(uuidNode[0:])
}

// 版本1，机器的时间戳和Node决定(默认是MAC地址)
func (this *UUID) V1() {
	// timestamp
	ts := uint64(time.Now().UTC().UnixNano())
	// time low
	binary.BigEndian.PutUint32(this[0:], uint32(ts))
	// time mid
	binary.BigEndian.PutUint16(this[4:], uint16(ts>>32))
	// time high and version
	binary.BigEndian.PutUint16(this[6:], uint16(ts>>48))
	// clock
	uuidClock++
	binary.BigEndian.PutUint16(this[8:], uuidClock)
	// node
	copy(this[10:], uuidNode[0:])
	// version & variant
	this.initVersionAndVariant(0x1f)
}

// 版本2，和版本1相同，但会把时间戳的前4位置换为POSIX的UID或GID
func (this *UUID) V2(id uint32) {
	// timestamp
	ts := uint64(time.Now().UTC().UnixNano())
	// id
	binary.BigEndian.PutUint32(this[0:], id)
	// time mid
	binary.BigEndian.PutUint16(this[4:], uint16(ts>>32))
	// time high and version
	binary.BigEndian.PutUint16(this[6:], uint16(ts>>48))
	// clock
	uuidClock++
	binary.BigEndian.PutUint16(this[8:], uuidClock)
	// node
	copy(this[10:], uuidNode[0:])
	// version & variant
	this.initVersionAndVariant(0x2f)
}

// 版本2，用的是gid
func (this *UUID) V2GID() {
	this.V2(uuidV2GID)
}

// 版本2，用的是uid
func (this *UUID) V2UID() {
	this.V2(uuidV2UID)
}

// 版本3，计算名字和名字空间的MD5散列值
func (this *UUID) V3(namespace, name []byte) {
	this.V3WithHash(namespace, name, md5.New())
}

// 版本3，计算名字和名字空间的MD5散列值
// 先写namespace在写name
// 里边会调用hash.Reset()
// 使用指定的hash算法，一般传的是md5.New()，但是可以传其他的
// 一般情况下用在hash pool的情况
func (this *UUID) V3WithHash(namespace, name []byte, hash hash.Hash) {
	hash.Reset()
	hash.Write(namespace)
	hash.Write(name)
	copy(this[0:], hash.Sum(nil))
	// version & variant
	this.initVersionAndVariant(0x3f)
}

// 版本4，用的是随机数，最好不要用这个版本，可能会重复
func (this *UUID) V4() {
	uuidV4Rand.Read(this[0:])
	// version & variant
	this.initVersionAndVariant(0x4f)
}

// 版本5，与版本3一样，但是用的是sha1
func (this *UUID) V5(namespace, name []byte) {
	this.V5WithHash(namespace, name, sha1.New())
}

// 版本5，与版本3一样，但是用的是sha1
func (this *UUID) V5WithHash(namespace, name []byte, hash hash.Hash) {
	hash.Write(namespace)
	hash.Write(name)
	copy(this[0:], hash.Sum(nil))
	// version & variant
	this.initVersionAndVariant(0x5f)
}

func (this *UUID) initVersionAndVariant(n byte) {
	// version
	this[7] = (this[7] & 0x0f) | n
	// variant
	this[8] = this[8]&0x1f | 0x80
}

// uuid其实是一个128位
// 所以它可以是两个64位的整数
// 可以用uuid字符串做主键
// 其实也可以用64位的整数，可能查询性能高一点
func (this *UUID) Int() (uint64, uint64) {
	return binary.BigEndian.Uint64(this[0:]), binary.BigEndian.Uint64(this[8:])
}

// 返回uuid的字符串
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (this *UUID) String() string {
	var buf [36]byte
	Hex(buf, this)
	return string(buf[0:])
}

// 返回uuid的字符串，但是没有'-'
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
func (this *UUID) StringWithoutHyphen() string {
	var buf [32]byte
	HexWithoutHyphen(buf, this)
	return string(buf[0:])
}

// 返回uuid的字符串
// 一般用于缓存
func Hex(buf [36]byte, uuid *UUID) {
	i, j := 0, 0
	for ; i < 4; i++ {
		buf[j] = hexTable[uuid[i]>>4]
		j++
		buf[j] = hexTable[uuid[i]&0x0f]
		j++
	}
	buf[8] = '-'
	j = 9
	for ; i < 6; i ++ {
		buf[j] = hexTable[uuid[i]>>4]
		j++
		buf[j] = hexTable[uuid[i]&0x0f]
		j++
	}
	buf[13] = '-'
	j = 14
	for ; i < 8; i ++ {
		buf[j] = hexTable[uuid[i]>>4]
		j++
		buf[j] = hexTable[uuid[i]&0x0f]
		j++
	}
	buf[18] = '-'
	j = 19
	for ; i < 10; i ++ {
		buf[j] = hexTable[uuid[i]>>4]
		j++
		buf[j] = hexTable[uuid[i]&0x0f]
		j++
	}
	buf[23] = '-'
	j = 24
	for ; i < 16; i ++ {
		buf[j] = hexTable[uuid[i]>>4]
		j++
		buf[j] = hexTable[uuid[i]&0x0f]
		j++
	}
}

// 返回uuid的字符串，没有'-'
// 一般用于缓存
func HexWithoutHyphen(buf [32]byte, uuid *UUID) {
	i, j := 0, 0
	for ; i < 16; i++ {
		buf[j] = hexTable[uuid[i]>>4]
		j++
		buf[j] = hexTable[uuid[i]&0x0f]
		j++
	}
}

// 看UUID.V1()
func V1() string {
	uuid := UUID{}
	uuid.V1()
	return uuid.String()
}

// 看UUID.V2()
func V2(id uint32) string {
	uuid := UUID{}
	uuid.V2(id)
	return uuid.String()
}

// 看UUID.V2GID()
func V2GID() string {
	return V2(uuidV2GID)
}

// 看UUID.V2UID()
func V2UID() string {
	return V2(uuidV2UID)
}

// 看UUID.V3()
func V3(namespace, name []byte) string {
	uuid := UUID{}
	uuid.V3(namespace, name)
	return uuid.String()
}

// 看UUID.V3()
func V3WithHash(namespace, name []byte, hash hash.Hash) string {
	uuid := UUID{}
	uuid.V3WithHash(namespace, name, hash)
	return uuid.String()
}

// 看UUID.V4()
func V4() string {
	uuid := UUID{}
	uuid.V4()
	return uuid.String()
}

// 看UUID.V5()
func V5(namespace, name []byte) string {
	uuid := UUID{}
	uuid.V5(namespace, name)
	return uuid.String()
}

// 看UUID.V5WithHash()
func V5WithHash(namespace, name []byte, hash hash.Hash) string {
	uuid := UUID{}
	uuid.V5WithHash(namespace, name, hash)
	return uuid.String()
}
