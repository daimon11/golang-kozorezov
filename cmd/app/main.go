package main

import (
	"fmt"
	"os"

	"homeworks/internal/app"
)

func main() {
	if err := app.Run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
