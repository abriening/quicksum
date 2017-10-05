package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const defaultChunkSize int64 = 1024

func main() {

	chunkSize := flag.Int64("size", defaultChunkSize, "Chunk size")
	flag.Parse()
	args := flag.Args()

	hash := sha1.New()
	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stat, err := f.Stat()
	l := stat.Size()

	if l <= (*chunkSize * 2) {
		_, err = io.Copy(hash, f)
	} else {
		b := make([]byte, *chunkSize)
		f.Read(b)
		_, err = hash.Write(b)
		f.ReadAt(b, l-(*chunkSize))
		_, err = hash.Write(b)
	}
	sizeb := make([]byte, 8)
	_ = binary.Read(bytes.NewReader(sizeb), binary.BigEndian, l)
	_, err = hash.Write(sizeb)
	fmt.Printf("%x\n", hash.Sum(nil))
}
