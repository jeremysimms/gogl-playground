package file

import (
	"image"
	_ "image/png" // enables PNG decoding
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type assetLocations string

const (
	base     assetLocations = "assets"
	textures assetLocations = "textures"
	shaders  assetLocations = "shaders"
)

func getFilePath(pathPrefix assetLocations, filename string) string {
	dir, error := os.Getwd()
	if error != nil {
		log.Fatal(error)
		panic(error)
	}
	return filepath.Join(dir, string(base), string(pathPrefix), filename)
}

// LoadFileAsString Loads a file as a string
func LoadFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(getFilePath(shaders, fileName))
	if err != nil {
		panic(err)
	}
	return string(data)
}

// LoadTexture loads an image of a recognized format and returns it
func LoadTexture(filename string) (image.Image, error) {
	data, err := os.Open(getFilePath(textures, filename))
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(data)
	if err != nil {
		panic(err)
	}
	return img, nil
}
