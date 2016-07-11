package helper

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// we want abstraction: members are lowercase
type File struct {
	dirParts            []string
	filename, extension string
}

// Constructor
func NewFromPath(absolute_filename string) *File {
	basename := filepath.Base(absolute_filename)
	f := new(File)
	f.dirParts = strings.Split(filepath.Dir(absolute_filename), "/")
	f.extension = filepath.Ext(basename)
	f.filename = strings.TrimSuffix(basename, f.extension)
	return f
}

// the dirname part of the path
func (f *File) Dir() string {
	return strings.Join(f.dirParts, "/")
}

//  the basename part
func (f *File) Base() string {
	return f.filename + f.extension
}

// for setting the last part of the directory which we just call collection
func (f *File) setCollection(collection string) {
	f.dirParts[len(f.dirParts)-1] = collection
}

// setting the Extension for whatever the filetype suggests
func (f *File) setExtension(extension string) {
	f.extension = extension
}

// Get Full Pathname
func (f *File) Path() string {
	return filepath.Join(f.Dir(), f.Base())
}

// Stringer
func (f *File) String() string {
	return f.Path()
}

// Read
func (f *File) Read() []byte {
	bytes, err := ioutil.ReadFile(f.Path())
	check(err)
	return bytes
}

func (f *File) beforeWrite() {
	_, err := os.Stat(f.Dir())
	if os.IsNotExist(err) {
		os.MkdirAll(f.Dir(), 0777)
	}
}

func (f *File) Write(content []byte) {
	f.beforeWrite()
	err := ioutil.WriteFile(f.Path(), content, 0644)
	check(err)
}

func (f *File) Delete() {
	os.RemoveAll(f.Path())
}

func (f *File) DeleteCollection() {
	os.RemoveAll(f.Dir())
}
