package main

import (
	"keychron_rgb_adapter/internal/server"

	"github.com/caseymrm/menuet"
	"github.com/mishamyrt/go-keychron/pkg/hid"
)

// Port on which server will be listening
const port = 17085

func main() {
	hid.Init()
	go server.Run(port)
	menuet.App().Label = "keychron-adapter.myrt.co"
	menuet.App().SetMenuState(&menuet.MenuState{
		Image: "bulb.pdf",
	})

	menuet.App().RunApplication()
}
