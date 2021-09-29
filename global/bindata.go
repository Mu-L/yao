// Code generated for package global by go-bindata DO NOT EDIT. (@generated)
// sources:
// .tmp/data/ui/index.html
// .tmp/data/xiang/.DS_Store
// .tmp/data/xiang/apis/README.md
// .tmp/data/xiang/apis/table.http.json
// .tmp/data/xiang/apis/user.http.json
// .tmp/data/xiang/apis/xiang.http.json
// .tmp/data/xiang/flows/README.md
// .tmp/data/xiang/flows/menu.flow.json
// .tmp/data/xiang/flows/table/get.flow.json
// .tmp/data/xiang/models/README.md
// .tmp/data/xiang/models/menu.json
// .tmp/data/xiang/models/user.json
// .tmp/data/xiang/tables/README.md
// .tmp/data/xiang/tables/user.json
package global

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _uiIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x48\xcd\xc9\xc9\x57\x28\xcf\x2f\xca\x49\x51\x30\xe2\x02\x04\x00\x00\xff\xff\x65\x70\x87\xe8\x0e\x00\x00\x00")

func uiIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_uiIndexHtml,
		"ui/index.html",
	)
}

func uiIndexHtml() (*asset, error) {
	bytes, err := uiIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/index.html", size: 14, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangDs_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x41\x4b\xc3\x30\x14\x80\x5f\xba\xaa\x29\x43\xec\x41\x70\xc7\x5e\xbc\xed\xb0\x0d\x11\x77\x10\x4a\x9d\x07\x6f\x42\x45\x0f\x53\xb6\x96\x0e\x57\xa8\xcd\x68\x3b\x8b\xd6\xca\xfe\x86\xfe\x23\x7f\x81\x7f\x47\x6a\x9e\x12\xe6\xf4\x26\x4e\x7d\x1f\x8c\x2f\x6b\x9a\xb4\xaf\x29\xe9\x4b\x00\x80\x39\xd3\xa0\x0d\x60\x02\x00\x07\x69\xcd\x80\x85\x70\xfc\x7d\x40\x53\xcc\xaa\x3e\xfc\x3c\x9d\xf8\x91\xf0\x01\xe0\x69\x71\x5f\xc4\x92\x51\x8d\x9d\x0e\x1e\x4c\x20\x84\x54\x1d\x3f\x7f\x12\x85\x69\xd6\x6a\x3d\x33\xad\xa6\xaf\xac\xae\x71\xce\xeb\x7c\x9d\x5f\xb8\x63\x91\xbb\x99\x97\x4d\x53\xc7\x4b\xfa\xd5\xbf\x63\x2f\x1b\xfb\x58\x3e\x11\x22\x7a\x2f\x7b\xfe\x69\x38\xca\x07\xe6\xe6\x81\x88\x33\x2f\x8c\x47\xc9\x6b\xe3\x30\x18\xf9\x5e\x72\x7e\x16\xc6\x81\xc8\x1d\x31\x8d\x83\xb4\xaf\x54\x18\x86\xc1\x8d\x81\xd9\x28\x8a\xbd\x4e\xab\x69\xed\x76\x3b\x65\xd3\x2a\xba\x55\x79\xa7\xdb\x29\x4b\x83\x6f\x6d\xb7\xf7\x8f\x86\x57\x37\xb7\xc5\x5d\x79\xff\x20\xe3\x60\x0c\x03\xda\x98\x0b\xf0\x51\x0d\xf0\xda\x4d\xe2\x48\xc4\x97\xf2\x85\x25\x08\x82\xf8\xaf\xe0\x14\xc8\xeb\x3f\x7d\x23\x04\x41\x2c\x1d\xd5\xfc\x60\xa1\x6d\xf4\x4c\x9a\x61\xbd\x86\xd6\x95\x36\x26\xda\x42\xdb\xe8\x99\x34\xc3\xf3\x34\xb4\x8e\xe6\x68\x13\x6d\xa1\x6d\xf4\x4c\x1a\x27\x2d\x86\x8b\x0f\x86\x57\x66\xb8\x42\x61\x26\xda\x42\xdb\xdf\xf3\x6c\x08\xe2\xb7\x53\x93\x32\xab\xef\xff\xe1\xe7\xeb\x7f\x82\x20\xfe\x30\x4c\xef\xb9\x3d\xe7\x8b\x3d\x11\x0d\x13\x81\xe1\x5b\x83\xb9\x44\x00\x94\x24\x40\x93\x9b\x85\x0d\xe5\x38\x25\x02\x04\xb1\x64\xbc\x04\x00\x00\xff\xff\x95\x71\x05\x5b\x04\x18\x00\x00")

func xiangDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_xiangDs_store,
		"xiang/.DS_Store",
	)
}

func xiangDs_store() (*asset, error) {
	bytes, err := xiangDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangApisReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x70\x0c\xf0\x54\x78\xd6\xb7\xf4\x69\xff\x62\x2e\x40\x00\x00\x00\xff\xff\xcc\x30\x5e\xae\x0d\x00\x00\x00")

func xiangApisReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_xiangApisReadmeMd,
		"xiang/apis/README.md",
	)
}

func xiangApisReadmeMd() (*asset, error) {
	bytes, err := xiangApisReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/apis/README.md", size: 13, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangApisTableHttpJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x96\xcd\xaa\xd3\x40\x14\xc7\xf7\x79\x8a\x61\xe8\x32\x4d\xa2\xcb\xec\x04\x3f\xe8\xca\x42\x15\x17\xa5\x8b\x63\xe6\x98\x8c\xa4\x93\x71\x66\x52\xad\xa5\x8f\xa0\xae\xdc\x2b\x22\xb8\x12\x1f\xc0\xd7\xb9\xe1\x3e\xc6\x65\x26\x69\xcb\xcd\xfd\x68\x69\xb8\x34\xdd\xb4\x9c\x4f\xce\xef\xff\x2f\x43\x57\x1e\x21\x54\xc0\x1c\x69\x4c\x68\xf5\xfd\x5f\xf5\xe5\xef\xe5\xcf\x3f\xd5\x8f\xff\xd5\xd7\xdf\x17\xdf\x7e\x51\xdf\xd6\x17\xa8\x34\x2f\x84\x6d\x79\x14\x44\x41\x54\x67\x19\xea\x44\x71\x69\x9a\xca\xcd\xe1\x27\xe3\x51\xdd\x99\xaa\xa2\x94\xb6\xe7\x13\x07\x91\x86\x06\xde\xe6\xd8\x54\x4a\x50\xcc\x56\xb8\x18\x4a\x55\x24\xa8\x75\x5d\x90\x60\x32\x4d\x63\x32\xf5\x08\x21\x64\xe5\x3e\x9b\xb4\x6d\x0f\x63\x7b\x73\xa8\x11\x54\x92\xb9\x09\x57\x9e\xa3\xc9\x0a\xb7\xef\xc5\xb3\x57\xbb\xf4\x66\xf3\xe6\x82\xc0\x5d\x10\x4c\x5a\xd3\xdc\x72\x4c\xe9\x40\x82\x82\x79\xe0\x44\xf1\x09\x8d\x3f\x94\xa8\x96\x43\x97\xb4\xf1\xc0\xc5\x81\x84\x14\x5b\xa1\xe6\x9f\x91\xce\xb6\xeb\x8a\xd2\xd0\x78\x7b\x3a\x21\x54\x1b\x30\xa5\xbd\xe2\x71\x14\xf9\xbb\xb4\x59\x4a\xa7\x3e\x48\x99\xf3\x04\xac\x9e\xe1\x7b\x5d\x08\xda\xb4\xac\xdd\xf7\xda\xbf\x57\x89\x77\x5c\xb0\x30\xe6\xec\x58\x2d\x9e\x73\xc1\xf6\x29\xd1\x84\x9c\x9d\x0a\x52\xc3\x02\xef\x82\x1c\xbf\x9c\x1c\xe0\x38\x2c\x70\xaf\xdf\x12\x96\x79\x01\x27\x83\x64\x98\xa3\xe9\x86\xf9\xd4\xad\xe8\xbd\x9d\x5c\x68\x54\xe6\x68\xca\x51\x6b\xbc\xd7\x76\x7e\xcc\x50\x61\x47\x43\xdf\x5c\xdf\x71\xd8\x73\xd5\x13\x7e\x2e\x3a\xc2\x8f\xc4\xde\xdf\x73\xfd\x12\x73\xa6\xfb\xc0\x5d\x4a\x06\x9d\x7d\x7f\xed\x96\x9c\xa5\xef\x0d\x7f\x07\xdf\x6b\xf8\x73\xf3\x5d\xa3\x31\x5c\xa4\xc7\xff\x27\x69\x8d\xdf\xc2\xfc\x60\x64\x1e\x21\x33\x6f\xed\x5d\x05\x00\x00\xff\xff\xe7\x21\x1f\xc4\x14\x0a\x00\x00")

func xiangApisTableHttpJsonBytes() ([]byte, error) {
	return bindataRead(
		_xiangApisTableHttpJson,
		"xiang/apis/table.http.json",
	)
}

func xiangApisTableHttpJson() (*asset, error) {
	bytes, err := xiangApisTableHttpJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/apis/table.http.json", size: 2580, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangApisUserHttpJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x92\xbf\x4e\xf3\x30\x14\xc5\xf7\x3c\xc5\x95\xf5\x8d\x69\x92\x8f\x31\x1b\x42\x08\x55\x42\xa2\x82\x6e\x55\x07\x93\xde\x36\xee\x9f\xd8\xd8\x0e\x10\x55\xdd\x99\x60\x61\x67\x40\x48\xec\x0c\x0c\x7d\x1e\xda\xd7\x40\x76\x4c\x42\xab\x22\xa4\x0a\xca\x92\x28\xe7\x1c\x9f\xdc\x9f\xed\xa9\x07\x40\x32\x3a\x41\x12\x03\x59\xbe\xcc\x97\xf3\x87\xe5\xfd\xf3\xe2\xe6\x75\x71\xfb\xf4\x76\xf7\x48\x7c\xe3\x5f\xa2\x54\x8c\x67\x26\xf2\x3f\x88\x82\xa8\x54\x7b\xa8\x12\xc9\x84\x76\xce\xe7\xc5\xfb\xad\x66\x99\x19\x48\x9e\x0b\xe3\x5e\x33\x9a\x0d\xc2\x5c\xa1\x74\x46\x4e\x65\xcf\x18\xe7\x48\x25\xca\xc6\xf0\x4a\x97\x86\xa0\x3a\x55\x24\x86\x8e\x07\x00\x30\xb5\x4f\x27\x9b\x78\x98\x50\xa1\x93\x94\xda\xb0\x75\x26\xa8\x53\x6e\xab\x8e\x0e\xdb\xb5\x5c\xfd\xa0\x51\x6b\x42\xf2\x04\x95\xaa\xe6\x09\xcc\x3c\xc1\xc1\x7a\x23\x33\x3c\x1d\x12\x5f\xe4\x28\x0b\xd2\xad\x74\x9e\x6b\x12\x57\x23\x01\x10\xa5\xa9\xce\x4d\xdd\x5e\x14\xf9\xb5\xac\x0b\x61\xb7\x93\x0a\x31\x66\x09\x35\x1b\x14\x0e\x15\xcf\x88\x8b\xcc\xec\x7b\xe6\x7f\x41\x38\xe6\x03\x96\x6d\xe2\x6b\x9d\x9c\x6d\x05\x78\xbc\x5a\xf8\x81\x27\x68\x31\xe6\xb4\xb7\x7b\x40\xcd\x47\xf8\xa3\x80\xed\xd5\x42\x07\xf8\xcf\x01\x06\x23\x2c\x88\x0f\xf5\xb7\xc2\x44\xa2\xde\x3d\x37\xcb\xfa\xfc\xdb\x7b\xbb\x99\xb0\xb9\xb2\xb4\x04\xfc\x9b\x73\x0b\x25\xf6\x25\xaa\x74\x4b\x10\x7b\x54\xa7\xeb\x15\xbf\x0c\xe4\x01\x74\xbd\x99\xf7\x1e\x00\x00\xff\xff\x78\xa3\x97\x3c\xec\x04\x00\x00")

func xiangApisUserHttpJsonBytes() ([]byte, error) {
	return bindataRead(
		_xiangApisUserHttpJson,
		"xiang/apis/user.http.json",
	)
}

func xiangApisUserHttpJson() (*asset, error) {
	bytes, err := xiangApisUserHttpJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/apis/user.http.json", size: 1260, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangApisXiangHttpJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x90\x3f\x4e\xc3\x30\x14\xc6\x77\x9f\xe2\xc9\x73\x9a\x06\xc6\x6c\x0c\x08\x75\xeb\xc0\x56\x75\x78\x4d\x4d\xea\x2a\xb5\x2d\xfb\x85\x3f\xaa\x72\x05\x38\x02\x03\xe2\x04\x8c\x3d\x0f\xf4\x1a\xe8\x39\x56\x42\x06\x36\x58\x6c\xf9\xf7\x7d\xfe\xf4\xbe\x77\x14\x00\xd2\xe0\x41\xc9\x12\xe4\xf9\xe3\x74\x3e\xbd\x7e\x3d\xbf\x7f\xbe\xbc\xc9\x8c\x95\x7b\xe5\x83\xb6\x86\xc5\x8b\xbc\xc8\x8b\x9e\x6e\x55\xa8\xbc\x76\x94\x94\x9f\xdf\xae\x96\x8b\xde\x53\x7b\xdb\x3a\x56\x1f\x35\x9a\x3a\xb1\x16\xfd\x96\xd9\x46\xa1\x57\x7e\xb6\x7f\xa0\x5e\x70\x48\xbb\x20\x4b\x58\x09\x00\x80\x63\x3c\x13\x66\xfb\xdc\xe9\x14\x11\xf1\x41\xd1\xce\xc6\x9c\x9b\xeb\xdb\x11\x0f\xe9\xb3\x91\x39\x6f\x2b\x15\xc2\x30\x47\x5e\x37\x76\x83\x4d\xbe\x9c\x04\x6a\xae\xb1\x5a\x0f\x6f\xdb\x92\x2c\x87\x31\x00\x64\x20\xa4\x96\x53\x2e\x8b\x22\x1b\x31\x3d\xb9\xb8\x37\x74\xae\xd1\x15\xf2\x3e\xe6\xfb\x60\x8d\x4c\x96\x2e\xde\x5d\xf6\x4b\x2b\x6d\xee\xec\x9f\xb6\x5a\x4c\x02\xff\xb9\x95\x00\x58\x8b\x4e\x7c\x07\x00\x00\xff\xff\x3e\xc1\xb8\x22\x41\x02\x00\x00")

func xiangApisXiangHttpJsonBytes() ([]byte, error) {
	return bindataRead(
		_xiangApisXiangHttpJson,
		"xiang/apis/xiang.http.json",
	)
}

func xiangApisXiangHttpJson() (*asset, error) {
	bytes, err := xiangApisXiangHttpJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/apis/xiang.http.json", size: 577, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangFlowsReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x78\xba\x7d\xe9\x93\xbd\x73\x9e\x6d\x6d\xe4\x02\x04\x00\x00\xff\xff\x8c\xe8\x04\x3a\x0c\x00\x00\x00")

func xiangFlowsReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_xiangFlowsReadmeMd,
		"xiang/flows/README.md",
	)
}

func xiangFlowsReadmeMd() (*asset, error) {
	bytes, err := xiangFlowsReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/flows/README.md", size: 12, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangFlowsMenuFlowJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x93\xb1\x6e\x83\x30\x10\x86\x77\x9e\xc2\x3a\x75\x8c\x10\xed\x98\x57\xa9\xa2\xca\x98\x53\xb0\x62\x6c\xea\x33\x69\x2b\xc4\xde\xa5\xea\x54\x75\xe9\xd2\xad\x2f\x16\x29\x8f\x51\xd9\x40\x82\x81\x30\x20\xf9\xff\xef\xce\xe7\xcf\xe7\x36\x61\x0c\x14\xcf\x51\xc1\x96\xc1\xf9\xf3\xe7\xf4\xf1\x75\x7a\xff\x3e\xff\xfe\xc1\xc6\x5b\x47\xb4\x24\x8d\xf6\xe6\x7d\x9a\xa5\x59\xaf\x16\x48\xc2\xca\xda\x0d\xce\x32\x4d\x9b\x02\x09\xb6\xec\x31\x61\x8c\xb1\x36\xfc\xbd\xcc\x2b\xf4\x09\x15\xea\x86\x42\x64\x90\x6b\x6b\x04\x12\x05\xc7\x14\xa8\x28\x7d\x95\x5c\xef\x53\x1f\x96\xee\xd1\x5d\x23\xb9\xdd\x5f\xcb\x4e\x4b\x07\x97\x50\xa1\x70\x91\x1f\x74\x59\x5c\x2a\x4c\x1b\x99\x69\x52\x18\x3d\xd7\x6a\x6e\x51\xbb\xa5\xea\xca\xb9\x96\x2b\x23\x0e\x34\x57\x8f\x92\x64\xae\xf0\xc9\x9f\x04\x26\xd6\x6e\x1a\x07\x2f\xd2\x95\xfe\x58\x6d\x9c\x2c\x4a\xa9\x0a\x8b\x7a\xe1\x30\x06\xcf\x0d\xda\xb7\x15\xc3\xdf\xa6\xac\xa4\x87\xf0\x90\x65\xd9\x66\x69\xdf\x80\x34\x42\x98\xa3\x1a\xf4\x15\x60\x63\xc6\x12\xdb\xe0\xac\xc2\xbb\x78\x0b\x84\x83\xb3\x0a\x72\xf0\x6e\xe1\xec\xbf\xdd\x4c\xe9\x92\x5b\xab\x2e\xc6\x5f\xa2\x45\x5a\x10\x69\x19\x08\xa3\x9a\x2a\x8c\x38\x39\xee\xfc\xc8\x32\x38\x72\xd5\x84\x21\x46\xcd\x73\x85\x05\xc4\xd5\x66\x79\x23\x03\x06\xa6\xf6\x6b\xdd\x28\x05\x51\x2f\xf1\x28\x4c\x2e\x2f\xd2\x8d\x2d\xd0\x86\x1e\xa3\xf2\x96\xeb\x43\x5f\x7c\x7c\x8b\x9c\x04\xb0\xee\x8a\x62\xdc\xaa\x57\xfc\x2a\x6c\x08\xa6\x71\x75\xe3\x77\x82\xb6\xbd\xb3\x48\xe1\xb1\x51\xd7\x41\xd2\x25\xff\x01\x00\x00\xff\xff\xfc\x81\x95\xc1\x15\x04\x00\x00")

func xiangFlowsMenuFlowJsonBytes() ([]byte, error) {
	return bindataRead(
		_xiangFlowsMenuFlowJson,
		"xiang/flows/menu.flow.json",
	)
}

func xiangFlowsMenuFlowJson() (*asset, error) {
	bytes, err := xiangFlowsMenuFlowJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/flows/menu.flow.json", size: 1045, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangFlowsTableGetFlowJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x4d\x4a\xc4\x40\x10\x85\xf7\x39\x45\x51\xb8\x0c\x21\xba\x9c\xab\xc8\x20\x9d\x74\x31\x36\x56\xd2\x43\xff\x8c\x8b\xa6\x41\x2f\xe1\xca\x8d\x27\xd0\x1b\x78\x19\xc9\x39\xa4\x3a\x33\x99\x04\x71\x13\xc2\x7b\xaf\xbe\x7e\x55\xa9\x02\x40\x56\x1d\x31\xee\x00\xa7\xf7\x97\xe9\xed\xeb\xe7\xfb\x63\x7a\xfd\xc4\x5a\xac\x13\x39\x6f\xec\x28\xe6\x6d\xd3\x36\xed\xac\x6a\xf2\xbd\x33\xc7\x70\x76\xfe\x8e\x8d\x56\x93\xc7\x1d\xdc\x57\x00\x00\xa9\x7c\x45\x56\x03\xc9\x40\xf4\xe4\x4a\xb0\xa8\x47\x67\x7b\xf2\x12\xc7\xc1\x6a\x62\xdf\x88\xdf\x1c\x28\x5c\x33\xca\x1d\xae\xbc\x35\xb3\xb8\x9e\x98\xfa\x20\x3e\x1a\x8d\xf5\xf9\xa1\x5a\x78\x9d\x61\xc2\x7d\xbd\x4e\xb3\x19\x8c\x84\xef\xda\x8d\x6c\x9d\x26\x57\x1e\x49\x80\xbd\xe5\x38\x94\xe5\x7a\x47\x2a\x90\x7e\x50\x41\x80\x76\x59\x5a\x6e\x80\x90\xb7\xe8\xe7\x47\x72\xb4\x2d\x5a\xca\xae\x81\x3e\xa8\x10\xbd\xc0\x4e\x8a\x63\xb9\x07\x8d\xaa\x63\xd2\x08\xb9\xfe\x7f\xee\xb2\xd3\x32\x95\xd2\x8d\x19\x9b\x36\xe7\xb9\x98\x48\x6c\x9e\x08\x21\xaf\x20\xfb\xe5\xff\xa2\x2e\x8d\xd1\xc6\x50\xba\x0a\xc9\xc6\x90\x33\xce\x69\x49\xee\xab\x5c\xfd\x06\x00\x00\xff\xff\x34\x26\x65\xb9\x1d\x02\x00\x00")

func xiangFlowsTableGetFlowJsonBytes() ([]byte, error) {
	return bindataRead(
		_xiangFlowsTableGetFlowJson,
		"xiang/flows/table/get.flow.json",
	)
}

func xiangFlowsTableGetFlowJson() (*asset, error) {
	bytes, err := xiangFlowsTableGetFlowJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/flows/table/get.flow.json", size: 541, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangModelsReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x78\x36\x75\xc3\xb3\xde\x75\xcf\x56\x2c\x7c\x3a\xaf\x9b\x0b\x10\x00\x00\xff\xff\xb5\xd1\x53\xc2\x0f\x00\x00\x00")

func xiangModelsReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_xiangModelsReadmeMd,
		"xiang/models/README.md",
	)
}

func xiangModelsReadmeMd() (*asset, error) {
	bytes, err := xiangModelsReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/models/README.md", size: 15, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangModelsMenuJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x98\xdd\x4e\x1b\x47\x14\xc7\xef\xfd\x14\xa3\xb9\x36\x05\x0c\xfd\xb0\x2f\x2b\x6e\xb8\xe8\x13\x54\x51\x34\xde\x3d\x5e\x8f\x98\x9d\x75\x77\x66\x69\x2c\x84\x44\xa9\x9a\x58\x6e\x69\x68\x45\x44\x55\xdc\x24\x54\xad\xca\x4d\x0b\x51\x23\x4a\x49\x4a\x5f\x86\xb5\xcd\x5b\x54\xfb\xbd\xf6\xce\xae\x81\x12\x19\xb8\xe1\xe3\xcc\xd7\xf9\x9d\x39\xf3\x3f\xb3\xb3\x56\x42\x08\x73\x62\x02\xae\x21\x3c\x7c\xda\x73\xb7\x9e\xe1\xb2\x67\x93\xa4\xce\x3c\xa3\xd7\x21\xd5\xe5\x11\x25\xdc\x78\x68\x02\x77\xfc\x6e\x08\x61\xcd\x32\x4d\xe0\x32\x19\x3f\xdc\x3f\x88\xda\x80\x1b\x94\xfb\xe3\x96\x39\xb7\x96\x3e\xc6\x25\x84\xd6\xfd\xe9\x35\x8b\x39\x26\x17\xb8\x86\x3e\xf5\xbb\xae\x25\x4b\x50\x1d\x97\x11\x96\xed\x56\x30\x70\xc9\xfb\x2f\xb5\xc8\xf2\x12\x0e\xe6\x40\xa1\x6f\x29\xef\x5a\xc4\xf6\x7a\x95\x23\x7b\x34\x49\x9d\x1a\xcb\x5c\x82\x01\x76\xd2\x96\x9a\x72\xd0\x39\x1e\x76\x37\x07\x9b\x7f\x27\xad\xdc\x61\x2c\x8c\x80\xb4\x1d\x88\xed\x94\xeb\xf0\x28\x34\xfa\xb6\x3c\x57\xfc\xdf\x19\x47\x84\xb4\x29\x37\x12\x3b\x03\x6e\xc8\x26\xae\xa1\xca\xdc\x9c\xca\x33\x77\x7b\x6b\xf0\xdb\x11\xbe\xfa\xf2\x54\xb3\xf8\x0d\x2c\xbf\x77\xd6\x7f\xf9\x24\x37\x2a\x85\x1e\xd4\x99\xa5\xad\x08\xc5\x5e\x58\x16\x03\xc2\x95\x1b\xe1\x8f\x71\x4f\xbe\x74\x5f\x6d\x24\xed\x3a\x34\x88\xc3\xbc\xf6\x06\x61\xe2\x1a\x5b\xb1\x4a\x05\xad\x33\x48\x67\xed\x25\xfd\xe9\xff\x70\x36\xf8\xe5\xf4\x06\x5d\x69\x11\xd9\xfc\xff\xdb\x32\xfc\xeb\x70\xb0\xf3\xea\x1a\x59\x61\x13\xbe\x92\x5d\x9e\x8e\x1f\x8d\x04\xb3\x5a\xad\x56\x95\x81\xf9\xf6\x7b\xb7\xb3\x7b\x0d\x0f\x84\x24\xd2\x51\x64\x05\x70\xc7\x54\x39\x80\x81\x7b\x09\xa7\x27\x6d\x56\x4b\x52\x8b\x7b\xc2\x91\xb4\x21\xac\x53\x11\xfc\xfd\x40\x19\x30\x5f\x98\x06\xdd\xe3\xfe\xc6\x17\x28\x1c\x85\xdc\xb7\x1b\xee\xf6\x61\x19\x45\x43\x91\xfb\xd5\x9f\x17\xbb\xbf\x17\x41\x95\x10\xf2\xe7\xc7\x36\x30\xe2\xb9\x21\x12\x85\xd4\x9a\x94\xe9\x36\xf0\xd8\x92\xa2\x6b\x12\xf1\x09\xe1\xed\x64\x6a\xd3\xd2\x81\xc5\x7a\xfa\xde\x68\x66\xae\x40\x5b\x25\x66\x0d\xcb\x06\x6a\xf0\x48\x22\x23\xf3\x67\x0e\xd8\xed\xd4\xa2\x08\x61\x01\x0c\x34\xe9\x87\x28\x90\xa1\x48\x0f\x92\x53\x19\xa5\x62\x94\x13\xf1\xce\x3c\x08\xa7\x59\x8f\x89\x03\xc1\x5e\x25\xcc\x81\x94\x5e\xc7\x41\xd2\x71\x0d\xcd\x97\xc7\xf7\xb9\xff\xec\xa8\xbf\xf5\xc7\xa0\xf7\x75\xff\xa7\x7f\x53\x11\xd5\x2c\x1e\xa9\xd3\x0c\xd1\x24\x5d\xa5\x32\x15\x15\xdf\xa5\x1a\xc2\xb3\x2b\x84\xd7\xd3\x47\x32\x0c\x45\x0d\x79\x02\x14\x5b\x47\xce\xf5\xf8\x91\x0c\x41\xc7\xcd\x3e\x6d\x0d\xcd\x2b\x93\xd4\x67\xa9\xe4\xb0\xb8\x9d\xc7\xfd\xe7\xdb\x39\x2c\x2d\x0a\x33\x5a\x93\xd8\x52\x01\x33\x66\x2f\xf4\x5a\x0d\x5a\xcc\x52\xc9\x67\x59\xc8\x61\x19\xee\x1f\xf4\x5f\xbe\xcd\x61\x31\xec\x74\x7a\xc5\x18\xc1\x85\xe0\xdd\x61\x2c\xe4\x63\x2c\x66\x30\xdc\xd3\x9d\xc1\xce\x81\x7b\xb2\xe9\xf6\x4e\xa7\x86\x31\x72\x31\x08\x29\x16\xf3\x29\xde\xcf\x50\x0c\xff\x39\x3c\x7f\x73\xdc\xef\x6d\xb9\xdd\xfd\xe0\xe7\xf0\xf5\xc1\xc5\x93\xa7\xe3\xa9\x36\x9a\x4a\xb3\x82\x91\x0c\xf3\xe5\x0f\x46\x0c\x55\xb9\xe4\xc6\x78\x35\x20\x9f\xea\x83\x42\x2a\xb7\x77\xe4\xbe\x78\x31\x81\x87\xd8\x70\x8b\x80\x3e\x2c\x06\x7a\xbe\x39\xd8\x7b\x7d\xf1\xf3\xe3\xc9\x58\x0d\x0a\x4c\x17\xb7\x07\xec\xa3\x42\xb0\xe1\xfe\x37\xe7\x27\x3f\x4e\x40\xa2\x5c\x77\x84\xb4\xdb\xb7\x07\xaa\x5a\x08\x75\x7e\xb6\x77\xd1\xdb\x98\x00\xc5\x01\x6e\xd3\x36\xcd\x2b\x8a\x69\x4a\x1b\x14\x18\x02\xec\x55\xaa\x81\x52\x17\x02\x2d\xd4\x2c\xb3\x45\x84\xb8\xaa\xec\x2d\xde\x0c\x50\xb6\xa2\x4e\x90\x85\x08\x48\x29\x0c\x21\x91\x6d\x09\xd1\x24\xd4\x9e\x12\x53\xb6\xb2\xa6\x95\xa1\x80\x29\x47\x15\x02\x2a\xd1\xb4\x5a\x2d\xca\x8d\x99\x3a\x31\xa6\x04\x96\xad\xb5\x13\x94\x21\x02\xcb\xd5\x86\x00\xed\x73\xca\xf5\x29\x21\x65\x0b\xef\x04\x5d\x88\x90\xd4\xca\x10\xde\x5e\xb9\xd6\xb4\xa6\x95\x7d\xd9\xa2\x7b\xc7\x25\x22\x5b\x74\xef\xbe\x44\x64\xeb\xed\x3d\x91\x08\x45\xcd\xbd\xe3\x12\x51\x99\xbb\x6f\x12\x51\x51\xdc\x22\xba\xbf\xba\xdf\x75\xdd\xce\x6e\xf2\x40\x9b\x26\xaa\x3b\xc6\x2c\xa3\x42\xe6\xc0\x8c\x36\x8d\xa1\x8c\x7c\x09\xc5\x24\x0b\x37\x43\xa2\xb8\x3e\x74\xf6\xdc\x37\xa7\x01\x4f\x0e\x89\x66\x03\x91\x90\xf7\xb5\xce\x9c\x2b\x0b\xdd\xd5\x61\xe2\xc7\x22\xff\x1d\x29\x78\x3b\x09\x0c\xf1\x13\xd6\x1a\xc2\x92\x9a\x20\x24\x31\x5b\xf1\x17\x25\xc2\xc2\x6a\xc8\x87\x3a\x30\x90\x10\x59\xd1\x7a\x69\xbd\xf4\x5f\x00\x00\x00\xff\xff\x2f\x9d\x83\xca\xad\x17\x00\x00")

func xiangModelsMenuJsonBytes() ([]byte, error) {
	return bindataRead(
		_xiangModelsMenuJson,
		"xiang/models/menu.json",
	)
}

func xiangModelsMenuJson() (*asset, error) {
	bytes, err := xiangModelsMenuJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/models/menu.json", size: 6061, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangModelsUserJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x6f\x53\xd4\xd6\x17\x7e\xcf\xa7\xb8\x73\x7f\xce\x6f\x6a\x5d\xca\x66\xf9\x23\xec\x2b\xb1\xd4\x19\xab\xb5\x54\x6a\xb5\x0a\x32\x61\x73\x77\x0d\x66\x93\x35\xc9\x2a\xb8\xdd\x19\xa8\x15\xb1\xb2\x03\x2a\x0c\x32\xa5\x8e\xce\x14\xa5\x55\x59\xa8\x8e\xec\xa0\x8c\x1f\xa6\xb9\xd9\xf0\xaa\x5f\xa1\x93\xdc\x24\x9b\x6c\xee\xfe\x83\xed\x8c\x55\x5e\xb1\x9c\xe4\x9e\xf3\x9c\xe7\x3c\xf7\xdc\x6c\xce\x66\x5a\x00\x80\x22\x9b\x44\x30\x0a\x60\x71\x7e\x55\x9f\xde\x84\x21\xd3\xa6\xb2\x23\x82\x69\x34\x6f\xf0\xdc\x32\xc6\xb3\x62\x62\x38\xad\x20\xd9\xba\x0d\x00\x18\x93\x92\x49\x24\xaa\xa5\xf5\xc6\xe3\x55\xe7\x1a\x12\x13\xbc\x68\xad\x3b\x2e\x8a\x52\xdf\x51\xd8\x02\x40\xd6\x72\x1f\x93\x84\x74\x52\x54\x60\x14\x5c\xb0\x6e\xcd\x00\x28\xb0\x23\x48\xb0\xee\xed\x83\xa1\x52\x44\x9e\x33\xff\x53\xc7\x53\xc8\xbe\x46\x3c\x00\x1b\x19\xf0\x2c\x2c\x6e\xbc\xc1\x0f\xef\xd8\xc1\x3d\xa0\xad\xc5\xae\xd5\x71\x85\xc4\x74\xb2\x64\x95\x52\x2a\x2f\x89\x26\x1e\xc8\x72\x49\x5e\x34\x83\x2a\x2a\x1b\x8f\x9b\x1f\x48\xbe\x00\xca\xd2\x88\xa4\xc2\x21\x77\x91\x27\x77\xe3\xd5\x13\x3c\xbb\x49\x10\x00\xcb\x03\x28\xae\x3d\x2e\xce\x4d\xe1\xbb\x0f\x42\xc0\xf2\x04\xf0\xdd\x07\x78\x73\x25\x04\x4c\x77\x80\x90\x15\x02\x96\x4f\xa0\x2f\x6f\xe1\xa5\x55\x6d\x6b\xab\x84\x88\x43\x71\x36\x2d\x58\xce\x3d\x7c\x03\x00\x79\x91\x43\x63\x30\x0a\x54\x39\x8d\x5c\xe3\x55\x56\xe0\x39\xd6\xcc\xa1\x44\xaa\x97\x25\xeb\xa6\x24\x52\x2f\x49\x9c\x43\x89\x14\x77\x7d\x5a\x57\x59\x39\x61\xad\x85\x8a\x2a\xf3\x62\xa2\x94\xa7\xbd\x56\x51\xd8\x84\xc5\x5c\x26\xc3\x8b\xa9\xb4\x9a\xcd\x92\x74\x77\xe6\x97\x8c\x7c\x3e\x04\x32\x19\xab\x12\xd9\x2c\xde\x9a\x37\xf2\x2b\x5a\x61\x0b\xbf\x58\x2c\x3e\x7f\xa2\x15\xfe\x84\xae\xa7\x6c\xa8\x16\x34\x5f\x5d\xfc\xc0\xea\xae\x4c\x45\xc4\x5a\x21\x87\x97\x57\x8d\xb5\x02\x9e\xcd\x1b\x33\x37\xf0\x2f\xaf\x68\xb8\x49\x01\xdb\xac\x28\x6d\x66\x88\x36\xe2\xbf\x94\x85\xfd\x69\xa8\xc5\x93\x53\x50\x90\x3b\x3f\xae\x15\xd7\x36\x82\x82\x44\x49\x96\x17\x82\x8a\xb4\x79\x77\xed\x02\x12\x13\xea\x25\x18\x05\x9d\x61\x9a\xe4\xca\xbd\x53\x75\x21\xa6\x05\xc1\xde\xca\x1f\xaa\x5e\x7c\x64\xfa\x90\xd5\xa5\x07\xfd\x41\x9e\x30\x89\x97\xd7\xf1\xaf\x13\x0d\x17\x59\xbf\x7d\xc7\xdc\xbd\xb3\x9b\xc1\x3a\x27\xa5\x11\x5e\xa0\xb4\x9e\x06\x0b\x4d\x89\xd0\x50\xad\x63\xf2\x78\xca\x72\xd4\xfb\xc5\x00\xdc\xad\x02\xca\x72\x69\x90\x67\xfd\xd1\x5b\xfc\x76\x96\x54\xbe\x61\x86\x8b\x4b\x6f\xf0\xf6\x02\xce\x4f\x15\x1f\x4d\x06\x49\x4e\xb1\x8a\x72\x4d\x92\xb9\x06\x68\x8e\x74\x76\xd1\x78\xa6\xc7\x71\xd9\xeb\xef\x1d\x18\x38\xfb\xf5\xe9\xbe\x8f\x7e\xc7\x25\x79\xf1\x24\xa1\x92\x8a\xae\xab\x0a\x2c\x5f\xfc\xe2\xfc\x46\x57\x2b\xd3\xad\x6d\xe7\xfe\x7e\x3b\x83\x7f\x7b\x8a\xd7\x67\xf1\xd4\x12\x7e\xb1\xa8\xe7\x73\x7f\x4d\x4c\xea\x0b\xeb\xf8\xc5\x22\xbe\x37\x53\x7c\x6e\x9e\xad\xfa\xc3\x9f\xf4\xe9\xb9\x86\x70\xb2\x63\xd5\x70\x32\xdd\xef\x0b\xd0\x14\xab\xaa\x48\x16\x2b\x14\xfb\x42\xb8\xb5\x67\xe8\x50\xb5\x62\xfb\xc0\x1a\xb7\x5e\xe2\xf5\xbb\x78\xe6\x26\x9e\x7b\xa6\x15\x26\xb4\xc2\x1f\x04\x5f\x13\x01\xf5\xb6\x9e\xdf\x13\x20\x93\x42\x87\xbf\x26\xc2\x62\x5b\xaf\xef\x0d\x96\xa7\xac\x4d\x84\x75\xe4\x7f\x07\xfe\xff\xe9\xe0\xe0\xa1\x3d\x61\x23\xd2\x6a\xfc\x70\xba\x9f\xd3\xb6\x97\x6b\xb5\xce\x88\xd8\x84\xee\x49\x0f\xb5\xdf\x3d\xcb\xa1\xed\x77\xcf\xfd\xee\xb9\xdf\x3d\xff\x13\xdd\x13\x3f\xbd\x8f\xe7\x72\xc1\xbe\x69\xfd\xad\xbf\x61\x76\x53\x9f\xea\xcb\x9d\xef\xb7\x43\x1a\xba\x48\xed\xa2\x93\x72\xef\x2c\x4f\x18\x4f\x26\x23\xa6\x10\x1b\xdb\xad\x35\xfa\x5c\x47\xb8\x36\x02\xad\x90\x33\x6e\x6c\x1b\xaf\x6f\x1a\xef\x6e\x45\xc2\x01\x08\xf5\x89\xcd\xd8\x7a\xa6\xbd\xd9\x36\xf2\x93\x78\x76\x93\x7a\x5a\xf3\x5c\x8c\x6d\xc6\xd7\x9c\x8a\x81\xe8\xdf\x13\x2b\x49\xf0\x03\x7f\x0d\x55\xbd\x2d\x5d\xfc\x64\x70\x90\xcb\x30\xdd\xd9\x83\x3f\x90\x4f\x1d\xd9\x73\x07\x0f\xd4\xd1\xa1\xf6\xf4\x45\xb8\xb7\xff\x38\x38\x81\xc6\x83\xd2\xb8\xec\x35\xee\x56\x17\x41\xef\x15\x2a\x9f\x16\xf9\x2b\xe9\x7f\xb5\x25\xf1\xa2\x8a\x12\x48\x6e\x42\xed\x9b\x7d\x7a\x5f\xb4\x9e\x27\x32\x4c\x38\x5b\xa5\xda\xe5\x38\x8a\xf3\x1b\x4c\x58\xdb\xce\x11\x30\x81\x27\x9c\xfa\x8b\x8f\xf3\x53\x3b\xf7\x56\x82\xf5\x57\x50\x4c\x46\xea\x2e\x25\xd0\xd8\x6b\xa3\x32\xc1\x94\x23\xfa\xb8\xbb\x82\x29\x8e\xde\xd6\xf3\x6c\xeb\x75\xeb\xb1\x65\x28\xd3\x1e\xa9\xa6\x93\xa0\x4c\xda\x23\xbb\x7d\x4c\xae\x4f\x44\xfa\xed\xdf\xf1\xc6\x82\xf6\xee\xb1\x3e\x99\xa7\xbc\x97\x1e\x53\x65\x36\xa8\xa2\x51\x45\x12\xa9\x02\xa8\xe0\xcd\xaf\xa7\xea\x80\x8a\x3f\xbf\xd6\x27\x26\xa9\xde\xc9\x84\x84\xdc\x00\x90\x68\xba\xe4\x80\xbe\x7c\x5b\x5f\x98\x0e\x01\x8e\x57\x6c\xc3\xe2\x23\x7d\x61\x9a\xb2\x25\x54\x56\x4d\x2b\xb5\xc6\x3e\x9e\x21\x8b\x1d\x81\x3a\x12\x72\xaf\x01\xe8\x04\xf6\x8c\x81\x3e\x70\xcd\x57\x19\xc8\x54\xe5\xa5\x22\xca\x3a\x87\x30\xb6\xf3\x36\xd7\x73\x25\xb1\xb7\x00\x60\x05\x85\x32\x12\x5c\xc6\x33\x64\xd0\x78\x95\x15\xd2\xc8\x33\x67\x2c\x97\x89\x3b\xa4\x0b\x2a\xc5\x9e\x36\x39\x66\x32\x67\x70\x66\xa0\x47\xf8\x2b\x97\xd9\xcf\x62\x92\x47\x4b\xf6\x1b\xf2\x28\x30\x37\x80\x6b\x75\x5f\x4f\x9b\x0d\x93\x89\xb4\x77\x74\x76\xa5\x0e\x95\x16\x99\x27\x77\x14\xc0\xee\xf6\x70\x47\x4f\xa4\x93\x39\xdc\x55\xba\x64\x37\xf5\x28\x80\xe7\xbe\xfa\x96\x3b\x75\xfa\x3b\x3e\x31\x92\x38\xc3\xf7\xf6\x73\xfc\x97\x9f\xc7\xd9\xb3\x89\xb3\xb1\xeb\x91\x7e\xf6\x9b\x73\xd7\x3c\x6b\x88\xea\x69\x72\x26\x9b\x3b\x0a\x32\xa6\xe7\x31\xb2\xc1\x36\xa1\xcd\x65\xf9\xfe\x74\xe8\x21\x23\x4b\xda\xc1\x62\x0d\xe0\x02\xdc\x58\xf6\x5d\x71\x33\x50\x89\x1b\xdf\x0a\x97\x13\xbf\xb5\x91\xac\xf1\xca\xcb\x5a\x59\x7b\x06\xe1\xbe\xac\xfd\x73\x58\x27\x69\x1f\x14\x37\x51\xc8\xb4\xf7\x84\xc3\xe1\x30\xc3\x30\x0c\xa4\x66\x7c\xe6\x3d\xca\x98\x32\x7f\x76\x92\x26\xa3\xcf\x3a\xb3\xae\x54\x5e\x9f\xdd\xd6\x7c\x7b\x77\x47\x27\xd3\xd3\x15\x3e\x1c\xa1\x69\xfe\xda\x51\xf4\xfd\xe8\xc9\xc3\xc7\x46\x47\x62\x57\x53\xbd\xdc\x51\x79\x4c\xed\x3b\x16\x1f\x1d\xe7\x14\xa9\xff\xc4\xa5\xd3\xa7\x76\xa9\x79\x2f\x17\x6e\xdf\xb0\x1a\x38\xad\x4b\x78\x4f\x23\xd2\x5a\xc9\xa4\x70\x3e\xaf\x15\x26\xe8\x3f\x34\x18\xb6\x08\x1a\xb6\x9f\x91\x3d\xe7\x9a\xfb\x83\x07\xfb\xf7\x08\x0e\x95\x43\x41\x99\x91\xb5\xd4\x82\x05\x10\xb9\xc3\xc1\xaa\xa0\x48\x81\xea\x42\x65\xd7\xb2\x06\x2c\x97\x3b\xf7\x94\xcc\x00\xa8\xf2\x49\xa4\xa8\x6c\x32\xa5\x38\xa7\x21\x80\x8a\x14\x57\x87\x39\x24\x20\x15\x39\x56\x90\x6d\xc9\xb6\xfc\x13\x00\x00\xff\xff\x2b\x52\x7c\x27\x7a\x22\x00\x00")

func xiangModelsUserJsonBytes() ([]byte, error) {
	return bindataRead(
		_xiangModelsUserJson,
		"xiang/models/user.json",
	)
}

func xiangModelsUserJson() (*asset, error) {
	bytes, err := xiangModelsUserJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/models/user.json", size: 8826, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangTablesReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x78\x36\x75\xc3\xb3\xde\x75\x2f\x16\xae\x78\xb6\x60\x0f\x17\x20\x00\x00\xff\xff\x11\x5e\x34\x14\x0f\x00\x00\x00")

func xiangTablesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_xiangTablesReadmeMd,
		"xiang/tables/README.md",
	)
}

func xiangTablesReadmeMd() (*asset, error) {
	bytes, err := xiangTablesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/tables/README.md", size: 15, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _xiangTablesUserJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4d\x6b\xdb\x40\x10\xbd\xeb\x57\x0c\x7b\x4e\x4d\x0a\xc5\x07\xff\x8d\x1e\x7a\x08\x3e\xac\xb5\x63\x67\x60\xb5\x5a\x76\x57\x4e\x4d\x30\xb4\xb7\x96\x50\x5a\x68\x29\x3d\xf4\xcb\xa7\xf6\x52\xda\x9c\x02\x4e\x21\x7f\xc6\x92\xfb\x33\xca\x4a\xb2\xb4\x72\x44\x13\x43\x75\xd3\x9b\x37\xb3\x3b\xfb\xde\x3b\x8f\x00\x98\xe2\x09\xb2\x11\xb0\xed\xbb\xef\xc5\x8b\x2b\x76\xe4\xb1\x39\x1a\x4b\xa9\xf2\xf0\xc3\xc1\xf1\xe0\xb8\x42\x05\xc6\x86\xb4\xab\x0b\x7f\x2e\x57\x9b\xdf\x5f\xc3\xae\x09\x29\xc1\x46\xe0\xa7\x02\xb0\x24\x15\x28\x3d\xf1\x29\x71\x35\x1b\x64\x16\x0d\x8b\x00\x96\x25\x95\x6b\xb2\x9e\x5a\xfd\xc5\xa9\xcc\x12\x15\x00\x53\x92\x0e\x4d\x00\x48\xb2\xae\x1d\xad\x0d\x25\xdc\x2c\xfc\x70\x12\xe5\xd9\x9e\xc2\x17\x69\xd6\x92\x3a\x63\x4f\x6a\x08\xe0\xbc\x5d\x38\xff\xf6\x36\x7f\xf3\x8a\x1d\x01\x3b\x23\xe1\x4e\xd9\x08\x86\xd5\xed\x6e\x31\xb7\x97\xd7\xf9\xe7\x8b\xfb\x30\x8b\x97\x17\xc5\xc7\x75\xfe\xfa\x2a\x24\x3f\x82\x65\xcd\x1d\xef\x9a\x82\x0d\x4f\xee\x6a\x1f\xc2\x72\x5c\xb6\xd5\x47\x32\x1e\x7b\x0d\x6c\x67\x55\x83\xdc\x61\x80\x00\x30\xb7\xd0\xe5\xd0\x49\xe6\x5c\xaa\x58\x7b\x5f\xa6\x4d\xaa\x6d\x87\x5c\xbe\xdf\xa4\xd2\xab\x78\xff\x2b\xbf\x5e\x07\xc2\x36\x14\x8a\x2b\xe9\xa7\xdc\xc2\x94\x3f\xd0\x32\xb3\xac\xa9\xef\x76\x6c\x1e\x86\xcd\x09\xcf\x1a\x09\x4b\x04\x05\xb9\x2e\x42\x89\x4e\xcd\x1e\x96\x69\x51\x2f\xd3\x62\x02\x25\xee\x63\xa4\x2c\xf6\xf7\x3e\x39\x45\xd3\x3b\xa0\xa7\x50\x75\x3c\x46\x89\xb1\xeb\x6b\xe9\xab\x68\x3e\x23\xc5\xeb\x24\x04\xb8\x45\xe7\x48\xcd\x4a\x30\xda\x3d\x4a\xe5\xe0\xdd\xea\x07\x39\x78\x4a\x28\x85\x45\xd7\xb5\x70\x28\x89\x23\x27\x2b\x37\x7f\x59\x6f\x57\xcf\x36\x37\xab\xe2\xf9\xcf\xae\x6a\x02\x6d\x98\xdb\x6e\xb1\x2f\x23\x70\x50\x4e\xe0\xa0\xac\xc0\x3d\xf2\x32\x6c\xbc\xe4\xbf\xf1\x2d\x87\xdd\x9d\x06\xae\xe2\xd2\xcb\xa1\x34\x7c\xfe\x9f\xf2\xb1\xb9\xf9\x94\xff\xf8\xf0\x2f\xe3\x07\x66\xdd\x73\xc1\x9e\x65\xdb\x88\x44\xcb\xe8\x6f\x00\x00\x00\xff\xff\x42\x62\x48\x5d\x8e\x05\x00\x00")

func xiangTablesUserJsonBytes() ([]byte, error) {
	return bindataRead(
		_xiangTablesUserJson,
		"xiang/tables/user.json",
	)
}

func xiangTablesUserJson() (*asset, error) {
	bytes, err := xiangTablesUserJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "xiang/tables/user.json", size: 1422, mode: os.FileMode(420), modTime: time.Unix(1632926333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"ui/index.html":                   uiIndexHtml,
	"xiang/.DS_Store":                 xiangDs_store,
	"xiang/apis/README.md":            xiangApisReadmeMd,
	"xiang/apis/table.http.json":      xiangApisTableHttpJson,
	"xiang/apis/user.http.json":       xiangApisUserHttpJson,
	"xiang/apis/xiang.http.json":      xiangApisXiangHttpJson,
	"xiang/flows/README.md":           xiangFlowsReadmeMd,
	"xiang/flows/menu.flow.json":      xiangFlowsMenuFlowJson,
	"xiang/flows/table/get.flow.json": xiangFlowsTableGetFlowJson,
	"xiang/models/README.md":          xiangModelsReadmeMd,
	"xiang/models/menu.json":          xiangModelsMenuJson,
	"xiang/models/user.json":          xiangModelsUserJson,
	"xiang/tables/README.md":          xiangTablesReadmeMd,
	"xiang/tables/user.json":          xiangTablesUserJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"ui": {nil, map[string]*bintree{
		"index.html": {uiIndexHtml, map[string]*bintree{}},
	}},
	"xiang": {nil, map[string]*bintree{
		".DS_Store": {xiangDs_store, map[string]*bintree{}},
		"apis": {nil, map[string]*bintree{
			"README.md":       {xiangApisReadmeMd, map[string]*bintree{}},
			"table.http.json": {xiangApisTableHttpJson, map[string]*bintree{}},
			"user.http.json":  {xiangApisUserHttpJson, map[string]*bintree{}},
			"xiang.http.json": {xiangApisXiangHttpJson, map[string]*bintree{}},
		}},
		"flows": {nil, map[string]*bintree{
			"README.md":      {xiangFlowsReadmeMd, map[string]*bintree{}},
			"menu.flow.json": {xiangFlowsMenuFlowJson, map[string]*bintree{}},
			"table": {nil, map[string]*bintree{
				"get.flow.json": {xiangFlowsTableGetFlowJson, map[string]*bintree{}},
			}},
		}},
		"models": {nil, map[string]*bintree{
			"README.md": {xiangModelsReadmeMd, map[string]*bintree{}},
			"menu.json": {xiangModelsMenuJson, map[string]*bintree{}},
			"user.json": {xiangModelsUserJson, map[string]*bintree{}},
		}},
		"tables": {nil, map[string]*bintree{
			"README.md": {xiangTablesReadmeMd, map[string]*bintree{}},
			"user.json": {xiangTablesUserJson, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
