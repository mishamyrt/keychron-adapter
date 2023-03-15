package main

import (
	"keychron_rgb_adapter/internal/server"

	"github.com/sstallion/go-hid"
)

// Port on which server will be listening
const port = 17085

func main() {
	hid.Init()

	server.Run(port)
}
