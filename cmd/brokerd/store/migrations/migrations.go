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

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xcd\x2b\xc8\xcc\x8b\xcf\xca\x4f\x2a\xb6\xe6\x42\x12\x4e\x49\x4d\xcc\x41\x15\x29\x2e\xc9\x2f\x4a\x4c\x4f\x8d\x2f\x4a\x2d\x2c\x4d\x2d\x2e\x41\x95\x4c\x4a\x2c\x49\xce\x48\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\xbc\x58\x40\x6e\x5a\x00\x00\x00")

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

	info := bindataFileInfo{name: "001_init.down.sql", size: 90, mode: os.FileMode(436), modTime: time.Unix(1628018891, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x55\x41\x6f\xdb\x3c\x0c\xbd\xe7\x57\xf0\xd6\x06\xe8\xe1\xbb\xf7\xe4\x2f\x55\x06\x63\xa9\x53\x38\x2a\xd0\x9e\x04\xda\x62\x52\xad\xaa\xe5\xc9\x52\xd7\xec\xd7\x0f\xb1\xe3\x34\x8e\xad\xba\xcb\x65\x03\x76\xd5\x23\x29\xf2\x91\x7c\x9c\xa5\x2c\xe2\x0c\x78\xf4\xff\x82\x41\x3c\x87\x64\xc9\x81\x3d\xc4\x2b\xbe\x82\x0c\x5d\xfe\x44\x15\x5c\x4e\x00\x00\x94\x04\x47\x6f\x0e\xee\xd2\xf8\x36\x4a\x1f\xe1\x2b\x7b\xbc\xaa\x81\xca\xa1\xf3\x55\x03\xee\xbc\x93\xfb\xc5\xa2\x41\x2c\x95\x62\x8d\xb9\x33\x16\x54\x71\x0a\x4a\x42\x2d\xa4\xb7\xe8\x94\x29\x06\xf0\x12\xb7\xda\xa0\x14\x79\xfb\xf1\x09\xac\x28\xa7\x8f\xc1\x4a\xfd\x24\xc8\xd4\xa6\x1f\x3b\x47\x2b\xbc\xd5\x5d\x57\xb8\x61\xf3\xe8\x7e\xc1\xe1\xe2\xe2\xdd\x4a\x95\xeb\xaa\xff\x4b\xd8\x14\xa5\xb4\xd5\x88\xb1\x54\x15\x6a\x6d\x7e\x08\x4b\x35\xc5\xaa\xd8\x40\x66\x8c\x26\x2c\xfa\x4e\xf3\x68\xb1\x62\x8d\xdf\x5a\x69\x41\xa5\xc9\x9f\x84\x24\x94\x5a\x15\x81\xea\xc8\x5a\x63\x47\x72\x30\x76\xe7\x39\xc4\x5d\x6e\x09\x1d\x49\x81\x0e\x78\x7c\xcb\x56\x3c\xba\xbd\xeb\xc7\x99\xdd\xa7\x29\x4b\xb8\x38\x98\x34\xce\xbe\x94\xe7\x38\xd7\xbe\xd3\xeb\xc9\x64\x6c\x1a\xc5\x0b\x16\x6a\x4d\x95\x6b\xa7\xb2\x79\x0d\xce\x66\x6b\x0e\xd9\xd6\x11\x1e\x32\xf9\xfc\x7f\x0e\x37\xbd\xaf\x38\x7b\xe0\x27\xa4\x3d\xd3\x76\xe8\xf9\x15\xb5\xa7\x21\xe0\x6c\x92\x6b\xef\xa3\x32\x2f\xdb\xac\xae\x76\x39\x4c\x9b\xe8\xb3\x65\xb2\xe2\x69\x14\x27\x1c\xd6\xcf\xe2\xbd\x10\x71\x28\x61\xbe\x4c\x59\xfc\x25\xe9\x44\x98\x42\xca\xe6\x2c\x65\xc9\x8c\x1d\x76\xff\x52\xc9\xe9\x64\x7a\x3d\x01\xf8\x90\xab\xca\x19\x8b\x1b\x12\x96\xbe\xfb\xa3\xe6\x04\xdb\x22\xd1\x61\x68\x77\x3b\x1d\xed\x28\x4c\xf5\x82\x5a\xf7\xc7\x7d\x3f\xca\x35\xcb\x85\x71\x50\x78\xad\x5b\xfd\x69\x82\xe5\xc6\x17\xae\x23\x31\x07\x8a\xff\x3b\x5a\x19\x91\xa3\xaf\x68\x6c\xd3\xff\xc0\x7a\x84\xbb\xfa\xbb\xad\xfc\xc4\xdc\xef\x74\x79\x78\xbb\xba\xb4\xa3\xcf\x77\xca\x1d\x40\x33\x25\x03\x48\x3b\x2b\xa5\x35\xaf\x4a\x92\x0d\x98\xd5\xe7\x41\xc9\x61\x85\xab\x41\x7a\x2b\xd5\xfe\x7a\x84\x65\xf0\xdf\xe9\xe9\xbe\xa5\x71\x72\xc3\x1e\x86\x5a\x2a\x86\x88\x5f\x26\x6d\xbf\x07\xd0\x91\x41\xf1\x45\xa9\x0a\xf1\xcd\x64\xa3\xeb\x4e\x6f\x94\x7b\x77\x7c\xe6\x06\xae\x5b\x40\x0e\xdc\xb6\xa4\xd0\xe6\x5b\x42\xb9\x3d\x8f\xf8\xbf\xf5\xca\x99\x92\x9a\x99\x3e\x22\xb5\x56\xb6\x1e\xa9\x67\x56\xd0\x26\xf1\x2b\x00\x00\xff\xff\xa5\xf5\x14\x36\xf8\x09\x00\x00")

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

	info := bindataFileInfo{name: "001_init.up.sql", size: 2552, mode: os.FileMode(436), modTime: time.Unix(1628018891, 0)}
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
