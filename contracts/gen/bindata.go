// Code generated by go-bindata.
// sources:
// compiled/Authority.abi
// compiled/Authority.bin-runtime
// compiled/Energy.abi
// compiled/Energy.bin-runtime
// compiled/Params.abi
// compiled/Params.bin-runtime
// DO NOT EDIT!

package gen

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

var _compiledAuthorityAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4d\x6b\x02\x41\x0c\xfd\x2f\x39\xef\x49\x6f\x7b\x2b\xf4\xd2\x43\xa1\x97\x9e\x44\x24\x3a\xb1\x0d\xac\x99\x61\x92\xb1\x9d\x8a\xff\xbd\xac\xac\x8e\x5a\xb5\x5f\x5b\xa4\xc7\x81\x97\xe4\xbd\x37\x8f\x37\x5a\xc1\xcc\x8b\x1a\x8a\x41\x6d\x31\x51\x05\x2c\x21\x99\x42\x3d\x5a\x81\xe0\x82\xa0\x86\x09\x3a\x17\xa1\x02\xcb\xa1\x7d\xb6\x2f\x52\x85\xf5\xb8\xda\x22\x42\xf4\xc1\x2b\xb5\x20\x9f\xec\x78\x5e\x0d\x2d\x69\x59\x90\x58\x6c\x38\x80\x75\xb5\x43\xb0\x23\x31\xb6\x5c\x30\x6a\x91\xe5\x69\x73\x23\x60\xc6\x69\x43\x50\xcf\xb1\x51\xaa\x36\xeb\xe8\x3e\x19\x4e\xb9\x69\x67\x6a\x58\x32\xbd\x94\xd1\x79\x92\x99\xb1\x97\xcd\x81\xa2\xae\x9b\xfe\x89\x3c\x47\x98\xec\xd9\x47\x7e\xa3\x03\x85\x5f\x22\x27\x5e\xb6\xa0\xde\x29\x16\x07\x27\x17\x2d\xec\x40\xd7\x56\x41\x32\xf3\x8e\x5c\xd9\x30\xcd\x46\x3a\x1c\x8c\xc6\xfb\x2c\x35\xeb\x63\x70\x68\xff\xdb\x6b\xcd\xfa\x10\x49\xc9\xae\xa4\x62\xe9\xad\x65\x74\x29\xd6\x9a\xf5\x4e\xd8\x18\x9b\xbf\x8e\xc4\x51\xb3\x7c\x2c\x0e\x3d\xdd\x1c\xe7\x92\xd2\x53\x25\x9c\xa5\xb5\x33\xef\x22\xa7\x7d\x4b\x7f\xc5\x08\xc5\x4b\x5e\xf8\xa4\xa7\x3e\x94\xc5\xd1\x2b\xb9\x2d\xdb\xcf\x53\xba\x1b\xe8\x76\x7d\x27\xb5\x37\x7b\x0d\xd1\x81\x68\x49\x62\x7d\xd3\x2c\x17\x6f\x0f\xba\xf5\xf0\xe6\xf8\x3d\x00\x00\xff\xff\xcf\x89\x9f\xfd\x9f\x06\x00\x00")

func compiledAuthorityAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledAuthorityAbi,
		"compiled/Authority.abi",
	)
}

func compiledAuthorityAbi() (*asset, error) {
	bytes, err := compiledAuthorityAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Authority.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledAuthorityBinRuntime = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x6f\xb6\xdb\xbe\x0a\xdc\x12\x20\x81\x60\x39\xb2\x6c\xed\x7f\x09\xef\x00\x72\xe2\xc4\xb9\x7f\xda\xbe\x0f\xbf\xe4\xdc\x36\x91\x6d\x04\xc3\x30\xa0\x08\xf8\xbb\x02\x93\x00\xd4\x22\x08\x82\x00\xba\x73\x93\x32\xd7\xab\x0d\x40\xf8\xcb\x97\x00\x40\x61\xa8\x28\x52\x80\xeb\xd4\xb1\x4d\x45\xac\xbe\x8b\x11\x37\x05\x29\xd4\x86\xb5\xd9\x39\x56\xb1\x70\xae\xf2\x01\x43\x77\xd5\x5c\x65\xc9\x55\xdd\x66\xdd\xc7\xd8\x73\x75\x5b\xf7\xf6\xb6\x19\x0d\xb1\x58\x25\xa8\xb9\xba\x15\x23\xd9\xb4\xe7\xaa\x94\x5c\x1d\x7d\x9b\x0a\x96\x76\x49\x97\x0f\x73\x1c\x38\xc6\x58\xf7\x1e\xca\x8d\x37\xf7\x5d\x61\xee\xbc\x95\x8a\x1c\x1e\x3b\x2e\x8f\x55\x5f\xd9\x50\x00\x50\xa0\x7b\xa4\x04\x1d\x4a\xe0\xc8\x28\x82\x50\xb0\xb1\xb8\x9d\x0a\x8c\x4f\x3c\xb5\xa0\x28\x3a\xe2\x15\x04\x08\x94\x00\x15\x7d\xc5\xf2\x73\x89\x6f\x98\xf7\x10\x00\x1a\x32\x28\x30\x1a\xac\xef\xa0\xa0\x45\x7d\x2f\xe0\xcd\xef\x0f\x88\x10\x60\x9a\xc7\xe3\x56\xd8\xed\xf8\xff\xcb\x86\x5f\x3d\x1c\xc7\x8d\x21\xdf\x16\x7f\x8a\x10\x76\x71\xa2\x68\x5a\x41\x92\x65\xa5\xf8\xae\x1e\xa1\xef\x45\x00\x4e\x10\xe7\x42\x87\x82\x76\x46\x71\x7a\xe8\x96\xad\x9c\xd6\x33\x6a\x05\x43\x28\x06\xb3\x3c\x50\xc4\x0a\x6f\x28\x22\xd7\xef\x50\x9c\xc3\xbd\xf6\x48\x4f\x0b\x82\x1f\x2d\x54\x85\xc2\xef\x96\x50\x3c\xba\x5a\x05\x28\xae\x9f\x11\x6b\x71\xae\x69\x20\xea\xd1\x67\x16\xcc\xad\x38\x87\x30\xfd\xcf\xbb\xb3\x42\x56\x8e\xc8\x56\x1e\xb4\x46\xae\x8a\x23\xa5\xb5\x34\x06\xab\x26\xec\x15\x54\x75\x37\x3e\x91\x88\xb7\x3c\xbd\x1f\x9f\xe3\x4f\x0f\xdd\x27\xbc\x7b\x18\x6b\x9f\xbd\x02\x2b\xe9\x93\x5e\x18\x93\xec\x20\x00\xfa\xe0\x1d\x77\xfa\xd2\x3b\x82\xf9\x1f\xc7\x56\x68\xfb\xda\x7b\x39\xfe\x88\x5b\xd2\xcb\xf5\x69\x7d\xaf\x6f\x32\x5f\x97\x49\x97\x3a\x76\xef\x40\xc9\x7d\xce\x08\x7e\x53\xaf\x8f\x5c\x7c\xaa\x5a\xda\xeb\x37\x55\x4b\xdb\x78\xaf\x5a\x40\xa3\x1f\xaa\x8c\x66\x79\x8f\x65\xfa\xbf\x8d\xf1\x1a\xcb\x2b\x32\x06\x86\xd7\xaa\xbe\x5b\x77\x8b\x5e\x93\xe4\x9e\xd9\x98\x61\x0b\xa1\x14\x0a\x95\xa8\x4f\x95\xf3\xbd\x04\x50\x72\xb7\x40\x1d\xf4\x03\x83\xee\x2b\x17\x1c\xe1\xea\x0b\x24\xf2\x90\x19\xd5\xf0\xde\xa5\x60\x69\x12\x2d\x6b\x14\x3c\x0b\xde\xc1\xca\x96\x21\x54\x78\x89\xe9\xad\x82\xee\xf5\x43\xbf\xdd\x27\xd6\xcb\x31\xa2\x8f\xb8\x8a\x06\x3e\xc3\x15\x2a\xee\xf7\x9c\x73\xf5\xfd\xb5\x30\x19\x9e\x0c\x59\x4f\x09\x6f\x9e\xef\xe4\x3d\x40\xee\x9f\xb1\x9a\xb3\x04\xb9\x66\x67\x70\x2f\x1e\x4c\x42\xb7\x8a\x61\x63\x4c\x6e\xea\xbe\x94\xa5\xe1\x61\xed\xc9\x97\xa5\xf3\xf8\xf2\x97\x39\x01\xae\xa5\xdc\x39\xa0\x88\xe2\x3c\x88\x5e\x58\xb1\xbd\xf1\xa8\x92\x06\x2e\xf4\x9a\xeb\xad\x65\xae\x91\x03\x91\x5a\x5e\xf9\xf7\xb6\x8b\xef\xe1\x57\xcf\x0c\x43\xa8\x55\xf6\x44\x45\x0a\x0d\xe5\x1a\x28\x59\x76\x2d\x85\x3b\x53\xda\x6c\x76\xec\x47\xd1\x62\x63\x6f\xbd\xed\xba\xed\xad\x1d\xd6\x0e\x1e\x6a\x2e\x0c\xf3\xa8\xb3\xc9\xd8\xb7\xed\xe8\xbd\x48\x95\x2a\x07\x69\x57\x31\xee\x4d\xb5\x9e\x75\x70\xe5\x79\xa7\xd4\x94\xdf\x22\xd4\xf5\x2d\x52\x8f\x09\x23\x3b\x75\x93\x77\xf4\x46\x5b\xfd\xf4\x8a\x9e\x76\x7c\x45\x6f\xa7\x6f\xd1\xa3\xaf\xd1\xb3\xec\x16\x1e\x15\x42\x9d\x9b\x3d\x39\xe7\xa8\x76\xd2\x54\x13\xbd\x55\x7f\xe0\x79\x74\xa4\x9d\x4b\xb3\xbd\x57\xdb\x58\xb0\xb4\x42\x6d\x1e\xa5\xf6\x71\x8c\x43\xb9\x1f\xa6\x7d\x1c\x73\xe8\x31\xa6\x1d\x3a\x67\x11\x56\x02\x82\x7d\x53\x56\xfa\x17\x95\xfc\x42\x1f\x59\xbe\xd3\x47\xae\x37\x7d\xfc\x38\xd5\xb0\xe1\x5f\x4f\x35\x1f\xd5\xb6\xd3\xd9\x7d\x62\xca\xfd\xc0\xcf\x3b\x77\x9c\x33\x2b\xc7\x3c\xea\x35\xc7\x0c\x2b\xee\xc0\x68\x45\x2e\xd4\x42\x47\x04\x9d\x25\x3c\x29\x7a\xe2\xba\xc3\x57\x8e\xc2\x6d\x1e\xbc\xbd\x77\x1a\x47\xca\xab\xb2\x1e\x4b\xa1\x9d\x64\x55\xeb\xe5\x59\x01\xfc\xf6\x59\xc3\xba\x62\x4b\x75\xba\xf2\xd5\x68\x3e\x94\x1d\x33\x0b\x43\x1f\x59\xf0\x8e\xfa\xa7\x68\x48\xdd\xbe\xac\x22\x61\x7b\xab\x22\x91\xfe\x8b\x2a\x92\xc6\xff\x97\x2a\x12\x3b\xbe\xaa\xa2\xbf\x8f\x77\xbc\xeb\x63\x74\xb6\xf6\x88\xe7\xa7\x97\x33\xf5\xa6\x4c\x64\x60\xe4\xda\x14\x9d\xbc\x19\x30\x9f\xf9\x9f\xfd\xd2\xb1\x63\x16\xf1\xfe\x96\x55\xfa\xe5\x64\xf9\x45\x5f\x5c\x15\xd5\x6a\xa8\x5b\x4c\x36\xbf\xe8\x61\xc1\x24\xaf\x9d\x7b\x17\x6b\x05\xfd\xc4\xf5\xda\xb3\x9c\x99\x4f\x1d\xfe\x38\x27\xc8\xcf\x9d\x85\x40\x8b\x5d\x3a\x4b\x7a\x9f\xfe\x2b\xbf\xf1\xaa\xf5\x5d\xeb\xe2\x74\xf0\xcb\xed\xb9\xf7\xf8\x78\xc6\xc6\xb3\x6a\xde\x63\x8d\x69\xf3\x51\x39\x4f\xf5\x88\x4f\xf0\xaf\xde\x1e\xeb\xdc\xe0\xd6\xa2\xe3\x3e\x2c\xc6\xb4\x9c\x18\x01\xda\x9a\xc1\xb3\xca\xdb\xac\xdf\xfb\xfb\x43\xbc\x8a\xeb\xf9\xdc\xd9\x15\x3d\xcf\xac\xb4\x74\xc1\x90\x63\xaa\x53\x4e\x1f\xb4\x18\x69\x0b\x26\x7a\x66\x95\xe9\xe4\x05\x9a\xca\xc5\x33\xad\xf8\x23\x92\xa7\x7e\x7d\xec\xc0\x0f\x5f\xe2\xbc\x61\x3e\xbd\x96\x55\xbd\xa8\x94\x13\xa3\x05\x8a\x56\x0d\xac\x32\xab\xa0\xc4\x14\x06\x2a\x16\xa7\x0e\x82\xe5\xfd\x42\x30\xaa\x5b\x4d\x1e\x1e\x97\x50\xb7\xde\x25\xeb\x3d\x38\x5c\x9f\xa7\xdc\x93\x9f\xa1\x3e\x7b\x74\x2f\x04\xd3\x9a\x2a\xe4\xab\x5b\xe7\x96\xda\xee\x57\xe0\xc8\x89\x2f\x31\xcf\x5a\x5a\x9f\x49\x03\xaf\x1d\x6c\x69\x4c\xee\x69\xf8\x09\x9d\x73\xcf\x43\x63\xa6\x0e\xdb\x3d\x4e\x29\x99\x91\xb3\x5f\x65\x8d\x7d\x3a\x05\x69\xb9\x31\x8f\xed\xa2\x7e\xcc\xbc\x5d\x99\xdb\xea\xef\x7f\x05\xf2\xc9\xd8\x3e\xf3\x9d\x6e\xbb\xd6\x97\x5d\x9f\x7c\x37\xae\x6f\xd5\x69\x7c\x9c\x08\x77\x7a\x70\x9f\x5f\xb8\x5f\x2f\x0c\xb3\x46\x9f\xb8\xeb\x08\x39\x7b\x99\xdf\xb2\x78\x3f\x99\x5c\xfc\x5c\x5c\xbf\x78\xea\x7e\xae\x28\x1f\xaa\x8d\xf2\x27\x38\x01\xdd\x98\x9d\xca\x2d\xa1\xdd\xe1\x15\xc1\x5d\x99\x57\x77\xa4\xbc\x87\x7f\x79\x26\x8a\xa7\x88\xf9\x72\x4e\xe9\xd0\x82\x9d\xde\xd1\xb2\x53\x3c\x4e\x4d\x99\xc3\x3b\xf7\xd6\x64\x15\x4f\x9c\x4c\x1d\xab\x3a\x92\xdb\xbf\xf3\xe6\x93\xed\xd7\xdd\x63\x27\xf2\x5a\x39\xbd\x95\xf8\x75\xcb\x67\xb7\xec\x7f\x5a\xd4\x77\x70\x1d\xf7\x3d\xba\xc9\xe9\x41\xea\x99\xb3\x83\x59\x09\xcf\xab\x2d\x95\x21\xee\x41\xc4\xe7\xba\x12\x3b\x3f\xf8\x79\x3e\xbb\x9c\xb8\x10\x7a\xdb\x56\xfd\x7b\x5d\x52\xd6\xc5\x35\xf2\xcc\x44\x6a\xa7\x16\x8e\xdf\x70\x7c\x7c\xc9\xe9\x61\xf1\xb8\x3c\xb8\x04\x71\xc2\xf1\xb9\xe2\xf2\x5b\x4f\x76\xa0\x7a\xf8\xde\x7e\xde\x51\x3a\x3d\xec\xe7\x69\x60\x69\x6d\x6a\xd2\x18\xfe\xa5\xfb\xc8\x43\xad\xb7\xde\xa8\x80\x4f\xe2\xb0\x6d\x4d\xec\x20\xc4\x26\xfb\xf0\x1e\x54\xf4\xd8\xf6\x42\x2a\x07\x30\xb7\xda\x2b\xd0\x01\x53\x61\x02\x58\xc5\x52\x9a\x28\x4e\xdc\xb7\x3a\x60\x8e\xe6\x09\xfa\x5f\x00\x00\x00\xff\xff\x50\xb9\xa7\xdf\x18\x16\x00\x00")

func compiledAuthorityBinRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_compiledAuthorityBinRuntime,
		"compiled/Authority.bin-runtime",
	)
}

func compiledAuthorityBinRuntime() (*asset, error) {
	bytes, err := compiledAuthorityBinRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Authority.bin-runtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledEnergyAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x41\x6f\xdb\x3c\x0c\xfd\x2f\x3a\xfb\xd4\xef\x5b\x31\xe4\xd6\x16\xe8\x4e\x5b\x81\xb6\xdb\x0e\x45\x51\xd0\x32\xd3\x08\x93\x25\x43\xa2\x9c\x7a\x45\xff\xfb\x60\xc7\xb1\x9d\xc6\xb2\xe2\xc5\x4d\xb2\x53\x0b\x98\x94\xf8\xa8\xc7\xc7\x97\x87\x57\xc6\xb5\xb2\x04\x8a\xd8\x8c\x8c\xc3\x88\x09\x95\x39\xb2\x6c\xf6\xf0\x18\x31\x05\x29\xb2\xd9\xea\x4f\xc4\xb4\xa3\xfa\xd3\xeb\xfa\x0b\x8b\x18\x15\x59\xf9\x9f\x25\x23\xd4\x33\x7b\x7b\x8c\x58\x06\x05\xc4\x12\xd9\x6c\x0e\xd2\x62\xc4\x2c\x01\xe1\x57\x47\x10\x0b\x29\xa8\x60\x33\x96\x39\x83\x6d\xea\xdc\x29\x4e\x42\x2b\xf6\x16\x75\xcb\xa9\xb3\x9b\x7a\x9a\x4b\x9f\x0c\x72\x81\x39\x9a\xf6\x08\x48\x12\x83\xd6\x56\x27\xac\xa3\x20\xd5\x4e\x51\x1b\xe3\x84\xa2\xb3\x4f\xe7\x55\x89\x75\x0c\x64\x99\xd1\xb9\x07\x9b\x75\x9c\x97\x67\x36\x07\xc4\x5a\xcb\x1d\x01\x2a\xad\xd6\x41\x1f\x0d\x53\x8a\x54\xf4\xa1\xec\x84\x70\x83\x89\xa0\x2f\x46\x2f\x69\x71\x0b\x84\x9b\xd1\xe7\xff\x6f\x04\xe3\x4b\x26\x4c\x4f\x48\xdb\x35\xbb\x00\x83\xf7\x7a\xa3\x6b\x53\x77\xc5\xcb\x45\xd2\x04\xf2\xce\x65\x99\x2c\x42\x94\xec\x3e\x78\xb8\xb8\x5c\xe0\x72\x8f\xc7\xe2\x5a\x91\x01\x4e\x17\x49\x12\x78\x30\x85\xcb\x9b\xa5\xea\x7d\xd6\x4e\x8f\x91\xbe\xd5\x71\xd7\xda\x5c\xd5\x87\x7f\x68\xcb\xfd\xd8\xe6\x46\xa7\xc3\x98\x48\xef\x3d\x8b\x64\x40\xd9\x39\x9a\xeb\xd5\x65\xc7\x1b\x48\x2f\xf5\x12\xe4\x22\x05\x69\x77\xe1\xdd\xe7\xc3\x28\xe1\xee\xac\xd3\xbb\x50\xee\x44\xf8\x16\xae\xd5\x65\x09\x10\x5e\x82\x04\xc5\x83\xab\x69\x9c\x0e\x4c\x82\x80\x83\x94\x21\xe1\xae\x62\x70\xef\xb9\x29\xcb\x71\xbe\xfd\xfc\x54\x7f\xf5\x74\xf3\x30\x73\x23\x51\xd1\xe2\x66\x7e\x8b\xb9\xb0\x42\xab\x9d\xe6\x67\x62\xdd\x7e\x57\xdb\x18\x69\x1b\xd8\xc2\x2d\xc4\x67\xa4\x8b\x1c\x84\x2c\x6b\xbd\xaa\xb6\xed\x29\xa1\xcc\xd1\x94\x8d\x1f\xa4\x51\x2c\x0c\x2d\x7e\x0a\x5a\xfc\xa8\xa0\xfe\x8b\x8b\x75\x32\x5b\x58\xe9\xcf\xc5\x09\x78\x43\x2f\x6d\xc3\x0a\x19\xaf\xb4\xf1\x66\xde\x0f\x20\x6e\xa4\xf3\x60\x74\xec\x6c\x9a\x22\x8d\xb5\x3c\xfc\x2f\x0a\x5f\x3b\x3b\xcf\x58\x10\xda\xff\xce\xb6\xdc\x6e\x32\x38\xd4\x61\x0b\xee\x33\xd5\x6d\xc4\x18\x8f\xbe\x8a\xb5\xdf\x2d\x26\x83\xb7\x5a\x05\xd9\xbd\x48\xfb\xcd\xfc\x87\x0f\xee\x64\xe3\xc8\x17\x60\x9e\xf1\x48\x5e\x64\x42\x67\x7b\x92\x52\xb2\xcb\x6e\x20\x91\xe2\x29\xae\x86\x5c\x53\xa9\x10\x43\x32\x28\x94\x20\x01\x52\xfc\x3e\x16\x7f\xaa\xbd\x3a\xd8\x5d\x8b\x54\x1b\xd9\xcb\x3a\xf6\x60\x3f\xab\xc3\x0b\xa5\xc3\x74\x9b\xa1\x4a\x02\x5b\x07\xa4\xd4\x4b\xbf\x27\x37\x98\x82\x50\x1b\x6f\x76\xc8\xbd\xd3\xf0\x65\x90\xc2\xe3\x0c\xf2\x40\x45\x4d\xf2\xaa\xa0\x9e\xa5\xb5\xf5\x6a\x20\x65\x0c\xfc\x57\x95\x0e\x4a\xab\x22\xd5\xce\xf6\xd1\x4b\xa8\x04\x5f\x30\x59\x9f\x1d\xb6\xb3\x9e\x04\x9f\xbc\x35\xe1\xf5\xd5\xcd\xc8\x81\x74\x1e\xd7\x50\x87\xdc\xb7\x6a\x57\x07\x61\x8e\x8a\xfe\x1a\xd2\x00\x33\x3d\x19\x7e\xa2\xee\x05\x6c\x65\x08\x41\x4e\x04\x0c\x5f\x90\x3b\xd2\xe3\xca\xa4\xad\x6d\xbe\x5e\xf9\x9e\x84\xb0\xfc\xdc\x6d\xc9\xcf\x14\xe8\xc6\x12\x71\x2c\x0f\x07\x2c\x97\x27\x65\xd8\x60\x05\x92\xfb\xfc\xdb\xfb\x3e\x96\x3e\xf1\x7a\x03\x76\xdd\xc1\xc7\x3f\x01\x00\x00\xff\xff\xcc\xfe\x21\xc8\x66\x17\x00\x00")

func compiledEnergyAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledEnergyAbi,
		"compiled/Energy.abi",
	)
}

func compiledEnergyAbi() (*asset, error) {
	bytes, err := compiledEnergyAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Energy.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledEnergyBinRuntime = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x0d\x76\xeb\xa0\x8e\xde\x92\x7e\x41\x2c\x07\x0c\xec\x7f\x09\x73\x04\xd8\x71\x6c\xa7\x4d\xef\x9b\x79\x67\x92\x73\x6f\x5b\x62\x40\x12\xd2\xa7\x4f\x22\x01\xfc\x2d\xa0\x14\x00\x84\x03\x42\x40\x40\xae\x1a\x03\xf7\xf5\x8a\x1b\x20\xfc\xe3\x2b\x00\x00\x2b\x08\x86\xc0\x10\x7a\xad\x0d\xd8\x10\xc5\x77\x11\xd2\x68\x10\x18\x92\xb6\x1c\x0b\xcf\xd1\x6d\x5b\xa3\x8d\x1b\xd9\x36\x9f\x25\x58\xcf\xa2\x61\x80\x5a\xeb\x1c\xe5\x34\x47\x89\x31\x91\xee\xcf\x6a\xdb\x47\x8b\x45\xda\x9f\x35\x9e\xa3\x8c\xbc\x35\x0d\x71\x8e\xe6\xb2\x46\x43\x07\xb1\x38\x65\xa0\x2a\x73\x54\xa0\x98\x80\xe6\x39\xda\xd7\x6e\xb2\xe5\xa4\x29\xc3\x18\x65\xb4\x39\xaa\xac\xb9\x06\x9d\xbb\xb1\x2e\x2d\x34\x48\x02\xe5\x3e\x47\x43\x5f\xa3\x4d\x29\x43\x9c\xf2\x72\x5a\xbb\x85\x2a\x01\x74\x5b\xcf\xe6\x3c\x47\x23\x64\x30\xe2\xf5\x6c\x5d\x76\x48\x5a\x2d\x15\x59\xa3\x1d\xd7\x68\xea\x00\x22\x73\x05\x81\xb5\x6e\xe6\xde\x73\xda\xa6\x64\xa2\x4b\x8b\x9c\x40\xd3\x56\xca\x1c\x8d\xcb\x0e\x25\x6e\x0d\xab\x4c\x8d\x25\xd5\x39\xba\x49\x0d\xa1\x36\x9b\xa3\x65\x59\x72\xd3\x12\x3a\x26\x9d\xa3\xbb\x64\xb5\x06\x6a\x95\xdb\x1c\x6d\xcb\x3a\x7d\x6b\xb8\x6d\xdb\x5c\x57\xa1\x6a\xd4\xe2\xbe\x61\xd0\xab\x16\x16\xd4\xe1\x11\xee\x77\xc7\xa8\x8f\xa8\x8f\x2b\x81\x06\x7f\x5e\x40\x31\x00\x81\x81\x91\x92\x61\x02\x43\x40\x63\x43\x43\xf5\x7f\xe4\x9f\x02\x26\x54\x30\x50\x4c\xb0\xfe\xf6\x19\x6c\xec\x6b\x6b\xf1\xe7\x7d\x69\x5f\x3f\x0d\xcb\x19\x81\xcf\xe7\xf1\x73\xad\xe1\x9f\xc6\xe4\xfb\x2a\xcc\x77\x1a\xff\x7c\x47\x5f\x17\x3b\x06\x5b\xab\x94\xb6\x56\x61\x1b\x12\xba\x4c\x63\x15\x0f\x28\x8f\x9d\x0c\x8c\x09\xc3\xbb\x84\xbe\x72\xa2\xb9\xf6\xd4\xcd\x20\xb9\x3f\x40\xe7\x97\x4d\x6a\xbc\xda\xa4\x35\xdf\x21\x40\xf6\xe8\x22\x5f\x7b\xc4\xae\x62\x08\x40\xc2\xc3\x62\x81\x5e\x16\x4b\xae\x9c\xbe\xf6\xfe\xb8\x17\xc1\xd5\xfe\xc4\xf1\x97\xbd\x62\xbf\xbc\x02\x88\xb0\x1a\x62\x70\x2b\x85\xf9\x2c\x82\x7a\xd0\x69\x71\xfb\xef\xbb\x89\x5c\x77\x93\x2d\x20\x44\x6b\x67\xd9\xbf\x91\x3b\xa4\x6f\xe5\xde\xe5\xa2\x5d\xae\x98\xd8\x77\xdb\x57\xb2\xf6\xad\xb5\x2f\x2b\x0d\xad\x03\x82\x61\x3f\xaf\x57\xc2\x55\xb2\xd2\x02\x42\xea\x72\xf6\xe8\xde\x13\x24\x3c\xfb\xc7\x67\x5d\x6b\xff\x67\x5d\x53\x4f\x27\xd9\x18\x9e\xec\xff\xe9\xb4\x11\x72\xda\xce\xb3\x89\x2f\xb3\x59\xe0\x4f\x96\xda\x44\xcf\x36\x78\x9f\xf9\xb5\x3d\x38\x5c\xe3\x63\x7a\x51\x8d\xed\x2c\x6d\xcc\xdf\xea\x7a\xb3\x5a\x35\x39\xaf\x94\x6e\xf6\x1f\x2b\x4d\x9d\x5a\xb2\xf3\xb3\x45\xff\x43\x6f\x6a\x2d\x9f\xd7\xab\xb7\xf8\xfc\xf1\xc4\x7a\x2e\xe7\xd9\x7d\x7b\x42\x57\xc4\xad\x9e\x9e\x12\xb8\xea\x27\x98\x77\xfd\x90\xe0\x1c\x9b\x62\xaa\x74\x47\x80\xc4\x26\x03\x1f\xc0\x82\xe3\x69\x42\xe3\xa1\x93\xa3\x70\x70\xec\xf3\xd3\xf5\x93\x0d\x60\xfa\xfa\xdb\x4f\x3c\x0c\xa4\x9e\x28\x9c\xc1\xc7\x1d\x6d\x1f\x4f\x5e\xc2\xcd\x7f\x7f\x45\x2b\x44\x52\x3c\x6b\x6b\xd7\xf8\xfc\x02\x5d\x91\xf2\x39\xc6\x25\xdb\x67\x8f\x40\xae\xe7\x98\x93\x72\xc5\x97\x9f\x64\x46\x14\x3a\x9f\xbf\xdc\xcf\x7f\xcc\x9e\x3b\x49\x7c\xdb\xa9\x5f\xe3\xf3\x6b\x8f\x47\x7d\xf3\x78\xc5\xab\x7e\x1e\xe9\x88\x3a\xad\x30\xf2\xb3\x05\xc4\x20\x79\xf7\x0d\x5b\xc8\xbd\xe2\xd6\xb9\x65\xf6\x58\x8e\x5d\x83\xa8\x0a\x81\x68\x68\x41\x23\x85\x18\xd3\xdf\xf8\x24\xcd\x6c\x4d\x9e\x8d\xb5\xa4\xc5\x0c\xde\xf5\x62\xcf\xf0\x61\xc8\xbc\x58\x01\x0c\x1c\xd9\xb3\xc9\xf0\x44\x26\x48\x62\xd1\xb3\xb8\x89\x52\x92\xe9\x85\x9e\xc7\x09\x4c\x13\xa8\x26\x48\xe4\xde\x19\xbb\x6d\x5a\x1a\x36\x6d\xa5\x6d\xb1\x6a\xa9\x28\x3d\xa2\x50\xac\xd8\x4c\x3a\xd7\x0a\x8c\xb2\x41\x8f\x85\x28\x61\xd3\x42\x00\x79\xb3\x2d\x6e\x5c\x12\x69\x9a\x2b\xfe\x90\xcf\x32\xeb\xc0\x4f\xdf\x71\x70\x82\xb0\x33\x24\xb7\x70\xcd\xcc\x6e\xf3\x32\xe2\xcf\xf3\xb9\x8f\xb6\x2b\xbe\x68\x6b\x16\xce\xcf\xf9\x58\x7f\x3b\x3f\x79\x88\x58\x8f\x57\x74\xc6\x13\xe0\xcd\x3f\x99\x07\x93\xd2\xab\x7d\x31\x18\xdd\xc7\x62\x4f\xce\xfc\x18\xa8\x41\xe8\x5b\xc7\xad\xb7\xa0\xcc\xc0\x8d\x5a\x2e\xc1\x5a\xeb\xba\x61\xa7\x98\x10\x7b\x4e\x35\x49\x27\x13\xb0\xde\x92\x28\xf5\x68\x16\xfd\xbd\xd0\xc5\x19\xfb\x13\xba\xa0\xd1\x42\x17\x71\x3f\x38\x90\x85\xdf\x91\xe5\x8e\x1d\xd9\x79\x62\x58\xf9\xe6\x87\x9a\xe6\x96\x8b\x06\x9e\x99\xb3\x03\xd2\xe1\x0f\x89\x9c\x86\x06\x40\x39\x18\xa3\x5d\xcf\x93\x06\x5f\x9c\xeb\xc8\x5b\x3c\x4c\x1f\xb0\xe0\xab\x3d\x9c\x05\x58\x58\xfa\xbd\xa1\xa3\x45\xd7\xdb\xd8\xc7\x86\x47\xc3\x8e\x9f\xce\x83\x95\x84\xa6\x5c\x61\xcd\x53\x32\x19\x3b\x00\x78\x09\xe2\x3e\x3d\x39\xb2\x1e\x31\x34\x99\x2a\x01\x1a\xa8\xdc\x25\x71\xc6\xfa\x60\xff\xb5\xcf\x94\x0b\xa3\xc7\xc9\x38\x81\xc9\xa1\xc7\x9a\xae\xa7\xab\x31\xb9\xaf\xa1\x8a\xe1\x7d\x25\x20\xdf\x21\x81\xf1\xc3\xde\x01\x68\xae\xad\x53\xcf\xd7\xea\x34\xf2\x16\xd8\x4b\x07\x9e\xb2\xfc\x51\x8f\x9b\x16\x3b\xcf\x5f\xef\x15\x7f\x0b\x61\x9c\x29\x6e\x46\xb7\x18\x8c\xf9\xca\x3f\x62\x01\xc3\x4b\x0c\xc6\xf2\xce\x3f\x2e\x1e\x66\xb4\xe3\xd5\x8e\x0f\x30\xea\x87\x89\x55\xee\x4b\x2a\x4c\x36\xe4\xc5\x80\x13\x97\x0d\x8b\xde\x72\xc7\xef\x2b\xb9\x8d\x22\x5f\xad\xf1\xe9\x35\x4e\x68\xe4\xef\x37\x4b\x95\x0b\x42\xb9\x37\x22\x18\xd5\x1b\xfa\x98\x3e\x49\x09\x8e\xbb\xdf\xeb\x7b\x5a\xe7\x65\x67\x0b\x62\x7a\xdd\x2d\x5c\xf3\xab\xc5\x68\x27\xf6\xea\x78\xe6\xa3\x46\x3e\xf7\x18\x05\xf5\xfa\x90\x9c\x12\xad\x9a\xce\x32\x69\xd4\x89\xbe\x96\x61\x20\xc7\x61\x35\xdf\xb5\x09\x86\xb1\x2f\x6a\xd9\xe7\x54\xf1\x39\x5f\x68\xfb\x9c\x8d\xd8\xbc\x46\x64\x25\x8f\x31\x8f\x58\x15\x93\x04\x83\xa8\x4d\xfd\x52\x0b\x0f\xd6\x0c\x97\xd5\xf1\x6a\x4b\x9f\x09\xe4\x39\xe8\x5d\x87\xce\x53\x87\x07\x89\xed\xf3\x9a\x33\x43\x26\xe7\xe8\xec\x39\x72\xe4\x50\x9c\xf2\xfa\x4e\x1c\xef\x3b\xed\xd6\xfa\x6c\x9b\x23\x53\xe3\xa3\x6d\x3c\x17\xab\x6a\x22\xcb\x7b\x5d\x30\x7a\x3a\xe4\xd5\x3c\x41\x1a\xb3\xd3\x49\x8a\x98\xff\x45\xdf\x5f\xf9\x42\xb9\xf2\x05\x7f\x66\xa0\xfd\x61\x8f\x84\xb1\xd7\xda\x49\x29\x57\x2c\x8d\x36\x4b\x25\xa4\x8d\x0a\x04\xeb\x1b\x47\xab\x39\x27\xa5\x92\x63\xc7\xc0\x9b\x64\xc4\x40\xd6\x55\xb3\xd4\xae\xc4\x85\x5b\x77\xab\xfe\xc6\x17\x26\x5b\x18\xfe\x9c\x5a\x39\xf0\x8a\x66\x67\x41\x5d\x22\x3a\x63\x19\xd2\xc4\x32\xbe\x79\x28\x06\xe6\xfb\x98\xe3\xcc\xf0\xba\x8c\xd7\xfa\x3a\x13\x3d\x60\x61\xa6\x2b\x3f\xcd\x1c\x6e\x58\x98\x05\x7e\xc4\x42\x98\xd5\xc2\x4f\xd8\x80\x61\xed\xf8\x88\x81\x3b\x43\xf8\xdf\xc5\xc0\x99\x2f\x16\xfe\x3d\xf9\xd0\x15\x77\x67\xa7\x6a\xe4\xe6\x99\x1f\xfd\x0c\x3d\x13\x26\x78\xca\x77\x4b\xa3\x02\xf4\x10\xe1\x57\x8b\x5c\xd6\x76\x9f\xf3\xd5\x3f\xe7\xbf\x27\xd6\x77\x68\x53\xc6\xbe\xc5\xf8\x55\x33\xde\x24\x90\xa7\x28\x7d\x59\x74\xb1\x94\xa1\x01\x63\x1a\xee\x83\x01\x9e\x2d\x25\xdf\x6b\x13\x6f\x52\x5f\x5f\x57\xf6\x36\xb2\xbd\xdd\xd9\x5d\xf8\xf5\x9c\x67\x04\x63\x98\x0c\x07\x68\xcf\x78\x7f\xb7\xa8\x5b\xa1\xd0\xab\x0b\x31\x99\xa0\x5b\x0f\x71\x9c\xf2\xb6\xba\x27\xff\x7c\xca\xbf\xdb\xe5\x2e\x5b\x7c\xb0\xca\xef\xf5\x96\x5b\xd4\x99\xf6\x8f\x1c\xed\x7e\xc6\x8f\xba\x9c\x63\xf8\xe8\x80\x62\xba\x71\x09\x83\x3f\x63\xd4\x16\xae\xfc\x60\x8b\x6a\xc1\xec\xd5\xc1\x71\x34\x34\x31\x1e\x77\x1c\xd5\x2b\x1d\xd3\xbd\xda\xf8\xbf\xab\x07\x66\x7e\x98\x15\xc1\x5b\x04\xc9\xbb\x45\x1c\xa5\xec\x47\x0e\xfb\xa3\xc7\x4d\x5e\xec\x39\x69\x66\x84\xad\x4b\x82\xc9\xc6\x9a\x1c\xd6\xe5\x8f\x7b\x0f\x56\x6d\x9a\x78\xcc\xae\x51\x16\xeb\xad\xe8\x15\xd9\x3b\x53\x92\xa3\x7f\x5e\xe9\xda\x3f\xa9\x8e\xf8\x57\xae\xf4\x31\xef\x3e\xb2\x8c\xab\xc7\x3b\x9e\xab\xfe\xab\x65\x2c\xb8\x4e\x7b\x17\xff\x9a\x13\xc1\x3d\xf1\xe6\x7d\xa6\xa6\xff\x8f\xfc\x02\xdd\x2f\x3c\x62\x92\xcc\x1b\x83\x27\x6d\x85\x16\xb6\xd4\xbe\xfa\x62\xeb\x2c\x9b\x8d\x3e\x8d\x57\x61\x93\x33\xf9\x5a\x23\x06\x1a\xa4\xa7\x27\xf7\xaa\x70\x3c\xcf\xf3\x79\xff\x54\x5a\x7a\x40\x81\x27\xac\x1d\xd5\x98\x8c\xbe\x12\x8e\x99\x94\x50\x66\x7f\xef\xd1\x2f\x12\xbc\x8d\x07\x0a\x87\xbf\x2c\x79\x47\x76\x71\xfc\x1c\x72\x06\xd5\xe8\xfa\xdc\xe4\xa6\x25\xb7\x8f\x6d\x69\x7c\x16\xdb\xa8\x30\x1f\xf7\x35\xf6\x95\x3f\xc8\x74\xf6\x9a\x4b\x4f\x26\xcc\xfb\xa1\x96\xe5\x85\xf1\x5e\x7b\xad\x13\x68\x05\x2f\x0c\xc7\x39\xcb\xc4\xa9\xb6\xd1\xaa\x11\x86\x74\x4d\x5f\xbe\xa7\x62\x9c\x0e\x1d\x2b\x68\xec\x6d\xd8\xc6\xe7\x0f\x0e\xe8\xbf\x39\x2b\x1d\x1e\x47\xb0\x6a\xed\xe1\xc1\xe0\xb5\xcd\x19\x49\x7d\x8d\xae\x26\x37\x96\xd6\x01\x2e\x31\xdb\x21\x19\x5f\x58\x5a\x47\xfe\x85\xa5\xe9\x9f\x2b\xd6\x9e\xe1\xb1\x16\xd4\x2f\xb9\xb7\x7d\xe2\xde\xb3\x5b\xf7\xdf\xec\xd2\x0d\x6d\x64\xde\x0e\xce\x5b\xc5\x1b\xb2\xfc\x13\x3b\x7c\x8a\x6c\x43\x83\xdb\x3b\xef\xa7\xd9\x56\xde\x4b\x2e\x09\x22\x96\xba\x90\xbb\x77\x7e\x79\xe7\x4d\x9e\xf6\x6b\x7e\x86\x94\x15\x92\x29\x58\x9a\x7b\xa1\x8f\x47\xb3\xcb\x4e\xd7\x95\xeb\x77\x2c\x46\xe5\x09\x39\x0c\x41\x7e\x67\x6a\x18\x5c\xae\x7f\xe1\x33\xc9\xf1\xf5\xc1\xc6\x29\x28\x58\x76\x1d\x13\x2f\xce\xb3\xf9\x29\xef\xfd\x7b\x8f\x6c\x73\xad\x21\xef\xf1\x8b\xe6\xda\x46\x94\x19\x31\x6e\x70\x8d\xfe\xb3\x0c\x9e\x62\x65\xdc\xfe\xae\x1b\xa1\xa4\xc3\x66\xd0\xc7\x93\xcd\xfc\xff\x5a\xfc\xff\xad\x09\x59\xff\x80\x4f\x1f\x91\xd1\xec\xe3\x27\xf9\x8a\x68\x7b\xed\x3d\xce\xf2\x72\x7a\x16\x3c\xa7\x68\xb1\x64\x3a\x73\xfa\xf8\xc4\xf5\x40\x1c\xfd\x95\x43\x03\x9e\xf3\xb0\xdb\xe9\x5e\x63\x70\xaa\x60\x3a\x23\x1c\x71\xdc\xc9\x20\x22\xeb\xec\x1a\x9f\x6d\x38\x7b\x2f\x88\x5a\x0f\x1b\xcc\x11\xd8\x9a\xe9\x47\x1b\x84\x4f\x9a\xba\x3e\x43\x5e\x3a\xb2\x96\x9f\xc8\x21\x8b\xd9\x3c\x8f\x93\xcd\x8f\xfd\x84\x2c\x7e\xe0\x28\x88\xe3\x66\xeb\x26\x3d\xcc\xdd\x72\x7f\x90\x5e\x3e\xad\xe6\x32\x4e\xd9\x12\x24\x01\x4c\x8e\xb3\x08\x7c\xfe\x5e\xc1\xfb\xfb\x84\xe2\x88\x58\xf5\xd7\x9b\x15\x59\x37\x2b\x22\x26\x5d\xe9\xaf\x11\xf1\x62\x33\xe7\x9b\x95\x83\x13\x1d\x6c\x04\x4e\x88\xbc\xd7\xd0\x6f\x4c\x61\x65\xea\xd1\x8f\x85\xc4\x8b\xb1\x3c\x77\xb8\xbd\xca\xe2\x0f\xcc\x81\x17\x73\x20\x0c\x76\x64\xc7\xbf\xd6\x03\x48\xf1\x72\xff\x85\x14\xb7\x51\xe1\x9e\x98\xac\x8f\x5a\x33\xbc\xf2\x83\x4f\xdd\xaa\x59\x1f\xfe\x84\x9b\xb3\x76\x5c\xbd\xe4\x4b\x7f\x14\xa9\xd4\x6b\xc7\x12\x69\x8b\x57\xae\x60\xe2\x29\x13\xc7\xb7\x90\xde\x35\xa8\x95\xf9\xda\xcb\x44\x6a\xf6\x73\x2f\x13\x19\x6c\xef\x65\x22\xc3\x03\x3f\x3f\xf7\x32\xe7\x8c\x6d\xbb\x67\x6a\xfe\xb5\xa6\xf3\x99\x1c\xbe\xef\xbb\xf1\x87\xbc\xf8\x4d\x9f\x11\x39\x94\xff\xbc\xcf\xb8\x4b\x7f\xe9\xe1\x21\xaf\x6f\x20\x24\xf8\x6f\xf7\xf2\x90\x2b\x3e\xf4\xf2\x1e\x19\x28\x72\xd3\x27\x06\x8a\xdc\xe9\x13\x03\x45\x01\xfe\x86\x81\xa2\x20\x7e\xc5\x40\x61\x67\xa0\x2f\xe4\xfa\x73\xb4\x8a\x5c\xef\x47\xf4\xef\x3d\xba\x5b\xb4\x3e\xdc\x54\x2d\x5d\x75\xe0\xd4\x3a\xf7\xbb\x64\xc3\x06\xc7\xdd\x03\x4a\x4a\xb7\x18\xe5\x11\xa3\x92\xc3\x8d\x25\x23\x4a\x81\xf3\xa9\xac\x38\x95\x32\xef\x8f\x4e\xf9\xc4\x40\x68\x65\x2b\xe9\xce\xad\xe8\xa8\xd2\x57\x8d\x3d\x3f\xab\xe1\xdb\x5a\x60\xf0\xac\xc9\x35\x94\xcb\xf9\x64\xf7\x3a\xd2\x3f\x71\xac\x1e\xb2\x04\x5d\x48\xf2\xf3\x09\x3f\x65\x1e\xa1\xdd\x9b\x47\xd7\xc2\xad\x4b\xd3\x03\xee\xf7\x8c\xfb\xad\xe0\x71\xc7\x56\x9e\xbc\x21\xf6\x02\x5b\x65\xa1\x5a\x40\xb7\xaa\x5b\x63\x2c\x5c\x1a\x6e\xc6\x94\x36\xeb\x2d\x26\xa2\xa4\x5e\x20\x50\xa3\x2a\xb1\xf0\x16\x29\x6e\x5a\x51\x1b\x62\xf5\xac\xbe\xea\xf6\xd1\x33\xa0\x57\x4e\xdb\xa3\xfc\x72\x23\xfc\x03\x63\x7f\xe8\x22\x7f\xb8\xb5\xf1\xcc\x8e\xe1\xa8\xf9\x17\x42\x1d\x19\x74\xfa\xda\xf5\xfe\x1c\xdf\xb2\xc3\xad\x77\xb6\x4e\x5e\xeb\x8a\x8b\x75\x6f\xb5\x71\xdf\x99\xc3\x94\x8c\x0b\xe2\xa9\x7f\x31\xea\xdf\xd3\xf7\x1b\xbd\x46\x19\x9e\xb3\x7f\xba\xf0\x41\x9d\xab\xfa\xb8\x67\x7a\xb7\xc4\x81\x2b\xbf\x46\xda\xb8\x85\xa6\x11\x17\x81\xe8\xb5\xca\x9b\x0c\x64\xbc\xf8\x72\x90\x79\x8f\x36\x6e\xa3\x0d\x97\x5e\x81\xd7\xfe\x20\x93\x3f\x5f\xa5\xa5\xdb\xad\xf9\xe4\x20\x36\xf9\xc9\xea\x67\x78\xe6\x56\x75\x44\x47\x1d\xf1\x18\x6c\x7e\xd7\xc1\x25\xc0\xf5\x93\x6f\xbe\x3d\x38\x4b\x9a\xb1\x10\xcc\xfc\xd4\xc6\x6f\xf5\xc4\xca\xd6\x37\x5b\xb4\x0f\x9f\x2e\x9e\x8b\x96\xed\x86\xd6\xab\x4f\x44\xa3\xf3\x36\x39\xd0\xea\xc1\xf9\x13\x89\xfd\x8f\xec\xa9\x88\x62\x8e\x39\x12\x83\x9a\xfb\xca\x46\xd9\x7a\xdc\x42\x8a\x05\x68\xeb\xb9\x57\x8c\x85\x3c\x30\x9a\x65\xc4\x5e\xa5\xf4\x2d\x77\x2b\xb1\xc6\xc2\x09\x38\xf6\x9e\x88\x42\xb0\xdc\xa2\x00\x50\xfa\x9f\x00\x00\x00\xff\xff\x05\x7d\xa5\x51\xb2\x2d\x00\x00")

func compiledEnergyBinRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_compiledEnergyBinRuntime,
		"compiled/Energy.bin-runtime",
	)
}

func compiledEnergyBinRuntime() (*asset, error) {
	bytes, err := compiledEnergyBinRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Energy.bin-runtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledParamsAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x92\xcf\x6a\xc3\x30\x0c\x87\xdf\x45\x67\x9f\x06\xdb\xc1\x6f\x30\xc6\x60\xb0\x63\x29\x43\x6d\xd4\x22\xe6\xca\x21\x92\xb3\x79\x25\xef\x3e\x5a\x9c\xa4\x0b\xed\x5a\xf6\x87\x1e\x13\xe9\xa7\x7c\xf9\xa4\xd9\x16\x96\x51\xd4\x50\x0c\xbc\x35\x89\x1c\xb0\xd4\xc9\x14\xfc\x6c\x0b\x82\x1b\x02\x0f\x2f\xaf\x94\xc1\x81\xe5\x7a\xf7\xa4\xd6\xb0\xac\xa1\x9b\xbb\xbe\xbe\x26\x03\x07\x31\xd9\x34\x38\x86\x12\x8b\xdd\xdc\xde\xed\x53\x35\x66\x5c\x04\x02\xbf\xc2\xa0\xe4\x40\x0d\x8d\x1e\x93\xe1\x82\x03\x5b\x06\x0f\x2d\xd3\xdb\x98\x5d\x25\x59\x1a\x47\x81\xce\x1d\xd2\x96\xf4\xa5\xb8\x6e\xac\xb7\x18\x12\x1d\x67\x2b\x2d\x3a\xf9\xa3\x8b\xa8\x25\x4a\xdf\xf4\x73\xf6\x36\xda\x0e\x77\x18\x80\x55\xd5\x90\xea\x17\xb8\xac\xf7\xc2\xc6\x18\xf8\x83\xae\x84\xf9\x07\x8a\xb3\x3e\x35\xf4\xdf\xa2\x27\x27\x3d\x7e\x7e\xf0\xfc\xed\xd9\x1e\xda\xff\xd5\xd9\xa2\x44\xc9\x9b\x98\xf4\x98\x54\x96\x8a\xde\xa9\xea\x69\x7b\x87\xe5\xf5\xc3\x09\xd5\x43\xaa\x0c\x3c\xb7\x9a\x53\xfd\xe7\x57\xf5\xbc\x5f\x52\xa9\x53\x4b\x62\xd0\xcd\x3f\x03\x00\x00\xff\xff\x48\x39\xff\xdd\x3a\x04\x00\x00")

func compiledParamsAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledParamsAbi,
		"compiled/Params.abi",
	)
}

func compiledParamsAbi() (*asset, error) {
	bytes, err := compiledParamsAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Params.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledParamsBinRuntime = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x6d\x16\x23\x29\x08\xbc\x52\x01\x82\x70\x1c\xb5\xf5\xfe\x47\xd8\x87\x9d\xce\xe7\xec\x6e\x66\xf6\xed\xbc\xe9\xfe\x11\x43\x14\xab\x0a\x8a\x18\xf2\x2d\x50\x36\xa0\x88\x11\x8c\x00\x1b\x5a\x4d\xd6\xed\xa9\x03\x84\x5f\x7c\x0c\x80\x28\x0a\x99\x89\x85\xcc\xe1\x3a\x9d\xa8\xe4\x2d\x95\xb4\x3a\x4c\xbc\x15\x9e\x7d\xc6\x19\x3d\xca\x19\xed\x12\x6c\xdd\xdb\x8e\x12\xc7\x19\x1d\xa3\xd7\xe8\x95\xcf\x68\xf1\x33\xba\xc6\xa4\x31\xc6\x6d\x6f\x74\xad\xda\xf3\x66\xc7\x3a\xb4\x4b\x21\xdd\xf7\x25\xab\x7b\x34\x23\x63\xb3\x36\x70\x71\x12\x75\xf2\x3c\x0d\x17\xe4\x37\x18\x18\x06\x5a\xce\x20\xa7\x00\x8a\x13\x18\xb4\xf5\xca\x7d\x74\x5b\x73\x9e\x54\x0e\x0e\x8a\x7d\xca\x0b\xc8\xc5\xc5\xe1\xec\x45\xaa\x22\x4a\x98\xa6\xb6\x34\x5a\xa8\xe2\xe9\xb5\xc4\x7a\x65\xcc\x3a\x30\xae\x3b\x1c\x41\x90\xc0\x92\x07\x8b\x63\xbd\xb1\x20\xae\xbf\x97\x85\x22\x44\x34\x64\xf3\x61\x89\xe0\x07\x13\xe0\x8e\x94\xa4\xfc\x10\x29\x19\x5a\xf6\x05\xa3\x41\x76\xd7\x29\x99\x11\xa4\x1f\x99\xe1\x3a\xad\xf2\x47\xf1\x2c\x18\xcf\x3c\x2f\x94\xcd\xde\x51\xb6\x99\xbb\xc3\x1f\x75\x7d\xe7\x1c\x08\x22\xfb\xa7\x5a\xdb\xe9\x1c\x72\xbe\x7e\x73\xf6\xcd\xe2\x3c\x11\x70\x24\xe2\xdc\xc9\x70\xd9\xae\xa5\x35\xd3\x0d\x4a\xce\x99\x99\x16\x45\x20\x18\x14\x94\xbb\x82\x7c\xaf\x41\xb9\xf7\x58\x27\xbe\x64\xbe\x73\xa6\xf5\xd3\xe5\x0d\xe2\xa0\x70\xd6\x8d\xd1\xcb\xfe\x84\x33\x55\x57\xe5\xe7\xce\x0d\x68\x62\xa1\x73\xfd\x77\x7c\x18\x5a\x72\x67\xaa\x7d\x75\x3b\xa0\x45\xe4\x53\x17\x27\xb2\xd4\x66\xfb\x98\xf5\xa5\x7f\x1c\x17\xda\xaf\x15\x61\xaf\x5f\x2b\xc2\xe6\xff\xb7\x22\x19\x09\xdc\x34\x51\xa7\x9f\x61\x32\xc7\xf7\x4c\xc6\xf1\x33\x4c\xa8\x66\xa6\x8b\x47\x56\x29\xf3\x64\xaf\x27\xbe\x93\xcb\x73\x35\xeb\x9a\x58\x36\x46\xa7\xe1\xed\x18\xad\x5a\x94\x36\xd0\xb8\x48\xad\x4a\xa2\x73\xe8\xc1\x55\x46\xaf\x2c\x1e\xf3\x98\x8b\x55\xdc\x6c\xba\x86\x0c\x1e\x9c\xec\xee\xdc\x6f\xac\x5d\x1e\x6a\x39\x3b\x41\x92\xa9\x97\xed\x55\x7a\xfc\x96\x78\x7e\xa4\xd6\xee\xa9\x9e\x59\x72\x74\xe4\x34\xa9\xdb\x0d\xa9\x8b\x66\xf6\xfc\xbc\xaa\x41\x10\xb3\xd4\xe8\xb9\x6e\xe7\x84\x88\x3d\x43\xb2\xbe\x67\x96\xd6\x6f\x59\xc4\x6f\x4e\xfe\xd4\x34\x8b\xf0\x8a\x30\x33\xe7\xdc\x38\xdf\x57\x3f\x34\xbe\xcf\x10\xbc\x3b\x80\xec\xd3\x15\xe9\x06\xd2\x13\xcd\x71\xbc\x4c\x9b\xbd\xd2\x52\xef\xff\xd4\xff\xf6\x24\xd2\x0f\xd7\x71\xd6\x3f\x7d\xb7\xe7\x52\x0d\xa8\xfe\x12\xba\xc2\xe3\x3f\xfa\xb5\xe8\xfa\xba\xcb\x4b\xc1\xef\xf5\xab\xbe\x4c\xaf\x4f\x1d\x9c\xd4\x80\x46\xa6\xc6\xb5\xd5\x56\x59\xa0\xce\x98\x47\x13\x35\xcb\x3f\xfd\xe2\xea\x9d\xe6\xb1\xbc\x55\x19\xae\x3a\x63\x16\x6e\x46\xd1\xc4\x7b\x8f\x3a\xfb\x24\x5b\x45\x8b\x93\x6a\x59\x2d\x8e\x45\x00\xc7\x5f\x01\x00\x00\xff\xff\xca\x9e\xb6\xf0\xa6\x09\x00\x00")

func compiledParamsBinRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_compiledParamsBinRuntime,
		"compiled/Params.bin-runtime",
	)
}

func compiledParamsBinRuntime() (*asset, error) {
	bytes, err := compiledParamsBinRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Params.bin-runtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"compiled/Authority.abi": compiledAuthorityAbi,
	"compiled/Authority.bin-runtime": compiledAuthorityBinRuntime,
	"compiled/Energy.abi": compiledEnergyAbi,
	"compiled/Energy.bin-runtime": compiledEnergyBinRuntime,
	"compiled/Params.abi": compiledParamsAbi,
	"compiled/Params.bin-runtime": compiledParamsBinRuntime,
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
	"compiled": &bintree{nil, map[string]*bintree{
		"Authority.abi": &bintree{compiledAuthorityAbi, map[string]*bintree{}},
		"Authority.bin-runtime": &bintree{compiledAuthorityBinRuntime, map[string]*bintree{}},
		"Energy.abi": &bintree{compiledEnergyAbi, map[string]*bintree{}},
		"Energy.bin-runtime": &bintree{compiledEnergyBinRuntime, map[string]*bintree{}},
		"Params.abi": &bintree{compiledParamsAbi, map[string]*bintree{}},
		"Params.bin-runtime": &bintree{compiledParamsBinRuntime, map[string]*bintree{}},
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

