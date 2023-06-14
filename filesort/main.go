package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"

	"github.com/paulsnow/stuff/filesort/collect"
)

// E
// Cheap throw an error because programmer is lazy
func E(err error, s string) {
	if err != nil {
		panic(s)
	}
}

// tx
// Junk transaction to generate
type tx struct {
	hash   [32]byte
	length int
	data   [2000]byte
}

// fill
// create data for a transaction
func (t *tx) fill(i int) {
	t.length = rand.Int()%500 + 100
	randData := sha256.Sum256([]byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)})
	copy(t.data[:], randData[:])
	t.hash = sha256.Sum256(t.data[:t.length])
}

func main() {
	numberTx := 10000
	start := time.Now()
	c, err := collect.NewCollect("./snapshotX") // Get a collect object
	E(err, "didn't create a Collect Object")

	t := new(tx) // Generate transactions for test
	fmt.Println("writing transactions")
	for i := 0; i < numberTx; i++ { //
		if i%1000000 == 0 && i > 0 {
			fmt.Printf("Transactions processed %d in %v", i, time.Since(start))
		}
		t.fill(i)            //
		c.WriteTx(t.data[:]) //

	} //
	fmt.Printf("sorting Indexes %v\n", time.Since(start))
	err = c.SortIndexes() //              Sort all the transactions indexes
	E(err, "failed to sort transactions")
	fmt.Printf("clean up %v\n", time.Since(start))
	err = c.Close()
	E(err, "failed on close")
	fmt.Printf("done in %v\n", time.Since(start))

	c.Open()
	c.TestIndex()
	
	sum := 0
	for i := 0; i < numberTx; i++ {
		tx, hash, err := c.Fetch(i)
		E(err, "failed to fetch")
		h := sha256.Sum256(tx)
		if !bytes.Equal(h[:], hash) {
			fmt.Printf("%8d failed\n", i)
		}

		fmt.Printf("%3d Search for: %x\n", i, hash)

		tx2, h2, err2 := c.Fetch(hash)
		sum += c.GuessCnt

		E(err2, "failed to fetch hash")
		if !bytes.Equal(h2[:], hash) || !bytes.Equal(tx, tx2) {
			fmt.Printf("%x failed\n", h[:8])
		}
	}
	fmt.Printf("average guesses: %6.3f\n", float64(sum)/float64(numberTx))
	fmt.Printf("test complete %v\n", time.Since(start))
}
