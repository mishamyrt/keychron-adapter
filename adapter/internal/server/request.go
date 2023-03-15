package server

type EffectRequest struct {
	Effect     uint8  `json:"effect"`
	Brightness uint8  `json:"brightness"`
	Color      string `json:"color,omitempty"`
	Speed      uint8  `json:"speed,omitempty"`
	Direction  uint8  `json:"direction,omitempty"`
}
