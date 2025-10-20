package backend

import (
	"fmt"
	"log"
	"os"
	"path"
	"tinierFaaS/pkg/manager"
	"tinierFaaS/pkg/util"

	"github.com/docker/docker/client"
	uuid "github.com/google/uuid"
)

const (
	TmpDir = "./tmp"
)

type dockerHandler struct {
	name       string
	uniqueName string
	filePath   string
	client     *client.Client
	network    string
	container  string
	handlerIP  string
}

type DockerBackend struct {
	client *client.Client
	id     string
}

func New(id string) *DockerBackend {

	// Creating Docker Client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("creating docker backend failed with error: %s", err)
		return nil
	}

	return &DockerBackend{
		client: cli,
		id:     id,
	}
}

func (db *DockerBackend) Create(name string, filedir string) (manager.Handler, error) {
	// 1. Create a unique name
	id, err := uuid.NewRandom()
	if err != nil {
		log.Printf("creating uniqueName for function: %s failed with err: %s", name, err)
		return nil, err
	}

	uniqueName := fmt.Sprintf("%s-%s", name, id.String())

	// Create dockerHandler
	handler := &dockerHandler{
		name:       name,
		uniqueName: uniqueName,
		client:     db.client,
		container:  "",
		handlerIP:  "",
	}

	// 2. Create a folder for the function (mkdir <folder>)

	handler.filePath = path.Join(TmpDir, handler.uniqueName)

	err = os.MkdirAll(handler.filePath, 0777)
	if err != nil {
		log.Printf("create ./tmp failed with error: %s", err)
		return nil, err
	}

	// 3. copy Docker specific stuff into the folder (cp runtimes/<env>/* <folder>)
	srcPath := "runtimes/sh/*"
	err = util.CopyAll(srcPath, handler.filePath)
	if err != nil {
		log.Printf("copying %s to %s failed with error: %s", srcPath, TmpDir, err)
		return nil, err
	}
	// 4. copy Function(fn.sh) into a subfolder called fn (cp <file> <folder>/fn)

	// 5. build docker image (docker build -t <image> <folder>)
	// 6. create docker network (docker network create <network> -> <network> = uniqueName)
	// 7. create container (docker run --detach --network <network> --name <name> <image>)
	// 8. delete the folder (rm -rf <folder>)

}

func (db *DockerBackend) Stop() error {
	return nil
}

func (h Handler) IPs() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Start() error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Destroy() error {
	//TODO implement me
	panic("implement me")
}
