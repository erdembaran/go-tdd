package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "8")
}

func main() {
	Countdown(os.Stdout)
}