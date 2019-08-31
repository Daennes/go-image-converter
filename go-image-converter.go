package converter

import (
	"fmt"
	"log"
	"path/filepath"
)

type image struct {
	basename string
	path     string
	data     []byte
}

func ConvertImages(images []byte, format string) ([]byte, error) {
	return nil, nil
}

func ConvertImagesFromPathToPath(path string, targetPath string, format string) ([]byte, error) {
	loadImages(path)
	return nil, nil
}

func loadImages(path string) []image {
	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
	return nil
}

// // Decoders -----------
// func decodeImage(data []byte, format string) (image.Image, error) {
// 	format = strings.ToLower(format)
// 	switch format {
// 	case "image/png", "png":
// 		return decodePNG(data)

// 	case "image/jpeg", "jpeg", "jpg":
// 		return decodeJPEG(data)

// 	case "image/gif", "gif":
// 		return decodeGIF(data)

// 	case "image/bmp", "bmp", "bitmap":
// 		return decodeBMP(data)
// 	default:
// 		return nil, fmt.Errorf("Format not supported")
// 	}
// }

// func decodePNG(data []byte) (image.Image, error) {
// 	return png.Decode(bytes.NewReader(data))
// }

// func decodeJPEG(data []byte) (image.Image, error) {
// 	return jpeg.Decode(bytes.NewReader(data))
// }

// func decodeGIF(data []byte) (image.Image, error) {
// 	return gif.Decode(bytes.NewReader(data))
// }

// func decodeBMP(data []byte) (image.Image, error) {
// 	return bmp.Decode(bytes.NewReader(data))
// }

// // Encoders -----------
// func (I *Image) saveImgToPNG(path string) error {
// 	fullOutputPath := []string{path, "/", I.filename, ".png"}

// 	if _, err := os.Stat(path); os.IsNotExist(err) {
// 		os.MkdirAll(path, 0777)
// 	}

// 	img, err := decodeImage(I.data, I.imagetype)
// 	if err != nil {
// 		return err
// 	}

// 	f, err := os.OpenFile(strings.Join(fullOutputPath, ""), os.O_WRONLY|os.O_CREATE, 0777)
// 	if err != nil {

// 		return err
// 	}

// 	return png.Encode(f, img)
// }

// func (I *Image) saveImgToJPEG(path string) error {
// 	fullOutputPath := []string{path, "/", I.filename, ".jpeg"}

// 	if _, err := os.Stat(path); os.IsNotExist(err) {
// 		os.MkdirAll(path, 0777)
// 	}

// 	img, err := decodeImage(I.data, I.imagetype)
// 	if err != nil {
// 		return err
// 	}

// 	f, err := os.OpenFile(strings.Join(fullOutputPath, ""), os.O_WRONLY|os.O_CREATE, 0777)
// 	if err != nil {
// 		return err
// 	}

// 	return jpeg.Encode(f, img, nil)
// }

// func (I *Image) saveImgToGIF(path string) error {
// 	fullOutputPath := []string{path, "/", I.filename, ".gif"}

// 	if _, err := os.Stat(path); os.IsNotExist(err) {
// 		os.MkdirAll(path, 0777)
// 	}

// 	img, err := decodeImage(I.data, I.imagetype)
// 	if err != nil {
// 		return err
// 	}

// 	f, err := os.OpenFile(strings.Join(fullOutputPath, ""), os.O_WRONLY|os.O_CREATE, 0777)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	return gif.Encode(f, img, nil)
// }

// func (I *Image) saveImgToBitmap(path string) error {
// 	fullOutputPath := []string{path, "/", I.filename, ".bmp"}

// 	if _, err := os.Stat(path); os.IsNotExist(err) {
// 		os.MkdirAll(path, 0777)
// 	}

// 	img, err := decodeImage(I.data, I.imagetype)
// 	if err != nil {
// 		return err
// 	}

// 	f, err := os.OpenFile(strings.Join(fullOutputPath, ""), os.O_WRONLY|os.O_CREATE, 0777)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	return bmp.Encode(f, img)
// }
