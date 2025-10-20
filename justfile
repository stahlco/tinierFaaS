GO_FILES := `find . -name '*.go' | grep -v /vendor/ | grep -v /ext/ | grep -v _test.go`
SUPPORTED_ARCH := arm64
RUNTIMES := `find pkg/backend/runtimes -name Dockerfile`

