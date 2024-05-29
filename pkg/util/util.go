package util

import (
	"compress/gzip"
	"os"
	"strings"
)

func GetDirectoryArg() string {
	args := os.Args
	for i := 1; i < len(args)-1; i++ {
		if args[i] == "--directory" {
			return args[i+1]
		}
	}

	return ""
}

func GzipCompress(data string) string {
	var b strings.Builder
	w := gzip.NewWriter(&b)
	w.Write([]byte(data))
	w.Close()

	return b.String()
}
