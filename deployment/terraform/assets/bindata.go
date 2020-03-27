// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/cluster.tf (9.56kB)
// assets/outputs.tf (327B)
// assets/variables.tf (589B)

package assets

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _clusterTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x6f\x1b\xb9\xf1\x7f\xaf\x4f\x31\xff\x4d\x80\xfc\x5b\x94\xbb\xf2\x43\x1c\xfb\x0e\x4e\x91\x5e\x5b\xb4\xc0\x5d\x1b\xa0\x01\xfa\x22\x67\x2c\x28\x72\xb4\xe2\x89\x4b\xee\x91\x5c\xd9\x4a\xea\x7e\xf6\x82\xe4\xae\x76\xb5\x7a\xb0\xac\x9c\x93\x4b\x51\x05\x88\x05\x72\x38\x33\xfc\x71\xe6\xc7\x19\xaa\x32\x7a\x21\x38\x1a\x48\xe8\xad\x4d\xe0\xe3\x08\xa0\x32\x7a\x2a\x24\xc2\x35\x24\x1c\xa7\xb4\x96\x2e\x19\x01\x18\x2c\x84\x56\xe0\x47\x6b\x4b\x90\x5a\x47\x4e\xfd\xf8\x02\x8d\xf5\x13\xd7\x90\xfc\xfb\x35\x9c\xa6\xe7\xaf\x92\xd1\xfd\x68\x64\xd0\xea\xda\x30\x0c\x7a\xf3\x39\x2e\xf3\x8a\x0a\x93\x40\x32\xc7\x65\x34\xe3\xc7\x14\x2d\x11\x82\xce\xe7\x1f\x17\xd4\xa4\x4c\xd6\xd6\xa1\x09\xe3\xf7\x64\x8e\xcb\xb0\xc8\xfb\x54\x4f\xa4\x60\x5e\x0f\x5c\x83\xf7\xee\xff\xbd\xb8\xb5\xb3\xbc\x9b\xf9\xcd\xa6\x5d\xa1\xac\xa3\x8a\x61\x02\x09\xad\xaa\xdc\xa2\x59\xa0\x89\xe6\x1d\x2d\x2c\x5c\x87\xaf\x00\x7f\xf3\x7e\xec\xf0\x82\x56\x15\x79\xfe\x91\xe9\x5a\xb9\x54\x28\x8e\x77\xf7\xde\xa1\xfb\xd1\x08\x80\x69\xa5\x90\x39\xbf\xfd\xa8\xe7\x19\xbc\x9b\x21\x34\xa8\x41\x6d\xd1\x84\x1d\x4e\xb5\x01\x5d\x1b\x78\xf3\xc3\x5f\x83\x98\x5b\x56\xc1\x9c\xb5\xb3\x24\x0c\x78\xc9\x80\xec\xa4\x56\xae\x8e\x63\x33\x6d\x1d\x5c\x83\x45\x39\x4d\x9b\x4d\x8a\xaa\xb5\x4c\x4b\x01\xdd\xe7\x1a\x12\x5a\x0a\x32\x9e\xb2\xd3\x31\xe7\x27\x9c\x9e\x8f\x2f\x5e\x5d\x8e\x27\x09\x3c\x83\x93\xcb\x74\x7c\x0e\xdf\xbf\xfb\xc7\x08\xa0\x85\x23\x6f\x1d\x60\x2f\xd3\x3b\x49\x4d\x81\x09\xf4\x3f\xcf\xe0\x8d\xbc\xa5\x4b\xeb\xfd\x82\x95\x4c\xd8\x06\xd3\xca\x0a\xeb\x50\xb1\xe5\xfa\x19\x06\x37\xfa\x67\x9d\xce\x71\x99\x0a\x1e\x60\xaa\x95\xeb\x39\xeb\x31\xf6\xc7\xb1\x72\x27\x08\xf8\x58\xaa\x58\x6e\x91\xd5\x46\xb8\x65\x5e\x18\x5d\x57\xb9\xe0\xfe\x94\xde\x07\x44\x92\xe7\x1f\xbd\x81\x75\x09\xaf\x29\x15\xfc\x3e\xf9\xdd\x7e\x99\xbc\xd0\xd6\x8a\x28\xea\x11\xbc\x19\xc5\x50\x5f\x08\x1f\xbe\x3e\x01\x7c\x58\x25\xcd\x41\x36\x41\xd4\x73\xb8\xa4\xce\xa1\x29\xb5\x75\xb9\x14\x0c\x95\xc5\xdc\x2f\x08\xd2\x1c\xad\x13\x8a\xba\x26\x0f\xb2\x99\x2e\x31\x8b\x67\x99\x75\xeb\x7a\x2a\x48\xa3\x62\x15\x49\x6b\x8e\x18\x2c\xb5\x43\x82\x77\xc8\x5a\x7f\x84\x92\x42\xe1\x0a\x09\x80\xe4\x76\xe6\x73\xf4\x3d\xfc\x1f\x90\x29\x64\x0b\x6a\x32\x29\x26\x19\x93\xba\xe6\x59\x0b\x6c\x36\xd1\xda\x91\xa9\x50\xc2\xce\x90\xc3\xcd\xb7\xc0\x35\x20\x9b\x69\x78\xf1\x4f\x2a\x9c\x50\x45\x3c\x53\xbf\x88\x08\x25\x5c\x9a\xa6\x2f\xbe\x05\x2b\x11\x2b\x38\xf1\xd2\x0a\x1b\x5c\xbd\xc5\x02\x1d\x10\xa2\x34\x61\x33\x64\x73\xc2\xd0\x38\x31\x15\x8c\x3a\x04\xf2\xf3\xdf\x81\xc0\xcc\xb9\xca\x7e\x93\x65\xf6\x8c\x60\x4d\x6e\xd1\x3a\x72\x92\xd2\x92\x7e\xd0\x8a\xde\xda\x94\xe9\x32\xe3\x38\x49\x8d\x9e\xd4\xd6\x55\x68\x18\x56\x1e\xb3\x54\xe8\xec\xfc\xe4\x4f\x7f\xfe\xee\xea\xea\x8f\x69\x51\x15\xf0\x2f\xb0\x35\xd7\x40\x2b\xe7\x39\x00\x28\xe7\x40\x3a\x3f\x56\x73\xc1\x9f\x25\xd4\x15\xa7\x0e\x77\xcc\x07\x28\xa4\xf4\x72\x3f\xfd\xfc\xb0\x4c\x65\x74\x89\x6e\x86\xb5\x25\x4a\x73\x7f\x06\x95\x36\x0e\xcd\x00\x85\xc8\x13\xbd\x90\xe0\xfa\x56\x49\x4d\x79\x5e\x1b\x79\xdf\x09\x3b\x6a\xe0\xee\xc3\x14\x7a\x27\xff\xdb\xd4\x51\x93\x16\x1f\x06\xbe\x94\x8b\x9e\x10\x64\xba\x72\xd9\x50\x62\xce\x85\x01\x52\xc5\xc9\x4e\x38\xe3\xd4\xd1\x48\x19\x37\x21\x9c\xf6\xb1\x60\x89\xce\x08\x66\x8f\x63\xc2\x66\xf1\x7f\x0f\xfd\xb9\xd3\x34\xb2\xdf\x81\x54\x76\x2c\x45\x35\xc0\xad\x68\x6a\x93\x7b\xfe\x97\xf2\x4f\x92\xf2\x5d\x3a\x0f\x64\xed\xd2\x3a\x2c\x99\x93\x80\x8a\x4e\x24\xee\x96\xdc\xa2\x95\x72\x1e\x22\x56\x8a\xc9\x54\x2b\xc7\xb4\x9a\x8a\xe2\x64\x80\x5a\x0b\x0c\x97\x69\x61\xe8\x94\x2a\x1a\xc0\xd0\xd6\x66\x06\x25\x52\x8b\x59\x33\x9e\x5f\xa4\x17\xe9\x69\x4e\x4b\x7e\x71\x9e\x72\x9c\x0c\x1c\xe0\xd5\xbc\x00\x22\xe0\x30\xe9\x6e\x63\x9c\x62\xa9\x15\x31\xe8\x79\xe9\xa1\xed\x37\xca\x49\xc3\x0b\x03\x69\x34\x0b\xc1\x86\x42\x60\x1d\x35\xee\x50\xe2\xa9\x8c\xbe\x5b\x1e\x47\x3b\x61\x69\x24\x9d\x61\xe6\xaf\x7f\x76\xf2\xc0\x46\xee\x6f\xac\x2b\xcf\xdb\x42\xc8\xdb\xb0\x56\x33\x41\x1d\xe6\x2b\xe2\xc9\x29\xe7\x06\xad\x77\xd8\x99\x1a\x8f\x25\x82\xb0\x95\x58\x82\x44\xd0\xd6\x59\x67\x2f\xf3\x7c\x39\xb6\xfd\x1a\x98\xea\x58\x86\x50\x85\x50\x77\xbf\x4c\x0e\x6d\x53\x65\xca\x80\x00\x3a\x96\x85\xe9\xcc\x0a\x87\x96\xc4\x15\x3c\x6b\xfb\xab\xf5\x45\x52\x01\x99\xda\xcd\x55\x74\x41\x85\xf4\x0b\xb3\x7e\xa9\xb0\x4b\x77\x27\x33\xc8\xd1\x41\x92\x1a\x6e\xf3\x36\xe1\x7a\x09\x3b\x1c\x6a\x9a\xc3\xf5\x5a\xbe\x89\x58\x9f\xb3\x7c\xb2\x59\xcf\x0b\x8e\xca\xdf\x17\x68\xfa\xb9\xb6\x2d\xc7\xf9\x64\x4b\x8f\xb5\x72\xa1\x53\x13\xb3\xa3\xe7\xb2\xb7\xdb\x7e\x0d\xcd\x46\xe7\x84\xa4\xd6\xee\x72\xd0\xcf\x8d\x00\xd0\x03\x87\x0f\xec\x26\x0a\x79\x5e\xa8\x2a\xb9\xcc\x45\x59\x22\xf7\xec\x20\x97\xd0\xd2\xc1\x06\xf3\xf5\x3c\x4c\x20\xe9\x7c\x6c\x40\xdc\xdc\xd8\x6e\x60\x3c\x12\xbe\xc4\x9b\x50\x8b\x7d\xae\xd8\xbe\x20\xca\x97\x34\x8c\xac\x38\x61\x6d\x63\xed\x68\x27\x57\x51\x6b\x6f\xb5\xe1\xeb\x72\xed\xe8\x08\xc0\xce\x45\x95\x4f\x85\xa2\x32\xb7\x8a\x56\x76\xa6\x5d\xc7\x84\x5b\x60\xe9\x26\xb7\x20\xbc\x0f\xe2\xf8\x25\x6f\x1f\x15\xd6\xc5\xd7\x27\xdf\x6f\x57\x72\xb3\xbf\x4a\xdb\xce\xcb\x7c\x12\x48\xf9\x66\xef\x15\xe6\x39\xc0\xa1\x75\x39\x2d\x50\xb9\xc7\xbe\x22\xf8\x35\x07\xbf\x23\x7c\xb6\x0a\x79\x47\x55\xec\x8f\xb2\x2e\x0f\x2d\x8b\xb7\x77\xf8\xeb\x70\x35\x9c\x70\xc4\xd9\x84\xe5\xcd\xf1\x1c\x70\x19\x7d\x1d\xd7\xd1\x46\x89\xa8\x75\x21\x31\x54\x88\x85\xce\x0a\x1d\x43\xa8\xd0\x6d\xb0\xdf\xa7\x52\xa8\xfa\x8e\xc4\xb2\x6f\x6b\x07\xe9\x7b\x4d\xf2\x1d\x64\xb5\x35\x99\xd4\x8c\x4a\x20\xbe\xf5\x3c\x4a\x97\x9d\x01\x61\xf0\x63\x12\xf7\x1f\x9b\x60\x78\xfb\xe6\xdd\x5f\xae\x9f\xfb\xff\xbf\xe9\x8c\x78\x77\x27\x42\xbd\x80\xd7\xaf\xe3\x55\xd4\xbc\x1e\xfe\x98\xec\xd8\x6c\x21\xdc\xac\x9e\x84\x9d\xf6\x7a\xd8\xfe\xd3\x88\xa6\x9c\xf8\xc0\x21\xaa\xc8\xa8\x61\x33\xb1\xc0\x2c\xee\x61\x15\x52\x31\x43\x73\xa6\x39\xe6\x06\xa7\xf7\x1b\xbb\x08\x60\xf8\xed\x3f\x76\xe1\x5a\x1f\xbe\xe6\x0b\x79\x40\xd5\xae\x75\x9d\x6e\x56\xed\x92\xc9\x62\x0f\x91\x59\x51\x56\x12\x99\x56\xce\x68\x29\xfd\xed\x16\x0b\x84\xf4\x27\xab\xd5\xa3\x17\xfb\x45\x8f\xb1\x1e\xff\x3c\xca\x66\xb3\xa4\x67\xe9\x06\xb6\xf7\x02\xeb\x59\x1d\x1f\x64\x63\xce\xf6\xeb\xdf\x3d\x6f\xb0\xad\x02\x12\x15\x8c\xc2\xb3\x1b\x33\xa2\x6a\x9f\xdd\xde\x54\x15\xb4\x42\x10\x84\x42\xc6\xb6\x07\xd6\xde\xba\xb0\xc5\x40\x78\x11\x14\xaa\x08\x25\x7e\xe4\x91\xa9\xd1\x65\x1e\xe2\xde\x7b\x75\x7a\x1a\x79\x59\xb7\x43\xbd\xc1\xca\x68\xa7\x99\x96\x8d\xff\x8e\x55\x91\x9f\x99\xe0\x26\x9f\x48\xcd\xe6\x91\xdd\xc6\x69\xf8\x97\x8d\x93\x9b\xa6\x99\xd9\x67\xf1\x72\x7c\xf1\x72\x8b\xcd\xd5\xf0\x2f\x6f\x35\x28\x7f\x35\xb0\xd9\x1b\xec\x2c\xf6\xed\x3d\x83\x1f\xe8\x72\x82\x60\xd0\x3a\x23\x98\x03\xad\xe4\x32\x68\x85\xb7\xab\x9e\x1a\x9a\x96\xf1\xf7\xcd\x92\x3f\xd4\x0e\x66\x54\xf1\x25\x44\x6a\x76\x74\xee\x09\xb6\xe1\x0e\x0b\xb7\xc2\xcd\x74\xed\xa0\xa4\xaa\xa6\x52\x2e\xc1\xda\x19\xf1\x12\x42\x39\x0d\x6e\x86\x8d\xc2\xf4\x93\x81\x8e\xf0\x5d\x9d\x8c\xc7\x1b\x60\x0f\xa6\xfa\x80\x0f\x41\x5f\x0f\xee\x3d\x97\x59\xff\x25\xa8\x75\x0e\x77\x07\xc1\xa6\x53\xed\xd8\xe0\xf8\xc9\xc9\x61\xa7\x7f\x50\x5a\x36\x4f\xe5\x9f\x90\x9d\xa4\xd1\x70\x70\x92\x46\xf9\x03\x73\xf5\xe1\xf3\xbc\x1c\xbf\x3a\xdf\x71\x9e\xab\xa9\x2d\xe7\x59\xf3\xc7\x9e\x67\xfb\xe3\xc3\xa1\x81\x76\xa4\x63\x8f\x0f\xb4\x23\x1c\xdb\xa4\x9b\xc1\xd4\x97\x42\xec\x18\xc7\x3e\x1d\xb1\xcf\x9c\x96\x0f\xe5\x25\x9f\xf4\xf2\x71\x77\x17\x3d\xbc\x25\x0f\x40\xf8\xec\x6c\x7c\xb1\x03\xe1\xd5\xd4\x13\x20\x7c\x80\x67\x2f\xcf\xcf\x36\xef\xdd\xc1\xd4\x13\x78\x76\x00\x49\x76\x1d\xe0\x21\xfc\x18\x7a\xbf\x07\xeb\x97\xef\x5b\xf2\x0b\xf2\x5f\x55\x29\xb3\xb7\xaa\x38\x1f\x6f\x5c\xae\xbd\xc1\xed\x55\x85\x6f\x6a\x5b\x4c\x9b\x97\x84\x60\xe5\xd7\x75\x57\xb6\x3f\x88\x3d\x98\x98\x8d\xe0\xe3\xb2\xf3\x4b\x54\x9d\x57\xe3\xab\x6d\x38\xae\x86\x9f\xc6\xea\xd9\x66\x84\xac\x0d\x1f\x6f\xf5\xd7\x15\x2f\xf1\x97\x8c\x43\x69\x23\x48\x3f\x4c\x1b\x6f\xbd\xd8\x93\xb1\xc5\xe5\x36\x98\x2e\x3f\xed\x58\xbe\x0c\x41\x7d\xde\x50\xf8\x4f\x00\x00\x00\xff\xff\x5b\xb3\xcb\x62\x58\x25\x00\x00")

func clusterTfBytes() ([]byte, error) {
	return bindataRead(
		_clusterTf,
		"cluster.tf",
	)
}

func clusterTf() (*asset, error) {
	bytes, err := clusterTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd8, 0x4e, 0x5, 0x94, 0xc7, 0xda, 0x14, 0xdb, 0xbc, 0xda, 0x21, 0x3f, 0xc0, 0x81, 0x51, 0x8, 0xa7, 0xd1, 0xc, 0xe4, 0x82, 0x26, 0xd4, 0xf0, 0xe, 0xb1, 0x88, 0x69, 0xfd, 0x96, 0x32, 0x5a}}
	return a, nil
}

var _outputsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x10\x5c\xb9\xe8\x0d\x5c\x79\x04\x0f\x10\xa6\xcd\x20\x85\x98\x84\x99\x49\x55\x4a\xef\x2e\xc4\x06\xac\x48\xdd\x05\xfe\xfb\x2f\x7f\x52\xd1\x5c\x14\xec\x18\x45\x31\x0e\x24\x16\x66\x03\x30\x61\x28\x04\x27\xb0\x87\x19\xef\xe2\x5a\xda\x61\xce\x4e\x88\x27\xe2\xee\xb8\x58\xb3\x18\xd3\x04\xbe\x3f\x87\x22\x4a\xfc\x53\xc0\x5e\xdc\xf0\xce\x3b\xdf\xb7\xe7\xd6\x80\x57\x8a\xfa\xe7\xff\x90\xd0\x2b\x89\xba\x0a\x7f\x6f\xb8\x91\xf2\x38\xc8\xa5\x0e\xdc\x17\xad\xe8\x7a\xcc\x56\x93\x39\x3d\x9e\xfb\xf5\x8a\x7c\x94\x5f\x01\x00\x00\xff\xff\x6f\xd3\xe1\x13\x47\x01\x00\x00")

func outputsTfBytes() ([]byte, error) {
	return bindataRead(
		_outputsTf,
		"outputs.tf",
	)
}

func outputsTf() (*asset, error) {
	bytes, err := outputsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "outputs.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd5, 0xd7, 0xa7, 0xff, 0x49, 0x7d, 0xf7, 0xf2, 0xf0, 0xad, 0x3d, 0x63, 0x4d, 0x52, 0x40, 0x3f, 0x63, 0xb2, 0x30, 0x39, 0xfe, 0xb0, 0xbb, 0x6b, 0x82, 0xfe, 0xd3, 0xf7, 0xce, 0xb2, 0x30, 0x67}}
	return a, nil
}

var _variablesTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x41\x6b\xf3\x30\x0c\x86\xef\xf9\x15\xc2\xf7\xcf\x7c\xb4\x6c\x63\x87\xfc\x16\xa1\xc4\x6a\x66\xe6\xd8\x9e\x24\xb7\x94\xd1\xff\x3e\x9a\x1e\xd6\x92\x6c\xf3\xd1\xcf\x23\xf1\xa2\xf7\x48\x12\x69\x48\x0c\x6e\x4c\x4d\x8d\x05\x33\xcd\xec\xe0\xb3\xbb\x74\xdd\x37\xa4\x5a\x31\x66\x35\xca\x23\xe3\x58\x5a\xb6\x95\x92\x0a\x05\x63\x35\xa4\x89\xb3\xfd\x20\x85\xe1\xaf\x35\xf7\x06\xe7\x29\xe6\x75\x98\x87\x25\x89\x54\xb7\x8c\xdb\x2c\x1e\x59\x34\x96\x7c\x35\x00\xec\x5c\x19\x7a\x98\xa9\x76\x00\x81\x0f\xd4\x92\x41\xbf\x20\x00\x47\x4d\x8a\xd0\xbf\xf9\xac\x1f\xc9\xc1\xf2\x7a\x70\x4f\xfe\xc5\x2f\x5f\x78\xe3\x7e\xe7\xff\xef\xfd\xce\x3d\xcc\xd4\xa2\x36\x09\x2f\x83\x3d\xb8\x57\xff\xec\xf7\x57\xe3\xb2\x8a\xd5\x94\x65\xf3\xc2\x61\xc0\x4a\xaa\xa7\x22\x61\xc5\x54\xdf\xb0\xb6\x21\xc5\x11\xdf\xf9\xbc\xc2\x33\x99\xb1\xcc\x45\x0d\x43\x39\xe5\x6b\x11\xd8\x24\xfd\xe6\xa5\x38\x72\x56\xc6\x43\x4c\xeb\x28\x53\xb9\x3f\xdb\x76\xc9\x5a\x9a\x2c\x1d\x06\x46\xe1\xc3\x4d\xfc\x0a\x00\x00\xff\xff\x72\x3f\x4a\x9d\x4d\x02\x00\x00")

func variablesTfBytes() ([]byte, error) {
	return bindataRead(
		_variablesTf,
		"variables.tf",
	)
}

func variablesTf() (*asset, error) {
	bytes, err := variablesTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "variables.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdd, 0xb1, 0xf7, 0x16, 0xdf, 0x7a, 0x8c, 0x5, 0x77, 0xbb, 0x8, 0x28, 0x99, 0xe8, 0x71, 0xf9, 0x17, 0x60, 0x85, 0x7c, 0xeb, 0x55, 0x28, 0xd7, 0xfe, 0xa4, 0x7a, 0x88, 0x59, 0xe2, 0xa8, 0xc2}}
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
	"cluster.tf":   clusterTf,
	"outputs.tf":   outputsTf,
	"variables.tf": variablesTf,
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
	"cluster.tf":   &bintree{clusterTf, map[string]*bintree{}},
	"outputs.tf":   &bintree{outputsTf, map[string]*bintree{}},
	"variables.tf": &bintree{variablesTf, map[string]*bintree{}},
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
