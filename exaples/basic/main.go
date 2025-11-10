package main

import (
	"fmt"

	// Importing shim triggers any future init() registrations automatically.
	_ "github.com/r2pq-suite/r2pq-crypto/shim"
)

func main() {
	fmt.Println("R2PQ wiring OK (shim linked).")
}
