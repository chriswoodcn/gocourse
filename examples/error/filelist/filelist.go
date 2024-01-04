package filelist

import (
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type PathError string

func (p *PathError) Error() string {
	return p.Message()
}

func (p *PathError) Message() string {
	return string(*p)
}

func HandleFileList(w http.ResponseWriter, r *http.Request) error {
	if strings.Index(r.URL.Path, prefix) == -1 {
		pathError := PathError("path must start with " + prefix)
		return &pathError
	}
	path := r.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	_, _ = w.Write(all)
	return nil
}
