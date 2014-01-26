package ide

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	// "io"
	"log"
	"runeos/fsnotify"
	"sync"
)

type ConfigKind int

const (
	JSON             = `{}`
	XML              = `<>`
	INI              = `[]`
	DIR              = ``
	YAML             = ``
	CONFIG_FILE_NAME = "/runeos."
)

type ConfigType struct {
	Name string
	Type string
	Re   string
}

var (
	cts = [...]*ConfigType{
		{Name: "json", Re: "{}"},
		{Name: "xml", Re: "<>"},
		{Name: "ini", Re: "[]"},
		{Name: "dir", Re: ""},
	}
	_        = fmt.Errorf("format")
	confpath string
)

type Configer interface {
	Open() error
	Load() error
	Get()
	Set()
	All()
	Write() error
	Read()
	Flush()
	Event()
	Close()
}

type config struct {
	lock     sync.RWMutex
	path     string
	file     string
	ext      string
	conftype ConfigType
	data     map[string]interface{}
}

func NewConfig(path string, format string) {

}

func Watcher() (err error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	done := make(chan bool)

	// Process events
	go func(err error) {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case e := <-watcher.Error:
				err = e
			}
		}
	}(err)

	err = watcher.Watch("testDir")
	if err != nil {
		return err
	}

	<-done

	/* ... do stuff ... */
	watcher.Close()
	return nil
}

func toXml() error {
	wx, err := xml.MarshalIndent(ws, "", "\t")
	if err != nil {
		return err
	}
	err = WriteFile(confpath+CONFIG_FILE_NAME+"xml", wx)
	if err != nil {
		return err
	}
	return nil
}

func toJson() error {
	wj, err := json.MarshalIndent(ws, "", "\t")
	if err != nil {
		return err
	}
	err = WriteFile(confpath+CONFIG_FILE_NAME+"json", wj)
	if err != nil {
		return err
	}
	return nil
}

func fromJson() error {
	data, err := ReadFile(confpath + CONFIG_FILE_NAME + "json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &ws)
	if err != nil {
		return err
	}
	return nil
}

func fromXml() error {
	data, err := ReadFile(confpath + CONFIG_FILE_NAME + "xml")
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, &ws)
	if err != nil {
		return err
	}
	return nil
}
