package components

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io/ioutil"
)

var (
	inter font.Face
)

func FontInter() font.Face {
	return inter
}

func init() {
	fontBytes, err := ioutil.ReadFile("assets/font/Inter-Bold.ttf")
	if err != nil {
		panic(err)
	}
	fontParsed, err := opentype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	face, err := opentype.NewFace(fontParsed, &opentype.FaceOptions{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	inter = face
}
