package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
)

// GetMultiFormFirstValue returns the first string associated with the given key from the form data
func GetMultiFormFirstValue(form *multipart.Form, key string) string {
	val := form.Value[key]
	if len(val) == 0 {
		return ""
	}
	return val[0]
}

// GetMultiFormFirstFile returns the first file associated with the given key from the form data
func GetMultiFormFirstFile(form *multipart.Form, key string) *multipart.FileHeader {
	val := form.File[key]
	if len(val) == 0 {
		return nil
	}
	return val[0]
}

// SaveMultiFormImage saves the image contained in the fileheader
func SaveMultiFormImage(fileHeader *multipart.FileHeader, dstPath string, maxSize int64, allowedFormats map[string]struct{}) error {
	if fileHeader.Size == 0 || (maxSize > 0 && fileHeader.Size > maxSize) {
		return fmt.Errorf("image size violation. max size %d", maxSize)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer file.Close()

	buff := make([]byte, fileHeader.Size)
	if _, err = file.Read(buff); err != nil {
		return err
	}

	filetype := http.DetectContentType(buff[:512])
	if _, ok := allowedFormats[filetype]; !ok {
		allowedFormatString := ""
		i := 0
		for k := range allowedFormats {
			if i != 0 {
				allowedFormatString += ", "
			}
			if i != 0 && i == len(allowedFormats)-1 {
				allowedFormatString += "and "
			}
			allowedFormatString += k
			i++
		}
		return fmt.Errorf("only %s formats are allowed", allowedFormatString)
	}

	savedImageFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer savedImageFile.Close()

	_, err = savedImageFile.Write(buff)
	if err != nil {
		return err
	}

	return nil
}
