package aab

import (
	"fmt"
	"image"

	"github.com/xmxu/aab-parser"
)

type AAB struct {
	manifest aab.Manifest
	icon     image.Image
	label    string
	size     int64
}

func (a *AAB) Name() string {
	return a.label
}

func (a *AAB) Version() string {
	return a.manifest.VersionName
}

func (a *AAB) Identifier() string {
	return a.manifest.Package
}

func (a *AAB) Build() string {
	return fmt.Sprintf("%v", a.manifest.VersionCode)
}

func (a *AAB) Channel() string {
	return ""
}

func (a *AAB) Icon() image.Image {
	return a.icon
}

func (a *AAB) Size() int64 {
	return a.size
}
