package main

import "os"

func main() {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	f.ReadByte()
	f.Close()
}