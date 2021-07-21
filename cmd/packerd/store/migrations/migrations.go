// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package migrations generated by go-bindata.// sources:
// migrations/001_init.down.sql
// migrations/001_init.up.sql
package migrations

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x4a\x2c\x49\xce\x88\x4f\x49\x4c\x2f\xe6\x42\x17\x4d\x2d\xb6\x86\x8a\x45\x06\xc0\x85\xe2\x8b\x4b\x12\x4b\x4a\x8b\xad\xb9\x00\x01\x00\x00\xff\xff\x51\x7a\x82\xf6\x44\x00\x00\x00")

func _001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initDownSql,
		"001_init.down.sql",
	)
}

func _001_initDownSql() (*asset, error) {
	bytes, err := _001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.down.sql", size: 68, mode: os.FileMode(436), modTime: time.Unix(1626890806, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x92\xbf\x6e\xb3\x30\x14\xc5\x77\x9e\xe2\x6e\x80\xc4\xf0\xed\x99\x1c\x72\x89\xac\x0f\x4c\x64\x8c\x94\x4c\x96\x1b\x3b\x2d\x6a\x14\x52\x6c\xa4\xb6\x4f\x5f\x11\x70\x52\xda\xaa\x95\xba\x74\x61\xf0\x39\xe7\xfe\xf9\x5d\x52\x8e\x44\x20\x88\xdd\x06\xe1\x4e\xb9\xfd\x83\xb4\x4e\xb9\xde\x02\xa9\x00\x59\x5d\x40\x14\xb6\x67\x73\x0a\x93\xb0\x33\x4a\xbf\x84\x49\xa8\xdb\x93\x09\xe3\x45\x10\xf8\x28\x59\xe6\x08\x34\x03\x56\x0a\xc0\x2d\xad\x44\x35\x56\x32\x16\xa2\x00\x00\xa6\xba\x8d\x06\x81\x5b\x01\x1b\x4e\x0b\xc2\x77\xf0\x1f\x77\xc9\x45\x9e\x1a\xce\xba\x0f\xb5\x58\x9d\xe7\xa0\xcd\x41\xf5\x47\x07\xd3\x14\x97\x80\x6b\x9d\x3a\x4a\xdb\xbc\x1a\x58\xd2\x35\x65\xe2\x66\x5f\x61\x46\xea\x5c\xc0\xbf\xd1\xb9\xef\x8c\x72\x46\x4b\xe5\x40\xd0\x02\x2b\x41\x8a\xcd\x67\x73\x5a\x73\x8e\x4c\xc8\xab\x65\x0c\xf7\x67\xfd\x9b\x70\x10\x2f\x3c\x1a\xca\x56\xb8\xf5\x30\xa6\xd5\x64\xa3\x9f\xa1\x64\xfe\x35\x1a\x5f\x93\x77\xa3\xfe\xc0\xd6\xba\xb6\x53\xf7\x46\x76\xe6\xa9\x37\xd6\x79\xc8\xed\xd9\x74\xca\x35\xed\xe9\x3b\xd0\xb3\xe4\xd5\xe8\x77\x1a\x5d\x5a\x39\x25\xf7\x5f\x6b\xf3\x4b\xce\xb5\xbf\x60\x3d\x86\x2f\x9f\xb4\x64\x95\xe0\x64\xf8\x1b\x0e\x8f\xd2\xaf\xaa\x8d\x3a\x0e\xf3\x66\x25\x47\xba\x66\x03\x8b\xe8\x83\x16\x03\xc7\x0c\x39\xb2\x14\x6f\x70\x07\xcd\x46\x8d\x8e\x87\x6b\x06\xc1\x5b\x00\x00\x00\xff\xff\x52\x15\x9f\x01\x26\x03\x00\x00")

func _001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initUpSql,
		"001_init.up.sql",
	)
}

func _001_initUpSql() (*asset, error) {
	bytes, err := _001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.up.sql", size: 806, mode: os.FileMode(436), modTime: time.Unix(1626897230, 0)}
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
	"001_init.down.sql": _001_initDownSql,
	"001_init.up.sql":   _001_initUpSql,
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
	"001_init.down.sql": &bintree{_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql":   &bintree{_001_initUpSql, map[string]*bintree{}},
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
