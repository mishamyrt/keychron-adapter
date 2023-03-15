package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/mishamyrt/go-keychron"
	"github.com/mishamyrt/go-keychron/pkg/effect"
	"github.com/mishamyrt/go-keychron/pkg/hid"
)

func setEffect(w http.ResponseWriter, r *http.Request) {
	var req EffectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	k, err := openKeyboard()
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
	}

	m, err := effect.Get(req.Effect)
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
	}

	if len(req.Color) > 0 {
		color, err := ParseHex(req.Color)
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
		}
		m.Color = color
	}
	m.Brightness = MapBrightness(req.Brightness)
	if req.Speed != 0 {
		m.Speed = req.Speed
	}

	log.Printf("Applying mode: %v", m)
	k.Set(m)
	k.Close()
}

func checkConnection(w http.ResponseWriter, r *http.Request) {
	k, err := openKeyboard()
	var available string
	if err == nil {
		available = "true"
	} else {
		available = "false"
	}
	log.Printf("Keyboard connected status: %v", available)
	io.WriteString(w, available)
	k.Close()
}

func openKeyboard() (keychron.Backlight, error) {
	return keychron.Open(hid.K3V2Optical)
}

func Run(port int) {
	hid.Init()

	http.HandleFunc("/status", checkConnection)
	http.HandleFunc("/apply", setEffect)

	p := strconv.Itoa(port)

	log.Printf("Starting server on %s", p)
	err := http.ListenAndServe("0.0.0.0:"+p, nil)
	if err != nil {
		panic(err)
	}
}
