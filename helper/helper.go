package helper

import (
	"errors"
	"fmt"
	"strings"
)

func CheckFileExtension(filename string) (err error) {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])

	if extension != "jpeg" {
		return errors.New("jpeg only")

	}

	return nil
}

func CheckFileSize(size int64) error {
	if size == 0 {
		return fmt.Errorf("illegal file size")
	}
	if size > 1097152*2 {
		return fmt.Errorf("file size too big")
	}

	return nil
}
