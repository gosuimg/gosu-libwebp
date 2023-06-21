// Package main is an example implementation of WebP decoder.
package main

import (
	"github.com/gosuimg/gosu-libwebp/test/util"
	"github.com/gosuimg/gosu-libwebp/webp"
)

func main() {
	var err error

	// Read binary data
	data := util.ReadFile("cosmos.webp")

	// Decode
	options := &webp.DecoderOptions{}
	img, err := webp.DecodeRGBA(data, options)
	if err != nil {
		panic(err)
	}

	util.WritePNG(img, "encoded_cosmos.png")
}
