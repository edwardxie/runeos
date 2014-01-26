package service

import (
	// "fmt"
	"net/http"
	"runeos/ide"
)

func WebStart(port string) error {
	if rr, ok := ide.GetEnv("RUNEOS_ROOT").(string); ok == true {
		http.Handle("/static/", http.StripPrefix("/static/",
			http.FileServer(http.Dir(rr+"/static"))))
		if err := http.ListenAndServe(port, nil); err != nil {
			return err
		}
	}
	return nil
}

func Webapp() {

}
