package manager

import (
	"log"

	"github.com/docker/docker/client"
)

type ManagerService struct {
	id               string
	backend          DockerBackend
	functionHandlers map[string]Handler
}

type Backend interface {
	Create(name string, filedir string) (Handler, error)
	Stop() error
}

type Handler interface {
	IPs() ([]string, error)
	Start() error
	Destroy() error
}
