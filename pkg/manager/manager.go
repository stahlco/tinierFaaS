package manager

// ManagerService is a Service that gets embedded by a server that handles incomming requests from the RProxy
// the ManagerService handles the creation of a FunctionHandler which is a Container that holds the function.
// ManagerService has 2 major functions, first the upload of a function -> creation, and secondly the deletion of a function.
type ManagerService struct {
	id      string
	backend Backend
}

// Handler is responisble for the handling of the specific function, when uploading a new function, the ManagementService
// will return a Handler for a specific Backend -> I think it will be Docker, but might be Podman.
type Handler interface {
	Start() error
	Stop() error
}

// Backend references to a specific type of Backend, e.g. Docker, Podman, ...
// In the Backend are several Handlers running, which each contain *a single function* -> no multi threading
type Backend interface {
	Create() error
	Destroy() error
}

func New() *ManagerService {
	return nil
}

func (ms *ManagerService) uploadFunction(name string, functiondir string) (Handler, error) {
	return nil, nil
}

func (ms *ManagerService) deleteFunction(name string) error {
	return nil
}
