package tests

import (
	"fmt"
	"os"
	"testing"
)

func TestFiles(t *testing.T) {
	for i := 1; i <= 20; i++ {
		t.Run("", func(t *testing.T) {
			a, err := os.ReadFile(fmt.Sprintf("./%.2d", i) + ".a")
			if err != nil {
				t.Fatal(err)
			}
			b, err := os.ReadFile(fmt.Sprintf("./%.2d", i) + ".b")
			if err != nil {
				t.Fatal(err)
			}

			aa := string(a)
			bb := string(b)
			if aa != bb {
				t.Fatal(fmt.Sprintf("not equal file %.2d, \n`%s`\n`%s`", i, aa, bb))
			}
		})

	}
}
