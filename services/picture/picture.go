package picture

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/nfnt/resize"
	"github.com/spf13/viper"
)

var (
	profileOutput string
)

func Configure() {
	profileOutput = viper.GetString("profile_pictures")
	if err := os.MkdirAll(profileOutput, 0755); err != nil {
		log.Fatal(err)
	}
}

func ProfilePictureFromReader(r io.Reader, id, contentType string) error {
	var img image.Image
	var err error

	switch contentType {
	case "image/png":
		img, err = png.Decode(r)
	case "image/jpeg":
		img, err = jpeg.Decode(r)
	default:
		return errors.New(fmt.Sprintf("unsupported image format '%s'", contentType))
	}
	if err != nil {
		return err
	}

	m := resize.Resize(400, 400, img, resize.Bicubic)
	destPath := filepath.Join(profileOutput, id+".jpg")
	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()
	return jpeg.Encode(out, m, nil)
}

func ProfileFromUrl(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return ProfilePictureFromReader(resp.Body, id, resp.Header.Get("Content-Type"))
}
