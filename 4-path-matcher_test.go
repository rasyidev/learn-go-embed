package learngoembed

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed files/*.txt
var files2 embed.FS

func TestPatchMatcher(t *testing.T) {
	dirEntries, _ := files2.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			filename := entry.Name()
			content, err := files2.ReadFile("files/" + filename)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))
		}
	}
}

/*
=== RUN   TestPatchMatcher
Dua, dua, juga sayang ayah
Satu, satu, aku sayang ibu
Tiga, tiga, sayang adik kakak
Satu, dua, tiga, sayang semuanya
--- PASS: TestPatchMatcher (0.00s)
PASS
ok      learn-go-embed  0.735s
*/
