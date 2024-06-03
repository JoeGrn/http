package util_test

import (
	"os"
	"testing"

	"github.com/joegrn/http/pkg/util"
)

func TestGetDirectoryArg(t *testing.T) {
	testCases := []struct {
		desc     string
		args     []string
		expected string
	}{
		{
			desc:     "No args",
			args:     []string{"cmd"},
			expected: "",
		},
		{
			desc:     "No --directory arg",
			args:     []string{"cmd", "--file", "file.txt"},
			expected: "",
		},
		{
			desc:     "With --directory arg",
			args:     []string{"cmd", "--directory", "dir"},
			expected: "dir",
		},
		{
			desc:     "With --directory arg in the middle",
			args:     []string{"cmd", "--flag", "value", "--directory", "dir"},
			expected: "dir",
		},
		{
			desc:     "With multiple args including --directory",
			args:     []string{"cmd", "--directory", "dir", "--flag", "value"},
			expected: "dir",
		},
		{
			desc:     "With multiple --directory args, returns first",
			args:     []string{"cmd", "--directory", "dir1", "--directory", "dir2"},
			expected: "dir1",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tC.args

			result := util.GetDirectoryArg()

			if result != tC.expected {
				t.Errorf("expected %q, got %q", tC.expected, result)
			}
		})
	}
}

func TestGzipCompress(t *testing.T) {
	testCases := []struct {
		desc     string
		data     string
		expected string
	}{
		{
			desc:     "Compresses a string",
			data:     "test",
			expected: "\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff*I-.\x01\x04\x00\x00\xff\xff\f~\x7f\xd8\x04\x00\x00\x00",
		},
		{
			desc:     "Compresses an empty string",
			data:     "",
			expected: "\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := util.GzipCompress(tC.data)
			if result != tC.expected {
				t.Errorf("expected %q, got %q", tC.expected, result)
			}
		})
	}
}
