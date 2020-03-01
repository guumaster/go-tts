package cache

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/gosimple/slug"
)

type FileCache struct {
	BaseDir string
}

func (f *FileCache) Open(filename string) (io.ReadCloser, error) {
	path := f.FullPath(filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *FileCache) Save(path string, content bytes.Buffer) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(content.Bytes())

	return err
}

func (f *FileCache) FullPath(filename string) string {
	return fmt.Sprintf("%s/%s", f.BaseDir, filename)
}

func (f *FileCache) Exist(filename string) bool {
	path := f.FullPath(filename)
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func (f *FileCache) Key(in, lang string) string {
	h := md5.New()
	h.Write([]byte(slug.MakeLang(in, lang)))
	return hex.EncodeToString(h.Sum(nil))
}
