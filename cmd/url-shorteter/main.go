package main

import (
	"fmt"
	"github.com/saintcriminal/REST_API/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
