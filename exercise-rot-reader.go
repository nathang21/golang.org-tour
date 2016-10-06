// https://tour.golang.org/methods/23
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(p []byte) (int, error) {
	n, err := rot13.r.Read(p)
    for i := 0; i < n; i++ {
        if p[i] < 'A' || p[i] > 'z' { // ignore non-ASCII letters
            continue
        }

        if p[i] > 'a' && p[i] < 'a'+13 || p[i] > 'A' && p[i] < 'A'+13 {
            p[i] += 13
        } else {
            p[i] -= 13
        }
    }
    return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
