// Code generated by go-bindata.
// sources:
// ../../resource/test_certs/ca.crt
// ../../resource/test_certs/ca.key
// ../../resource/test_certs/node.crt
// ../../resource/test_certs/node.key
// DO NOT EDIT!

package securitytest

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
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

var _test_certsCaCrt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x92\x4b\x73\x9b\x3e\x14\x47\xf7\x7c\x8a\xff\xde\xf3\x9f\xa8\xd8\xad\x61\xd1\xc5\xd5\x03\x90\x89\x34\x11\x60\x13\x65\x07\x76\x10\x0f\x3f\x62\x70\x2c\xe2\x4f\x5f\xec\x45\xa7\x9d\xde\xe5\x59\xfc\xe6\xcc\x99\xfb\xff\xfd\x30\x0b\xb9\xfc\x8f\xb0\x24\xe3\x01\x27\x90\xb1\x07\x75\x04\xe7\xf8\x9d\x12\x02\x69\x6a\xc0\x72\x0c\x86\x2b\xd9\x55\x9b\xb1\x63\x55\x1d\x3d\xbf\x17\x68\xb6\x2a\xab\x65\x74\xa4\xf0\x8c\x4d\x77\xae\xbb\x26\xf4\x2d\xc2\xa0\x06\x1b\xdf\x58\xea\x08\x0c\x21\x7c\x5b\x33\x52\x8b\x95\x72\xfd\xb6\x98\xaf\xae\x3a\x97\x27\x91\xac\x2d\xb3\x9a\x6e\x94\xa2\xcc\x8e\xb4\x74\x65\xbf\x3d\xf8\xb5\x76\x8d\x51\x88\xd9\xa8\xde\x4a\x91\xad\xad\x23\x5b\x58\x88\xcc\x20\x49\x95\xcd\x1f\x50\xdb\x89\xcd\x7f\xb3\x16\x7a\x91\x70\xcb\xe0\x31\x16\x33\xbb\xff\x73\x6c\x74\x82\x0c\x32\x6c\xe4\x06\x83\xc8\x28\x93\x57\xed\x5e\xbe\x4a\x37\x68\x0b\x82\xa9\xca\xf0\x56\x00\x0a\x49\x7a\x0e\x53\x5e\xce\xa9\x62\x93\xfa\x1a\x00\x0d\x20\x98\x21\x4a\x11\xcf\x39\xcf\xca\xc1\x3d\x86\x2f\x33\x8d\x9b\x03\x2a\xfc\xe8\x74\x19\xcf\xd7\x83\xcf\xf5\x17\x1d\xcb\xdd\x65\x6f\xb1\xee\xcb\xb7\xbd\xa1\x88\x45\xfa\x13\x1d\xfa\x9d\xa9\x8a\x62\xd9\xd0\x3d\x39\xc4\x59\xe8\x8c\x30\xcb\xdf\xc7\xe3\x76\xa4\xed\x79\x28\x87\x8f\x20\x78\xab\xc0\x4c\x65\x20\x6c\x57\x2d\x74\x02\x16\xf7\x46\x3b\x6a\x19\x7e\xb2\x8a\x4d\xa1\xe1\x3b\x85\xf4\x6e\x1d\x25\x02\x3b\x50\x79\x8c\xd0\x29\xa4\x8a\x9e\xc0\x30\x22\x60\xf8\x4b\x99\x58\x39\x8d\xc9\xe4\xf3\xe8\x56\x97\xa5\x3e\xed\x4a\x76\xf3\xf3\xb1\x7b\xdd\x91\x38\x7c\xc9\x13\xa7\xf6\x3a\xf7\x99\xef\xf1\x38\x17\xd7\x79\x93\xe2\x45\xa4\x92\xfe\xd8\x10\xab\xed\x2d\x8e\x84\xd7\x7a\x3f\xf2\x3e\x8a\x25\x9a\xe5\x1f\x4b\xaf\x06\xf4\xba\xb2\xd5\xba\xd9\x5c\x82\xc5\x4f\xe7\xf1\x09\x4c\xd2\x7f\xbf\xe3\x57\x00\x00\x00\xff\xff\xcb\x20\xf7\xeb\x3a\x02\x00\x00")

func test_certsCaCrtBytes() ([]byte, error) {
	return bindataRead(
		_test_certsCaCrt,
		"test_certs/ca.crt",
	)
}

func test_certsCaCrt() (*asset, error) {
	bytes, err := test_certsCaCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/ca.crt", size: 570, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsCaKey = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd1\xcb\x96\x63\x40\x00\x80\xe1\xbd\xa7\xe8\xbd\x33\xc7\x25\xa4\x59\xcc\xa2\x8a\x62\x4a\xa4\x94\x12\xd7\x1d\xe2\x1e\x89\x16\xa2\xe5\xe9\xe7\x4c\xaf\xe7\xdf\xfe\xcb\xef\xd7\xbf\x20\xb2\x31\xf9\x60\x01\xf8\xa0\x0c\x47\xe0\x82\x3e\x4e\x28\xfd\x39\xdc\x19\x63\xe8\xf9\x18\x02\xe0\x40\xe0\xee\xb3\xba\xbe\xf3\x32\x15\xd4\xc6\x3e\x15\x6c\x16\xab\xee\x60\x7f\xf1\x85\xd8\xf6\x98\x5a\x8b\x2c\x27\xc0\xea\x96\x65\x8c\x01\xf5\xfd\x6b\xf7\x19\xe4\xdc\x22\x43\x7d\x9a\x57\xec\x85\xa7\x74\x3a\x17\xc8\x54\xb3\x4f\xfb\xfa\x46\x9e\xf7\xf5\xbd\xee\x43\x18\xdd\x34\x03\x6c\x08\x00\xdf\x01\x9e\x62\x23\x5a\x51\xba\x26\x65\x26\x51\xa1\x89\x8b\xd1\xe3\xda\x24\xb2\xaa\x5d\xaf\x1f\xb0\x96\x78\x7b\x27\x5b\x3c\xa5\xcd\xac\xe6\xe2\x81\x54\xd1\xc1\x2c\x74\x26\x1b\x97\xab\x7f\xad\xe3\x71\x14\xcc\x6e\x4d\xf3\xe3\x30\x46\x42\x99\x14\xa5\x06\x5d\x3e\xe4\x1e\x3e\x6e\x81\x67\x91\x7c\x7e\x0e\xf5\x74\x91\xd9\x1e\xb7\xc0\xa9\x87\xed\xd1\xeb\xc5\xc4\x0c\xc2\x63\xd8\xbf\x4b\xdc\x98\x86\x9e\x86\xae\x03\x3a\x04\xa4\xd1\x74\x21\x5e\x27\xbf\xee\x77\x0e\x66\x2b\xb9\x16\x34\x4b\x14\x6d\xa9\x6c\xf1\xbb\x95\x54\x9a\x14\x19\x1c\x1a\x30\xf2\x7c\x69\x60\x43\x1f\x9e\x47\x73\x0d\x6e\x53\x52\xc3\xd7\xc0\xac\xf8\x95\xc8\xc1\x6d\xdd\xb4\x46\x5d\x98\xca\x69\x0a\x1c\xff\xc8\xa8\x8d\xc2\x8c\x81\x0e\xc6\x0e\x1d\xe8\xac\x2c\x93\xdb\x6b\xf9\x5b\xc3\x24\x09\xdc\x29\x18\x8d\x33\x59\xf6\x9e\xe8\xab\x77\x2f\x2c\x6b\xa4\x0d\x8b\x37\xdc\x54\x34\x45\x9c\x2f\x3c\x7d\x69\x31\x98\x78\xbc\x4b\x0f\xa5\x57\x24\x45\x15\x9d\x53\xe8\x93\x78\x3f\x6c\x4c\x7c\x1d\x85\x39\x38\xde\xc5\xdf\xdc\x0f\x1a\x22\xe6\xff\x31\xff\x06\x00\x00\xff\xff\x6b\x78\xcb\x8b\xed\x01\x00\x00")

func test_certsCaKeyBytes() ([]byte, error) {
	return bindataRead(
		_test_certsCaKey,
		"test_certs/ca.key",
	)
}

func test_certsCaKey() (*asset, error) {
	bytes, err := test_certsCaKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/ca.key", size: 493, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsNodeCrt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x96\x4d\x57\xe2\x3c\x18\x86\xf7\xf9\x15\xef\x9e\xf3\x1e\x4a\x41\x84\xc5\x2c\xf2\xd5\x52\x24\x4f\x08\x14\xb1\xee\x00\xb5\x58\x40\x70\x18\x49\xe9\xaf\x9f\xf2\x31\x02\xe6\x61\xa3\xe7\x3e\x79\x4a\xaf\xbb\xbd\x48\xfe\x3f\x7c\x98\x0c\x23\xf8\x8f\xcb\x41\x1c\x05\x11\xa7\xb1\x3c\xa6\x44\x45\x51\x67\x19\x73\xce\xbc\x30\xa5\x36\x62\x34\x8d\x06\x34\x5a\xf5\xec\x62\x57\x4d\x05\xff\xf0\x2a\xab\x50\xb6\x9a\x55\x36\xf6\x2c\xb7\x49\xf7\x61\xfd\x1c\xcd\x77\x33\xa0\x46\xf6\x14\xdf\xe6\x44\x66\xd4\xb0\x14\x1e\x19\x5d\xc7\x7c\x04\xbb\xc4\xff\xb3\x9f\xfa\x41\x36\x11\xf2\x51\x31\x15\xd2\xda\x48\xd2\x5c\x29\xe3\xb7\xb3\x49\xbd\xbb\x4b\xc6\xb0\x8e\x24\x30\xc5\x1a\x4f\x22\x96\x35\xa2\x44\x62\xb5\x90\x0d\x10\xc6\x83\x60\x7d\x08\xfd\x43\x06\xc5\x77\x66\x7b\xb1\x1c\x2a\x46\x8f\x17\xe3\x73\xd5\xbd\xbe\x18\x51\x83\x99\x0d\x4c\x22\x1e\x8d\x11\xd2\xde\x89\xa9\x0f\xbf\x67\xab\xf6\x3c\xf1\xd3\x34\x5e\xb5\x17\xcf\x31\x9b\x29\xea\x85\x7c\xf8\x19\x0e\xa3\x69\x5d\x18\xc9\xa8\x19\x51\xea\x6d\xa9\x92\x29\x27\xc6\xf0\xcd\xfd\x73\x33\xba\x1f\x2c\xe9\xa3\x1f\x2e\x6a\xa3\x6e\x4e\x47\xf5\x90\x4f\x5e\x57\x9b\x76\x7f\xe7\xc7\x4d\x2f\x58\xbd\x54\x26\xad\x36\x1f\x73\xb3\x59\x56\x1a\xed\xb5\x32\x0f\xbd\x8a\x4e\x93\x20\x5b\x91\xa5\x7d\xe3\x95\x71\xb1\xb8\x17\xad\x51\xff\x43\xdf\x2d\x17\x33\xe6\xd7\x1b\x5d\x9a\x96\x37\x4d\xc3\x2c\x4d\x93\x4a\x59\x74\xa8\x33\xaa\x0f\x5d\x75\x4c\x8b\xd1\xb7\x56\x79\x1b\x8a\xd3\x07\x6a\x49\xe7\x70\xfb\x03\x6f\xc9\x58\x62\x03\x9a\x44\x0f\x36\x61\xcc\x8c\x3a\xd4\xca\x90\xf3\x6d\x58\xde\x6e\xc0\xac\xe2\x8a\xb6\x0e\x1d\xbc\x48\x2b\x59\xd5\x9a\x40\x1d\x2e\x60\x22\x4b\xd2\xf4\xc9\x3f\x5e\x79\x20\x65\xf9\xff\x9f\xf2\xdb\x82\x66\x12\x75\xa7\x61\x3b\x4b\xc6\xf9\x7a\x5a\x07\x2f\x4d\xf3\xf3\x93\xa1\x5f\xa7\x7c\xcf\x45\xe7\x5f\x89\xc3\xbb\x6d\x59\xdc\x7c\x1a\x45\xea\x5c\xde\xbe\xb7\xca\xcb\xc7\x19\x6c\x2f\x83\xca\x19\x04\x7e\x9a\x23\x57\x83\x35\x77\x30\x71\x07\xf7\xb7\x5f\x48\xca\xc1\x86\x3b\xb8\xb8\x0c\x3e\x9d\x06\x55\x4c\x91\x4c\x7e\x11\x37\x8c\x90\x85\x0a\xc9\xcc\x31\x23\xb7\xe1\x08\x59\x98\x20\xd9\xec\x3b\x23\x97\x30\x45\x16\x22\x28\xd9\x2d\x0a\x39\x85\x12\x59\x88\xa0\x64\x2e\x0a\x51\x99\x41\x16\x22\x28\x19\x82\x92\xcd\x90\x12\x33\x04\x25\x43\x50\x0a\x8a\x94\x58\x20\x28\x05\x82\x52\x28\xa4\xc4\x02\x41\x29\x10\x94\x22\x41\x4a\x2c\x66\xc8\x42\x04\xa5\x70\x51\x08\x08\xf7\x0d\x03\xe1\xa2\x80\x70\x51\xa0\x74\xc4\x29\xb1\xfc\x19\x43\x16\xba\x28\x70\xf6\x84\xdc\x86\x2e\x0a\x08\x17\x05\xae\x5c\xf9\x2e\x11\x10\x59\x20\x46\x50\x7e\xb8\x42\x4e\xa1\xfb\x86\x41\x8c\xa0\x20\xae\x10\x40\x64\x81\x18\x41\x41\x5c\x81\xd2\x15\xb7\xc4\x0c\x41\x41\x5c\x81\xb3\x2b\xb7\x25\x22\xb2\x00\xe2\x0a\x5c\xb9\x72\x29\x11\x91\x05\x32\x04\xe5\x87\x2b\xa7\x12\x11\x59\xa0\x40\x50\x10\x57\x08\x20\xb2\x40\x81\xa0\x20\xae\x40\xe9\x8a\x5b\x62\x81\xa0\x20\xae\xc0\xd9\x95\xdb\x12\x11\x59\x34\xe2\x8a\xbe\x72\x85\x5c\x42\x17\x45\x0b\x17\x45\xff\x70\x85\x9c\x42\xf7\x0d\xd3\xc2\x45\xd1\x88\x2b\x44\x23\xb2\x68\x64\x63\xd1\x88\x2b\x1a\xdb\x58\x34\xb2\xb1\x68\xc4\x15\x8d\x6d\x2c\x1a\x91\x45\x23\xae\x68\x6c\x63\xd1\x88\x2c\xfa\xb2\xb1\x74\xca\xd3\x46\xf9\xe1\x1d\xd6\x69\x95\x7f\xe5\xf7\x89\x8d\xfc\x3b\xb2\x51\x4f\xd2\x8f\xcf\x45\xf5\x6e\x3c\x9f\xd5\x35\xed\x8d\x7f\xaf\x65\x6b\x95\x6f\xbe\x5e\xd3\x56\x97\x7d\xf9\xd3\xb8\xfa\xfe\x2a\x9e\x52\x10\x9f\xac\xf7\xbe\x8f\x27\x9e\xf2\x5f\xf2\x7e\x9f\xf4\x96\x1d\x6f\x6c\x36\xfc\xe3\x2d\x1c\x6e\xb6\x73\xbf\x51\xed\x37\x6c\x4d\x0e\x3f\x37\xef\xdb\xe7\x9c\xf9\x4d\xfb\xeb\x17\x39\x1e\x29\x25\x08\xf7\x98\xf9\x37\x00\x00\xff\xff\xba\xf6\x7b\x9b\x83\x0a\x00\x00")

func test_certsNodeCrtBytes() ([]byte, error) {
	return bindataRead(
		_test_certsNodeCrt,
		"test_certs/node.crt",
	)
}

func test_certsNodeCrt() (*asset, error) {
	bytes, err := test_certsNodeCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/node.crt", size: 2691, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsNodeKey = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd1\xbb\xb2\xa2\x30\x00\xc6\xf1\x9e\xa7\xb0\x77\x76\xc8\x01\x24\x5a\x26\x10\x20\x51\x6e\x0a\x18\xec\x02\xca\x4d\x45\x17\x14\xa3\x4f\xbf\xb3\xd6\xe7\x6b\xbf\xe6\x3f\xf3\xfb\xf3\x7f\x98\xb8\x34\x98\x6d\x77\x68\x16\x6d\x69\x86\x12\x32\x5b\x93\xfc\xfb\x28\x3e\xa5\x38\xac\x29\x46\x88\x61\xb4\xee\x1f\xfd\xad\x7b\xb8\x29\xe6\xb9\x48\xb2\xb8\x27\x38\x29\x73\x76\x37\x45\x0f\xe6\xea\x21\x1a\xe2\xfd\x41\x5f\x0c\x1f\xc0\x72\x66\x5d\x39\xec\xb4\x5a\x91\xe8\x36\x19\x26\xae\xf7\xa1\xe0\x78\xb9\x59\x14\x61\x32\x46\x32\x9e\x97\xf0\xba\xdf\xbe\xbc\xa2\xaa\xcf\x16\x7a\x11\x84\x62\x86\x42\xf9\x93\x41\x6a\x7a\x5e\x46\xe9\x34\x19\xc1\xb5\xc5\x86\xa2\x86\xd0\x30\xc7\x9b\x7c\x1e\x26\x61\x6a\x3f\x83\x7e\xa8\x1d\xf9\x5e\x4e\xdc\xd6\x97\x03\x72\xde\x0b\xb3\x10\xd1\x3a\x7f\x5e\x5a\xa3\x39\x81\xb3\xaf\x16\x16\x4d\x85\xb8\x67\x16\x9f\xcf\x2b\xd8\x2a\x6a\x4c\x1b\x14\x1c\x1b\x1e\x96\xbd\xe6\xac\x4f\x97\x8c\xe9\xce\x91\xf6\x89\x61\x6b\x62\xa1\x06\xb9\xf5\x21\xb2\x8c\x37\xe0\x74\x87\xbe\xce\x06\xd4\x12\xf4\xae\xfa\x48\x5f\x81\x76\x33\x96\x99\x62\x56\x70\x65\xc5\x6a\x12\x0c\x81\xd6\xbb\xc6\xad\xe3\x92\xef\x26\x24\xad\xbb\x7b\x5c\x5d\x46\x8b\x22\xe9\xa6\xc9\xab\xda\x75\x3c\xf1\x4f\xce\x79\x5b\x8f\x9c\x78\x5e\xeb\x66\xb7\x8a\x82\x92\x29\xc6\xea\x08\xcb\xa8\xfe\x10\xd0\x7a\xa8\xc5\x61\xf8\x78\x49\x3f\xd5\x52\xba\x69\x80\x5f\x9c\x50\xfa\xea\x52\x4f\x42\x66\x12\x7b\xf9\xb1\xb7\x9d\xe8\xd2\xb4\x7d\x32\x9f\x7c\xcb\x45\xac\x18\x16\x78\x5c\x1b\xcd\xc6\xfd\x05\x3c\x56\x45\xe0\xf0\xc5\x5f\x36\x4f\xc6\x65\x63\x88\x01\xc4\x70\xbd\x6b\xb3\x0e\x0e\x1b\xe5\x8b\x46\x02\xfb\x77\xcc\x7f\x01\x00\x00\xff\xff\xd6\x94\xd6\xdb\xed\x01\x00\x00")

func test_certsNodeKeyBytes() ([]byte, error) {
	return bindataRead(
		_test_certsNodeKey,
		"test_certs/node.key",
	)
}

func test_certsNodeKey() (*asset, error) {
	bytes, err := test_certsNodeKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/node.key", size: 493, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
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
	"test_certs/ca.crt":   test_certsCaCrt,
	"test_certs/ca.key":   test_certsCaKey,
	"test_certs/node.crt": test_certsNodeCrt,
	"test_certs/node.key": test_certsNodeKey,
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
	"test_certs": {nil, map[string]*bintree{
		"ca.crt":   {test_certsCaCrt, map[string]*bintree{}},
		"ca.key":   {test_certsCaKey, map[string]*bintree{}},
		"node.crt": {test_certsNodeCrt, map[string]*bintree{}},
		"node.key": {test_certsNodeKey, map[string]*bintree{}},
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
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
		err = RestoreAssets(dir, path.Join(name, child))
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
