package main

import (
	"fmt"
	"crypto/rand"
	"github.com/singularian/mdhash/ultraLargeHash"
)

func main() {

	h := ultraLargeHash.New(0, 512)

	a, _ := GenerateRandomBytes(5000)
	h.Write(a)
	fmt.Printf("\n\nultrahash bytes %x\nhash hex %x\n", a, h.Sum(nil))

	b, _ := GenerateRandomBytes(5000)
	h.Write(b)
	fmt.Printf("\n\nultrahash hex %x\nhash 2 %x\n", b, h.Sum(nil))

	c, _ := GenerateRandomBytes(50)
        h.Write(c)
        fmt.Printf("\n\nultrahash hex %x\nhash 2 %x\n", c, h.Sum(nil))

}

func GenerateRandomBytes(n int) ([]byte, error) {
        b := make([]byte, n)
        _, err := rand.Read(b)

        // Note that err == nil only if we read len(b) bytes.
        if err != nil {
                return nil, err
        }

        return b, nil
}

