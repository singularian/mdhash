// This is a fork of murmur3 32 hash 
// http://en.wikipedia.org/wiki/MurmurHash
//
// https://github.com/zhangxinngang/murmur/blob/master/murmur.go
// 

package murmur332

import (
	"hash"
	"encoding/binary"
_	"errors"
)

// The size of a murmur3 checksum in bytes.
const Size = 4 

// The blocksize of murmur3 in bytes.
const BlockSize = 4

const (
	c1 = 0xcc9e2d51
	c2 = 0x1b873593
	c3 = 0x85ebca6b
	c4 = 0xc2b2ae35
	r1 = 15
	r2 = 13
	m  = 5
	n  = 0xe6546b64
)

// digest represents the partial evaluation of a checksum.
type digest struct {
	start    int
	end      int
	seed  uint32
	hash  uint32
	hashBytes []byte
}

func init() {
}

func (d *digest) Reset() {
	d.hashBytes = d.hashBytes[:]
}

func (d *digest) Size() int { return Size }
func (d *digest) BlockSize() int { return BlockSize }

// Write the bytes array and comput the murmur3 hash
func (d *digest) Write(p []byte) (nn int, err error) {

	d.hash = d.murmur3(p)
	binary.BigEndian.PutUint32(d.hashBytes, d.hash)
	return

}

// New returns a new hash.Hash computing the murmur3 checksum. 
func New(start int, end int, seed uint32) hash.Hash {
	d := new(digest)

	r := make([]byte, 4)
	d.hashBytes = r

	d.seed = seed

	d.start = start 
	d.end   = end


	return d
}

// return the murmur3 sum
func (d *digest) Sum(in []byte) []byte {

	return d.hashBytes[d.start:d.end]
}

// compute the murmur3 of a byte array
func (d *digest) murmur3(key []byte) uint32 {
	var hash uint32 = d.seed

	iByte := 0
	for ; iByte+4 <= len(key); iByte += 4 {
		k := uint32(key[iByte]) | uint32(key[iByte+1])<<8 | uint32(key[iByte+2])<<16 | uint32(key[iByte+3])<<24
		k *= c1
		k = (k << r1) | (k >> (32 - r1))
		k *= c2
		hash ^= k
		hash = (hash << r2) | (hash >> (32 - r2))
		hash = hash*m + n
	}

	var remainingBytes uint32
	switch len(key) - iByte {
	case 3:
		remainingBytes += uint32(key[iByte+2]) << 16
		fallthrough
	case 2:
		remainingBytes += uint32(key[iByte+1]) << 8
		fallthrough
	case 1:
		remainingBytes += uint32(key[iByte])
		remainingBytes *= c1
		remainingBytes = (remainingBytes << r1) | (remainingBytes >> (32 - r1))
		remainingBytes = remainingBytes * c2
		hash ^= remainingBytes
	}

	hash ^= uint32(len(key))
	hash ^= hash >> 16
	hash *= c3
	hash ^= hash >> 13
	hash *= c4
	hash ^= hash >> 16

	// 出发吧，狗嬷嬷！
        return hash
}

