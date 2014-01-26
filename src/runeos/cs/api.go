package cs

import (
	"fmt"
	"runeos/cs/service"
)

func Server(mode string, port string) error {
	switch mode {
	case "web":
		return service.WebStart(port)
	default:
		return fmt.Errorf("No server start mode! ")
	}

}
