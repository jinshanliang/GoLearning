package main

import (
	"fmt"
)

func main() {
	const (
		Debug = iota
		Info
		Warning
		Error
		Fatal
	)
	fmt.Print(Debug, Info, Warning, Error, Fatal)
}
