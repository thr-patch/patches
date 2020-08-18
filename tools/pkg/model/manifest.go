package model

type Manifest []Patch

type Patch struct {
	Filename string   `json:"filename"`
	Name     string   `json:"name"`
	Tone     Tone     `json:"tone"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
