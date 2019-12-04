// Code generated by go-bindata.
// sources:
// template/default.tmpl
// DO NOT EDIT!

package deftmpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x08\xfb\x12\x0b\x98\x7a\x2f\xb0\x0d\xc5\xb0\xf5\x12\x0c\x43\xb2\xec\xb2\x06\x81\x1a\x33\xae\x5a\x89\xca\x64\xba\x0d\xe0\xe8\xdf\x07\xc9\x86\x1b\x37\xd9\x6d\xbd\x49\x14\xf9\x1e\x1f\x1f\xd5\x75\x50\xe1\x4e\x13\x42\xbe\xd9\x34\xed\xfd\x23\x6e\x39\x87\x10\x7e\x77\x1d\xc8\x25\x2b\x6e\x1b\x38\x02\xbb\xd5\x7e\x8f\x1e\x42\xe8\x3a\xd0\x3b\xc0\x3f\xe3\x63\xbe\xd3\x5e\x53\x1d\x6b\xae\x63\xcd\x8d\x41\xcf\x8d\xfc\x96\xa2\x70\x04\x83\xd4\x97\x21\x55\x10\xc2\x1a\x62\xd2\xad\x77\xed\x7e\xae\xee\xd1\x34\x72\xe9\x3c\x63\xf5\x43\x69\xdf\xc8\x5f\xca\xb4\x18\x09\x1f\x9d\x26\xc8\x21\xa2\x42\x4f\x59\x33\xcc\x22\x96\xfc\xe2\xac\x75\xd4\x17\x97\x43\xec\x04\xaf\x84\x10\x66\x5d\x07\x2f\x9a\x1f\xa6\xc9\x72\x81\xd6\x3d\xe3\x94\xfd\xbb\xb2\xd8\xf4\x0d\x5e\x64\x1f\x1b\x2f\xc7\xd3\x78\xc8\x26\xc3\x53\x51\xb8\x55\xa4\x6a\xf4\xab\xc5\x7c\x28\x96\x5f\x0f\x8c\x9e\x94\x59\x2d\xe6\x10\xc2\x55\x71\x95\xf2\x9a\xcf\x1e\xb7\xa8\x9f\xd1\x7f\x8c\x49\x8b\xe1\x32\x41\x9f\xc2\x33\x1e\xb8\xe7\xd8\x18\xdd\xf0\x00\xef\x15\xd5\x08\x32\xa6\x0b\xd1\x4b\x12\x22\x7b\x7d\x38\x9f\x31\x84\xf0\x09\x3e\x24\x17\xa2\xf6\x68\x1b\x8c\xe2\xe1\x08\x56\xf9\xa7\xca\xbd\x10\x1c\xe1\x81\xad\x19\x64\x0e\x2d\x09\x71\x43\xe4\x58\xb1\x76\x34\x25\x3a\x89\xff\x47\xb6\xa5\x6b\xfd\x16\xaf\x85\x80\xb4\x8f\xb7\x48\xe8\x15\x3b\xdf\x0f\x73\x3d\xbb\x10\x2c\xb3\xec\x82\x53\xa7\xb3\xac\x34\xd5\xd2\x68\x7a\x92\xac\xd9\xe0\x30\x49\x46\xbb\x37\x8a\xa7\xff\x40\xfe\xcb\xee\x57\x8c\xad\x23\x46\x4a\x7e\x14\x45\x51\xc0\xdd\x7b\xfd\x9c\xbb\x35\x08\x11\xc1\x35\x55\x78\x98\x6c\x31\xe4\x69\x31\x48\xd9\xa4\x26\xcd\xe5\x54\xcf\xd9\x6a\x46\x5d\xa5\x10\x59\xa2\x64\x6d\xb1\x77\x25\x5d\x7f\xea\xe4\x53\xf6\x06\xe3\x6c\xff\xde\xf4\x3b\x71\xee\x6f\x00\x00\x00\xff\xff\x7d\xe8\xc9\x81\x56\x04\x00\x00")

func templateDefaultTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDefaultTmpl,
		"template/default.tmpl",
	)
}

func templateDefaultTmpl() (*asset, error) {
	bytes, err := templateDefaultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/default.tmpl", size: 1110, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/default.tmpl": templateDefaultTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"default.tmpl": &bintree{templateDefaultTmpl, map[string]*bintree{}},
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