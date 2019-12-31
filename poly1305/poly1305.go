// poly1305 wrapper
// poly1305 is just a one time hash
package poly1305

import (
	"hash"
	"golang.org/x/crypto/poly1305"
)

// TagSize is the size, in bytes, of a poly1305 authenticator.
const Size = 16

// The blocksize of SHA-1 in bytes.
const BlockSize = 64

// digest represents the partial evaluation of a checksum.
type digest struct {
	poly1305 *poly1305.MAC
	// key 
	key [32]byte
	start    int
	end      int
}

func init() {
}

func (d *digest) Reset() {
//	d.poly1305.Reset()
}

func (d *digest) Size() int { return Size }
func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(p []byte) (nn int, err error) {
	d.poly1305.Write(p)
	return

}

func New(key [32]byte) hash.Hash {
	d := new(digest)

        var keyx [32]byte
        copy(keyx[:], key[:32])


	h  := poly1305.New(&keyx)
	d.poly1305 = h
	copy(d.key[:], key[:32])

	return d
}

func (d *digest) Sum(in []byte) []byte {
	// Make a copy of d so that caller can keep writing and summing.
	sl := d.poly1305.Sum(in)

	h  := poly1305.New(&d.key)
	d.poly1305 = h

	return sl
}

