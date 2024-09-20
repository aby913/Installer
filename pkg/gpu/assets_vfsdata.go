// Code generated by vfsgen; DO NOT EDIT.

package gpu

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// assets statically implements the virtual filesystem provided to vfsgen.
var assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2024, 9, 14, 6, 33, 21, 724900006, time.UTC),
		},
		"/device-plugin.yaml": &vfsgen۰CompressedFileInfo{
			name:             "device-plugin.yaml",
			modTime:          time.Date(2024, 9, 14, 6, 33, 21, 721697115, time.UTC),
			uncompressedSize: 3594,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x56\xdf\x6f\xdb\x38\x0c\x7e\xf7\x5f\x41\x24\x40\xb1\x01\x75\xdc\xf6\x5e\x0e\xbe\xa7\xac\xcd\x6d\xc1\xf5\xd2\xa0\x69\x3b\x0c\xc3\x30\xc8\x32\x13\x13\x95\x25\x9d\x44\x27\x31\xb0\x3f\xfe\x20\xc7\xd9\x9c\x1f\xbb\x64\x4f\xe7\x27\x43\x22\xa9\x8f\x9f\x3e\x91\xec\xc3\xad\xb1\xb5\xa3\x45\xc1\xf0\x46\xbe\x85\x9b\xab\x9b\xdf\xe0\x3d\x1a\xb7\x20\xe3\x61\xa8\x70\x6d\xac\xa9\x94\xf1\x51\x3f\xea\xc3\x3d\x49\xd4\x1e\x73\xa8\x74\x8e\x0e\xb8\x40\x18\x5a\x21\x0b\xdc\xee\x5c\xc2\x0b\x3a\x4f\x46\xc3\xcd\xe0\x0a\xde\x04\x83\x5e\xbb\xd5\x7b\xfb\x47\xd4\x87\xda\x54\x50\x8a\x1a\xb4\x61\xa8\x3c\x02\x17\xe4\x61\x4e\x0a\x01\xd7\x12\x2d\x03\x69\x90\xa6\xb4\x8a\x84\x96\x08\x2b\xe2\xa2\x39\xa6\x0d\x32\x88\xfa\xf0\xa9\x0d\x61\x32\x16\xa4\x41\x80\x34\xb6\x06\x33\xef\xda\x81\xe0\x06\x70\xf8\x0a\x66\x9b\x26\xc9\x6a\xb5\x1a\x88\x06\xec\xc0\xb8\x45\xa2\x36\x86\x3e\xb9\x1f\xdf\x8e\x26\xb3\x51\x7c\x33\xb8\x6a\x5c\x9e\xb5\x42\xef\xc1\xe1\x3f\x15\x39\xcc\x21\xab\x41\x58\xab\x48\x8a\x4c\x21\x28\xb1\x02\xe3\x40\x2c\x1c\x62\x0e\x6c\x02\xde\x95\x23\x26\xbd\xb8\x04\x6f\xe6\xbc\x12\x0e\xa3\x3e\xe4\xe4\xd9\x51\x56\xf1\x0e\x59\x5b\x74\xe4\x77\x0c\x8c\x06\xa1\xa1\x37\x9c\xc1\x78\xd6\x83\x77\xc3\xd9\x78\x76\x19\xf5\xe1\xe3\xf8\xe9\xc3\xc3\xf3\x13\x7c\x1c\x3e\x3e\x0e\x27\x4f\xe3\xd1\x0c\x1e\x1e\xe1\xf6\x61\x72\x37\x7e\x1a\x3f\x4c\x66\xf0\xf0\x27\x0c\x27\x9f\xe0\xaf\xf1\xe4\xee\x12\x90\xb8\x40\x07\xb8\xb6\x2e\xe0\x37\x0e\x28\xd0\x88\x79\xe0\x6c\x86\xb8\x03\x60\x6e\x36\x80\xbc\x45\x49\x73\x92\xa0\x84\x5e\x54\x62\x81\xb0\x30\x4b\x74\x9a\xf4\x02\x2c\xba\x92\x7c\xb8\x4c\x0f\x42\xe7\x51\x1f\x14\x95\xc4\x82\x9b\x95\x83\xa4\x06\x51\x24\x2c\xb5\xd7\x9f\x06\xce\x7c\xb2\xbc\x8e\x5e\x49\xe7\x29\xdc\x09\x2c\x8d\x9e\x21\x47\x25\xb2\xc8\x05\x8b\x34\x02\xd0\xa2\xc4\x14\xf4\xd2\x17\xc2\x61\x9c\xe3\x92\x24\xc6\x56\x55\x0b\xd2\xed\xae\xb7\x42\x76\x4c\x7c\xed\x19\xcb\x28\xa0\x0e\xfe\x1e\x15\x4a\x36\x2e\xfc\x03\x94\x82\x65\x71\x2f\x32\x54\x7e\xb3\x70\xea\x00\xc6\xd2\x2a\xc1\xd8\xba\x77\x80\x85\x4f\xed\x44\x3a\x15\x0b\x60\x0b\x2a\x7c\xae\xd2\x4c\x25\xde\x2a\xe1\xfd\xa4\xf5\xa3\x9c\x04\xf4\x61\xb4\x0e\x5a\x22\x56\x75\x23\x31\xf4\xdc\x90\xd8\x7a\xb4\xfe\xd6\x91\x71\xc4\x75\x27\xc0\x26\xf5\x58\x9b\x1c\x63\x19\x04\x27\x85\x6a\xad\x49\x13\xdf\x1a\x1d\x5e\x03\xba\x0e\xe2\xb8\xc5\x1c\xf6\xe3\x9c\xdc\xf7\x0d\x00\x2a\xc5\x02\x53\xc8\x2a\x5f\x67\x66\x9d\x5e\x0f\x6e\x7e\xef\xec\x2e\x8d\xaa\x4a\xfc\xdb\x54\x9a\x3b\xe1\xba\x21\x0b\xe3\x39\x5e\x0a\x17\xbb\x4a\xc7\x2d\x27\x3b\x86\x00\x65\x70\x9f\x0a\x2e\x52\x48\x96\xc2\x25\xae\xd2\xc9\xa1\xa5\x34\x65\x29\x74\xbe\x7f\x8a\x2f\xf6\x16\x62\xb9\xb7\xd0\xfb\x0c\x71\x7e\x10\x39\x51\x94\xb5\xbf\x03\x6f\xe0\x0b\x5c\x5c\x80\x2b\x21\x76\xf3\x53\xa6\xdf\xbe\x01\xbb\x0a\x7b\xd1\x16\xd7\xcf\xf9\xdc\x6a\x40\x51\x76\x84\xd2\x9a\x91\x9d\xc8\x71\x7b\x4e\xfa\xe3\x9c\x33\xf2\x56\x88\x76\x6f\x8d\xf4\x3c\xdc\x60\xdd\x59\x56\x34\x47\x59\x4b\x85\xbb\xfe\xd6\x78\x9e\xb1\x70\x9c\xee\xdd\x05\xae\x7f\x68\xf3\x04\x86\xef\xfc\x26\x19\xe9\xc4\x17\xbd\xe3\xbb\xb1\xfc\xc9\x06\x07\x45\xc7\x73\x48\x8e\x49\xe4\x08\xe7\xa6\x92\xc5\x59\xc6\x17\x17\x1b\x49\x41\xbc\x84\x38\xce\x48\xe7\xb0\x67\x71\x46\x94\x5d\xd0\xd6\xe1\x8c\x8d\xfd\x1f\xd9\xaa\xbe\xa7\x74\x26\x05\x5b\x31\xff\x5a\xaa\x1e\x65\xd5\xd4\x13\xa3\x19\xd7\x7b\xf2\xe8\xc3\x04\x25\x7a\x2f\x5c\xdd\xf4\x85\x06\x92\x0f\xbd\x6d\x65\xdc\xeb\x60\x8f\x31\x5a\x92\xc2\x05\xe6\x69\xf3\x5c\xce\x2c\x1a\xdd\x5a\x70\x46\xe9\x38\xbb\xc6\xf4\x61\x08\x19\xe5\xe4\x50\x86\xae\x24\x54\x2b\x11\xd4\xbe\x72\xe8\x81\x0b\xc1\xed\x92\x35\x14\xb2\x12\xca\x1b\xf0\x85\x59\x41\x65\x43\xdf\xe5\xe2\x30\x66\x38\x78\x00\x1f\x11\x74\xd3\xe3\xc3\x84\x92\xa1\x14\x61\x5c\x39\x5a\xfe\xa1\x34\x39\xcd\xa9\x39\x6f\xd3\x52\xfd\x41\xcc\x76\x3c\x99\x9a\xbc\x45\xb5\x2d\xfe\x6d\x44\x58\x92\xe3\x4a\x28\x78\x3f\x7d\x6e\xda\x2d\x88\x3c\xf7\x20\x1a\x30\x81\xb9\x83\x88\x9b\xb4\xc2\x85\xfd\x77\x65\x3b\x5a\x96\x9d\xb1\x62\xd1\x74\xf2\x14\xde\x75\x09\xfc\x69\xb1\x3b\x6c\x78\x27\xca\xde\x59\x8e\xd3\x4a\xa9\xa9\x51\x24\xeb\x14\xc6\xf3\x89\xe1\xa9\x43\x8f\x9a\x3b\x76\xa8\x97\xc7\x7b\xd0\xe4\x65\xf6\x61\xf8\x38\xfa\xfa\x32\x7e\x7c\x7a\x1e\xde\x7f\xbd\x1b\xbd\x8c\x6f\x47\xb3\xbd\x7c\x97\x42\x55\x98\x42\xef\xfa\xea\xec\x07\x21\x94\x32\xab\xe9\x56\xe9\x23\x2f\x85\x6a\xa9\x9a\x0b\xe5\x77\xf5\x22\x85\x15\x19\x29\x62\x42\xbf\x5f\x1a\x72\x67\x6c\x0a\x9f\x7b\xc3\xfb\xfb\xde\x97\x5f\x6a\xaf\x3b\x9c\xc5\xde\xc8\x57\xe4\x13\xed\x55\x51\x96\xbc\x56\x19\x2a\xe4\x64\xc7\xbd\xab\x45\x87\xde\x54\x4e\xee\x43\x6d\xe6\xba\x03\xf8\x9b\x91\x65\x20\x4d\x99\x2c\x6c\x95\xc2\x75\xd4\x85\x7f\xa4\x2f\x9e\x78\xb0\x5b\x25\xef\xf5\xac\x53\x13\x02\x00\xd7\x16\x53\xb8\x6b\x44\x6a\x5c\xfd\xe0\x6e\x1d\x0a\xc6\xe8\x17\x59\x3b\x79\xfe\x49\x0a\xd9\x28\x74\x9b\xf9\xf7\x47\x90\x3e\x8c\x35\x78\x53\x22\x48\xe1\xd1\x5f\x86\x17\x0c\x61\x4c\xf3\x50\x88\x25\x86\xc9\x7e\x97\x49\x08\x83\x05\x87\xfa\xea\x2a\x0d\x46\xab\xba\x13\x2b\x38\x87\xb2\xab\x8c\xc8\xfd\x00\x9e\x36\x27\xe2\xa6\x66\x34\x8e\x83\x4e\xde\xaf\x58\xa7\x7b\xd1\x3b\xc9\x19\x1b\x7c\x8d\x4b\x61\xb4\x26\xcf\x5d\x21\xe0\x7c\x8e\x92\x53\x98\x98\x99\x2c\x30\xaf\x14\x46\xff\x06\x00\x00\xff\xff\xac\xb1\x97\x89\x0a\x0e\x00\x00"),
		},
		"/nvidia-device-plugin.yml": &vfsgen۰CompressedFileInfo{
			name:             "nvidia-device-plugin.yml",
			modTime:          time.Date(2024, 9, 14, 6, 33, 21, 722111715, time.UTC),
			uncompressedSize: 2089,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x54\xc1\x8e\xdb\x36\x10\xbd\xeb\x2b\x06\xf6\x25\x01\x2c\x69\x37\x87\x22\x55\x4e\xee\xae\x83\x0a\x75\xec\xc0\xde\x24\x08\x8a\x22\x18\x93\x63\x89\x30\x45\xb2\xe4\xc8\x5e\xfd\x7d\x41\x59\xde\xb5\xb6\xed\xb6\xba\xd8\x12\xe7\xbd\x37\x33\x6f\x86\x53\xb8\xb3\xae\xf3\xaa\xaa\x19\xde\x88\xb7\xf0\xee\xe6\xf6\xe7\x19\xac\xbe\x96\xf7\xe5\x1c\xee\xd6\x9b\xcf\xeb\xcd\xfc\xa1\x5c\xaf\x32\x80\xb9\xd6\xd0\x07\x06\xf0\x14\xc8\x1f\x49\x66\xc9\x34\x99\xc2\x52\x09\x32\x81\x24\xb4\x46\x92\x07\xae\x09\xe6\x0e\x45\x4d\x97\x93\x19\x7c\x25\x1f\x94\x35\xf0\x2e\xbb\x81\x37\x31\x60\x32\x1c\x4d\xde\x7e\x48\xa6\xd0\xd9\x16\x1a\xec\xc0\x58\x86\x36\x10\x70\xad\x02\xec\x95\x26\xa0\x47\x41\x8e\x41\x19\x10\xb6\x71\x5a\xa1\x11\x04\x27\xc5\x75\x2f\x33\x90\x64\xc9\x14\xbe\x0f\x14\x76\xc7\xa8\x0c\x20\x08\xeb\x3a\xb0\xfb\xeb\x38\x40\xee\x13\x8e\x4f\xcd\xec\x8a\x3c\x3f\x9d\x4e\x19\xf6\xc9\x66\xd6\x57\xb9\x3e\x07\x86\x7c\x59\xde\x2d\x56\xdb\x45\xfa\x2e\xbb\xe9\x21\x5f\x8c\xa6\x10\x0b\xff\xb3\x55\x9e\x24\xec\x3a\x40\xe7\xb4\x12\xb8\xd3\x04\x1a\x4f\x60\x3d\x60\xe5\x89\x24\xb0\x8d\xf9\x9e\xbc\x62\x65\xaa\x19\x04\xbb\xe7\x13\x7a\x4a\xa6\x20\x55\x60\xaf\x76\x2d\x8f\x9a\x75\xc9\x4e\x85\x51\x80\x35\x80\x06\x26\xf3\x2d\x94\xdb\x09\xfc\x32\xdf\x96\xdb\x59\x32\x85\x6f\xe5\xc3\xaf\xeb\x2f\x0f\xf0\x6d\xbe\xd9\xcc\x57\x0f\xe5\x62\x0b\xeb\x0d\xdc\xad\x57\xf7\x65\x34\x6a\x0b\xeb\x8f\x30\x5f\x7d\x87\xdf\xca\xd5\xfd\x0c\x48\x71\x4d\x1e\xe8\xd1\xf9\x98\xbf\xf5\xa0\x62\x1b\x7b\xeb\x60\x4b\x34\x4a\x60\x6f\xcf\x09\x05\x47\x42\xed\x95\x00\x8d\xa6\x6a\xb1\x22\xa8\xec\x91\xbc\x51\xa6\x02\x47\xbe\x51\x21\x9a\x19\x00\x8d\x4c\xa6\xa0\x55\xa3\x18\xb9\xff\xf2\xb7\xa2\xb2\x04\x9d\x1a\xdc\x2f\xc0\x58\x49\xd9\xe1\x7d\xc8\x94\xcd\x8f\xb7\xc9\x41\x19\x59\xc0\xa6\x35\xac\x1a\xba\xd3\x18\x42\xd2\x10\xa3\x44\xc6\x22\x01\x30\xd8\x50\x01\xe6\xa8\xa4\xc2\xa4\x46\x23\x35\xf9\xa7\xf7\x24\x4d\xd3\x11\x37\x3a\x17\x9e\x49\xef\x91\x1a\x6b\xb6\xc4\xff\xca\x98\x4a\x3a\x2a\x41\xa9\xd3\x6d\xa5\x4c\x2a\x7b\x40\x20\x1e\xc2\x82\x43\x41\x05\x1c\xda\x1d\xa5\xa1\x0b\x4c\x4d\x12\xdb\x12\x59\x02\x69\x12\x6c\x7d\xfc\x0f\xd0\x20\x8b\x7a\x89\x3b\xd2\xe1\xfc\xe1\x75\x99\x90\x00\xb4\x4e\x22\xd3\x96\x3d\x32\x55\xdd\x19\xc5\x9d\xa3\x02\x36\x56\x6b\x65\xaa\x2f\x7d\x40\x02\xc0\xd4\x38\x8d\x4c\x83\xd4\x55\x29\xf1\xd1\x23\xd5\xff\xd6\x05\xb8\x94\x10\x1f\x7f\xd5\xf7\xd5\x15\x12\xa6\xb0\x78\x8c\xa3\xad\x58\x77\xfd\xc4\x53\xe0\xde\xd3\x01\x31\xe0\xd9\x6a\xf2\x67\xdf\x9f\x53\x48\xe1\x40\xdd\x85\x29\x13\xb6\xc9\x2b\xd7\x3e\x9d\x02\x58\x17\x31\xd6\x17\xb0\x78\x54\x81\xc3\xd5\x11\xed\xf7\x24\xb8\x80\x95\xdd\x8a\x9a\x64\xab\x2f\x42\x53\xf8\x84\xfe\x70\xbe\x14\x9c\x95\x80\x21\xae\x77\x5c\x2f\x81\x1a\x50\xca\xd4\x9a\x0f\x70\xaa\xc9\x00\x99\xb8\x8f\x72\xd6\xa7\xfb\x22\xe4\x89\x2d\x0c\xf4\xfe\x72\x8f\xf5\x17\x9a\x6d\xbd\xa0\xd0\xef\xc0\x0b\x60\x14\x0d\x10\x2c\x70\x8d\x7d\x23\x3a\x10\xf8\x4c\xb7\xa3\x08\x1f\x38\x25\xe0\x9e\xc9\x03\xc2\x1e\x95\x6e\x3d\x65\x4f\x71\x71\xdb\xe2\xad\x13\x8a\x3c\x8f\x63\xe5\x0d\x31\xf5\xab\x20\xad\x08\x39\x63\x38\x84\x1c\x65\xa3\x8c\x0a\x4c\x3e\x15\xba\x8d\xbf\x79\xd5\xa2\x47\xc3\x44\x32\x1d\x54\x94\xa9\xd2\x4b\x8e\x29\x4a\x69\x4d\x1a\x53\xcc\x07\x29\xe7\x95\xf5\x8a\xbb\x2b\x67\x27\xe7\x11\x4e\xe3\x02\x3e\x41\x27\x43\xbc\xb0\x26\xde\x99\xe4\x47\x3e\xaa\x06\xab\x7e\x26\x84\x8f\x39\x9e\x1d\xcd\x0f\xef\xc3\x78\xb2\x8a\xe3\x4d\x76\xfb\x53\x76\x7b\x65\xe4\x2b\x63\x28\xd8\x5f\x3b\x6e\x8e\xc5\xd5\x6b\x54\x3d\x63\x3f\xce\xcb\xe5\x8f\xf5\xea\x47\xb9\x2a\x1f\x7e\x2c\x36\x9b\xf5\x66\x14\x06\x70\x44\xdd\xc6\xb2\xf6\xa8\x03\x4d\xae\x0e\x03\x89\xb6\xaf\xdd\x1a\xa6\x47\x1e\xd3\xa3\xd6\xf6\xf4\xd9\xab\xa3\xd2\x54\xd1\x22\x08\xd4\xfd\xfc\x16\xd0\xf3\x8c\x62\x05\x3a\xdc\x29\xad\x58\x51\x28\x5e\xa8\x4b\x6f\x5d\x01\xbf\x4f\xe6\xcb\xe5\xe4\x8f\xab\xb3\xa3\xd5\x6d\x43\x9f\x6c\x6b\x38\xfc\x73\x61\xa3\x6e\xbc\x60\x6d\x22\xee\x33\x72\x5d\x40\x7e\x44\x9f\x6b\xb5\xeb\xe7\x44\x13\xe7\x23\xdc\x65\x69\xce\x72\x23\xd3\x5e\x57\xa9\x6d\x38\x0b\x8c\x94\xdd\xff\x92\xfc\x2b\x00\x00\xff\xff\x0d\x10\x99\xc7\x29\x08\x00\x00"),
		},
		"/nvshare-system-quotas.yaml": &vfsgen۰CompressedFileInfo{
			name:             "nvshare-system-quotas.yaml",
			modTime:          time.Date(2024, 9, 14, 6, 33, 21, 724491396, time.UTC),
			uncompressedSize: 1320,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x92\x4f\x8f\xdb\x36\x10\xc5\xef\xfc\x14\x0f\xd6\x25\x01\xd6\xde\xed\xf6\x52\xb8\x27\x75\xb3\x6d\x85\x04\x76\xba\xda\x34\x08\x8a\x00\x3b\xa6\xc6\x12\x5b\x8a\xc3\x92\x94\xb5\xfe\xf6\x05\x65\xc5\xa9\x51\xec\xa9\x97\xf0\x38\xf3\xf8\xe6\x37\x7f\x0a\xdc\x89\x3f\x06\xd3\x76\x09\xaf\xf4\x6b\xdc\xde\xdc\x7e\x8f\x5f\x58\x42\x6b\x24\xa2\xb4\xfc\x2c\x5e\x06\x2b\x51\x15\xaa\xc0\x3b\xa3\xd9\x45\x6e\x30\xb8\x86\x03\x52\xc7\x28\x3d\xe9\x8e\xbf\x64\xae\xf0\x3b\x87\x68\xc4\xe1\x76\x75\x83\x57\x59\xb0\x98\x53\x8b\xd7\x3f\xaa\x02\x47\x19\xd0\xd3\x11\x4e\x12\x86\xc8\x48\x9d\x89\xd8\x1b\xcb\xe0\x67\xcd\x3e\xc1\x38\x68\xe9\xbd\x35\xe4\x34\x63\x34\xa9\x9b\xca\xcc\x26\x2b\x55\xe0\xd3\x6c\x21\xbb\x44\xc6\x81\xa0\xc5\x1f\x21\xfb\x7f\xeb\x40\x69\x02\xce\xaf\x4b\xc9\xaf\xaf\xaf\xc7\x71\x5c\xd1\x04\xbb\x92\xd0\x5e\xdb\x93\x30\x5e\xbf\xab\xee\xee\x37\xf5\xfd\xf2\x76\x75\x33\x7d\xf9\xe0\x2c\xc7\x88\xc0\x7f\x0f\x26\x70\x83\xdd\x11\xe4\xbd\x35\x9a\x76\x96\x61\x69\x84\x04\x50\x1b\x98\x1b\x24\xc9\xbc\x63\x30\xc9\xb8\xf6\x0a\x51\xf6\x69\xa4\xc0\xaa\x40\x63\x62\x0a\x66\x37\xa4\x8b\x61\x7d\xa1\x33\xf1\x42\x20\x0e\xe4\xb0\x28\x6b\x54\xf5\x02\x3f\x95\x75\x55\x5f\xa9\x02\x1f\xab\xc7\x5f\xb7\x1f\x1e\xf1\xb1\x7c\x78\x28\x37\x8f\xd5\x7d\x8d\xed\x03\xee\xb6\x9b\x37\xd5\x63\xb5\xdd\xd4\xd8\xfe\x8c\x72\xf3\x09\x6f\xab\xcd\x9b\x2b\xb0\x49\x1d\x07\xf0\xb3\x0f\x99\x5f\x02\x4c\x1e\x23\x37\x79\x66\x35\xf3\x05\xc0\x5e\x4e\x40\xd1\xb3\x36\x7b\xa3\x61\xc9\xb5\x03\xb5\x8c\x56\x0e\x1c\x9c\x71\x2d\x3c\x87\xde\xc4\xbc\xcc\x08\x72\x8d\x2a\x60\x4d\x6f\x12\xa5\x29\xf2\x9f\xa6\x56\x2a\x23\x33\xfa\x21\x26\xe8\xc0\x94\x4e\x25\x9f\x1e\x38\xca\x10\x34\xff\x36\x48\xa2\x27\xc8\xee\x4f\xd6\xe9\x4c\xe0\xa8\xe7\xe8\x49\x73\x9e\xa4\x84\x6c\x3a\xa7\x54\x81\xb7\x3f\x44\x94\xef\x2b\x44\x0e\x87\x5c\x4d\x40\xd6\xca\x78\xb2\xcf\x57\x26\x7b\x84\xd9\x3e\x7e\x3d\x96\xa7\x78\x8c\x89\xfb\xa5\x93\x86\x97\x3a\x6f\x47\x93\x7d\x52\x45\x6e\xe3\x9c\xd4\x76\x88\x89\xc3\xd7\x3c\xde\x07\x23\xc1\xa4\xe3\x9d\xa5\x18\x39\x66\xa0\xe9\x3c\xcf\x88\x2b\xa5\xc8\x9b\xf9\xc2\xd7\x38\x7c\xa7\xfe\x32\xae\x59\xe3\xa2\x43\xd5\x73\xa2\x86\x12\xad\x15\xa6\xaf\x6b\x78\x69\xe2\xf2\x85\xb2\xb3\x68\xf2\x5f\xc3\x1d\x62\x47\x81\x67\xb1\xca\xeb\xc9\x36\x51\x8b\xe7\x9a\x2d\xeb\x24\x21\x07\x80\x9e\x92\xee\xee\x4f\xcb\xce\x0b\x39\x45\x81\x25\xc4\x73\xa0\xac\x43\xe5\xe6\xe0\xec\xb0\x99\x68\x2e\xda\x3c\x0b\x0e\x64\x07\x8e\x6b\xfc\xb1\x78\x01\x74\xf1\x59\x2d\x97\xcb\xff\x3f\x80\x8b\xa5\x7c\xb3\xdd\x5f\x50\x2e\x3e\xab\x7f\x02\x00\x00\xff\xff\xa9\xe3\x1d\x56\x28\x05\x00\x00"),
		},
		"/nvshare-system.yaml": &vfsgen۰CompressedFileInfo{
			name:             "nvshare-system.yaml",
			modTime:          time.Date(2024, 9, 14, 6, 33, 21, 724755636, time.UTC),
			uncompressedSize: 653,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x64\x91\xc1\x6e\xdb\x30\x0c\x86\xef\x7a\x8a\x1f\xf1\xa5\x05\x52\xa7\xcb\x6e\xd9\xc9\x4b\xb3\xcd\x58\xe1\x00\x71\xba\xa2\x47\xc6\x66\x6c\x62\xb2\xa4\x49\x72\x1c\xbf\xfd\xe0\x34\x05\x16\x4c\x47\xf1\x13\xf5\xf1\x67\x82\xb5\x75\xa3\x97\xa6\x8d\xb8\xab\xee\xb1\x7c\x5c\x7e\xc6\x77\xb6\xbe\x11\x1b\x90\x69\x3e\x5b\x67\x7b\x6d\x83\x4a\x54\x82\x67\xa9\xd8\x04\xae\xd1\x9b\x9a\x3d\x62\xcb\xc8\x1c\x55\x2d\x7f\x54\xe6\xf8\xc5\x3e\x88\x35\x58\xa6\x8f\xb8\x9b\x80\xd9\xb5\x34\xbb\xff\xa2\x12\x8c\xb6\x47\x47\x23\x8c\x8d\xe8\x03\x23\xb6\x12\x70\x14\xcd\xe0\x73\xc5\x2e\x42\x0c\x2a\xdb\x39\x2d\x64\x2a\xc6\x20\xb1\xbd\x7c\x73\x6d\x92\xaa\x04\x6f\xd7\x16\xf6\x10\x49\x0c\x08\x95\x75\x23\xec\xf1\x5f\x0e\x14\x2f\xc2\xd3\x69\x63\x74\xab\xc5\x62\x18\x86\x94\x2e\xb2\xa9\xf5\xcd\x42\xbf\x83\x61\xf1\x9c\xaf\x37\x45\xb9\x79\x58\xa6\x8f\x97\x27\x2f\x46\x73\x08\xf0\xfc\xa7\x17\xcf\x35\x0e\x23\xc8\x39\x2d\x15\x1d\x34\x43\xd3\x00\xeb\x41\x8d\x67\xae\x11\xed\xe4\x3b\x78\x89\x62\x9a\x39\x82\x3d\xc6\x81\x3c\xab\x04\xb5\x84\xe8\xe5\xd0\xc7\x9b\xb0\x3e\xec\x24\xdc\x00\xd6\x80\x0c\x66\x59\x89\xbc\x9c\xe1\x6b\x56\xe6\xe5\x5c\x25\x78\xcd\xf7\x3f\xb6\x2f\x7b\xbc\x66\xbb\x5d\x56\xec\xf3\x4d\x89\xed\x0e\xeb\x6d\xf1\x94\xef\xf3\x6d\x51\x62\xfb\x0d\x59\xf1\x86\x9f\x79\xf1\x34\x07\x4b\x6c\xd9\x83\xcf\xce\x4f\xfe\xd6\x43\xa6\x18\xb9\x9e\x32\x2b\x99\x6f\x04\x8e\xf6\x5d\x28\x38\xae\xe4\x28\x15\x34\x99\xa6\xa7\x86\xd1\xd8\x13\x7b\x23\xa6\x81\x63\xdf\x49\x98\x96\x19\x40\xa6\x56\x09\xb4\x74\x12\x29\x5e\x6e\xfe\x1b\x2a\x55\x8a\x9c\x5c\xd7\xbf\xc2\xe9\x93\xfa\x2d\xa6\x5e\xa1\xa0\x8e\x83\xa3\x8a\x55\xc7\x91\x6a\x8a\xb4\x52\x80\xa1\x8e\x57\x30\xa7\xd0\x92\xe7\x87\x30\x86\xc8\x9d\xfa\x1b\x00\x00\xff\xff\xa0\xb7\x47\xa7\x8d\x02\x00\x00"),
		},
		"/scheduler.yaml": &vfsgen۰CompressedFileInfo{
			name:             "scheduler.yaml",
			modTime:          time.Date(2024, 9, 14, 6, 33, 21, 725003694, time.UTC),
			uncompressedSize: 1925,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x54\xc1\x6e\xdb\x46\x10\xbd\xf3\x2b\x1e\x24\x20\x48\x00\x4b\x4a\xdc\x4b\xc1\x9e\x54\x59\x6d\x85\xba\x92\x61\x39\x0d\x82\x20\x87\xd5\x72\x48\x0e\xbc\xdc\xdd\xee\x0e\x25\x11\xc8\xc7\x17\xa4\x28\x83\x72\x1a\xfb\xd8\x3d\x09\x3b\xef\xbd\x79\x7c\x33\xab\x31\x16\xce\x37\x81\x8b\x52\xf0\x56\xbf\xc3\xf5\xfb\xeb\x9f\xf0\x3b\xb9\x50\xb0\x8b\x98\x1b\x3a\x3a\xef\x6a\xe3\x62\x32\x4e\xc6\xb8\x65\x4d\x36\x52\x86\xda\x66\x14\x20\x25\x61\xee\x95\x2e\xe9\x5c\xb9\xc2\xdf\x14\x22\x3b\x8b\xeb\xe9\x7b\xbc\x6d\x01\xa3\xbe\x34\x7a\xf7\x4b\x32\x46\xe3\x6a\x54\xaa\x81\x75\x82\x3a\x12\xa4\xe4\x88\x9c\x0d\x81\x8e\x9a\xbc\x80\x2d\xb4\xab\xbc\x61\x65\x35\xe1\xc0\x52\x76\x6d\x7a\x91\x69\x32\xc6\xe7\x5e\xc2\xed\x44\xb1\x85\x82\x76\xbe\x81\xcb\x87\x38\x28\xe9\x0c\xb7\xa7\x14\xf1\xe9\x6c\x76\x38\x1c\xa6\xaa\x33\x3b\x75\xa1\x98\x99\x13\x30\xce\x6e\x57\x8b\xe5\x7a\xbb\x9c\x5c\x4f\xdf\x77\x94\x8f\xd6\x50\x8c\x08\xf4\x4f\xcd\x81\x32\xec\x1a\x28\xef\x0d\x6b\xb5\x33\x04\xa3\x0e\x70\x01\xaa\x08\x44\x19\xc4\xb5\x7e\x0f\x81\x85\x6d\x71\x85\xe8\x72\x39\xa8\x40\xc9\x18\x19\x47\x09\xbc\xab\xe5\x22\xac\xb3\x3b\x8e\x17\x00\x67\xa1\x2c\x46\xf3\x2d\x56\xdb\x11\x7e\x9d\x6f\x57\xdb\xab\x64\x8c\x4f\xab\x87\x3f\x36\x1f\x1f\xf0\x69\x7e\x7f\x3f\x5f\x3f\xac\x96\x5b\x6c\xee\xb1\xd8\xac\x6f\x56\x0f\xab\xcd\x7a\x8b\xcd\x6f\x98\xaf\x3f\xe3\xcf\xd5\xfa\xe6\x0a\xc4\x52\x52\x00\x1d\x7d\x68\xfd\xbb\x00\x6e\x63\xa4\xac\xcd\x6c\x4b\x74\x61\x20\x77\x27\x43\xd1\x93\xe6\x9c\x35\x8c\xb2\x45\xad\x0a\x42\xe1\xf6\x14\x2c\xdb\x02\x9e\x42\xc5\xb1\x1d\x66\x84\xb2\x59\x32\x86\xe1\x8a\x45\x49\x77\xf3\xdd\x47\x4d\x93\x44\x79\xee\xc7\x9f\xb6\x99\xc5\xd9\xfe\x43\xf2\xc8\x36\x4b\x71\xa3\xa8\x72\x76\x4b\x92\x54\x24\x2a\x53\xa2\xd2\x04\xb0\xaa\xa2\x14\x76\x1f\x4b\x15\x68\x12\x75\x49\x59\x6d\x28\xf4\x95\xe8\x95\x1e\x96\x9b\x28\x54\x25\xad\xe3\x96\x1b\xc9\x90\x16\x17\xda\xdf\x40\xa5\x44\x97\xb7\x6a\x47\x26\x9e\x2e\x5e\x12\x17\xaa\xbc\x51\x42\x3d\x75\x60\xa8\x3d\xe6\x42\xe5\x25\x1d\xe0\x6c\xa6\x3d\xa1\xb6\xc2\x15\x2d\x8c\x8a\x71\xdd\x73\x38\x63\x85\x31\x96\xc7\x76\x7f\x58\x4c\xd3\xad\x15\x45\xe9\x82\xeb\x19\x3d\xdf\x07\x76\x81\xa5\x19\x08\x9c\x3e\x79\x62\x5d\x46\x13\xdd\x2e\x99\x56\xa6\x47\xb3\x65\x59\x38\xdb\xbe\x00\x0a\x03\xb7\x93\xde\x6f\x5b\x9f\x64\x1c\x9e\x0a\x00\x57\xaa\xa0\x14\xbb\x3a\x36\x3b\x77\x4c\x3f\x4c\xaf\x7f\x1e\x54\xf7\xce\xd4\x15\xfd\xe5\x6a\x2b\x03\xb9\xa1\xe4\x53\x04\x4e\x3f\x52\x27\xde\x0d\xa0\xb9\x00\x03\x55\x2b\x71\xa7\xa4\x4c\x31\xdb\xab\x30\x0b\xb5\x9d\xf5\xd4\x01\x52\xbb\xaa\x52\x36\x7b\xde\x29\x96\xcf\x2e\x26\xfa\xd9\xc5\xe8\x0b\x26\xd9\x77\xca\xb3\xa7\xb9\x4c\x5b\x7b\xf8\x8a\x37\x6f\x10\x2a\x4c\x42\xfe\x2a\xf6\xdb\x37\x48\xa8\x69\x94\x9c\x9d\xfd\x38\xd5\xff\xde\x82\xcb\x78\x1b\x21\x09\x2a\xa3\x73\xbb\xf4\x55\xd2\x5d\x6d\xcc\x9d\x33\xac\x9b\x14\xab\x7c\xed\xe4\x2e\x50\x24\x2b\x03\x5c\x24\x5d\x77\xcb\xe1\xac\xd0\x51\x2e\x63\x53\xc6\xb8\xc3\x5d\xe0\x3d\x1b\x2a\x68\x19\xb5\x32\xdd\x23\x4d\x91\x2b\x13\xe9\x02\xab\x95\x57\x3b\x36\x2c\x4c\xcf\xc6\x0c\x64\xc1\xf9\x14\x5f\x46\xf3\xdb\xdb\xd1\xd7\xff\x67\x35\x4e\xad\x5e\x4a\xfe\xc7\x1d\x4a\x17\x4f\xe2\x17\x5d\xfd\x6b\x9b\x08\x48\xe3\x29\xc5\xcd\x59\x72\x13\x16\x81\x94\x9c\x31\xe2\x0c\x85\xd3\x9f\xde\xd0\xd6\x23\x35\xe7\x17\x3e\xd5\xae\x9a\x15\xbe\x1e\x88\x3a\xdf\x72\x5c\x48\xb1\x3c\x72\x94\x38\x28\x51\x9e\x93\x96\x14\x6b\xb7\xed\x17\x22\xf9\x37\x00\x00\xff\xff\x63\x08\xd3\x9d\x85\x07\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/device-plugin.yaml"].(os.FileInfo),
		fs["/nvidia-device-plugin.yml"].(os.FileInfo),
		fs["/nvshare-system-quotas.yaml"].(os.FileInfo),
		fs["/nvshare-system.yaml"].(os.FileInfo),
		fs["/scheduler.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(io.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
