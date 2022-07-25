package main

import "testing"

func Test_printResult(t *testing.T) {
	tests := []struct {
		name string
		res  word
	}{
		{
			name: "",
			res: word{
				val:          "конспект",
				translations: "synopsis, abstract",
				children: []word{
					{
						val:          "synopsis",
						translations: "онспект, краткий обзор; синопсис",
					},
					{
						val:          "abstract",
						translations: "абстракция, отвлечённое понятие",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printResult(tt.res)
		})
	}
}
