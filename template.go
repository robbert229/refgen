// Code generated by go-bindata.
// sources:
// refgen.tmpl
// DO NOT EDIT!

package refgen

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

var _refgenTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x41\x8b\xe2\x30\x14\xc7\xef\xf9\x14\x8f\x1e\x96\x04\x25\xee\x59\xd7\xe3\xee\x65\x19\x11\x9c\x9b\x08\x66\xda\xd7\xfa\x46\x93\x94\x34\x65\x18\x42\xbe\xfb\x90\xa4\x5a\xbd\x14\xfa\xfe\xbf\xbc\xfe\xfa\x4f\xaf\xea\xab\xea\x10\xb4\x22\xc3\x18\xe9\xde\x3a\x0f\x9c\x01\x00\x54\x21\xc8\x7d\x89\x63\xac\xca\x08\x4d\x6d\x1b\x32\xdd\xea\x73\xb0\x66\x9a\x75\xe4\x2f\xe3\x87\xac\xad\x5e\xb5\xca\xd3\x65\x35\x78\x37\xd6\x7e\x98\x62\x3b\x54\x4c\x30\xe6\xbf\x7b\x84\xbd\x72\x4a\xa3\x47\x07\x85\x81\x90\x99\x9d\xd2\x98\x26\x64\x3a\x38\xa7\xcd\xeb\xca\x28\x8d\xd5\x39\xa7\xef\xe9\xe4\x6b\x9a\x96\x55\x67\x16\x19\x6b\x47\x53\x67\x79\x2e\xa6\x65\x04\xeb\x2d\xfc\x0a\x41\x1e\x2e\xd6\xf9\xc7\x0f\xc8\x10\xe4\x9b\x6d\xf0\x16\x63\x88\x2c\x93\x2d\xe1\xad\x19\x12\x3e\x19\xcb\x7f\x79\xc2\x49\xe4\xbc\x4f\xb6\x39\xd7\xea\x8a\xfc\x78\x7a\xe8\x2f\xe1\x86\x86\x97\xf3\xa2\xc0\xad\x75\xe5\xd3\xbf\x37\x40\xf0\xe7\x19\xd8\x00\x2d\x16\x93\xdd\xbc\xf7\x48\x27\xd8\xce\x8d\xcc\xf1\xbd\x91\xf5\x24\x78\xa4\x93\x4c\xef\x5c\x2c\x5f\x98\xd4\xcb\x33\xf3\x9f\x4c\xc3\x85\x3c\xe4\xa6\x9e\xe1\xc8\xe6\x27\x9a\x3a\x49\xa6\x1a\xe5\x0e\xbf\xfe\xa6\xfb\x44\xc7\xed\x20\x0f\xbe\xb1\xa3\x17\x77\x4a\x96\x88\x17\x59\xc1\xe2\x4f\x00\x00\x00\xff\xff\x76\x2a\x8a\x17\x2a\x02\x00\x00")

func refgenTmplBytes() ([]byte, error) {
	return bindataRead(
		_refgenTmpl,
		"refgen.tmpl",
	)
}

func refgenTmpl() (*asset, error) {
	bytes, err := refgenTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "refgen.tmpl", size: 554, mode: os.FileMode(420), modTime: time.Unix(1493268868, 0)}
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
	"refgen.tmpl": refgenTmpl,
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
	"refgen.tmpl": &bintree{refgenTmpl, map[string]*bintree{}},
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
