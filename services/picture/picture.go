package picture

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
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

func ProfileFromUrl(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var img image.Image
	contentType := resp.Header.Get("Content-Type")
	switch contentType {
	case "image/png":
		img, err = png.Decode(resp.Body)
	case "image/jpeg":
		img, err = jpeg.Decode(resp.Body)
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
