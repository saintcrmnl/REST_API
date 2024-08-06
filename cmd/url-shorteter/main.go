package main

import (
	"fmt"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
