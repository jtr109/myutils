package jzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Unzip all directories and files in the zip file in src path into the dst directory path
func Unzip(src, dst string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		readCloser, err := file.Open()
		if err != nil {
			return err
		}
		defer readCloser.Close()

		path := filepath.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return err
			}

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, readCloser)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
