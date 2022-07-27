package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ruLang = 1033
	enLang = 1049

	bearerURL   = "https://developers.lingvolive.com/api/v1.1/authenticate"
	minicardURL = "https://developers.lingvolive.com/api/v1/Minicard?text=%s&srcLang=%d&dstLang=%d"
)

type Translator struct {
	bearer string
}

func NewTranslator(appKey string) (*Translator, error) {
	t := &Translator{}
	bearer, err := t.getBearer(appKey)
	if err != nil {
		return nil, fmt.Errorf("cant get bearer for appKey %s; err: %w", appKey, err)
	}
	t.bearer = bearer
	return t, nil
}

func (t *Translator) getBearer(key string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, bearerURL, nil)
	if err != nil {
		return "", fmt.Errorf("GET Request `%s` err: %w", bearerURL, err)
	}
	req.Header.Set("Authorization", "Basic "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cant readAll: %w", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return "", fmt.Errorf("Response err: `%s`", string(b))
	}

	return string(b), nil
}

func (t *Translator) TranslateEng(text string) (LingvoliveResponse, error) {
	return t.request(text, ruLang, enLang, t.bearer)
}

func (t *Translator) request(text string, srcLang int, dstLang int, bearer string) (LingvoliveResponse, error) {
	hostURL := fmt.Sprintf(minicardURL, text, srcLang, dstLang)
	req, err := http.NewRequest(http.MethodGet, hostURL, nil)
	if err != nil {
		return LingvoliveResponse{}, fmt.Errorf("GET Request `%s` err: %w", hostURL, err)
	}
	req.Header.Set("Authorization", "Bearer "+bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return LingvoliveResponse{}, fmt.Errorf("cant readAll: %w", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return LingvoliveResponse{}, fmt.Errorf("Response err: `%s`", string(b))
	}

	r := LingvoliveResponse{}
	if strings.Contains(string(b), "No translations found for text") {
		return r, nil
	}
	//err = json.NewDecoder(resp.Body).Decode(&r)
	err = json.Unmarshal(b, &r)
	if err != nil {
		return r, fmt.Errorf("Unmarshal `%s` err: %v", string(b), err)
	}

	return r, nil
}

type LingvoliveResponse struct {
	SourceLanguage int    `json:"SourceLanguage"`
	TargetLanguage int    `json:"TargetLanguage"`
	Heading        string `json:"Heading"`
	Translation    struct {
		Heading        string `json:"Heading"`
		Translation    string `json:"Translation"`
		DictionaryName string `json:"DictionaryName"`
		SoundName      string `json:"SoundName"`
		Type           int    `json:"Type"`
		OriginalWord   string `json:"OriginalWord"`
	} `json:"Translation"`
	SeeAlso []interface{} `json:"SeeAlso"`
}
