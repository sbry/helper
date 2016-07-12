package helper

import (
	"testing"
)

func TestBasic(t *testing.T) {
	filename := "/tmp/test/test2/test.txt"
	f := NewFromPath(filename)
	f.SetCollection("test3")
	f.SetExtension(".html")
	if f.Path() != "/tmp/test/test3/test.html" {
		t.FailNow()
	}
}

func TestReadWriteDelete(t *testing.T) {
	filename := "/tmp/test/test2/test.txt"
	f := NewFromPath(filename)
	wcontent := []byte("Hello")
	f.Write(wcontent)
	rcontent := f.Read()
	if string(wcontent) != string(rcontent) {
		t.FailNow()
	}
	f.Delete()
}
