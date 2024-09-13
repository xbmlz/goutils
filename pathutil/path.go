// refrence from https://github.com/docker/buildx/blob/master/util/osutil/path.go

package pathutil

import (
	"os"
	"path/filepath"
)

// GetWd retrieves the current working directory.
//
// On Windows, this function will return the long path name
// version of the path.
func GetWd() string {
	wd, _ := os.Getwd()
	if lp, err := GetLongPathName(wd); err == nil {
		return lp
	}
	return wd
}

func ToAbs(path string) string {
	if !filepath.IsAbs(path) {
		path, _ = filepath.Abs(filepath.Join(GetWd(), path))
	}
	return SanitizePath(path)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func MkdirIfNotExist(path string) error {
	if Exists(path) {
		return nil
	}
	return os.MkdirAll(path, 0755)
}
