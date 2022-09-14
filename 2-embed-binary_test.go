package learngoembed

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed rasyidev.png
var image []byte

func TestCopyBinaryFile(t *testing.T) {
	// copy image jadi rasyidev_copy.png
	err := ioutil.WriteFile("rasyidev_copy.png", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

/*
=== RUN   TestCopyBinaryFile
--- PASS: TestCopyBinaryFile (0.00s)
PASS
ok      learn-go-embed  0.051s
*/
