package learngoembed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

/*
=== RUN   TestEmbed

--- PASS: TestEmbed (0.00s)
PASS
ok      learn-go-embed  0.051s
*/

/* Sudah diisi
=== RUN   TestEmbed
Hello World
--- PASS: TestEmbed (0.00s)
PASS
ok      learn-go-embed  0.049s
*/
