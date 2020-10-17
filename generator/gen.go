// Code generated for package generator by go-bindata DO NOT EDIT. (@generated)
// sources:
// generator/template/README.md.tmpl
// generator/template/app/.gitignore
// generator/template/app/README.md.tmpl
// generator/template/app/config-overrides.js
// generator/template/app/package.json.tmpl
// generator/template/app/public/index.html
// generator/template/app/public/robots.txt
// generator/template/app/src/App.tsx.tmpl
// generator/template/app/src/index.tsx.tmpl
// generator/template/app/src/integration.tsx
// generator/template/app/src/react-app-env.d.ts
// generator/template/app/src/styles.module.less
// generator/template/app/tsconfig.json
// generator/template/go.mod.tmpl
// generator/template/integration.go.tmpl
// generator/template/integration.yaml.tmpl
// generator/template/internal/root.go.tmpl
package generator

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

// Mode return file modify time
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

var _templateReadmeMdTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcd\x41\x0a\xc2\x30\x10\x85\xe1\xfd\x9c\xe2\x41\xd6\x7a\x00\x77\x62\x37\x82\x98\x22\xf1\x00\x53\x18\xec\x60\x4c\x24\x9d\x54\xa4\xf4\xee\x42\x11\x0a\xdd\xfe\xbc\xc7\xe7\x1c\xa6\x09\xfb\xa0\x16\xe5\xc4\x83\x5c\xf9\x25\x98\x67\xb4\x9a\xde\x59\x93\xe1\x9c\x4c\x1e\x85\x4d\x73\x22\x72\xce\xc1\x8f\x52\x46\x95\x0f\x51\xf0\x8d\x3f\x20\x70\x7c\x82\xbb\x5c\x0d\xdf\x5c\x0b\x74\x7b\x38\x56\xeb\x73\x21\xda\x2d\x52\x5b\xbb\xa8\x43\x2f\xe5\x2f\x6d\xf3\xfd\x76\x59\x6b\xc3\xb6\x6c\x7e\x01\x00\x00\xff\xff\x55\xc9\x8c\xa1\xa6\x00\x00\x00")

func templateReadmeMdTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateReadmeMdTmpl,
		"template/README.md.tmpl",
	)
}

func templateReadmeMdTmpl() (*asset, error) {
	bytes, err := templateReadmeMdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/README.md.tmpl", size: 166, mode: os.FileMode(420), modTime: time.Unix(1592296240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppGitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8c\x41\x6a\x04\x21\x10\x45\xf7\x9e\x42\x98\x5d\x60\x74\x9f\x75\x6e\xd0\x07\x18\x6c\xfd\x63\x1b\xb4\x4a\xca\xb2\x21\xb7\x0f\x66\x12\x32\x30\x9b\xa2\xde\xfb\x9f\x7f\xb1\x1b\x60\x0f\xd5\x3e\xde\xbd\x3f\x50\xbb\xcb\x45\x8f\xb9\xbb\xc8\xcd\x07\xd1\x12\x2b\x86\x2f\x99\x58\x0a\xe5\xeb\xbd\x2c\xb4\x77\x16\xdb\x58\x60\xc3\xce\x53\xed\x5f\x6c\x7f\x62\x67\xcc\xc5\x26\x74\x50\x02\xc5\x82\x61\x3c\x71\xc2\xad\x71\x9a\x75\x91\xeb\xd4\xcd\x3a\xee\x73\xac\xae\x62\x68\xa1\x6c\x7c\xe4\x13\x12\x32\x96\xec\xc2\x69\x46\x2d\x4c\xc6\xef\xb3\xd4\xb4\x64\x2b\x23\x1a\xf7\xb1\xdd\x36\x65\x81\x71\xa0\xd3\x55\x8e\xa1\x3e\xde\x84\x13\x95\x7b\x03\xe9\xb3\x5e\xfb\xcf\xfc\x3f\xfd\x6b\x0d\xf5\x76\x4d\xd8\x67\x76\x95\xf3\x9b\xf9\x0a\x42\x2f\x0c\x11\x96\x07\x7f\x07\x00\x00\xff\xff\x45\xd9\x06\x74\x36\x01\x00\x00")

func templateAppGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_templateAppGitignore,
		"template/app/.gitignore",
	)
}

func templateAppGitignore() (*asset, error) {
	bytes, err := templateAppGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/.gitignore", size: 310, mode: os.FileMode(420), modTime: time.Unix(1592296240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppReadmeMdTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\x41\x0a\xc2\x30\x10\x45\xf7\x73\x8a\x0f\x59\xeb\x01\xdc\x89\xdd\x14\xc4\x16\x69\x0e\x30\x85\xc1\x0e\xc6\x44\xd2\x49\x45\x4a\xef\x2e\x54\xa1\xe0\xf6\xf1\x1e\xcf\x39\xcc\x33\xf6\x9d\x5a\x90\x13\x8f\x72\xe1\x87\x60\x59\xd0\x6a\x7c\x26\x8d\x06\x5f\xa3\x8e\x26\xb7\xcc\xa6\x29\x12\x39\xe7\xd0\x4c\x92\x27\x95\x17\x51\xd7\x54\xcd\x01\x1d\x87\x3b\xb8\x4f\xc5\xf0\x4e\x25\x43\xb7\x00\xbe\xfe\x36\xc7\x62\x43\xca\x44\xbb\xf5\xd7\x96\x3e\xe8\x38\x48\xfe\xfd\xfe\xb1\xbf\x9e\x37\x5a\xb1\xad\xce\x27\x00\x00\xff\xff\xf8\xfd\xe5\xe1\xac\x00\x00\x00")

func templateAppReadmeMdTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateAppReadmeMdTmpl,
		"template/app/README.md.tmpl",
	)
}

func templateAppReadmeMdTmpl() (*asset, error) {
	bytes, err := templateAppReadmeMdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/README.md.tmpl", size: 172, mode: os.FileMode(420), modTime: time.Unix(1592296240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppConfigOverridesJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\xc1\x0d\x42\x21\x0c\x06\xe0\xb3\x9d\xa2\xb7\x07\xc9\xd3\x05\x08\x1b\xb0\x04\xa1\xff\x81\x44\xac\xb6\x60\x8c\xc6\xdd\xbd\x99\xbc\x05\xbe\xa6\x37\x9f\xfc\x61\x7d\xc2\xac\x0b\x76\xae\x22\x05\xee\x45\xab\xc0\xf8\xcb\x99\x0d\x8f\xd5\x0d\x61\x6b\xcb\xa7\x8e\xfe\xc6\xb9\x59\xdd\x62\x22\x1a\x2a\xeb\x8a\x0b\x5e\x77\xb5\xe9\x9c\xff\x4e\xa0\xd3\x01\x0a\x71\xa7\x98\x7e\x01\x00\x00\xff\xff\xbd\x27\x54\x77\x6e\x00\x00\x00")

func templateAppConfigOverridesJsBytes() ([]byte, error) {
	return bindataRead(
		_templateAppConfigOverridesJs,
		"template/app/config-overrides.js",
	)
}

func templateAppConfigOverridesJs() (*asset, error) {
	bytes, err := templateAppConfigOverridesJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/config-overrides.js", size: 110, mode: os.FileMode(420), modTime: time.Unix(1594144320, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppPackageJsonTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x93\xcd\x6e\xdb\x30\x10\x84\xef\x7e\x0a\x42\x40\x6f\xd1\xda\x94\xdc\x04\xc9\xa1\x08\x90\x53\x81\xa2\x2f\x50\x34\x05\x2d\xae\x1d\x06\x14\x29\xec\xd2\x76\x5c\x41\x7d\xf6\x42\x94\xac\x9f\xd8\xbe\xee\x37\x33\x9c\x95\xbd\xf5\x42\x88\xc4\xa9\x12\x93\x27\x91\x3c\xd7\x35\x7c\xd7\xe8\x82\xd9\x1a\xa4\xa6\x59\xd6\x35\xfc\xf0\x47\xa4\x17\xc5\xf8\x53\x95\xd8\x34\xc9\x5d\x6b\x38\x20\xb1\xf1\xae\xf5\x48\x58\xc1\xaa\x9b\x56\x64\x0e\x2a\xb4\x49\x81\xf6\x18\x47\x1a\x2b\x74\x1a\x5d\x61\x90\x93\x27\xd1\xbe\x26\x44\xf2\x1c\x90\x83\x71\xbb\xd4\x9a\x0d\x29\x3a\x2d\xdf\x91\x43\xaa\x7d\xd9\x06\xbe\xae\x21\x83\x75\x4c\xbc\x26\x25\x54\x45\x88\xba\x47\xc8\x21\xbb\xa9\xdb\x33\x52\x8a\x07\x74\x9d\xf8\x01\xe4\x54\x7c\xaa\x90\xe3\xab\x11\x66\xeb\x61\x89\x91\x3a\xaf\xe3\x47\x79\x95\xd9\x15\x3a\xd6\x90\xf7\xf0\x78\x1d\x0f\x1b\xcd\x25\x05\x73\x6a\xbd\xd2\x48\x11\xe6\xd3\xf4\x62\xcf\xc1\x97\xe6\x2f\xa6\x05\xa9\xce\x3b\xc5\x16\x99\x7b\x93\x5c\x41\x3e\x1d\x4f\x23\xbf\x4e\x3d\xb3\xa6\x32\x07\x39\x03\xa9\xaa\xaa\x94\xf0\x68\x08\x75\xf7\x2d\x40\xc2\xfd\x5c\x32\x59\xe3\xd2\xcf\x05\x99\x2a\xc4\x56\x39\xac\x47\xca\xe1\x64\x71\x5a\x4a\x82\x1c\x1b\xc7\xaf\x14\x9d\x2d\xfb\x97\xc3\x03\x64\xc9\x42\x88\x26\xfe\x6b\xc6\xcc\xfa\x1c\xa6\x28\x2a\x2f\x3a\x8b\x0e\xf5\xb1\x9b\xbd\xb1\xfa\xba\xae\x43\xe7\xe7\xfb\x5f\xfe\x52\x16\x49\xaf\xc2\x77\x2c\x6e\xc8\x3a\x34\x14\x46\xb6\xc6\x85\x17\xef\xb6\x66\x37\xb6\xc6\x8f\x80\x4e\xf3\x2c\x61\xf4\x6c\xc8\x1f\x19\x89\xad\x89\x5d\x7a\x4f\x45\x5e\xef\x8b\xd0\x1d\xd7\xaf\x38\x13\x22\xf9\xb6\x82\xec\x4b\x5f\xab\xbd\x57\x1f\x84\x46\xa5\xe7\x13\x5f\xfd\x29\x8d\x33\x42\x59\x9b\xc4\xf9\xef\x7e\x0f\x8d\x07\xb4\xbe\x2a\xbb\x5b\x18\x42\xad\xe2\x20\xa4\x28\xde\xc8\x97\x28\xce\x27\x7d\xf7\x09\x6f\x0d\xe1\xd6\x7f\xdc\xe4\xac\xb6\x8a\xcc\x80\xbb\x87\x87\x25\xdf\x7c\x89\x95\xda\xc5\x43\x82\x65\xb2\x68\x16\xff\x03\x00\x00\xff\xff\x22\x69\x6d\xa8\x71\x04\x00\x00")

func templateAppPackageJsonTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateAppPackageJsonTmpl,
		"template/app/package.json.tmpl",
	)
}

func templateAppPackageJsonTmpl() (*asset, error) {
	bytes, err := templateAppPackageJsonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/package.json.tmpl", size: 1137, mode: os.FileMode(420), modTime: time.Unix(1594144320, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppPublicIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x90\xbf\x4e\xc4\x30\x0c\x87\xf7\x7b\x0a\x93\x99\x12\xb1\x31\x38\x5d\x80\x85\x05\x24\x58\x6e\xf4\x25\x86\x58\x4a\x9d\x2a\x75\x73\xe2\xed\x51\xaf\x07\x62\xf2\x9f\x9f\xfc\x7d\x92\xf1\xe6\xe9\xf5\xf1\xe3\xf8\xf6\x0c\xd9\xa6\x32\x1e\x70\x2b\x50\x48\xbf\x82\x63\x75\xe3\x01\x00\x33\x53\xda\x1a\x00\x9c\xd8\x08\x62\xa6\xb6\xb0\x05\xb7\xda\xe7\xf0\xe0\xc0\xff\x0f\x95\x26\x0e\xae\x0b\x9f\xe7\xda\xcc\x41\xac\x6a\xac\x16\xdc\x59\x92\xe5\x90\xb8\x4b\xe4\xe1\x32\xdc\x82\xa8\x98\x50\x19\x96\x48\x85\xc3\xfd\x15\x85\xfe\xd7\x88\xa7\x9a\xbe\xaf\x74\xad\x4b\x6c\x32\xdb\x78\xac\x2b\x28\x73\x02\xab\xc0\x4a\xa7\xc2\xf0\x42\x9d\xde\x2f\xe9\xb6\x6c\xab\x82\x65\x59\x80\xe6\xf9\x0e\xfd\xdf\xe1\xce\x49\xd2\x41\x52\x70\xad\x56\x73\x23\xfa\x24\x7d\x97\xee\x2e\xf4\xfb\x23\x7e\x02\x00\x00\xff\xff\x19\x69\x1b\xad\x19\x01\x00\x00")

func templateAppPublicIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateAppPublicIndexHtml,
		"template/app/public/index.html",
	)
}

func templateAppPublicIndexHtml() (*asset, error) {
	bytes, err := templateAppPublicIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/public/index.html", size: 281, mode: os.FileMode(420), modTime: time.Unix(1592296240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppPublicRobotsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xc8\x28\x29\x29\x28\xb6\xd2\xd7\x2f\x2f\x2f\xd7\x2b\xca\x4f\xca\x2f\x29\x2e\xa9\x28\xd1\xcb\x2f\x4a\xd7\x47\xf0\x32\x4a\x72\x73\xb8\x42\x8b\x53\x8b\x74\x13\xd3\x53\xf3\x4a\xac\x14\xb4\xb8\x5c\x32\x8b\x13\x73\x72\xf2\xcb\xad\xb8\x00\x01\x00\x00\xff\xff\x55\x9d\xd1\x4a\x43\x00\x00\x00")

func templateAppPublicRobotsTxtBytes() ([]byte, error) {
	return bindataRead(
		_templateAppPublicRobotsTxt,
		"template/app/public/robots.txt",
	)
}

func templateAppPublicRobotsTxt() (*asset, error) {
	bytes, err := templateAppPublicRobotsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/public/robots.txt", size: 67, mode: os.FileMode(420), modTime: time.Unix(1592296240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppSrcAppTsxTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\x41\x8f\xdb\x2c\x14\x3c\xc3\xaf\x78\x5a\x7d\x5a\x67\xa5\xfd\xec\x9e\xed\x38\xea\xaa\xa7\x95\xa2\x76\x95\x26\xa7\xaa\x07\xd6\x3c\xc7\x68\x31\x20\xc0\xc9\x56\x16\xff\xbd\x02\xc7\x8d\xb3\xed\x0d\x0f\x33\xc3\xbc\x31\x88\xde\x68\xeb\x61\x87\xac\xf1\xd0\x5a\xdd\x43\x66\xe3\x3a\xab\xe8\x65\x6b\x84\xef\xa2\x1f\x24\xf3\xda\x3e\x2b\xe7\x99\x94\x68\x1f\xe1\x59\x79\x3c\x5a\xe6\x85\x56\x10\x2e\xc2\xcf\x46\x28\xe3\x0b\x76\x44\xe5\xf3\x33\xbe\x3a\xfe\x76\xb5\x59\x08\x0e\xcf\x17\x41\x5e\x88\x2b\x9a\x55\x94\xb6\x83\x6a\x92\xe5\x93\x31\xab\x07\x18\x29\x29\x0a\x68\x3a\x6c\xde\xc0\x6b\x70\x88\x20\x5a\x38\x23\x30\x8b\x60\x07\xa5\x84\x3a\x82\xd4\x0d\x93\xc0\x14\x07\x85\xc8\x23\xcf\x0e\x0a\x84\x02\x37\xc5\x8e\x76\xbd\xe6\x48\x89\x68\x61\x75\x16\x8a\xeb\x33\xd4\x75\x0d\xd3\x32\x37\xcc\xa2\xf2\x70\x7f\x3f\x03\xd1\x30\xaa\xf2\xce\x62\x9b\x0b\xc5\xf1\xfd\x5b\xbb\xca\xd2\x39\x9d\x76\x3e\x7b\x80\x0d\x7c\x4a\xe9\x48\xa3\x95\xf3\xb0\x98\xa2\xbc\x69\xa6\x4e\x24\xa2\x58\x8f\x25\x64\xe3\x08\xf9\xcb\xf0\x2a\x85\xeb\xd0\x7e\x65\x3d\x42\x08\xd9\x63\x24\x70\x74\x8d\x15\x66\x32\xc8\xf6\x9d\x70\x20\x1c\xf8\x0e\x21\x6a\xf6\xc2\x4b\xfc\xc2\x1c\x5e\x34\xcb\xf3\xa0\xd5\x16\x5e\x84\x32\x5a\x28\x3f\x99\x79\x76\x74\x25\xfc\x88\x4b\x32\x8e\xff\x83\x65\xea\x88\xf0\xdf\x89\x49\x28\x6b\xc8\x17\xf9\xf6\xbf\x0c\x3a\x08\x21\x51\x49\xcc\x97\x58\x73\x2a\x80\xa4\x47\xc5\x27\xca\xcf\x84\x8a\xcb\x25\xe0\x25\xb4\x4c\x3a\x4c\xa0\xc5\x36\x9a\xa5\x21\xf3\xad\x3e\xa3\x9d\xf3\xce\x66\xa2\x49\xb3\x4d\x1f\x66\x6e\xa1\x9c\x0a\xba\x36\x74\x5b\xd0\x2c\x26\xec\xc4\x3c\xb3\x1f\x18\x4f\x09\xfc\xc3\x19\xac\xfc\x40\x38\xec\xb6\x21\x64\x71\x33\x24\xca\x20\x0e\xbb\x6d\xf9\xcf\xff\x4c\x09\x09\x15\x8d\x83\xf8\xc1\x2a\x58\xff\x7d\xe7\x97\xad\xd7\xe3\xe2\x23\x40\xb1\xa9\x28\x09\x74\x16\xaf\x28\x21\x6b\x2e\x4e\xd0\x48\xe6\x5c\x9c\xa3\xbe\x7b\x32\xe6\x6e\x13\x33\xac\x6f\x1f\x42\x11\xc1\x75\xc1\xc5\x69\x43\xc9\x43\x45\x03\xa5\xf8\x9e\x1e\x0c\xc7\x96\x0d\xd2\xc7\xa7\x50\xd1\xdf\x01\x00\x00\xff\xff\x74\x25\x9b\x65\xa7\x03\x00\x00")

func templateAppSrcAppTsxTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateAppSrcAppTsxTmpl,
		"template/app/src/App.tsx.tmpl",
	)
}

func templateAppSrcAppTsxTmpl() (*asset, error) {
	bytes, err := templateAppSrcAppTsxTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/src/App.tsx.tmpl", size: 935, mode: os.FileMode(420), modTime: time.Unix(1592411432, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppSrcIndexTsxTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x4e\xeb\x30\x10\x86\xd7\xf6\x29\xac\x6e\xdc\x4a\xef\xd9\x07\x48\xa8\x08\x85\x05\x12\x85\x0a\xb8\x40\x1a\x4f\x8b\x45\x9d\x19\x4d\xa6\x94\x2a\xca\xdd\x51\xa2\x54\xa9\x04\xbb\xb1\xbf\xdf\xdf\xef\x89\x89\x90\xc5\xbc\x42\x59\x89\xd9\x31\x26\x63\xb9\x9f\x6d\xa6\xaf\xd1\xfd\xcb\xfa\x9a\xfe\x0f\x98\xa6\x44\x6b\x0a\xa2\x15\xd6\x02\xdf\x62\xba\x31\x77\x4b\xb1\x26\xf1\xe5\x1e\x6a\x71\x27\xd8\x36\xe1\x73\x7a\x51\x10\x8d\x31\xe7\x0b\x22\x9b\x69\x7d\xa9\x71\x0c\x75\x00\x9e\x6b\x95\x0f\x57\xee\x4d\x38\x56\xb2\xc6\x00\x4b\xad\x54\x3e\x55\xb9\x0d\xe3\x57\x0c\xc0\x86\x8e\xdb\x43\x6c\x3e\x80\x6f\x66\x6d\xeb\x36\x97\xd3\x73\x99\xa0\xeb\x66\x86\x61\xf7\x7e\x26\x18\xe0\x13\x9e\x80\x57\x65\x03\x23\xec\x9d\x83\xd4\xf8\x41\xef\xff\xf0\x2f\xb5\xca\xfd\xaf\xcf\xfc\xd3\x2a\x60\x75\x4c\xfd\x7e\x7b\x90\x87\x03\xf4\xe3\xdd\xf9\x31\xcc\x2d\x23\x8a\x5d\xe8\x45\xa6\x7f\x02\x00\x00\xff\xff\x23\x1f\x05\x22\x60\x01\x00\x00")

func templateAppSrcIndexTsxTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateAppSrcIndexTsxTmpl,
		"template/app/src/index.tsx.tmpl",
	)
}

func templateAppSrcIndexTsxTmpl() (*asset, error) {
	bytes, err := templateAppSrcIndexTsxTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/src/index.tsx.tmpl", size: 352, mode: os.FileMode(420), modTime: time.Unix(1592296240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppSrcIntegrationTsx = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xb1\x4e\xc4\x30\x10\x44\x6b\xef\x57\x6c\x97\xbb\x26\xf7\x01\xbe\xa4\xa1\x4a\x43\x81\x44\x41\x69\xc5\x9b\x60\xc9\xf6\xa2\xdd\x75\x04\x8a\xf2\xef\x08\x48\xc1\x75\x33\x7a\xd2\x9b\x49\xe5\x83\xc5\xf0\x85\xc2\x6c\xb8\x08\x17\xec\xe4\x27\x77\x1e\x4e\xa4\xf6\x95\x49\x4f\xd6\xdf\xfe\x6a\x5f\x38\xb6\x4c\x7d\x26\xd5\xce\x03\xcc\x5c\xd5\x70\xaa\x46\xab\x04\x4b\x5c\x71\xc0\xcb\x15\x87\x11\x77\x70\x42\xd6\xa4\xe2\x05\x9c\xbb\xc7\xb4\xe1\x9c\x83\xea\x73\x28\x34\xec\xa7\xed\x89\xab\x85\x54\x49\x8e\xf1\x8d\x9b\x60\xfa\x27\x7a\x9d\x70\x65\x52\x7c\x27\xa1\xfb\x2d\xa6\x6d\x04\x77\xf5\x70\x78\x00\xfa\xfc\x7d\x18\x69\x09\x2d\x3f\xcc\xfb\xef\x00\x00\x00\xff\xff\xa7\x4c\x64\x80\xd8\x00\x00\x00")

func templateAppSrcIntegrationTsxBytes() ([]byte, error) {
	return bindataRead(
		_templateAppSrcIntegrationTsx,
		"template/app/src/integration.tsx",
	)
}

func templateAppSrcIntegrationTsx() (*asset, error) {
	bytes, err := templateAppSrcIntegrationTsxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/src/integration.tsx", size: 216, mode: os.FileMode(420), modTime: time.Unix(1594144320, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppSrcReactAppEnvDTs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcb\x31\xaa\x83\x50\x10\x05\xd0\xda\x59\xc5\xc5\x46\xf8\xf0\xf3\x7a\x4d\xb2\x91\x90\xe2\x31\x5e\x83\xe4\x45\x65\x66\x84\x88\xb8\xf7\x54\x49\x77\x9a\x93\x12\x12\xce\xc6\x81\xc6\x49\x89\xd8\x16\xfa\xa5\x36\x66\x8d\x7f\x57\x1b\x97\xf0\x1a\xe9\x2a\xd2\x53\x4b\x36\xe2\x35\xf7\x6b\x21\x9a\xbf\x53\xa1\x7b\x83\x5d\x2a\x9d\x27\x0f\x68\xc9\xee\xf4\x16\x3b\x6e\x4f\x6e\x2d\x3c\x6c\x9c\x1e\xf7\x2f\x70\x74\x52\xf1\xbd\xcc\x16\xe8\x39\xe4\xb5\xfc\x4e\x27\x87\x7c\x02\x00\x00\xff\xff\xc1\x76\xa2\x73\x8a\x00\x00\x00")

func templateAppSrcReactAppEnvDTsBytes() ([]byte, error) {
	return bindataRead(
		_templateAppSrcReactAppEnvDTs,
		"template/app/src/react-app-env.d.ts",
	)
}

func templateAppSrcReactAppEnvDTs() (*asset, error) {
	bytes, err := templateAppSrcReactAppEnvDTsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/src/react-app-env.d.ts", size: 138, mode: os.FileMode(420), modTime: time.Unix(1594144320, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppSrcStylesModuleLess = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xc7\xc1\x0a\xc2\x30\x0c\x06\xe0\xf3\xf2\x14\xb9\x4d\x2f\x9d\x67\x77\x19\xf8\x24\x65\xfd\xad\x01\x9b\x86\x34\x82\x20\xfa\xec\xc2\x8e\xdf\x26\xcd\xba\x07\x9f\x1c\x77\x38\x74\xc7\x99\xe7\xdf\x66\xa2\x16\xcb\x4b\xf6\xa4\x78\xc7\x12\x0f\x34\xa4\x27\xc6\x98\x57\xa2\x74\xeb\x1a\x59\x14\xce\x1f\x9a\x5a\xf6\x2a\x7a\xe5\xcb\x4a\x93\xe5\x52\x44\xeb\x81\xef\x3f\x00\x00\xff\xff\x5a\x31\x9d\xb6\x5b\x00\x00\x00")

func templateAppSrcStylesModuleLessBytes() ([]byte, error) {
	return bindataRead(
		_templateAppSrcStylesModuleLess,
		"template/app/src/styles.module.less",
	)
}

func templateAppSrcStylesModuleLess() (*asset, error) {
	bytes, err := templateAppSrcStylesModuleLessBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/src/styles.module.less", size: 91, mode: os.FileMode(420), modTime: time.Unix(1594144320, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateAppTsconfigJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\xc1\x4e\xf3\x30\x10\x84\xef\x7d\x0a\x6b\xcf\xbf\xfe\x1b\x97\x5e\x0b\x48\xad\x28\x48\x70\x44\x3d\x38\xce\xb6\x5d\x6a\xef\x46\xde\x0d\x14\xa1\xbe\x3b\xb2\x93\x02\xc9\x71\xbe\x89\x66\x36\xe3\xaf\x85\x73\x10\x24\x75\x14\x31\x3f\x75\x46\xc2\x0a\x4b\x57\xb0\x73\x60\x3e\x1f\xd0\x60\xe9\x00\xf5\x06\xfe\x0d\x30\x52\x03\x4b\xf7\x5a\x85\x73\xd0\x4a\x1a\x9d\x41\xfc\x27\xc3\xec\x9b\x88\xbf\x14\x95\xf1\x6c\x50\xe5\x6e\x4c\xf1\x31\xca\xc7\xa6\x74\x59\xee\x71\x84\x7a\xa2\xee\x81\x9a\xd5\x11\xc3\x69\xea\xa0\x6e\xa5\xed\x23\xae\xd9\x30\x4b\x37\x35\x6b\xd6\xcb\x27\xdb\x11\x8d\xc2\x2d\xee\x7d\x1f\x6d\x9d\x3a\xc9\x36\x2f\xb0\x4c\xc1\xa6\x6c\x2f\x39\xe0\x4a\x58\x49\x0d\xd9\x56\x5e\x89\x0f\x6b\xbe\xa7\x88\x8f\x3e\xe1\x2c\x21\xd5\x33\x86\x49\xea\x5f\x4d\xf8\x33\xaa\xc4\xbe\xac\x58\xbe\x60\x69\xaf\x2b\x40\x2e\xce\x3b\x6e\x54\x78\x7b\x8d\xf8\x13\x4b\x2a\xd1\x1b\xb6\x83\x37\xeb\x64\xb9\x4b\x34\xbb\xfa\x4d\xcf\xa5\x22\xa3\x0f\x75\xd9\x4b\xe1\x40\x1c\x62\xdf\xe2\xcf\xfb\x80\xe6\x50\xdc\xdd\xe2\xb2\xf8\x0e\x00\x00\xff\xff\xc9\x20\xaf\x73\xeb\x01\x00\x00")

func templateAppTsconfigJsonBytes() ([]byte, error) {
	return bindataRead(
		_templateAppTsconfigJson,
		"template/app/tsconfig.json",
	)
}

func templateAppTsconfigJson() (*asset, error) {
	bytes, err := templateAppTsconfigJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/app/tsconfig.json", size: 491, mode: os.FileMode(420), modTime: time.Unix(1592296240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateGoModTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xcd\x4f\x29\xcd\x49\x55\xa8\xae\xd6\x0b\xc8\x4e\xaf\xad\xe5\xe2\x4a\xcf\x57\x30\xd4\x33\x34\xe1\xe2\x2a\x4a\x2d\x2c\xcd\x2c\x4a\x55\xd0\xe0\xe2\x4c\xcf\x2c\xc9\x28\x4d\xd2\x4b\xce\xcf\xd5\x2f\xc8\xcc\x2b\x28\xd1\x4f\x4c\x4f\xcd\x2b\xd1\x2f\x33\x01\xe9\x73\xcf\x2c\x09\x49\x04\x69\xd5\xe4\x02\x04\x00\x00\xff\xff\xab\x0b\x56\x3e\x4d\x00\x00\x00")

func templateGoModTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateGoModTmpl,
		"template/go.mod.tmpl",
	)
}

func templateGoModTmpl() (*asset, error) {
	bytes, err := templateGoModTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/go.mod.tmpl", size: 77, mode: os.FileMode(420), modTime: time.Unix(1601638489, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateIntegrationGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xb1\x6e\xc2\x40\x0c\x86\xe7\xf8\x29\xac\x0c\x15\x19\x38\x2f\x7d\x83\xd2\x81\xa1\xd0\x21\x2f\xe0\x06\x73\xb1\x48\x9c\xe8\xe2\xa0\x4a\xa7\x7b\xf7\x0a\x5a\xa9\xb0\x5a\x9f\xff\xef\x23\xc2\xdd\x11\x0f\xc7\x16\xdf\x77\xfb\x16\xb7\x5b\xf4\x5e\x17\xd4\x05\x19\xa3\x98\x24\x76\x39\xe1\x59\x07\x81\x99\xbb\x0b\x47\xc1\x91\xd5\x00\x74\x9c\xa7\xe4\xb8\x81\xaa\xce\x39\x7c\x5e\x62\x29\xa4\xe6\x92\x8c\x87\x1a\xaa\x3a\xaa\xf7\xeb\x57\xe8\xa6\x91\x66\xb5\xd9\x89\xa3\x98\xd3\xf5\x95\xd2\x6a\x26\xa9\x86\x06\x80\x08\xf7\xe6\x12\x13\xbb\x4e\x76\xb3\xae\x8b\x9c\xd0\x27\x94\xef\xfb\xbc\xf7\x82\xfa\x4f\xc0\x95\xd3\xf3\xc7\x9f\x31\xe4\x1c\x5a\xf5\x41\xde\x78\x91\x03\x8f\x52\xca\x03\x06\x70\x5e\xad\xbb\x87\x6f\x1a\xcc\x50\xfd\x26\x84\x8f\xdb\xe1\xe5\x01\x6c\xa0\xc0\x4f\x00\x00\x00\xff\xff\xf1\x1d\x67\xea\x12\x01\x00\x00")

func templateIntegrationGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateIntegrationGoTmpl,
		"template/integration.go.tmpl",
	)
}

func templateIntegrationGoTmpl() (*asset, error) {
	bytes, err := templateIntegrationGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/integration.go.tmpl", size: 274, mode: os.FileMode(420), modTime: time.Unix(1601385547, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateIntegrationYamlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xb1\x6e\x02\x31\x10\x44\x7b\x7f\xc5\xc8\x4a\xcb\x15\x29\x2d\xa5\x88\x68\x52\x44\x49\x0a\x7a\xb4\x70\x0b\xb7\xc2\x67\x9f\xec\x3d\x50\x64\xf9\xdf\x23\x1f\x10\x41\xe7\x9d\x37\xb3\x63\x6d\xa0\x91\x1d\x6c\x29\xdd\x46\xd4\xf3\x9a\x32\x7f\xd1\xc8\xb5\x5a\x93\xf8\xb0\xd5\xdf\xe9\x86\x3f\xe3\x85\xd3\x23\xee\x39\xef\x93\x4c\x2a\x31\x38\xd8\xcd\x20\x19\x92\xa1\x03\xa3\x14\x3c\x6f\x43\xad\x90\xa0\x7c\x4c\xd4\xec\x38\xc4\x84\x1f\x09\x53\x94\xa0\xd6\xd0\x99\x94\xd2\x76\x4e\xde\xc1\x5a\xb3\xa7\x89\x76\xe2\x45\x85\xb3\x33\xa5\xac\x90\x28\x1c\x19\x2f\x67\xf2\x70\x6f\xe8\xd6\x0f\x06\xd4\x6a\x80\x55\xab\x5c\x78\xad\x4b\x82\x43\x5f\xab\x91\x90\x95\xbc\x5f\x3a\x9d\x01\xc6\xd8\xb7\x95\x40\x4b\xec\x7d\x9c\xfb\xdb\x3b\xb3\x3f\x8c\x14\xe8\xc8\x4d\x59\xc8\xd5\x16\x58\x2f\x31\x9d\xae\x03\x30\xc4\xac\xed\x60\xf9\x2e\xb4\xb0\xb5\xcb\xf0\xfc\xed\x3b\xbc\xf0\xae\xfb\x88\xf1\xf4\x2f\xd0\xac\x43\xf7\xfd\x3e\xeb\xf0\x6a\xfe\x02\x00\x00\xff\xff\xd3\xa7\x76\x0a\x7c\x01\x00\x00")

func templateIntegrationYamlTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateIntegrationYamlTmpl,
		"template/integration.yaml.tmpl",
	)
}

func templateIntegrationYamlTmpl() (*asset, error) {
	bytes, err := templateIntegrationYamlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/integration.yaml.tmpl", size: 380, mode: os.FileMode(420), modTime: time.Unix(1592493632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateInternalRootGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4f\x6f\xe3\xb6\x13\x3d\x4b\x9f\x62\xa0\x93\xb4\xf0\xcf\xbe\xfc\x4e\x0b\xf8\xb0\x48\x0a\x34\xc0\xee\xa2\xd8\x04\xdd\x43\x51\x04\xb4\x34\x92\xd8\x50\x24\x97\x1c\xda\x1b\x04\xf9\xee\x05\xff\xd9\x72\xe2\xd4\x71\x81\xe6\x24\x8e\x39\xf3\xe6\xbd\x19\x0e\x19\xcd\xda\x07\x36\x20\x70\x49\x68\x24\x13\x65\xc9\x27\xad\x0c\x41\x5d\x16\xd5\xc0\x69\x74\x9b\x65\xab\xa6\x95\xe6\x52\xd3\x8a\x0d\x28\x69\xb5\xfd\xff\xca\x76\x0f\x55\xd9\x94\xe5\x6a\x05\x4f\x4f\xb0\xbc\xe3\x24\xf0\x8a\x59\xfc\xca\x26\x84\xe7\xe7\x1b\x49\x38\x18\x46\x5c\x49\xe0\x16\x98\x0c\xf1\xb3\xa5\x57\xe6\xa4\x57\x49\x8f\x1a\xcf\xc6\xb3\x64\x5c\x4b\xf0\x54\x16\x42\x0d\x03\x1a\x00\xdb\x3d\x2c\x3f\x87\xef\xb2\x68\x95\xec\xf9\x10\x6d\x57\xe1\xbb\x2c\x26\x26\x99\xdf\xe8\x6d\x5f\xe2\x77\x59\x18\xec\xef\x3c\x9c\x25\xc3\xe5\x50\x3e\x97\xe5\x96\x19\xb8\x0f\x9b\xe6\x70\x6b\xa8\x3f\x9c\x49\xa9\xa9\x25\x17\x51\x8c\x5b\x62\x86\x3c\xe5\x96\x09\x81\x1d\xec\x46\x94\x40\x23\x1e\xf1\xe7\x16\xac\xdf\xc7\xe5\x00\x4e\x97\xbd\x93\x2d\xd4\x03\x9c\x85\x89\xd1\xeb\x44\xfb\xc0\x7a\x01\x89\xf5\x81\xf4\x02\x4e\x90\x6e\x00\x8d\xf1\xda\x97\xc5\xb0\x4c\x51\xd6\x39\xce\x77\x4e\x63\x0a\xbd\x80\x4a\x3f\x0c\xd5\x02\x2a\x9f\xd1\x67\xb5\x43\x33\xcb\xa8\x6a\xbc\x7b\x42\x5c\x27\x68\x6f\xca\x88\xeb\x8c\xed\x8d\x59\xe6\xf5\x1b\xc1\xca\x22\xe1\xdf\xc8\x5e\xd5\x39\xad\x05\x54\x59\x21\x0f\x67\x90\x9c\x91\x20\xb9\xf0\x75\x5a\xad\xe0\x17\x69\x94\x10\x2f\x75\x66\x20\x71\x77\xac\xb4\xb4\xc4\x64\x8b\xa1\x09\xbb\x0e\xbb\x0b\xc4\x8e\x18\xf5\x3e\x44\x6c\x8c\xb8\x98\x29\xf9\x46\xfa\x18\x33\x94\x8a\x80\x4f\x5a\xe0\x84\x92\xb0\x3b\x49\xe6\x9a\xdb\x89\x5b\xfb\x8a\x8d\x04\xfc\xc9\x6d\xe8\x92\xb7\x38\x19\x9c\xd4\xf6\x22\x56\x09\xec\xdf\xd2\xea\x52\xae\xef\xe1\xf5\x1d\x37\xbf\x2a\xf5\xf0\xba\x4a\x3b\xdc\x8c\xe9\x07\x83\x2d\xf2\x2d\x76\xa0\x24\x6c\x70\x64\xa2\x07\xd5\xbf\x3c\x30\x17\xd0\x4b\x98\x75\x86\xf0\x34\x92\xed\x3c\xb9\xec\xf4\x1e\x72\x5f\x1c\xed\x0f\xf3\x31\xbb\x29\xff\x62\xf0\x87\x43\x4b\xff\x05\xcd\x8c\x5e\xef\xc1\xc2\x31\x4f\x8b\x06\xea\x0f\xf3\xf5\x37\xb4\x5a\x49\x8b\x8b\xa8\x40\xe3\x25\x38\x10\x5a\xcc\x58\xdd\x92\xd2\xef\x99\x5e\xa3\xa3\xd0\x97\x9d\xda\xc5\x71\xde\x0a\x64\xf2\xc2\x59\xa6\x74\x7d\xbe\x26\x96\x94\xd6\x6f\x8e\x81\x9f\xe1\xaa\x3a\x24\x4c\x0a\x08\x85\x78\x95\x33\x29\x30\x2e\x1d\x29\xef\x72\xc9\x18\x08\x0e\x75\xf4\x0b\x32\x47\xcb\x2c\xf5\x34\x4a\x3f\x1e\xcf\xd2\x19\x89\xd6\x59\x52\x13\x9a\x7b\xde\x55\x8b\x94\xc2\xf2\x2a\x19\x6f\xae\xeb\xa6\x29\x8b\x30\xef\xb0\xf3\x51\x88\x4f\xb8\xfc\xaa\x76\x75\x53\x1e\xeb\x72\x98\x2e\x29\x9b\xe8\x53\xf9\x7d\xab\x15\xfc\xc6\x35\xc2\xe4\x2c\xc1\x06\x67\x8a\x6c\x70\xe0\x33\xee\xc0\x64\x97\x1b\x12\x18\x68\xef\xe4\x4b\x68\x51\x76\xa1\xa6\x8c\x58\x59\x04\xf3\xc7\x75\x4e\xd6\x87\xae\x13\xcc\x2d\x31\x8a\x13\x15\x32\x31\xb0\x1a\x5b\xde\xf3\xd6\xa7\x44\x08\x6a\xf3\x17\xb6\x14\xc2\xd2\xc8\xed\x51\x2d\x3c\x7c\xf6\x0b\xb4\x69\x0e\x14\x82\x67\xa4\x83\x44\xb0\xe3\x42\x40\x6a\x00\x5f\xdd\x3d\x32\xef\x12\x0c\xe6\xda\x16\xed\xc1\xed\x10\x78\x2e\x77\x8a\x1e\xaf\xb0\xf0\x3e\x79\x3c\x41\x25\xde\x6c\x6e\xf6\x68\x09\x5c\x0e\xb9\xa7\x3b\x70\x86\x11\x0c\xf3\xb2\x5d\xe3\xc6\x0d\xa7\xeb\x16\x9b\x3a\x64\x12\xd6\xb9\x3c\x23\x1a\x0c\xd6\xff\xcd\xff\x82\xe5\x87\xe3\x84\x91\xbe\x9a\x34\x17\x68\x3e\x96\xc5\x3d\xac\x43\x11\xe3\x57\xd0\x33\x7e\x1e\x64\x48\xeb\x78\x53\x9f\xe9\xa9\x9e\x4b\x6e\x47\xf4\x8d\x5a\x75\x89\x7c\xb5\x88\x3d\x79\xcb\x65\x8b\x75\xea\xba\xe6\xd4\x89\xfc\xe4\x48\x5d\x25\xdd\xf0\xf5\x6c\x6c\x85\x72\xdd\x51\x37\x8c\xcc\xe6\x41\x99\xbb\x15\x98\x23\xb5\x57\xff\xa2\xdb\xed\x08\xbe\xf6\x71\x66\x2f\xa3\xa3\x1f\xf3\x8c\xcc\xcf\xa5\x73\x93\xf1\x77\x26\x78\x97\xfa\x3e\x71\xda\x60\xaf\x0c\xbe\xe7\xd5\xe1\x79\x75\x48\x68\x26\x2e\xd1\x47\xe3\x7d\xaa\x62\xee\xc0\xad\x0f\x1f\x8e\xc6\xcb\xd1\xd5\x32\x09\xda\x28\x8d\x46\x3c\xfa\xb2\x4f\x4e\xf2\xd6\x27\xb2\xe3\x34\xfa\xdd\x3e\x9e\x55\xce\xf8\xbb\xfc\xd1\x12\x4e\x4b\xb8\x1b\x11\x0c\x5a\x27\x68\x1f\x32\x0e\xab\x70\x8a\x36\x8a\x46\x2f\x73\x87\x82\x6f\xd1\xa4\xb1\x39\x22\x7c\xd2\x7a\xe9\xa3\x7d\x0b\x0a\xf8\x59\xc0\xbc\x02\xc9\xd7\x13\x57\xd2\xf2\x2e\xb8\x30\xb0\xae\x6d\xd1\xda\xde\x89\x98\x7d\x48\x77\xf9\xfe\x62\x65\x45\xeb\x6d\x96\xd6\xd7\x23\x5b\x1b\xa8\x13\x83\x89\xe9\x3f\xe2\x23\xfd\xcf\xf0\x4f\x4a\xcf\x5a\x7c\x7a\x0e\x15\xfb\xc7\xaa\xfd\x1d\x00\x00\xff\xff\x93\x5f\xd7\x10\xda\x0c\x00\x00")

func templateInternalRootGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateInternalRootGoTmpl,
		"template/internal/root.go.tmpl",
	)
}

func templateInternalRootGoTmpl() (*asset, error) {
	bytes, err := templateInternalRootGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/internal/root.go.tmpl", size: 3290, mode: os.FileMode(420), modTime: time.Unix(1601638984, 0)}
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
	"template/README.md.tmpl":             templateReadmeMdTmpl,
	"template/app/.gitignore":             templateAppGitignore,
	"template/app/README.md.tmpl":         templateAppReadmeMdTmpl,
	"template/app/config-overrides.js":    templateAppConfigOverridesJs,
	"template/app/package.json.tmpl":      templateAppPackageJsonTmpl,
	"template/app/public/index.html":      templateAppPublicIndexHtml,
	"template/app/public/robots.txt":      templateAppPublicRobotsTxt,
	"template/app/src/App.tsx.tmpl":       templateAppSrcAppTsxTmpl,
	"template/app/src/index.tsx.tmpl":     templateAppSrcIndexTsxTmpl,
	"template/app/src/integration.tsx":    templateAppSrcIntegrationTsx,
	"template/app/src/react-app-env.d.ts": templateAppSrcReactAppEnvDTs,
	"template/app/src/styles.module.less": templateAppSrcStylesModuleLess,
	"template/app/tsconfig.json":          templateAppTsconfigJson,
	"template/go.mod.tmpl":                templateGoModTmpl,
	"template/integration.go.tmpl":        templateIntegrationGoTmpl,
	"template/integration.yaml.tmpl":      templateIntegrationYamlTmpl,
	"template/internal/root.go.tmpl":      templateInternalRootGoTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"README.md.tmpl": &bintree{templateReadmeMdTmpl, map[string]*bintree{}},
		"app": &bintree{nil, map[string]*bintree{
			".gitignore":          &bintree{templateAppGitignore, map[string]*bintree{}},
			"README.md.tmpl":      &bintree{templateAppReadmeMdTmpl, map[string]*bintree{}},
			"config-overrides.js": &bintree{templateAppConfigOverridesJs, map[string]*bintree{}},
			"package.json.tmpl":   &bintree{templateAppPackageJsonTmpl, map[string]*bintree{}},
			"public": &bintree{nil, map[string]*bintree{
				"index.html": &bintree{templateAppPublicIndexHtml, map[string]*bintree{}},
				"robots.txt": &bintree{templateAppPublicRobotsTxt, map[string]*bintree{}},
			}},
			"src": &bintree{nil, map[string]*bintree{
				"App.tsx.tmpl":       &bintree{templateAppSrcAppTsxTmpl, map[string]*bintree{}},
				"index.tsx.tmpl":     &bintree{templateAppSrcIndexTsxTmpl, map[string]*bintree{}},
				"integration.tsx":    &bintree{templateAppSrcIntegrationTsx, map[string]*bintree{}},
				"react-app-env.d.ts": &bintree{templateAppSrcReactAppEnvDTs, map[string]*bintree{}},
				"styles.module.less": &bintree{templateAppSrcStylesModuleLess, map[string]*bintree{}},
			}},
			"tsconfig.json": &bintree{templateAppTsconfigJson, map[string]*bintree{}},
		}},
		"go.mod.tmpl":           &bintree{templateGoModTmpl, map[string]*bintree{}},
		"integration.go.tmpl":   &bintree{templateIntegrationGoTmpl, map[string]*bintree{}},
		"integration.yaml.tmpl": &bintree{templateIntegrationYamlTmpl, map[string]*bintree{}},
		"internal": &bintree{nil, map[string]*bintree{
			"root.go.tmpl": &bintree{templateInternalRootGoTmpl, map[string]*bintree{}},
		}},
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
