package ide

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	// "os"
)

var (
	_ = fmt.Errorf("format")
)

type File interface{}
type FileInfo struct{}

func WriteFile(file string, data []byte) error {
	if err := ioutil.WriteFile(file, data, 0777); err != nil {
		return err
	}
	return nil
}

func ReadFile(file string) (data []byte, err error) {
	return ioutil.ReadFile(file)
}

func Remove(fn string) error { return nil }

func Rm(path string) error { return nil }

func Rename(oldname, newname string) error { return nil }

func FileList() {}

func FileFilter() {}
