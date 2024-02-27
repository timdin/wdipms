package main

import (
	"fmt"

	"github.com/timdin/wdipms/config"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg.Dumps())
}
