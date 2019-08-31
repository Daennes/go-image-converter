package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
)

func ConvertImages(images []byte, format string) ([]byte, error) {
	return nil, nil
}

func ConvertImagesFromPathToPath(path string, targetPath string, format string) ([]byte, error) {
	imagePaths, err := listImages(path)
	if err != nil {
		return nil, err
	}
	splittedPath := strings.Split(path, string(os.PathSeparator))
	srcDirName := splittedPath[len(splittedPath)-1]
	fmt.Println(srcDirName)
	for _, imagePath := range imagePaths {
		img, _, err := readFile(imagePath)
		if err != nil {
			fmt.Println(err)
			continue
		}

		imagePathWithoutExt := strings.TrimSuffix(imagePath, filepath.Ext(imagePath))
		splittedImagePath := strings.Split(imagePathWithoutExt, string(os.PathSeparator))

		filePathFromSrcDir := make([]string, 0)
		foundBase := false
		for i := 0; i < len(splittedImagePath); i++ {
			if foundBase {
				filePathFromSrcDir = append(filePathFromSrcDir, splittedImagePath[i])
			}
			if splittedImagePath[i] == srcDirName {
				foundBase = true
			}
		}

		newImagePath := filepath.Join(targetPath, filepath.Join(filePathFromSrcDir...)) + "." + format
		err = saveImageToFile(img, newImagePath, format)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Saved succesfuly: ", filepath.Join(filePathFromSrcDir...))
	}
	return nil, nil
}

func listImages(path string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(path, "**/*.jpg"))
	if err != nil {
		return nil, err
	}
	fmt.Println(len(files))
	return files, nil
}

func readFile(path string) (image.Image, string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer reader.Close()
	return image.Decode(reader)
}

func saveImageToFile(data image.Image, path string, tagetFormat string) error {
	basePath := filepath.Dir(path)
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		os.MkdirAll(basePath, 0777)
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	defer f.Close()

	switch tagetFormat {
	case "png":
		err = png.Encode(f, data)
	case "jpeg":
		err = jpeg.Encode(f, data, nil)
	case "gif":
		err = gif.Encode(f, data, nil)
	case "bmp":
		err = bmp.Encode(f, data)
	default:
		return fmt.Errorf("Image format not supported")
	}

	return err
}
