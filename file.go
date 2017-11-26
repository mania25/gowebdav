package gowebdav

import (
	"fmt"
	"os"
	"time"
)

type File struct {
	path        string
	name        string
	contentType string
	size        int64
	modified    time.Time
	etag        string
	isdir       bool
}

func (f File) Name() string {
	return f.name
}

func (f File) ContentType() string {
	return f.contentType
}

func (f File) Size() int64 {
	return f.size
}

func (f File) Mode() os.FileMode {
	// TODO check webdav perms
	if f.isdir {
		return 0775 | os.ModeDir
	} else {
		return 0664
	}
}

func (f File) ModTime() time.Time {
	return f.modified
}

func (f File) ETag() string {
	return f.etag
}

func (f File) IsDir() bool {
	return f.isdir
}

func (f File) Sys() interface{} {
	return nil
}

func (f File) String() string {
	if f.isdir {
		return fmt.Sprintf("Dir : '%s' - '%s'", f.path, f.name)
	} else {
		return fmt.Sprintf("File: '%s' SIZE: %d ETAG: %s MODIFIED: %s", f.path, f.size, f.etag, f.modified.String())
	}
}
