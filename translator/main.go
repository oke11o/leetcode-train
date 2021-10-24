package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "unexpected exit: %s \n", err)
	}
	fmt.Println("DONE! Thank you for using! See you")
}

func run(args []string, stdout *os.File, stderr *os.File) error {
	cfg, err := readConfig()
	if err != nil {
		return fmt.Errorf("cant read config: %w", err)
	}
	bearer, err := getBearer(cfg.AppKey)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		if text == "" || text == "\n" {
			return nil
		}
		res, err := translateWord(strings.TrimSpace(text), bearer)
		if err != nil {
			return fmt.Errorf("cant translate word `%s` err: %w", text, err)
		}
		printResult(res)
	}
	return nil
}

func printResult(res word) {
	fmt.Printf("\ntext: %s\ntran: %s\n", res.val, res.translations)

	for _, child := range res.children {
		fmt.Printf("\t%s: %s\n", child.val, child.translations)
	}
	fmt.Print("\n")
}

type Config struct {
	AppName string `yaml:"appName"`
	AppKey  string `yaml:"appKey"`
}

const (
	ruLang = 1033
	enLang = 1049

	bearerURL   = "https://developers.lingvolive.com/api/v1.1/authenticate"
	minicardURL = "https://developers.lingvolive.com/api/v1/Minicard?text=%s&srcLang=%d&dstLang=%d"
)

type word struct {
	val          string
	translations string
	children     []word
}

func translateWord(text string, bearer string) (word, error) {
	response, err := request(text, enLang, ruLang, bearer)
	if err != nil {
		return word{}, err
	}
	out := make(chan word, 10)
	transes := strings.Split(strings.Replace(response.Translation.Translation, ";", ",", -1), ",")
	wg := &sync.WaitGroup{}
	wg.Add(len(transes))
	for _, tr := range transes {
		go func(wg *sync.WaitGroup, text, bearer string) {
			defer wg.Done()
			space := strings.TrimSpace(text)
			respons, er := request(space, ruLang, enLang, bearer)
			if er == nil {
				out <- word{val: space, translations: respons.Translation.Translation}
			} else {
				fmt.Println("Lingvolive error. Sorry!")
				// TODO: error handling
			}
		}(wg, tr, bearer)
	}

	closer := make(chan struct{})
	result := word{
		val:          text,
		translations: response.Translation.Translation,
	}
	go func() {
		defer close(closer)
		for second := range out {
			result.children = append(result.children, second)
		}
	}()
	wg.Wait()
	close(out)
	<-closer
	return result, nil
}

func request(text string, srcLang int, dstLang int, bearer string) (LingvoliveResponse, error) {
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

func getBearer(key string) (string, error) {
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

func readConfig() (Config, error) {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return Config{}, fmt.Errorf("read file err: %v", err)
	}
	c := Config{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return Config{}, fmt.Errorf("marshal err: %v", err)
	}

	return c, nil
}
