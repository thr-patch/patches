package model

type THRL6P struct {
	Data    Data   `json:"data"`
	Schema  string `json:"schema"`
	Version int64  `json:"version"`
}

type Data struct {
	Device        int64 `json:"device"`
	DeviceVersion int64 `json:"device_version"`
	Tone          Tone  `json:"tone"`
	Meta          Meta  `json:"meta"`
}

type Meta struct {
	Name string `json:"name"`
	TNID int64  `json:"tnid"`
}

type Tone struct {
	THRGroupAmp             THRGroupAmp             `json:"THRGroupAmp"`
	THRGroupCab             THRGroupCab             `json:"THRGroupCab"`
	THRGroupFX1Compressor   THRGroupFX1Compressor   `json:"THRGroupFX1Compressor"`
	THRGroupFX3EffectEcho   THRGroupFX3EffectEcho   `json:"THRGroupFX3EffectEcho"`
	THRGroupFX4EffectReverb THRGroupFX4EffectReverb `json:"THRGroupFX4EffectReverb"`
	Global                  Global                  `json:"global"`
}

type THRGroupAmp struct {
	Asset  string  `json:"@asset"`
	Bass   float64 `json:"Bass"`
	Drive  float64 `json:"Drive"`
	Master float64 `json:"Master"`
	Mid    float64 `json:"Mid"`
	Treble float64 `json:"Treble"`
}

type THRGroupCab struct {
	Asset   string `json:"@asset"`
	Enabled bool   `json:"@enabled"`
	Level   string `json:"Level"`
	Sustain string `json:"Sustain"`
}

type THRGroupFX1Compressor struct {
	Asset   string `json:"@asset"`
	Enabled bool   `json:"@enabled"`
}

type THRGroupFX3EffectEcho struct {
	Asset   string `json:"@asset"`
	Enabled bool   `json:"@enabled"`
}

type THRGroupFX4EffectReverb struct {
	Asset   string `json:"@asset"`
	Enabled bool   `json:"@enabled"`
}

type THRGroupGate struct {
	Asset   string `json:"@asset"`
	Enabled bool   `json:"@enabled"`
}

type Global struct {
	THRPresetParamTempo int64 `json:"THRPresetParamTempo"`
}
