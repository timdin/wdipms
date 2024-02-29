package main

import (
	"github.com/timdin/wdipms/config"
	"github.com/timdin/wdipms/storage"
)

func main() {
	cfg := config.NewConfig()
	storage.InitStorage(cfg.StorageConfig.Conn)
}
