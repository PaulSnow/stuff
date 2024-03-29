package collect

import (
	"encoding/binary"
	"fmt"
	"testing"
)

type LXRandom struct {
	cnt   uint64
	state uint64
	seed  [32]byte
}

// Randomize the state and seed of the LXRandom generator
func (r *LXRandom) spin() {
	r.cnt++
	if r.cnt == 1 { //                  First time this instance has been called, so the init
		r.state = 0x123456789ABCDEF0 // Give the state a value
		r.spin()                     // Spin through the generator a few times to get things shaken up
		r.spin()
		r.spin()
	}

	for i := 0; i < 32; i += 8 {
		r.state = r.cnt ^ r.state<<17 ^ r.state>>7 ^ binary.BigEndian.Uint64(r.seed[i:]) // Shake the state
		binary.BigEndian.PutUint64(r.seed[i:], r.state<<23^r.state>>17)
		r.seed[i] ^= byte(r.state) ^ r.seed[i] // Shake the seed
	}
	r.cnt++
}

// Hash
// Return a 32 byte array of random bytes
func (r *LXRandom) Hash() [32]byte {
	r.spin()
	return r.seed
}

// Int
// Return a random int
func (r *LXRandom) Uint() int {
	r.spin()
	i := int(r.state)
	if i < 0 {
		return -i
	}
	return i
}

// Byte
// Return a random byte
func (r *LXRandom) Byte() byte {
	r.spin()
	return r.seed[0]
}

// Slice
// Return a slice of the specified length of random bytes
func (r *LXRandom) Slice(length int) (slice []byte) {
	slice = make([]byte, length)
	for i := 0; i < length; i += 32 {
		r.spin()
		copy(slice[i%32:], r.seed[:])
	}
	return slice
}

func Test_LXRandom(t *testing.T) {
	var r LXRandom
	for i := 0; i < 10000; i++ {
		r.Hash()
	}
	fmt.Printf("%x\n", r.Hash())
}
