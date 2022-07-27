package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func scanFile(cfg *Config) error {
	text, err := os.ReadFile(cfg.TextFilepath)
	if err != nil {
		return fmt.Errorf("cant getFileText; err: %w", err)
	}
	words, err := parseText(text)
	if err != nil {
		return fmt.Errorf("cant parseText; err: %w", err)
	}

	states, err := scanState(cfg.WorkDir + "/" + cfg.StateFile)
	if err != nil {
		return fmt.Errorf("cant scanState; err: %w", err)
	}
	newWords := intersect(words, states)

	translator, err := NewTranslator(cfg.LingvoApiKey)
	if err != nil {
		return fmt.Errorf("cant create translator; err: %w", err)
	}

	translations, err := getTranslations(translator, newWords)
	if err != nil {
		return fmt.Errorf("cant getTranslations; err: %w", err)
	}
	err = saveTranslations(translations)

	return nil
}

func saveTranslations(translations []Translation) error {
	panic("TODO")
	return nil
}

func getTranslations(translator *Translator, words []string) ([]Translation, error) {
	result := make([]Translation, 0, len(words))
	for _, word := range words {
		resp, err := translator.TranslateEng(word)
		if err != nil {
			return nil, fmt.Errorf("cant translate word %s; err: %w", word, err)
		}
		if resp.Heading != "" {
			result = append(result, convertTranslation(resp))
		}
	}
	return result, nil
}

func convertTranslation(resp LingvoliveResponse) Translation {
	return Translation{
		Heading:        resp.Translation.Heading,
		Translation:    resp.Translation.Translation,
		DictionaryName: resp.Translation.DictionaryName,
		SoundName:      resp.Translation.SoundName,
		Type:           resp.Translation.Type,
		OriginalWord:   resp.Translation.OriginalWord,
	}
}

func intersect(words []string, states map[string]State) []string {
	result := []string{}
	for _, w := range words {
		if _, ok := states[w]; !ok {
			result = append(result, w)
		}
	}
	return result
}

func scanState(filepath string) (map[string]State, error) {
	f, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("cant open scanFile %s; err: %w", filepath, err)
	}
	var result map[string]State
	err = json.Unmarshal(f, &result)
	if err != nil {
		return nil, fmt.Errorf("cant unmarshal scanFile %s; err: %w", filepath, err)
	}

	return result, nil
}

func parseText(text []byte) ([]string, error) {
	paragraphs := strings.Split(string(text), "\n")
	res := make(map[string]struct{})
	for _, paragraph := range paragraphs {
		words := strings.Split(string(paragraph), " ")
		for _, w := range words {
			word := normalizeWord(w)
			if word != "" {
				res[word] = struct{}{}
			}
		}
	}

	result := make([]string, 0, len(res))
	for w := range res {
		result = append(result, w)
	}
	return result, nil
}

func normalizeWord(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimFunc(word, func(r rune) bool {
		if unicode.IsSpace(r) {
			return true
		}
		if unicode.IsSymbol(r) {
			return true
		}
		if unicode.IsDigit(r) {
			return true
		}
		for _, w := range []rune{'\'', ',', '.', '!', '?', '/', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '[', ']', '\\', ';', ',', '.', '/', '?', '>', '<', '|', '"', ':', '}', '{', '+', '_', '-', '='} {
			if w == r {
				return true
			}
		}
		return false
	})
	return word
}
