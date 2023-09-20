package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"

)

//go:embed version.txt
var version string

func Test(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/x.txt
//go:embed files/y.txt
//go:embed files/z.txt
var files embed.FS

func TestMultipleFile(t *testing.T) {
	x, _ := files.ReadFile("files/x.txt")
	fmt.Println(string(x))

	y, _ := files.ReadFile("files/y.txt")
	fmt.Println(string(y))

	z, _ := files.ReadFile("files/z.txt")
	fmt.Println(string(z))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
