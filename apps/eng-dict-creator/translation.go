package main

type Translation struct {
	Heading        string `json:"Heading"`
	Translation    string `json:"Translation"`
	DictionaryName string `json:"DictionaryName"`
	SoundName      string `json:"SoundName"`
	Type           int    `json:"Type"`
	OriginalWord   string `json:"OriginalWord"`
}
