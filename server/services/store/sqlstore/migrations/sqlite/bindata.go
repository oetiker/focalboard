// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sqlite_files/000001_init.down.sql
// sqlite_files/000001_init.up.sql
// sqlite_files/000002_system_settings_table.down.sql
// sqlite_files/000002_system_settings_table.up.sql
// sqlite_files/000003_blocks_rootid.down.sql
// sqlite_files/000003_blocks_rootid.up.sql
// sqlite_files/000004_auth_table.down.sql
// sqlite_files/000004_auth_table.up.sql
// sqlite_files/000005_blocks_modifiedby.down.sql
// sqlite_files/000005_blocks_modifiedby.up.sql
// sqlite_files/000006_sharing_table.down.sql
// sqlite_files/000006_sharing_table.up.sql
package sqlite

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

var __000001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xb6\xe6\x02\x04\x00\x00\xff\xff\x45\xbe\x01\x0f\x13\x00\x00\x00")

func _000001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initDownSql,
		"000001_init.down.sql",
	)
}

func _000001_initDownSql() (*asset, error) {
	bytes, err := _000001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.down.sql", size: 19, mode: os.FileMode(420), modTime: time.Unix(1603074564, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\xbf\x6e\x83\x30\x10\x87\x67\xfc\x14\xb7\x58\x18\x89\x4c\x95\x3a\xa4\x93\x93\x1c\x8d\x55\x20\x95\xb9\xb4\x61\x8a\x28\xbe\xa8\x56\x49\x8a\xc0\x1d\xfa\xf6\x15\x91\xca\x90\x6c\xf7\xfb\x4e\xdf\xfd\x59\x5b\xd4\x84\x40\x7a\x95\x23\x98\x0c\xca\x1d\x01\x1e\x4c\x45\x15\x7c\x74\xdf\xed\xd7\x08\x4a\x44\xde\xc1\x9b\xb6\xeb\xad\xb6\xea\xe1\x31\x49\x45\xe4\x2f\x23\x0f\xe1\xd8\x04\xd8\x68\x42\x32\x05\x5e\xc5\x72\x9f\xe7\xb0\xc1\x4c\xef\x73\x52\x15\xd9\x6c\xea\xa8\x58\xd6\x0b\x79\x5e\x48\x07\x72\xbb\x94\xc5\x52\x9e\xe2\x14\xe2\x72\xf7\x1e\x27\xd3\xac\xbe\x19\xf8\x12\x8e\x77\x3b\xc6\xf6\x93\xcf\x0d\xac\xcc\xb3\x29\x29\x15\x51\xf8\xed\x19\x08\x0f\xd7\xda\x87\x6e\x0e\x27\xcf\x9d\x1b\xff\x53\x3b\x70\x13\x78\x3a\x6d\x36\x7f\x7a\x77\x8b\x1c\x77\x7c\x83\x5e\xad\x29\xb4\xad\xe1\x05\x6b\x50\xde\xa5\x30\x3f\x99\x88\xe4\x49\xfc\x05\x00\x00\xff\xff\xa2\x33\x30\x8e\x29\x01\x00\x00")

func _000001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initUpSql,
		"000001_init.up.sql",
	)
}

func _000001_initUpSql() (*asset, error) {
	bytes, err := _000001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.up.sql", size: 297, mode: os.FileMode(420), modTime: time.Unix(1607029839, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_system_settings_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xae\x2c\x2e\x49\xcd\x8d\x2f\x4e\x2d\x29\xc9\xcc\x4b\x2f\xb6\xe6\x02\x04\x00\x00\xff\xff\x8b\x60\xbf\x1e\x1c\x00\x00\x00")

func _000002_system_settings_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_system_settings_tableDownSql,
		"000002_system_settings_table.down.sql",
	)
}

func _000002_system_settings_tableDownSql() (*asset, error) {
	bytes, err := _000002_system_settings_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_system_settings_table.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1603229117, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_system_settings_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x28\xae\x2c\x2e\x49\xcd\x8d\x2f\x4e\x2d\x29\xc9\xcc\x4b\x2f\xd6\xe0\xe2\xcc\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x34\x30\xd0\xd4\xe1\xe2\x2c\x4b\xcc\x29\x4d\x55\x08\x71\x8d\x08\xd1\xe1\xe2\x0c\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\x54\xd0\xc8\x4c\xd1\xe4\xd2\xb4\xe6\x02\x04\x00\x00\xff\xff\x1e\xfb\x02\xf2\x60\x00\x00\x00")

func _000002_system_settings_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_system_settings_tableUpSql,
		"000002_system_settings_table.up.sql",
	)
}

func _000002_system_settings_tableUpSql() (*asset, error) {
	bytes, err := _000002_system_settings_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_system_settings_table.up.sql", size: 96, mode: os.FileMode(420), modTime: time.Unix(1603229117, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_blocks_rootidDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xca\xcf\x2f\x89\xcf\x4c\xb1\xe6\x02\x04\x00\x00\xff\xff\x94\x1c\x55\xb9\x28\x00\x00\x00")

func _000003_blocks_rootidDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_blocks_rootidDownSql,
		"000003_blocks_rootid.down.sql",
	)
}

func _000003_blocks_rootidDownSql() (*asset, error) {
	bytes, err := _000003_blocks_rootidDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_blocks_rootid.down.sql", size: 40, mode: os.FileMode(420), modTime: time.Unix(1610349080, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_blocks_rootidUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xca\xcf\x2f\x89\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\xce\x60\x70\x4e\x33\x00\x00\x00")

func _000003_blocks_rootidUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_blocks_rootidUpSql,
		"000003_blocks_rootid.up.sql",
	)
}

func _000003_blocks_rootidUpSql() (*asset, error) {
	bytes, err := _000003_blocks_rootidUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_blocks_rootid.up.sql", size: 51, mode: os.FileMode(420), modTime: time.Unix(1610349080, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_auth_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\xb6\xe6\x42\x12\x29\x4e\x2d\x2e\xce\xcc\xcf\x2b\xb6\xe6\x02\x04\x00\x00\xff\xff\xa5\xe0\x77\xaa\x27\x00\x00\x00")

func _000004_auth_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_auth_tableDownSql,
		"000004_auth_table.down.sql",
	)
}

func _000004_auth_tableDownSql() (*asset, error) {
	bytes, err := _000004_auth_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_auth_table.down.sql", size: 39, mode: os.FileMode(420), modTime: time.Unix(1610481092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_auth_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd0\xc1\x4a\x03\x31\x10\x06\xe0\xf3\xe6\x29\xe6\xb8\x0b\x3d\xd4\x42\x4f\x9e\xd2\x12\x35\xa8\x55\xd2\x20\xed\x69\x19\x36\x23\x06\xbb\x9b\x25\x93\xd5\xd7\x97\x20\x58\x68\x56\x2f\xe6\xf8\xfd\x21\x93\xf9\xb7\x46\x49\xab\xc0\xca\xcd\x83\x02\x7d\x03\xbb\x27\x0b\xea\xa0\xf7\x76\x0f\x13\x53\x64\xa8\x45\xe5\x1d\xbc\x48\xb3\xbd\x93\xa6\xbe\x5a\x2e\x9b\x85\xa8\x72\x34\x60\x4f\x97\x4e\x3d\xfa\xd3\x0f\xae\xd6\xeb\x8c\x23\x32\x7f\x86\x58\x3c\xd2\xbf\x62\xcb\xd4\x45\x4a\x97\x09\x4e\xe9\xad\x65\x8a\x1f\xbe\x3b\x8f\x58\x9d\x23\x87\x09\x8b\x29\x31\x8c\x0c\xdf\xc7\xaa\x83\x5d\x88\xaa\x8b\x84\x89\x5a\x4c\xd9\x36\xfa\x56\xef\xb2\x4e\xa3\x9b\x51\x47\x27\x2a\xf5\xd9\xe8\x47\x69\x8e\x70\xaf\x8e\x50\x7b\xd7\x88\xe6\x5a\x88\x3f\x2a\x63\x62\xf6\x61\xf8\xa5\xb5\x14\xde\x69\x98\xab\xb2\x2d\xef\xfe\x73\x9d\xb9\x8f\x7f\x05\x00\x00\xff\xff\x82\xd7\x37\x19\xeb\x01\x00\x00")

func _000004_auth_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_auth_tableUpSql,
		"000004_auth_table.up.sql",
	)
}

func _000004_auth_tableUpSql() (*asset, error) {
	bytes, err := _000004_auth_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_auth_table.up.sql", size: 491, mode: os.FileMode(420), modTime: time.Unix(1610481092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000005_blocks_modifiedbyDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcd\x4f\xc9\x4c\xcb\x4c\x4d\x89\x4f\xaa\xb4\xe6\x02\x04\x00\x00\xff\xff\xe4\x42\x8b\x2e\x2c\x00\x00\x00")

func _000005_blocks_modifiedbyDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_blocks_modifiedbyDownSql,
		"000005_blocks_modifiedby.down.sql",
	)
}

func _000005_blocks_modifiedbyDownSql() (*asset, error) {
	bytes, err := _000005_blocks_modifiedbyDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_blocks_modifiedby.down.sql", size: 44, mode: os.FileMode(420), modTime: time.Unix(1610481092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000005_blocks_modifiedbyUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcd\x4f\xc9\x4c\xcb\x4c\x4d\x89\x4f\xaa\x54\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\xea\xb0\x5a\x65\x37\x00\x00\x00")

func _000005_blocks_modifiedbyUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_blocks_modifiedbyUpSql,
		"000005_blocks_modifiedby.up.sql",
	)
}

func _000005_blocks_modifiedbyUpSql() (*asset, error) {
	bytes, err := _000005_blocks_modifiedbyUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_blocks_modifiedby.up.sql", size: 55, mode: os.FileMode(420), modTime: time.Unix(1610481092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000006_sharing_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xce\x48\x2c\xca\xcc\x4b\xb7\xe6\x02\x04\x00\x00\xff\xff\xdd\x4c\x62\xe8\x14\x00\x00\x00")

func _000006_sharing_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_sharing_tableDownSql,
		"000006_sharing_table.down.sql",
	)
}

func _000006_sharing_tableDownSql() (*asset, error) {
	bytes, err := _000006_sharing_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_sharing_table.down.sql", size: 20, mode: os.FileMode(420), modTime: time.Unix(1610482438, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000006_sharing_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\xb1\xca\xc2\x30\x10\x00\xe0\x39\xf7\x14\x37\xb6\xd0\xa1\x3f\x3f\xb8\x38\x5d\xcb\xa9\xc1\xda\x4a\x1a\xc4\x4e\x25\xe5\xa2\x06\xb5\x15\x8d\x83\x6f\x2f\x2e\x0e\xee\x1f\x5f\x69\x98\x2c\xa3\xa5\xa2\x62\xd4\x0b\xac\x1b\x8b\xbc\xd7\xad\x6d\xf1\x71\x72\xf7\x30\x1e\x31\x01\x15\x04\x77\x64\xca\x15\x99\xe4\x7f\x96\x66\xa0\xfc\xe8\x86\x8b\x17\x2c\x9a\xa6\x62\xaa\x33\x50\x71\x3a\xfb\xf1\xab\xfe\xf2\xfc\xc3\xae\x93\x84\x43\xf0\xd2\x0f\xaf\x9f\xe0\x79\x13\x17\x7d\xef\x22\x16\x7a\xa9\x6b\x9b\x81\xda\x1a\xbd\x21\xd3\xe1\x9a\x3b\x4c\x82\xa4\x90\xce\xe1\x1d\x00\x00\xff\xff\x6c\x91\x98\xb6\x9f\x00\x00\x00")

func _000006_sharing_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_sharing_tableUpSql,
		"000006_sharing_table.up.sql",
	)
}

func _000006_sharing_tableUpSql() (*asset, error) {
	bytes, err := _000006_sharing_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_sharing_table.up.sql", size: 159, mode: os.FileMode(420), modTime: time.Unix(1610483328, 0)}
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
	"000001_init.down.sql": _000001_initDownSql,
	"000001_init.up.sql": _000001_initUpSql,
	"000002_system_settings_table.down.sql": _000002_system_settings_tableDownSql,
	"000002_system_settings_table.up.sql": _000002_system_settings_tableUpSql,
	"000003_blocks_rootid.down.sql": _000003_blocks_rootidDownSql,
	"000003_blocks_rootid.up.sql": _000003_blocks_rootidUpSql,
	"000004_auth_table.down.sql": _000004_auth_tableDownSql,
	"000004_auth_table.up.sql": _000004_auth_tableUpSql,
	"000005_blocks_modifiedby.down.sql": _000005_blocks_modifiedbyDownSql,
	"000005_blocks_modifiedby.up.sql": _000005_blocks_modifiedbyUpSql,
	"000006_sharing_table.down.sql": _000006_sharing_tableDownSql,
	"000006_sharing_table.up.sql": _000006_sharing_tableUpSql,
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
	"000001_init.down.sql": &bintree{_000001_initDownSql, map[string]*bintree{}},
	"000001_init.up.sql": &bintree{_000001_initUpSql, map[string]*bintree{}},
	"000002_system_settings_table.down.sql": &bintree{_000002_system_settings_tableDownSql, map[string]*bintree{}},
	"000002_system_settings_table.up.sql": &bintree{_000002_system_settings_tableUpSql, map[string]*bintree{}},
	"000003_blocks_rootid.down.sql": &bintree{_000003_blocks_rootidDownSql, map[string]*bintree{}},
	"000003_blocks_rootid.up.sql": &bintree{_000003_blocks_rootidUpSql, map[string]*bintree{}},
	"000004_auth_table.down.sql": &bintree{_000004_auth_tableDownSql, map[string]*bintree{}},
	"000004_auth_table.up.sql": &bintree{_000004_auth_tableUpSql, map[string]*bintree{}},
	"000005_blocks_modifiedby.down.sql": &bintree{_000005_blocks_modifiedbyDownSql, map[string]*bintree{}},
	"000005_blocks_modifiedby.up.sql": &bintree{_000005_blocks_modifiedbyUpSql, map[string]*bintree{}},
	"000006_sharing_table.down.sql": &bintree{_000006_sharing_tableDownSql, map[string]*bintree{}},
	"000006_sharing_table.up.sql": &bintree{_000006_sharing_tableUpSql, map[string]*bintree{}},
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

