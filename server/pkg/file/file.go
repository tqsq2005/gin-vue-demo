package file

import (
	"fmt"
	"os"
)
// CheckNotExist check if the file exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist {
		return MkDir(src)
	}
	return nil
}

// MkDir create a directory
func MkDir(src string) error {
	return os.MkdirAll(src, os.ModePerm)
}

// MustOpen maximize trying to open the file
func MustOpen(filename, filepath string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("os.Getwd() err: %v", err)
	}

	src := dir + filepath
	//校验权限
	perm := CheckPermission(src)
	if perm {
		return "", fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	err = IsNotExistMkDir(src)
	if err != nil {
		return "", fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	/*openFileName := src + filename

	_, err = os.OpenFile(openFileName, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0644)
	if err != nil {
		return "", fmt.Errorf("os.Open err:%v", err)
	}
	return openFileName, nil*/
	return src + filename, nil
}
