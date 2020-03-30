package main

import (
	"fmt"
	"github.com/singularian/mdhash/sdbm"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }
	sdbmo := sdbm.New(0, 8)
	sdbmo.Write(a)
	fmt.Printf("sdbm object hex %x\n", sdbmo.Sum(nil))

	sdbmo.Reset()
	sdbmo.Write(bytes)
	fmt.Printf("sdbm object hex %s %x size %d\n", str, sdbmo.Sum(nil), len(sdbmo.Sum(nil)))

	hash := sdbml(bytes)
	fmt.Printf("sdbm (%s) is %x uint64 %d\n", str, hash, hash)

}

func sdbml(s []byte) uint64 {
        var hash uint64 = 0 

        for _, c := range s {
		hash = uint64(c) + (hash << 6) + (hash << 16) - hash
        }

        return hash
}

