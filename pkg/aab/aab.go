package aab

import (
	"io"

	"github.com/xmxu/aab-parser"
	"github.com/xmxu/aab-parser/pb"
)

func Parse(readerAt io.ReaderAt, size int64) (*AAB, error) {
	pkg, err := aab.OpenZipReader(readerAt, size)
	if err != nil {
		return nil, err
	}
	defer pkg.Close()

	icon, err := pkg.Icon(&pb.Configuration{
		Density: 640,
	})
	if err != nil {
		return nil, err
	}

	return &AAB{
		icon:     icon,
		manifest: *pkg.Manifest(),
		size:     size,
		label:    pkg.Label(nil),
	}, nil
}
