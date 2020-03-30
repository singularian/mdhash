package main

import (
	"fmt"
	"github.com/singularian/mdhash/djb2"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }
	djb2o := djb2.New(0, 8)
	djb2o.Write(a)
	fmt.Printf("djb2 hex %x\n", djb2o.Sum(nil))

	djb2o.Reset()
	djb2o.Write(bytes)
	fmt.Printf("djb2 hex %s %x\n", str, djb2o.Sum(nil))

	hash := djb2l(bytes)
	fmt.Printf("djb2 Hash (%s) is %x uint64 %d\n", str, hash, hash)

}

func djb2l(s []byte) uint64 {
        var hash uint64 = 5381

        for _, c := range s {
                hash = ((hash << 5) + hash) + uint64(c)
                // the above line is an optimized version of the following line:
                // hash = hash * 33 + int64(c)
                // which is easier to read and understand...
        }

        return hash
}

