// Code generated by go-bindata.
// sources:
// assets/aws-node-1.10.yaml
// assets/aws-node.yaml
// assets/coredns.yaml
// DO NOT EDIT!

package defaultaddons

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

var _awsNode110Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdf\x6f\xdb\x36\x10\x7e\xd7\x5f\x71\x70\x1f\x0a\x0c\x95\x6c\x67\x59\x96\xe9\xcd\x75\xbc\xac\x58\xe2\x1a\xf1\x9a\x62\x28\x02\x83\xa6\x2e\x32\x67\x8a\xe4\xc8\xa3\x1d\xf7\xaf\x1f\x28\xf9\x87\x14\xcb\xee\x36\x0c\x18\x5f\x2c\xf3\xee\xbe\xe3\x7d\xfc\xc8\x63\x1c\xc7\x11\x33\xe2\x11\xad\x13\x5a\xa5\x60\xe7\x8c\x27\xcc\xd3\x42\x5b\xf1\x95\x91\xd0\x2a\x59\x5e\xbb\x44\xe8\xee\xaa\x1f\xbd\x81\xa5\x9f\xa3\x55\x48\xe8\x60\x55\x85\x38\x98\xe3\xb3\xb6\x08\xfd\xe4\x3a\xe9\x81\x5b\x68\x2f\x33\xf0\x0e\xcf\x42\xcd\x91\x58\x3f\x5a\x0a\x95\xa5\x30\x94\xde\x11\xda\x07\x2d\x31\x2a\x90\x58\xc6\x88\xa5\x11\x80\x62\x05\xa6\xc0\xd6\x2e\x56\x3a\xc3\xc8\x7a\x89\x2e\x8d\x62\x60\x46\xdc\x5a\xed\x8d\x0b\x4e\x31\x70\x9b\x95\xb8\xac\x60\x5f\xb5\x62\x6b\x97\x70\x5d\x44\x00\x16\x9d\xf6\x96\xe3\xd6\xad\xf3\x5d\xa7\xfc\x0d\xa8\xce\x30\x8e\x2e\x82\x50\xc3\xbc\x66\xaf\x63\xc3\x97\x4e\xe7\xe9\x18\xc6\xe8\xcc\x55\x38\x3a\x43\x77\x0a\x11\xbe\x74\xa4\x70\xd4\x79\x07\x9d\x35\x23\xbe\x08\x1f\x39\x52\xe7\xe9\x75\x0a\x7c\x21\x54\x25\x8d\x6d\xc9\x32\x86\x85\x56\x0e\xe9\x0c\xf2\x53\xf4\x7a\x0b\x57\x3b\x62\xa7\x68\x57\x82\xe3\x80\x73\xed\x15\x9d\xe3\x16\x0e\x45\xa4\xe5\x1e\xc7\x6e\xe3\x08\x8b\x23\xec\xff\x57\x1e\xef\x85\xca\x84\xca\xcf\xaa\x44\x4b\x7c\xc0\xe7\x60\xd9\x11\x7d\x66\xd5\x11\xc0\xb1\x06\x8f\x30\x9d\x9f\xff\x81\x9c\x4a\xf1\xb5\x32\xfb\xcf\xf8\xac\x20\x6e\xca\xbd\x9d\x22\x35\xf8\x65\xc6\xb8\xbf\x41\xe5\x4f\x4d\x2a\x0f\x2a\xda\x73\xf7\x6f\x36\x1b\x40\xb2\x39\xca\x52\x7c\x00\xcb\x6b\x17\x33\x63\xea\x3c\x18\xe4\xc1\xe6\x4d\xc6\x08\xa7\x64\x19\x61\xbe\xa9\xbc\x69\x63\x30\x85\x07\x2d\xa5\x50\xf9\xa7\xd2\x21\x02\x70\x28\x91\x93\xb6\x95\x4f\x11\x04\x7b\x57\x4b\xd1\x96\x04\x80\xb0\x30\x92\x11\x6e\x83\x6a\x85\x84\x21\x1b\xf1\xed\x08\x61\x30\xa5\x34\x95\x7b\x5d\x73\x76\x7c\x81\x99\x97\x68\x13\x26\xcd\x82\x25\x07\x92\x83\xee\xb8\x15\x24\x38\x93\xb1\xd1\x59\x0a\x6f\xdf\x96\x61\xbb\xa2\xcb\xef\xc6\xb6\x8f\x5f\xd3\x1a\xc6\x42\x3b\x1a\x23\xad\xb5\x5d\xa6\x40\xd6\xef\xe6\x49\x4b\xb4\xcd\xe5\xc4\xa0\x4d\x98\xd3\x36\x85\xd1\x8b\x70\xe5\x29\x0f\x83\x6b\x45\x4c\x28\xb4\x35\x57\x51\xb0\x1c\x53\xb8\xea\x5d\x5c\xf6\xfa\xfd\xcb\xef\x2f\x7f\xb8\x48\xb2\xa5\x4d\x90\xdb\xc4\xbb\x78\x8d\x8e\xe2\x8b\xe6\x1d\xd8\xad\xfe\xc5\x81\x21\xae\x44\xba\xea\x27\x97\x49\x6f\xcf\x45\x89\x38\xf1\x52\x4e\xb4\x14\x7c\x93\xc2\x40\xae\xd9\xc6\xed\xed\x46\x5b\xaa\x51\x17\x1f\x96\x35\xd1\x96\x52\xb8\xea\x5f\xfd\x78\xbd\x37\xef\x54\x56\x20\x59\xc1\x0f\x28\x47\xda\xab\x06\xaa\x55\x5a\x8b\x8d\xb7\x7e\x83\xcf\xd3\xd9\xe3\x64\x38\xfb\xf5\x7a\x3a\x1b\x8e\x3f\xcc\xee\x3e\xde\xde\x8d\x1e\x47\x77\x35\x57\x80\x15\x93\x1e\x53\xb8\x19\xbd\xff\x74\xdb\x82\x71\xff\xfb\x6c\xfc\xf1\x66\x34\x1b\x0f\xee\x47\xc7\x71\x3f\x5b\x5d\xa4\x8d\x69\x80\x67\x81\x32\xdb\x5e\x1a\x2d\x96\x09\xa3\x45\x5a\xea\x20\x09\x35\x84\x6d\x6f\x49\xfb\x79\xf0\xdb\xf0\x97\x32\xe9\x74\x32\x18\xfe\x97\x99\x77\x27\x20\xd9\x1f\xdb\xbd\x77\xa3\x5f\x1c\x26\xff\xf4\xe8\xc8\x35\x41\xb9\xf1\x29\xf4\x7b\xc5\xe1\x2c\x20\xf7\x56\xd0\x66\xa8\x15\xe1\x0b\xd5\xbd\x8d\x15\x2b\x21\x31\xc7\xac\xa1\x61\x80\x95\x96\xbe\xc0\xfb\xa0\xfe\x86\x34\x8a\x30\x53\xad\xb6\x1b\x4e\x40\x57\x1b\xea\x72\x25\xba\x73\xa1\x8e\x24\xc2\x95\x88\xe7\x42\xc5\x99\xb0\xe7\x20\x90\x78\x09\xa1\x90\x92\xac\x15\x44\x21\x7d\x0b\x64\xc5\x6c\x57\xea\xfc\x28\x5c\xea\xfc\x4c\x68\x88\xb2\x5e\x75\x33\xcd\x97\x68\x13\xa7\xf9\xf2\x08\xa1\xb2\xd5\x4c\x15\x37\xb5\x23\x7b\xba\xda\xb0\xb4\x32\x55\x9d\xf3\x2a\xf5\x31\x71\xf1\x99\x8a\xcf\x00\xb5\xd1\x17\x9f\xa8\xfe\x0c\x4c\x93\xc0\xf8\x54\xf1\xdf\xc4\x78\x4d\xe7\xeb\x87\x05\x33\xe2\xd0\xc5\x4e\x3c\x04\xbc\x23\x5d\x3c\x6c\x25\x7f\x83\xcf\x42\x89\x70\xa1\xb6\xf4\x3a\x54\x82\x6b\xf5\x2c\x72\x97\xb4\x3f\x0f\x77\xb7\xba\xe3\x3a\xf4\xad\x6d\xfb\x8f\x00\xf2\xea\xc5\x70\xea\x51\xf9\x06\xfa\x49\xbf\x17\x9a\xae\x83\xce\xb6\x2f\x77\xde\x01\x93\x12\x14\xae\xd1\x96\xed\x78\x67\x70\x9d\xea\xd9\xb6\x7b\x96\x95\x3d\xa7\xbf\xeb\xbf\x15\x51\x46\x7a\xcb\x64\x7d\xc5\x55\xd7\x11\x2a\xf7\x92\xd9\x9a\xa1\x6a\xca\x25\x13\xa3\xf1\x87\x61\x35\xf7\x57\x00\x00\x00\xff\xff\x41\x46\x08\xc0\xbf\x0b\x00\x00")

func awsNode110YamlBytes() ([]byte, error) {
	return bindataRead(
		_awsNode110Yaml,
		"aws-node-1.10.yaml",
	)
}

func awsNode110Yaml() (*asset, error) {
	bytes, err := awsNode110YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-node-1.10.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _awsNodeYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdf\x6f\xdb\x36\x10\x7e\xd7\x5f\x71\xf0\x1e\x0a\x0c\x93\x1c\x77\x59\x97\xe9\xcd\x75\xbc\xae\x58\xe2\x1a\xf1\x9a\x62\x28\x02\x83\xa6\x2e\x32\x67\x8a\xe4\xc8\xa3\x1c\xf7\xaf\x1f\x28\xf9\x87\x64\xcb\xee\x36\x0c\x18\x5f\x2c\xf3\xee\xbe\xe3\x7d\xfc\xc8\x63\x1c\xc7\x11\x33\xe2\x11\xad\x13\x5a\xa5\x60\x17\x8c\x27\xcc\xd3\x52\x5b\xf1\x85\x91\xd0\x2a\x59\xdd\xb8\x44\xe8\x7e\x39\x88\xbe\x81\x95\x5f\xa0\x55\x48\xe8\xa0\xac\x43\x1c\x2c\xf0\x59\x5b\x84\x41\x72\x93\x5c\x81\x5b\x6a\x2f\x33\xf0\x0e\x2f\x42\x2d\x90\xd8\x20\x5a\x09\x95\xa5\x30\x92\xde\x11\xda\x07\x2d\x31\x2a\x90\x58\xc6\x88\xa5\x11\x80\x62\x05\xa6\xc0\xd6\x2e\x56\x3a\xc3\xc8\x7a\x89\x2e\x8d\x62\x60\x46\xbc\xb3\xda\x1b\x17\x9c\x62\xe0\x36\xab\x70\x59\xc1\xbe\x68\xc5\xd6\x2e\xe1\xba\x88\x00\x2c\x3a\xed\x2d\xc7\xad\x5b\xef\xdb\x5e\xf5\x1b\x50\x9d\x61\x1c\x5d\x04\xa1\x86\x45\xc3\xde\xc4\x86\xcf\xbd\xde\xd3\x29\x8c\xd1\x99\xab\x71\x74\x86\xee\x1c\x22\x7c\xee\x49\xe1\xa8\xf7\x1d\xf4\xd6\x8c\xf8\x32\x7c\xe4\x48\xbd\xa7\xe3\x14\xf8\x42\xa8\x2a\x1a\xbb\x92\x65\x0c\x0b\xad\x1c\xd2\x05\xe4\xa7\xe8\x78\x0b\xcb\x1d\xb1\x33\xb4\xa5\xe0\x38\xe4\x5c\x7b\x45\x97\xb8\x85\x43\x11\x69\xb5\xc7\xb1\xdb\x38\xc2\xe2\x04\xfb\xff\x95\xc7\x5b\xa1\x32\xa1\xf2\x8b\x2a\xd1\x12\x1f\xf0\x39\x58\x76\x44\x5f\x58\x75\x04\x70\xaa\xc1\x13\x4c\xe7\x17\x7f\x20\xa7\x4a\x7c\x9d\xcc\xfe\x33\x3e\x6b\x88\xdb\x6a\x6f\x67\x48\x2d\x7e\x99\x31\xee\x6f\x50\xf9\x53\x9b\xca\x83\x8a\xf6\xdc\xfd\x9b\xcd\x06\x90\x6c\x81\xb2\x12\x1f\xc0\xea\xc6\xc5\xcc\x98\x26\x0f\x06\x79\xb0\x79\x93\x31\xc2\x19\x59\x46\x98\x6f\x6a\x6f\xda\x18\x4c\xe1\x41\x4b\x29\x54\xfe\xb1\x72\x88\x00\x1c\x4a\xe4\xa4\x6d\xed\x53\x04\xc1\xde\x35\x52\x74\x25\x01\x20\x2c\x8c\x64\x84\xdb\xa0\x46\x21\x61\xc8\x56\x7c\x37\x42\x18\x4c\x29\x4d\xd5\x5e\x37\x9c\x1d\x5f\x62\xe6\x25\xda\x84\x49\xb3\x64\xc9\x81\xe4\xa0\x3b\x6e\x05\x09\xce\x64\x6c\x74\x96\xc2\xab\x57\x55\xd8\xae\xe8\xea\xbb\xb5\xed\x93\x63\x5a\xc3\x58\x6a\x47\x13\xa4\xb5\xb6\xab\x14\xc8\xfa\xdd\x3c\x69\x89\xb6\xbd\x9c\x18\xb4\x09\x73\xda\xa6\x30\x7e\x11\xae\x3a\xe5\x61\x70\xad\x88\x09\x85\xb6\xe1\x2a\x0a\x96\x63\x0a\x6f\xae\x5e\x5f\x5f\x0d\x06\xd7\xdf\x5f\xff\xf0\x3a\xc9\x56\x36\x41\x6e\x13\xef\xe2\x35\x3a\x8a\x5f\xb7\xef\xc0\x7e\xfd\x2f\x0e\x0c\x71\x25\xd2\x72\x90\x5c\x27\x57\x7b\x2e\x2a\xc4\xa9\x97\x72\xaa\xa5\xe0\x9b\x14\x86\x72\xcd\x36\x6e\x6f\x37\xda\x52\x83\xba\xf8\xb0\xac\xa9\xb6\x94\xc2\x9b\xc1\x9b\x1f\x6f\xf6\xe6\x9d\xca\x0a\x24\x2b\xf8\x01\xe5\x44\x7b\xf5\x40\x55\xa6\x8d\xd8\x78\xeb\x37\xfc\x34\x9b\x3f\x4e\x47\xf3\x5f\x6f\x66\xf3\xd1\xe4\xfd\xfc\xee\xc3\xbb\xbb\xf1\xe3\xf8\xae\xe1\x0a\x50\x32\xe9\x31\x85\xdb\xf1\xdb\x8f\xef\x3a\x30\xee\x7f\x9f\x4f\x3e\xdc\x8e\xe7\x93\xe1\xfd\xf8\x34\xee\x67\xab\x8b\xb4\x35\x0d\xf0\x2c\x50\x66\xdb\x4b\xa3\xc3\x32\x65\xb4\x4c\x2b\x1d\x24\xa1\x86\xb0\xed\x1d\x69\x3f\x0d\x7f\x1b\xfd\x52\x25\x9d\x4d\x87\xa3\xff\x32\xf3\xee\x04\x24\xfb\x63\xbb\xf7\x6e\xf5\x8b\xc3\xe4\x9f\x1e\x1d\xb9\x36\x28\x37\x3e\x85\xc1\x55\x71\x38\x0b\xc8\xbd\x15\xb4\x19\x69\x45\xf8\x42\x4d\x6f\x63\x45\x29\x24\xe6\x98\xb5\x34\x0c\x50\x6a\xe9\x0b\xbc\x0f\xea\x6f\x49\xa3\x08\x33\xf5\x6a\xfb\xe1\x04\xf4\xb5\xa1\x3e\x57\xa2\xbf\x10\xea\x44\x22\x5c\x89\x78\x21\x54\x9c\x09\x7b\x09\x02\x89\x57\x10\x0a\x29\xc9\x3a\x41\x14\xd2\xd7\x40\x4a\x66\xfb\x52\xe7\x27\xe1\x52\xe7\x17\x42\x43\x94\xf5\xaa\x9f\x69\xbe\x42\x9b\x38\xcd\x57\x27\x08\xb5\xad\x61\xaa\xb9\x69\x1c\xd9\xf3\xd5\x86\xa5\x55\xa9\x9a\x9c\xd7\xa9\x4f\x89\x8b\x2f\x54\x7c\x01\xa8\x8b\xbe\xf8\x4c\xf5\x17\x60\xda\x04\xc6\xe7\x8a\xff\x2a\xc6\x31\x9d\xc7\x0f\x0b\x66\xc4\xa1\x8b\x9d\x79\x08\x78\x47\xba\x78\xd8\x4a\xfe\x16\x9f\x85\x12\xe1\x42\xed\xe8\x75\xa8\x04\xd7\xea\x59\xe4\x2e\xe9\x7e\x1e\xee\x6e\x75\xc7\x75\xe8\x5b\xdb\xf6\x1f\x01\xe4\xf5\x8b\xe1\xdc\xa3\x72\xd7\x8a\xeb\x2a\x77\x74\x94\x83\xaa\x9d\x0c\x1a\x6d\xe2\xe8\xe8\x38\xd2\xb6\xba\xc0\xb7\x73\xd5\x51\xae\x41\x8c\xf4\x96\xc9\xe6\x9a\xeb\xbe\x23\x54\xee\x25\xb3\x0d\x43\xdd\x96\x2b\x2e\xc6\x93\xf7\xa3\x7a\xee\xaf\x00\x00\x00\xff\xff\xb4\xcc\x66\x9b\xc1\x0b\x00\x00")

func awsNodeYamlBytes() ([]byte, error) {
	return bindataRead(
		_awsNodeYaml,
		"aws-node.yaml",
	)
}

func awsNodeYaml() (*asset, error) {
	bytes, err := awsNodeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-node.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x5b\x6f\xdb\x3a\x12\x7e\xf7\xaf\x20\xf4\xbc\x52\xec\x26\xe9\x66\xf9\xd6\xda\xde\xae\x81\xd6\x35\xe2\xa4\x2f\x8b\x45\x30\xa6\xc6\x16\xd7\x14\xc9\xe5\xc5\xb5\xbb\xe7\xfc\xf7\x03\xea\x66\x4a\x71\x8a\x14\x2d\x70\x8e\x9e\x24\x72\xf8\xcd\x70\x2e\xdf\x8c\x40\xf3\x2f\x68\x2c\x57\x92\x92\xc3\x64\xb4\xe7\x32\xa7\x64\x8d\xe6\xc0\x19\xbe\x63\x4c\x79\xe9\x46\x25\x3a\xc8\xc1\x01\x1d\x11\x22\xa1\x44\x4a\x98\x32\x98\x4b\xdb\x7c\x5b\x0d\x0c\x29\xd9\xfb\x0d\xa6\xf6\x64\x1d\x96\x23\x42\x04\x6c\x50\xd8\x70\x84\x10\xdc\xdb\x0c\x4a\xf8\xa6\x24\x7c\xb5\x19\x53\xe5\x15\x53\xa5\x56\x12\xa5\x8b\xb1\x08\xd9\xdf\xd9\x14\xb4\x6e\xb0\xc2\x6a\x9a\xa6\xa3\xd8\x46\xb3\x01\x96\x81\x77\x85\x32\xfc\x1b\x38\xae\x64\xb6\xbf\xb3\x19\x57\x57\x87\xc9\x06\x1d\xb4\x57\x98\x0a\x6f\x1d\x9a\x7b\x25\xb0\x67\x7f\x6c\x56\x50\x62\x24\x3a\xac\xce\x6f\x94\x72\xd6\x19\xd0\x9a\xcb\x5d\xad\x28\xcd\x71\x0b\x5e\x38\xfb\xb3\xb7\x68\xfd\x56\x7b\x87\xb6\xc2\xc6\x0b\xb4\x74\x94\x12\xd0\xfc\x83\x51\x5e\x57\x86\xa5\x24\x49\x46\x84\x18\xb4\xca\x1b\x86\xcd\x1a\xca\x5c\x2b\x2e\x2b\x5b\x52\x62\xeb\x08\xd5\x1f\x5a\xe5\xf5\x4b\x17\x8c\xf0\x79\x40\xb3\x69\xce\x0a\x6e\x5d\xf5\xf2\x15\x1c\x2b\x5e\xa7\x4f\xaa\x7c\x08\xb3\x43\xf7\x2b\xe2\xf1\x9e\xcb\x9c\xcb\x5d\x2f\x2c\x20\xa5\x72\xd5\xf1\x26\x36\x97\x70\x7b\xe1\x02\xef\x94\xd7\x39\x38\xa4\x24\x71\xc6\x63\xf2\x97\x8b\xae\x12\x78\x8f\xdb\xea\x7a\x8d\xbf\xbf\xe3\xaf\x11\x21\xcf\x33\xf7\x05\x64\xeb\x37\xff\x45\xe6\xaa\xd4\xb9\x58\xb1\xaf\xae\xd3\x61\x38\x3b\x0a\x98\x2a\xb9\xe5\xbb\x4f\xa0\xff\xd4\xea\x6f\xf5\x4e\x95\xc1\x2d\x17\x48\xc9\x6f\x95\x64\x46\x6f\xaf\xc9\xff\xab\xd7\x4a\x83\x31\xca\xd8\xee\xb3\x40\x10\xae\xe8\x3e\xcf\x89\x40\x58\xed\xdb\x4c\x28\x06\x22\x02\x20\x55\x0d\x11\x2e\x2d\x32\x6f\x30\x5a\xf7\xda\x3a\x83\x50\x46\x4b\x5b\x10\xc2\x15\x46\xf9\x5d\x41\xb8\x4c\x21\xcf\x4d\x06\x46\x03\xe1\xfa\x6d\xf5\xd2\xc9\xfe\xde\xbd\x69\xa3\x4a\x74\x05\x7a\x4b\xe8\x3f\x26\xb7\xd7\xf1\xc6\xf1\x44\x32\x72\x85\x8e\x5d\x85\x12\x14\x87\x8c\x29\xb9\xed\x04\x18\xb0\x02\xc9\xf5\x78\x54\x03\x0e\x03\x86\x47\x87\x32\xbc\xda\x41\xc1\xcd\x50\x0b\x75\x2a\xf1\x57\xf0\xf7\xa5\x94\x1f\x16\x58\x0d\x9c\x84\x48\xcd\x96\xeb\xe4\xf5\x91\xb7\x1a\x19\xad\xf8\x47\x0b\xce\xc0\x52\xf2\x66\x44\x48\xa8\x55\x87\xbb\x53\x6d\x80\x3b\x69\xa4\xe4\x5e\x09\xc1\xe5\xee\xb1\xaa\xfa\x9a\x25\xe2\x15\xda\xf8\xac\x84\xe3\xa3\x84\x03\x70\x01\x9b\x90\x32\x93\x00\x87\x02\x99\x53\xa6\x96\x29\x03\x0d\x7e\x8c\x2e\xf8\xd2\x15\x5f\x9d\xbc\x0e\x4b\x2d\x3a\x1b\x62\x87\x87\x47\xf4\x54\xbd\xac\xec\x07\x6a\xa5\xf5\x5a\xf5\xde\x2b\xfe\xe5\x20\xc2\x75\x96\x71\x65\xb8\x3b\x4d\x05\x58\xbb\x8c\x28\x25\x0d\x34\x9f\x32\xc3\x1d\x67\x20\x1a\x69\xd8\x6e\xb9\xe4\xee\x74\x36\x38\x48\xbd\x7b\xb6\x1a\x62\xf6\x3f\xcf\x0d\xe6\x33\x6f\xb8\xdc\xad\x59\x81\xb9\x0f\x01\x59\xec\xa4\xea\x96\xe7\x47\x64\x3e\x30\x5d\x7c\xb2\xc6\x5c\x37\x61\x79\x40\x53\xda\xfe\x76\x5a\x47\x69\x7e\xd4\x06\xad\x3d\x37\x86\x58\x62\x8f\x27\x4a\x92\x90\xf5\x83\xe6\xa0\x6c\x32\x10\x26\x44\x69\x34\x10\x52\x80\x2c\xe4\xb3\xcd\x03\x08\x8f\xcf\x34\xd4\xbd\x53\xfa\xe3\xeb\x35\x83\x61\xc5\x2f\xd3\x0d\x65\xfe\xf6\xa6\x59\x77\x4a\x04\x8c\xd8\x11\xad\x19\xd3\x26\x7c\xef\xf2\x5c\x49\xfb\x59\x8a\xd3\xd9\x82\xb3\xe6\x64\x7e\xe4\xd6\x75\x8e\x61\x4a\x3a\xe0\x12\x4d\x04\x37\x24\x87\xfa\xe1\x25\xec\x90\x92\xb7\xe3\x37\x37\xe3\xc9\xe4\xe6\xfa\xe6\xf6\x4d\x96\xef\x4d\x86\xcc\x64\xf7\xf3\x0f\x8b\xcf\xcb\x41\xca\xe2\xde\x5e\x35\x20\xf4\x30\xc9\x26\xd9\x75\x1f\x6b\xe5\x85\x58\x29\xc1\xd9\x89\x92\xc5\x76\xa9\xdc\xca\xa0\xc5\xaa\x6d\xd5\x4f\x6f\x14\x69\x1f\xc1\x4b\xee\x06\x6e\x2a\xb1\x54\xe6\x44\xc9\xe4\xef\xe3\x4f\x7c\x90\x97\x68\x87\xd2\x4c\x7b\x4a\x26\xe3\x71\x79\x11\xa3\x07\x01\x66\x67\x29\xf9\x37\x49\xd2\x40\xc6\xc9\xdf\x48\x52\x11\x74\x73\xab\xab\xb6\x1f\x25\xe4\x3f\xdd\x91\x83\x12\xbe\xc4\x4f\xa1\x04\x23\xbd\x67\xa7\x86\x7e\x9a\xd6\x42\x91\xfe\x32\xc8\xaf\xc0\x15\x94\xc4\x1a\x7a\x77\x81\x3c\xc4\x94\x92\x30\xe5\x9c\x1b\x87\x32\x7d\x3d\x5d\x40\x57\xca\x38\x4a\xa2\x1e\xd3\xb2\x7e\x1f\x57\x1b\xe5\x14\x53\x82\x92\xc7\xd9\xea\x47\x71\x52\xc7\xf4\x45\xac\x87\xe9\x77\xb0\x7a\x9d\xaf\x45\x2b\xd1\x19\xce\x2e\x5b\x16\xa3\x55\xad\x39\x70\x98\x92\x0e\x8f\x2e\x0e\x2d\x08\xa1\xbe\xae\x0c\x3f\x70\x81\x3b\x9c\x5b\x06\xa2\xaa\x14\x1a\x7a\xb5\x8d\xdd\xcd\x40\xc3\x86\x0b\xee\xf8\xb0\xe2\x20\xcf\x87\x04\xb4\x9c\x3f\x3c\xbd\x5f\x2c\x67\x4f\xeb\xf9\xfd\x97\xc5\x74\xde\xdb\xce\x8d\xd2\xc3\x03\x20\xc4\x85\xc0\xdd\x2b\xe5\xfe\xc9\x05\x36\x43\x5c\x3f\x8c\x82\x1f\x50\xa2\xb5\x2b\xa3\x36\x18\xe3\x15\xce\xe9\x0f\xe8\xfa\x2a\x74\x9d\x28\x83\x01\x87\x34\xe9\x40\xc9\xdd\xf8\x6e\xdc\x5b\xb6\xac\xc0\xe0\xe4\x7f\x3d\x3c\xac\xa2\x8d\x40\xe4\x1c\xc4\x0c\x05\x9c\xd6\xc8\x94\xcc\x6d\x28\xf0\x48\xc2\xf1\x12\x95\x77\xdd\xe6\x6d\xb4\x67\x3d\x63\x68\xed\x43\x61\xd0\x16\x4a\xe4\x75\x8b\x6d\x9f\x2d\x70\xe1\x0d\x46\xbb\xed\xd9\x5c\xda\xb6\xec\x67\xf5\xe8\xdd\x6c\xd4\x55\xf1\x03\x55\xc3\xda\xe9\x74\xd0\x52\x2e\xf2\x57\x75\x61\x87\xcf\x1b\x4c\xc5\x9e\x6d\x29\x0f\xe8\xb7\xf6\x74\xb7\xf9\xe2\x9c\xdc\x0c\xde\x17\x66\xac\xc1\xef\xc1\xc5\x21\xeb\xd9\x6f\xcf\x79\x4e\x0c\xcd\xa4\x0e\x6a\x12\xca\x26\xb9\xb0\x6d\x99\x01\xfd\xe2\xef\xcf\x2b\x66\xb6\x66\x1c\x4e\x9b\x01\x22\x42\xfa\xe9\xe9\xae\xd3\xda\x0e\x2a\xfd\x09\xec\x92\x75\x8d\x35\x8b\x15\x25\xb3\xe5\xfa\x69\xfa\xf1\x71\xfd\x30\xbf\x7f\x5a\x84\xc4\xed\xd8\x2e\x1d\x70\x99\x8e\x49\x6a\x48\x69\xe9\x05\xc2\x7a\xe1\x40\x60\x9a\x3f\x02\x00\x00\xff\xff\x17\xfe\xaa\x4f\x0c\x11\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"aws-node-1.10.yaml": awsNode110Yaml,
	"aws-node.yaml": awsNodeYaml,
	"coredns.yaml": corednsYaml,
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
	"aws-node-1.10.yaml": &bintree{awsNode110Yaml, map[string]*bintree{}},
	"aws-node.yaml": &bintree{awsNodeYaml, map[string]*bintree{}},
	"coredns.yaml": &bintree{corednsYaml, map[string]*bintree{}},
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

