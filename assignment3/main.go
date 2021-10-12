package assignment3

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logWriter struct{}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("No filename argument given. Exiting...")
	}

	f, err := os.Open(args[1])
	if err != nil {
		log.Fatal("Error occurred while opening file", err)
	}

	lw := logWriter{}
	_, err = io.Copy(lw, f)
	if err != nil {
		log.Fatal("Error occurred while writing file to stdout", err)
	}
}

func (l logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
