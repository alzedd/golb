package webserver

import (
	"fmt"
	"golb/fsutils"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var fs = afero.NewOsFs()

func TestGeneratePageFilename(t *testing.T) {
	testCases := []struct {
		name     string
		folder   string
		file     string
		expected string
	}{
		{"empty_index", "", "", "index"},
		{"folder_index", "index", "", "index"},
		{"file_index", "", "index", "index"},
		{"file/folder", "posts", "1985-09-12-todolist", "posts/1985-09-12-todolist"},
		{"onlyfolder", "posts", "", "posts/index"},
		{"onlyfile", "onlyfile", "", "onlyfile"},
	}

	files2create := []string{
		"onlyfile.md",
		"index.md",
	}
	for _, f := range files2create {
		_, err := fs.Create(fsutils.GetContentSrcFullPath(f))

		if err != nil {
			t.Fatal(err.Error())
		}
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.name), func(t *testing.T) {
			result := generatePageFilename(tc.folder, tc.file)
			assert.Equal(t, tc.expected, result)
		})
	}

	// file cleanup
	for _, f := range files2create {
		err := fs.Remove(fsutils.GetContentSrcFullPath(f))

		if err != nil {
			t.Fatal(err.Error())
		}
	}

}
