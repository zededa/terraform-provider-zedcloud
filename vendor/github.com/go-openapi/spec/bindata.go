// Code generated by go-bindata. DO NOT EDIT.
// sources:
// schemas/jsonschema-draft-04.json (4.357kB)
// schemas/v2/schema.json (40.248kB)

package spec

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _jsonschemaDraft04Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x3d\x6f\xdb\x3c\x10\xde\xf3\x2b\x08\x26\x63\xf2\x2a\x2f\xd0\xc9\x5b\xd1\x2e\x01\x5a\x34\x43\x37\x23\x03\x6d\x9d\x6c\x06\x14\xa9\x50\x54\x60\xc3\xd0\x7f\x2f\x28\x4a\x14\x29\x91\x92\x2d\xa7\x8d\x97\x28\xbc\xaf\xe7\x8e\xf7\xc5\xd3\x0d\x42\x08\x61\x9a\xe2\x15\xc2\x7b\xa5\x8a\x55\x92\xbc\x96\x82\x3f\x94\xdb\x3d\xe4\xe4\x3f\x21\x77\x49\x2a\x49\xa6\x1e\x1e\xbf\x24\xe6\xec\x16\xdf\x1b\xa1\x3b\xf3\xff\x02\xc9\x14\xca\xad\xa4\x85\xa2\x82\x6b\xe9\x6f\x42\x02\x32\x2c\x28\x07\x45\x5a\x15\x3d\x77\x46\x39\xd5\xcc\x25\x5e\x21\x83\xb8\x21\x18\xb6\xaf\x52\x92\xa3\x47\x68\x88\xea\x58\x80\x56\x4e\x1a\xf2\xbd\x4f\xcc\x29\x7f\x52\x90\x6b\x7d\xff\x0f\x48\xb4\x3d\x3f\x21\x7c\x27\x21\xd3\x2a\x6e\x31\xaa\x2d\x53\xdd\xf3\xe3\x42\x94\x54\xd1\x77\x78\xe2\x0a\x76\x20\xe3\x20\x68\xcb\x30\x86\x41\xf3\x2a\xc7\x2b\xf4\x78\x8e\xfe\xef\x90\x91\x8a\xa9\xc7\xb1\x1d\xc2\xd8\x2f\x0d\x75\xed\xc1\x4e\x9c\xc8\x25\x43\xac\xa8\xbe\xd7\xcc\xa9\xd1\xa9\x21\xa0\x1a\xbd\x04\x61\x94\x34\x2f\x18\xfc\x3e\x16\x50\x8e\x4d\x03\x6f\x1c\x58\xdb\x48\x23\xbc\x11\x82\x01\xe1\xfa\xd3\x3a\x8e\x30\xaf\x18\x33\x7f\xf3\x8d\x39\x11\x9b\x57\xd8\x2a\xfd\x55\x2a\x49\xf9\x0e\xc7\xec\x37\xd4\x25\xf7\xec\x5c\x66\xc7\xd7\x99\xaa\xcf\x4f\x89\x8a\xd3\xb7\x0a\x3a\xaa\x92\x15\xf4\x30\x6f\x1c\xb0\xd6\x46\xe7\x98\x39\x2d\xa4\x28\x40\x2a\x3a\x88\x9e\x29\xba\x88\x37\x2d\xca\x60\x38\xfa\xba\x5b\x20\xac\xa8\x62\xb0\x4c\xd4\xaf\xda\x45\x0a\xba\x5c\x3b\xb9\xc7\x79\xc5\x14\x2d\x18\x34\x19\x1c\x51\xdb\x25\x4d\xb4\x7e\x06\x14\x38\x6c\x59\x55\xd2\x77\xf8\x69\x59\xfc\x7b\x73\xed\x93\x43\xcb\x32\x6d\x3c\x28\xdc\x1b\x9a\xd3\x62\xab\xc2\x27\xf7\x41\xc9\x08\x2b\x23\x08\xad\x13\x57\x21\x9c\xd3\x72\x0d\x42\x72\xf8\x01\x7c\xa7\xf6\x83\xce\x39\xd7\x82\x3c\x1f\x2f\xd6\x60\x1b\xa2\xdf\x35\x89\x52\x20\xe7\x73\x74\xe0\x66\x26\x64\x4e\xb4\x97\x58\xc2\x0e\x0e\xe1\x60\x92\x34\x6d\xa0\x10\xd6\xb5\x83\x61\x27\xe6\x47\xd3\x89\xbd\x63\xfd\x3b\x8d\x03\x3d\x6c\x42\x2d\x5b\x70\xee\xe8\xdf\x4b\xf4\x66\x4e\xe1\x01\x45\x17\x80\x74\xad\x4f\xc3\xf3\xae\xc6\x1d\xc6\xd7\xc2\xce\xc9\xe1\x29\x30\x86\x2f\x4a\xa6\x4b\x15\x84\x73\xc9\x6f\xfd\x7f\xa5\x6e\x9e\xbd\xf1\xb0\xd4\xdd\x45\x5a\xc2\x3e\x4b\x78\xab\xa8\x84\x74\x4a\x91\x3b\x92\x23\x05\xf2\x1c\x1e\x7b\xf3\x09\xf8\xcf\xab\x24\xb6\x60\xa2\xe8\x4c\x9f\x75\x77\xaa\x8c\xe6\x01\x45\x36\x86\xcf\xc3\x63\x3a\xea\xd4\x8d\x7e\x06\xac\x14\x0a\xe0\x29\xf0\xed\x07\x22\x1a\x65\xda\x44\xae\xa2\x73\x1a\xe6\x90\x69\xa2\x8c\x46\xb2\x2f\xde\x49\x38\x08\xed\xfe\xfd\x41\xaf\x9f\xa9\x55\xd7\xdd\x22\x8d\xfa\x45\x63\xc5\x0f\x80\xf3\xb4\x08\xd6\x79\x30\x9e\x93\xee\x59\xa6\xd0\x4b\xee\x22\xe3\x33\xc1\x3a\x27\x68\x36\x78\x7e\x87\x0a\x06\xd5\x2e\x20\xd3\xaf\x15\xfb\xd8\x3b\x73\x14\xbb\x92\xed\x05\x5d\x2e\x29\x38\x2c\x94\xe4\x42\x45\x5e\xd3\xb5\x7d\xdf\x47\xca\x38\xb4\x5c\xaf\xfb\x7d\xdd\x6d\xf4\xa1\x2d\x77\xdd\x2f\xce\x6d\xc4\x7b\x8b\x4e\x67\xa9\x6f\xfe\x04\x00\x00\xff\xff\xb1\xd1\x27\x78\x05\x11\x00\x00")

func jsonschemaDraft04JsonBytes() ([]byte, error) {
	return bindataRead(
		_jsonschemaDraft04Json,
		"jsonschema-draft-04.json",
	)
}

func jsonschemaDraft04Json() (*asset, error) {
	bytes, err := jsonschemaDraft04JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jsonschema-draft-04.json", size: 4357, mode: os.FileMode(0640), modTime: time.Unix(1568963823, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe1, 0x48, 0x9d, 0xb, 0x47, 0x55, 0xf0, 0x27, 0x93, 0x30, 0x25, 0x91, 0xd3, 0xfc, 0xb8, 0xf0, 0x7b, 0x68, 0x93, 0xa8, 0x2a, 0x94, 0xf2, 0x48, 0x95, 0xf8, 0xe4, 0xed, 0xf1, 0x1b, 0x82, 0xe2}}
	return a, nil
}

var _v2SchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x4f\x93\xdb\x36\xb2\xbf\xfb\x53\xa0\x14\x57\xd9\xae\xd8\x92\xe3\xf7\x2e\xcf\x97\xd4\xbc\xd8\x49\x66\x37\x5e\x4f\x79\x26\xbb\x87\x78\x5c\x05\x91\x2d\x09\x09\x09\x30\x00\x38\x33\x5a\xef\x7c\xf7\x2d\xf0\x9f\x08\x02\x20\x41\x8a\xd2\xc8\x0e\x0f\xa9\x78\x28\xa0\xd1\xdd\x68\x34\x7e\xdd\xf8\xf7\xf9\x11\x42\x33\x49\x64\x04\xb3\xd7\x68\x76\x86\xfe\x76\xf9\xfe\x1f\xe8\x32\xd8\x40\x8c\xd1\x8a\x71\x74\x79\x8b\xd7\x6b\xe0\xe8\xd5\xfc\x25\x3a\xbb\x38\x9f\xcf\x9e\xab\x0a\x24\x54\xa5\x37\x52\x26\xaf\x17\x0b\x91\x17\x99\x13\xb6\xb8\x79\xb5\x10\x59\xdd\xf9\xef\x82\xd1\x6f\xf2\xc2\x8f\xf3\x4f\xb5\x1a\xea\xc7\x17\x45\x41\xc6\xd7\x8b\x90\xe3\x95\x7c\xf1\xf2\x7f\x8b\xca\x45\x3d\xb9\x4d\x32\xa6\xd8\xf2\x77\x08\x64\xfe\x8d\xc3\x9f\x29\xe1\xa0\x9a\xff\xed\x11\x42\x08\xcd\x8a\xd6\xb3\x9f\x15\x67\x74\xc5\xca\x7f\x27\x58\x6e\xc4\xec\x11\x42\xd7\x59\x5d\x1c\x86\x44\x12\x46\x71\x74\xc1\x59\x02\x5c\x12\x10\xb3\xd7\x68\x85\x23\x01\x59\x81\x04\x4b\x09\x9c\x6a\xbf\x7e\xce\x49\x7d\xba\x7b\x51\xfd\xa1\x44\xe2\xb0\x52\xac\x7d\xb3\x08\x61\x45\x68\x46\x56\x2c\x6e\x80\x86\x8c\xbf\xbd\x93\x40\x05\x61\x74\x96\x95\xbe\x7f\x84\xd0\x7d\x4e\xde\x42\xb7\xe4\xbe\x46\xbb\x14\x5b\x48\x4e\xe8\xba\x90\x05\xa1\x19\xd0\x34\xae\xc4\xce\xbe\xbc\x9a\xbf\x9c\x15\x7f\x5d\x57\xc5\x42\x10\x01\x27\x89\xe2\x48\x51\xb9\xda\x40\xd5\x87\x37\xc0\x15\x5f\x88\xad\x90\xdc\x10\x81\x42\x16\xa4\x31\x50\x39\x2f\x38\xad\xab\xb0\x53\xd8\xac\x94\x56\x6f\xc3\x84\xf4\x11\xa4\x50\xb3\xfa\xe9\xd3\x6f\x9f\x3e\xdf\x2f\xd0\xeb\x8f\x1f\x3f\x7e\xbc\xfe\xf6\xe9\xf7\xaf\x5f\x7f\xfc\x18\x7e\xfb\xec\xfb\xc7\xb3\x36\x79\x54\x43\xe8\x29\xc5\x31\x20\xc6\x11\x49\x9e\xe5\x12\x41\x66\xa0\xe8\xed\x1d\x8e\x93\x08\x5e\xa3\x27\x3b\xc3\x7c\xa2\x73\xba\xc4\x02\x2e\xb0\xdc\xf4\xe5\x76\xd1\xca\x96\xa2\x8a\x94\xcd\x21\xc9\x6c\xec\x2c\x70\x42\x9e\x34\x74\x9d\x19\x7c\xcd\x20\x9c\xea\x2e\x0a\xfe\x42\x84\xd4\x29\x04\x8c\x8a\xb4\x41\xa2\xc1\xdc\x19\x8a\x88\x90\x4a\x49\xef\xce\xdf\xbd\x45\x4a\x52\x81\x70\x10\x40\x22\x21\x44\xcb\x6d\xc5\xec\x4e\x3c\x1c\x45\xef\x57\x9a\xb5\x7d\xae\xfe\xe5\xe4\x31\x86\x90\xe0\xab\x6d\x02\x3b\x2e\xcb\x11\x90\xd9\xa8\xc6\x77\xc2\x59\x98\x06\xfd\xf9\x2e\x78\x45\x01\xa6\xa8\xa0\x71\x5c\xbe\x33\xa7\xd2\xd9\x5f\x95\xef\xd9\xd5\xac\xfd\xdc\x5d\xbf\x5e\xb8\xd1\x3e\xc7\x31\x48\xe0\x5e\x4c\x14\x65\xdf\xb8\xa8\x71\x10\x09\xa3\xc2\xc7\x02\xcb\xa2\x4e\x5a\x02\x82\x94\x13\xb9\xf5\x30\xe6\xb2\xa4\xb5\xfe\x9b\x3e\x7a\xb2\x55\xd2\xa8\x4a\xbc\x16\xb6\x71\x8e\x39\xc7\xdb\x9d\xe1\x10\x09\x71\xbd\x9c\xb3\x41\x89\xd7\xa5\x89\xdc\x57\xb5\x53\x4a\xfe\x4c\xe1\xbc\xa0\x21\x79\x0a\x1a\x0f\x70\xa7\x5c\x08\x8e\xde\xb0\xc0\x43\x24\xad\x74\x63\x0e\xb1\xd9\x90\xe1\xb0\x2d\x13\xa7\x6d\x78\xfd\x04\x14\x38\x8e\x90\xaa\xce\x63\xac\x3e\x23\xbc\x64\xa9\xb4\xf8\x03\x63\xde\xcd\xbe\x16\x13\x4a\x55\xac\x82\x12\xc6\xac\xd4\x35\xf7\x22\xd4\x3a\xff\x22\x73\x0e\x6e\x51\xa0\x75\x1e\xae\x8f\xe8\x5d\xc7\x59\xe6\xe4\x9a\x18\x8d\xd6\x1c\x53\x84\x4d\xb7\x67\x28\x37\x09\x84\x69\x88\x12\x0e\x01\x11\x80\x32\xa2\xf5\xb9\xaa\xc6\xd9\x73\x53\xab\xfb\xb4\x2e\x20\xc6\x54\x92\xa0\x9a\xf3\x69\x1a\x2f\x81\x77\x37\xae\x53\x1a\xce\x40\xc4\xa8\x82\x1c\xb5\xef\xda\x24\x7d\xb9\x61\x69\x14\xa2\x25\xa0\x90\xac\x56\xc0\x81\x4a\xb4\xe2\x2c\xce\x4a\x64\x7a\x9a\x23\xf4\x13\x91\x3f\xa7\x4b\xf4\x63\x84\x6f\x18\x87\x10\xbd\xc3\xfc\x8f\x90\xdd\x52\x44\x04\xc2\x51\xc4\x6e\x21\x74\x48\x21\x81\xc7\xe2\xfd\xea\x12\xf8\x0d\x09\xf6\xe9\x47\x35\xaf\x67\xc4\x14\xf7\x22\x27\x97\xe1\xe2\x76\x2d\x06\x8c\x4a\x1c\x48\x3f\x73\x2d\x0b\x5b\x29\x45\x24\x00\x2a\x0c\x11\xec\x94\xca\xc2\xa6\xc1\x37\x21\x43\x83\x3b\x5f\x97\xf1\x43\x5e\x53\x73\x19\xa5\x36\xd8\x2d\x05\x2e\x34\x0b\xeb\x39\xfc\x1d\x63\x51\x01\xbd\x3d\xbb\x90\x84\x40\x25\x59\x6d\x09\x5d\xa3\x1c\x37\xe6\x5c\x16\x9a\x40\x09\x70\xc1\xe8\x82\xf1\x35\xa6\xe4\xdf\x99\x5c\x8e\x9e\x4d\x79\xb4\x27\x2f\xbf\x7e\xf8\x05\x25\x8c\x50\xa9\x98\x29\x90\x62\x60\xea\x75\xae\x13\xca\xbf\x2b\x1a\x29\x27\x76\xd6\x20\xc6\x64\x5f\xe6\x32\x1a\x08\x87\x21\x07\x21\xbc\xb4\xe4\xe0\x32\x67\xa6\xcd\xf3\x1e\xcd\xd9\x6b\xb6\x6f\x8e\x27\xa7\xed\xdb\xe7\xbc\xcc\x1a\x07\xce\x6f\x87\x33\xf0\xba\x51\x17\x22\x66\x78\x79\x8e\xce\xe5\x13\x81\x80\x06\x2c\xe5\x78\x0d\xa1\xb2\xb8\x54\xa8\x79\x09\xbd\xbf\x3c\x47\x01\x8b\x13\x2c\xc9\x32\xaa\xaa\x1d\xd5\xee\xab\x36\xbd\x6c\xfd\x54\x6c\xc8\x08\x01\x3c\xbd\xe7\x07\x88\xb0\x24\x37\x79\x90\x28\x4a\x1d\x10\x1a\x92\x1b\x12\xa6\x38\x42\x40\xc3\x4c\x43\x62\x8e\xae\x36\xb0\x45\x71\x2a\xa4\x9a\x23\x79\x59\xb1\xa8\xf2\xa4\x0c\x60\x9f\xcc\x8d\x40\xf5\x80\xca\xa8\x99\xc3\xa7\x85\x1f\x31\x25\xa9\x82\xc5\x6d\xbd\xd8\x36\x76\x7c\x02\x28\x97\xf6\x1d\x74\x3b\x11\x7e\x91\xae\x32\xf8\x6c\xf4\xe6\x7b\x9a\xa5\x1f\x62\xc6\x21\xcf\x9a\xe5\xed\x8b\x02\xf3\x2c\x33\x33\xdf\x00\xca\xc9\x09\xb4\x04\xf5\xa5\x08\xd7\xc3\x02\x18\x66\xf1\xab\x1e\x83\x37\x4c\xcd\x12\xc1\x1d\x50\xf6\xaa\xbd\xfe\xe2\x73\x48\x38\x08\xa0\x32\x9b\x18\x44\x86\x0b\x6a\xc1\xaa\x26\x96\x2d\x96\x3c\xa0\x54\x65\x73\xe3\x08\xb5\x8b\x99\xbd\x82\xbc\x9e\xc2\xe8\x53\x46\x83\x3f\x33\x54\x2b\x5b\xad\x92\x79\xd9\x8f\x5d\x93\x98\xf2\xe6\xc6\x1c\xe6\x9a\x9e\xfc\x43\x82\x31\x66\x8e\x53\x77\xfe\x90\xe7\xf3\xf6\xe9\x62\x23\x3f\x10\x93\x18\xae\x72\x1a\x9d\xf9\x48\xcb\xcc\x5a\x65\xc7\x4a\x04\xf0\xf3\xd5\xd5\x05\x8a\x41\x08\xbc\x86\x86\x43\x51\x6c\xe0\x46\x57\xf6\x44\x40\x0d\xfb\xff\xa2\xc3\x7c\x3d\x39\x84\xdc\x09\x22\x64\x4f\x12\xd9\xba\xaa\xf6\xe3\xbd\x56\xdd\x91\x25\x6a\x14\x9c\x89\x34\x8e\x31\xdf\xee\x15\x7e\x2f\x39\x81\x15\x2a\x28\x95\x66\x51\xf5\xfd\x83\xc5\xfe\x15\x07\xcf\xf7\x08\xee\x1d\x8e\xb6\xc5\x52\xcc\x8c\x5a\x93\x66\xc5\xd8\x79\x38\x46\xd6\xa7\x88\x37\xc9\x2e\xe3\xd2\xa5\x7b\x4b\x3a\xdc\xa1\xdc\x9e\x29\xf1\x8c\x8a\x99\x16\x47\x8d\xd4\x78\x8b\xf6\x1c\xe9\x71\x54\x1b\x69\xa8\x4a\x93\x37\xe5\xb2\x2c\x4f\x0c\x92\xab\xa0\x73\x32\x72\x59\xd3\xf0\x2d\x8d\xed\xca\x37\x16\x19\x9e\xdb\x1c\xab\x17\x49\xc3\x0f\x37\xdc\x88\xb1\xb4\xd4\x42\xcb\x58\x5e\x6a\x52\x0b\x15\x10\x0a\xb0\x04\xe7\xf8\x58\x32\x16\x01\xa6\xcd\x01\xb2\xc2\x69\x24\x35\x38\x6f\x30\x6a\xae\x1b\xb4\x71\xaa\xad\x1d\xa0\xd6\x20\x2d\x8b\x3c\xc6\x82\x62\x27\x34\x6d\x15\x84\x7b\x43\xb1\x35\x78\xa6\x24\x77\x28\xc1\x6e\xfc\xe9\x48\x74\xf4\x15\xe3\xe1\x84\x42\x88\x40\x7a\x26\x49\x3b\x48\xb1\xa4\x19\x8e\x0c\xa7\xb5\x01\x6c\x0c\x97\x61\x8a\xc2\x32\xd8\x8c\x44\x69\x24\xbf\x65\x1d\x74\xd6\xe5\x44\xef\xec\x48\x5e\xb7\x8a\xa3\x29\x8e\x41\x64\xce\x1f\x88\xdc\x00\x47\x4b\x40\x98\x6e\xd1\x0d\x8e\x48\x98\x63\x5c\x21\xb1\x4c\x05\x0a\x58\x98\xc5\x6d\x4f\x0a\x77\x53\x4f\x8b\xc4\x44\x1f\xb2\xdf\x8d\x3b\xea\x9f\xfe\xf6\xf2\xc5\xff\x5d\x7f\xfe\x9f\xfb\x67\x8f\xff\xf3\xe9\x69\xd1\xfe\xb3\xc7\xfd\x3c\xf8\x3f\x71\x94\x82\x23\xd1\x72\x00\xb7\x42\x99\x6c\xc0\x60\x7b\x0f\x79\xea\xa8\x53\x4b\x56\x31\xfa\x0b\x52\x9f\x96\xdb\xcd\x2f\xd7\x67\xcd\x04\x19\x85\xfe\xdb\x02\x9a\x59\x03\xad\x63\x3c\xea\xff\x2e\x18\xfd\x00\xd9\xe2\x56\x60\x59\x93\xb9\xb6\xb2\x3e\x3c\x2c\xab\x0f\xa7\xb2\x89\x43\xc7\xf6\xd5\xce\x2e\xad\xa6\xa9\xed\xa6\xc6\x5a\xb4\xa6\x67\xdf\x8c\x26\x7b\x50\x5a\x91\x08\x2e\x6d\xd4\x3a\xc1\x9d\xf2\xdb\xde\x1e\xb2\x2c\x6c\xa5\x64\xc9\x16\xb4\x90\xaa\x4a\xb7\x0c\xde\x13\xc3\x2a\x9a\x11\x9b\x7a\x1b\x3d\x95\x97\x37\x31\x6b\x69\x7e\x34\xc0\x67\x1f\x66\x19\x49\xef\xf1\x25\xf5\xac\x0e\xea\x0a\x28\x8d\x4d\x7e\xd9\x57\x4b\x49\xe5\xc6\xb3\x25\xfd\xe6\x57\x42\x25\xac\xcd\xcf\x36\x74\x8e\xca\x24\x47\xe7\x80\xa8\x92\x72\xbd\x3d\x84\x2d\x65\xe2\x82\x1a\x9c\xc4\x44\x92\x1b\x10\x79\x8a\xc4\x4a\x2f\x60\x51\x04\x81\xaa\xf0\xa3\x95\x27\xd7\x12\x7b\xa3\x96\x03\x45\x96\xc1\x8a\x07\xc9\xb2\xb0\x95\x52\x8c\xef\x48\x9c\xc6\x7e\x94\xca\xc2\x0e\x07\x12\x44\xa9\x20\x37\xf0\xae\x0f\x49\xa3\x96\x9d\x4b\x42\x7b\x70\x59\x14\xee\xe0\xb2\x0f\x49\xa3\x96\x4b\x97\xbf\x00\x5d\x4b\x4f\xfc\xbb\x2b\xee\x92\xb9\x17\xb5\xaa\xb8\x0b\x97\x17\x9b\x43\xfd\xd6\xc2\xb2\xc2\x2e\x29\xcf\xfd\x87\x4a\x55\xda\x25\x63\x1f\x5a\x65\x69\x2b\x2d\x3d\x67\xe9\x41\xae\x5e\xc1\x6e\x2b\xd4\xdb\x3e\xa8\xd3\x26\xd2\x48\x92\x24\xca\x61\x86\x8f\x8c\xbb\xf2\x8e\x91\xdf\x1f\x06\x19\x33\xf3\x03\x4d\xba\xcd\xe2\x2d\xfb\x69\xe9\x16\x15\x13\xd5\x56\x85\x4e\x3c\x5b\x8a\xbf\x25\x72\x83\xee\x5e\x20\x22\xf2\xc8\xaa\x7b\xdb\x8e\xe4\x29\x58\xca\x38\xb7\x3f\x2e\x59\xb8\xbd\xa8\x16\x16\xf7\xdb\x79\x51\x9f\x5a\xb4\x8d\x87\x3a\x6e\xbc\x3e\xc5\xb4\xcd\x58\xf9\xf5\x3c\xb9\x6f\x49\xaf\x57\xc1\xfa\x1c\x5d\x6d\x88\x8a\x8b\xd3\x28\xcc\xb7\xef\x10\x8a\x4a\x74\xa9\x4a\xa7\x62\xbf\x0d\x76\x23\x6f\x59\xd9\x31\xee\x40\x11\xfb\x28\xec\x8d\x22\x1c\x13\x5a\x64\x94\x23\x16\x60\xbb\xd2\x7c\xa0\x98\xb2\xe5\x6e\xbc\x54\x33\xe0\x3e\xb9\x52\x17\xdb\xb7\x1b\xc8\x12\x20\x8c\x23\xca\x64\x7e\x78\xa3\x62\x5b\x75\x56\xd9\x9e\x2a\x91\x27\xb0\x70\x34\x1f\x90\x89\xb5\x86\x73\x7e\x71\xda\x1e\xfb\x3a\x72\xdc\x5e\x79\x88\xcb\x74\x79\xd9\x64\xe4\xd4\xc2\x9e\xce\xb1\xfe\x85\x5a\xc0\xe9\x0c\x34\x3d\xd0\x43\xce\xa1\x36\x39\xd5\xa1\x4e\xf5\xf8\xb1\xa9\x23\x08\x75\x84\xac\x53\x6c\x3a\xc5\xa6\x53\x6c\x3a\xc5\xa6\x7f\xc5\xd8\xf4\x51\xfd\xff\x25\x4e\xfa\x33\x05\xbe\x9d\x60\xd2\x04\x93\x6a\x5f\x33\x9b\x98\x50\xd2\xe1\x50\x52\xc6\xcc\xdb\x38\x91\xdb\xe6\xaa\xa2\x8f\xa1\x6a\xa6\xd4\xc6\x56\xd6\x8c\x40\x02\x68\x48\xe8\x1a\xe1\x9a\xd9\x2e\xb7\x05\xc3\x34\xda\x2a\xbb\xcd\x12\x36\x98\x22\x50\x4c\xa1\x1b\xc5\xd5\x84\xf0\xbe\x24\x84\xf7\x2f\x22\x37\xef\x94\xd7\x9f\xa0\xde\x04\xf5\x26\xa8\x37\x41\x3d\x64\x40\x3d\xe5\xf2\xde\x60\x89\x27\xb4\x37\xa1\xbd\xda\xd7\xd2\x2c\x26\xc0\x37\x01\x3e\x1b\xef\x5f\x06\xe0\x6b\x7c\x5c\x91\x08\x26\x10\x38\x81\xc0\x09\x04\x76\x4a\x3d\x81\xc0\xbf\x12\x08\x4c\xb0\xdc\x7c\x99\x00\xd0\x75\x70\xb4\xf8\x5a\x7c\xea\xde\x3e\x39\x08\x30\x5a\x27\x35\xed\xb4\x65\xad\x69\x74\x10\x88\x79\xe2\x30\x52\x19\xd6\x04\x21\xa7\x95\xd5\x0e\x03\xf8\xda\x20\xd7\x84\xb4\x26\xa4\x35\x21\xad\x09\x69\x21\x03\x69\x51\x46\xff\xff\x18\x9b\x54\xed\x87\x47\x06\x9d\x4e\x73\x6e\x9a\xb3\xa9\xce\x83\x5e\x4b\xc6\x71\x20\x45\xd7\x72\xf5\x40\x72\x0e\x34\x6c\xf4\x6c\xf3\xba\x5e\x4b\x97\x0e\x52\xb8\xbe\x8b\x79\xa0\x10\x86\xa1\x75\xb0\x6f\xec\xc8\xf4\x3d\x4d\x7b\x86\xc2\x02\x31\x12\x51\xbf\x07\x94\xad\x10\xd6\x2e\x79\xcf\xe9\x1c\xf5\x1e\x31\x23\x5c\x18\xfb\x9c\xfb\x70\xe0\x62\xbd\xf7\xb5\x94\xcf\xf3\xf6\xfa\xc5\x4e\x9c\x85\x76\x1d\xae\x37\xbc\xde\xa3\x41\xcb\x29\xd0\x5e\x70\x67\x50\x93\x6d\x98\xa8\xd3\x67\x0f\x68\xb1\xeb\x38\x47\x07\x10\x1b\xd2\xe2\x18\x68\x6d\x40\xbb\xa3\x40\xba\x21\xf2\x8e\x81\xfb\xf6\x92\x77\x2f\x70\xe8\xdb\xb2\x36\xbf\x30\x91\xc5\x21\xe7\x45\xcc\x34\x0c\x48\x8e\xd0\xf2\x9b\x7c\x3c\xbd\x1c\x04\x3e\x07\xe8\x7c\x2f\x84\x7a\x48\x4d\x1f\xba\xe1\x76\x45\x7b\x60\xe0\x01\xca\xee\x04\xca\x31\xbe\x73\x5f\xa3\x70\x0c\xad\x1f\xa5\xf5\x76\xd5\xbb\xd2\x7e\xfb\x30\x90\xcf\xfa\x67\x7a\xe6\xc3\x37\x42\x19\xe2\xc9\x9c\x61\x4c\xe7\xd1\x77\x55\x86\x6e\x8f\x7b\x85\x42\x33\xa3\xaa\x57\xae\xfd\xd5\xcc\x9c\x56\x68\xe2\xde\x0e\xa8\x2c\xa9\xb0\x7d\xf0\x54\x2d\x80\xf2\x48\x39\x3d\x98\x1a\x6d\x0b\x9d\xba\x53\xfb\xce\xf8\xd1\x7e\xbb\x60\x4f\x06\xf5\xce\xda\xab\xeb\xca\xcb\xd5\xac\x20\xda\x72\x3b\xa2\x4b\x38\xd7\xb5\x89\xbe\x42\xd9\xb9\x73\xc4\x0c\x6d\xb7\xd9\xf8\x8d\xbd\x3e\x9c\xf5\x53\x68\x48\x14\x36\x8f\x09\xc5\x92\xf1\x21\xd1\x09\x07\x1c\xbe\xa7\x91\xf3\x6a\xc8\xc1\x57\xb0\xdd\xc5\xc6\x1d\xad\x76\x1d\xa8\x82\x0e\x4c\x38\xfe\xa5\x8c\xc5\x0a\x40\x5d\xa1\xbb\x98\xd1\xfb\x74\x61\xed\x1a\x98\xaf\x3c\x8c\x1e\xe3\xc2\x92\x29\x74\x3e\x99\xd0\xf9\x41\x50\xd0\x38\x4b\x57\x7e\x5b\x7a\x0e\xe6\xce\x4e\xd7\x19\x35\x57\xbb\x3c\x3c\xd2\x5e\x4f\x4b\x4c\xf7\x0f\x4d\x2b\x91\x5d\x94\xa6\x95\xc8\x69\x25\x72\x5a\x89\x7c\xb8\x95\xc8\x07\x80\x8c\xda\x9c\x64\x7b\xb7\x71\xdf\x57\x12\x4b\x9a\x1f\x72\x0c\x13\x03\xad\x3c\xd5\x4e\xde\x8e\x57\x13\x6d\x34\x86\xcf\x97\xe6\xa4\x68\xc4\xb0\xf6\xc9\xc2\xeb\x8d\x0b\xd7\xcd\xfe\xba\xa6\xf5\x30\xeb\x30\x33\xbe\xc7\x56\x27\xab\x08\xd9\x6d\xbb\x09\xee\x7c\x2d\xcf\xee\x87\x38\xac\xc8\xdd\x90\x9a\x58\x4a\x4e\x96\xa9\x79\x79\xf3\xde\x20\xf0\x96\xe3\x24\x19\xeb\xba\xf2\x53\x19\xab\x12\xaf\x47\xb3\xa0\x3e\xef\x9b\x8d\x6d\x6d\x7b\xde\x3b\x3b\x1a\xc0\x3f\x95\x7e\xed\x78\xfb\x76\xb8\xaf\xb3\xdd\xc5\xeb\x95\xed\x5a\x62\x41\x82\xb3\x54\x6e\x80\x4a\x92\x6f\x36\xbd\x34\xae\xde\x6f\xa4\xc0\xbc\x08\xe3\x84\xfc\x1d\xb6\xe3\xd0\x62\x38\x95\x9b\x57\xe7\x71\x12\x91\x80\xc8\x31\x69\x5e\x60\x21\x6e\x19\x0f\xc7\xa4\x79\x96\x28\x3e\x47\x54\x65\x41\x36\x08\x40\x88\x1f\x58\x08\x56\xaa\xd5\xbf\xaf\xad\x96\xd7\xd6\xcf\x87\xf5\x34\x0f\x71\x93\x6e\x26\xed\x98\x5b\x9f\x4f\xcf\x95\x34\xc6\xd7\x11\xfa\xb0\x81\x22\x1a\xdb\xdf\x8e\xdc\xc3\xb9\xf8\xdd\x5d\x3c\x74\xe6\xea\xb7\x8b\xbf\xf5\x6e\xb3\x46\x2e\x64\xf4\xab\x3c\x4e\xcf\x36\x1d\xfe\xfa\xb8\x36\xba\x8a\xd8\xad\xf6\xc6\x41\x2a\x37\x8c\x17\x0f\xda\xfe\xda\xe7\x65\xbc\x71\x2c\x36\x57\x8a\x47\x12\x4c\xf1\xbd\x77\x6b\xa4\x50\x7e\x77\x7b\x22\x60\x89\xef\xcd\xf5\xb9\x0c\x97\x79\x0d\x2b\x35\x43\xcb\x3d\x24\xf1\x78\xfc\xf8\xcb\x1f\x15\x06\xe2\x78\xd8\x51\x21\xd9\x1f\xf0\xf5\x8f\x86\xa4\x50\xfa\xb1\x47\x43\xa5\xdd\x69\x14\xe8\xa3\xc0\x86\x91\xa7\x81\x50\xb4\x7c\xc0\x81\x80\x77\x7a\x9f\xc6\xc2\xa9\x8c\x05\x33\xb0\x3b\x31\xa4\xf4\xd7\x1b\x26\x55\x97\x7c\x65\xf8\x69\x1a\x84\x8e\x41\x78\xd9\xec\xc5\x11\x16\x1e\x74\x91\xf5\x56\xf5\x57\x49\x47\x5c\x92\xa9\x1e\x99\x36\xf4\xdb\xb1\x0e\xd3\x78\x02\xb0\x9b\x25\xcb\xe9\xe9\x1d\x0d\x44\x01\x42\x08\x91\x64\xd9\xdd\x37\x08\x17\xef\xf9\xe5\x0f\xbd\x46\x91\xf5\xf9\x89\x92\x37\xdd\x89\x59\x44\x1f\x9c\xee\x34\x1e\xbe\x47\x83\x32\x72\x8e\x37\xdf\xac\x69\x38\xef\x75\xb0\xda\xdb\xac\x83\x94\x2f\x39\xa6\x62\x05\x1c\x25\x9c\x49\x16\xb0\xa8\x3c\xc7\x7e\x76\x71\x3e\x6f\xb5\x24\xe7\xe8\xb7\xb9\xc7\x6c\x43\x92\xee\x21\xd4\x17\xa1\x7f\xba\x35\xfe\xae\x39\xbc\xde\xba\x69\xd9\x8e\xe1\x62\xde\x64\x7d\x16\x88\x1b\xed\x29\x11\xfd\x4f\xa9\xff\x99\x90\xc4\xf6\xf4\xf9\x6e\xe9\x28\x23\xd7\xca\xe5\xee\xee\x9f\x63\xb1\x5b\xfb\x10\xd7\x2f\x1d\xf2\xe3\xbf\xb9\xb5\x6f\xa4\x6d\x7d\x25\x79\xfb\x24\x31\xea\x56\xbe\x5d\x53\xcd\x2d\x36\xa3\x6d\xdf\xab\x1c\xb8\x6d\x6f\xc0\x98\xa7\xdd\xaa\x86\x8c\x1d\x39\xa3\x9d\x70\x2b\x9b\x68\xd9\xfd\x33\xfe\xa9\xb6\x4a\x2e\x63\x0f\xcf\x68\x27\xd9\x4c\xb9\x46\x6d\xcb\xbe\xa1\xa8\xd6\x5f\xc6\xd6\x9f\xf1\x4f\xf4\xd4\xb4\x78\xd0\xd6\xf4\x13\x3c\x3b\xac\xd0\xdc\x90\x34\xda\xc9\xb4\x9a\x1a\x8d\xbd\x93\x87\xd4\xe2\x21\x1b\xb3\x2b\xd1\xbe\xe7\x69\xd4\x53\x67\xd5\x40\xa0\xe3\x19\x3f\x6d\x1a\xbc\x0e\x86\x3c\x10\xb4\x3d\x2a\xcd\x78\x32\xe6\xab\xbd\x36\xc9\xf4\x3a\x58\xae\xc3\xf4\x47\xea\xbf\xfb\x47\xff\x0d\x00\x00\xff\xff\xd2\x32\x5a\x28\x38\x9d\x00\x00")

func v2SchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_v2SchemaJson,
		"v2/schema.json",
	)
}

func v2SchemaJson() (*asset, error) {
	bytes, err := v2SchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v2/schema.json", size: 40248, mode: os.FileMode(0640), modTime: time.Unix(1568964748, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xab, 0x88, 0x5e, 0xf, 0xbf, 0x17, 0x74, 0x0, 0xb2, 0x5a, 0x7f, 0xbc, 0x58, 0xcd, 0xc, 0x25, 0x73, 0xd5, 0x29, 0x1c, 0x7a, 0xd0, 0xce, 0x79, 0xd4, 0x89, 0x31, 0x27, 0x90, 0xf2, 0xff, 0xe6}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"jsonschema-draft-04.json": jsonschemaDraft04Json,

	"v2/schema.json": v2SchemaJson,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"jsonschema-draft-04.json": {jsonschemaDraft04Json, map[string]*bintree{}},
	"v2": {nil, map[string]*bintree{
		"schema.json": {v2SchemaJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
