package service

import (
	"fmt"
	"net/http"
)

func AddRouter(router string, handler http.Handler) error {
	if _, exist := wr[router]; exist != true {
		return fmt.Errorf("Router name <%s> already exist.", router)
	} else {
		wr[router] = handler
		return nil
	}
}
