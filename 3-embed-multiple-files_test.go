package learngoembed

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed files/satu.txt
//go:embed files/dua.txt
//go:embed files/tiga.txt
var files embed.FS

func TestEmbedMultipleFiles(t *testing.T) {
	satu, err := files.ReadFile("files/satu.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(satu))

	dua, err := files.ReadFile("files/dua.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dua))

	tiga, err := files.ReadFile("files/tiga.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tiga))
}

/*
$ go test -v -run TestEmbedMultipleFiles
=== RUN   TestEmbedMultipleFiles
Satu, satu, aku sayang ibu
Dua, dua, juga sayang ayah
Tiga, tiga, sayang adik kakak
Satu, dua, tiga, sayang semuanya
--- PASS: TestEmbedMultipleFiles (0.00s)
PASS
ok      learn-go-embed  0.692s
*/
