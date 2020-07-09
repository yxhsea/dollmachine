package cut_path

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

// check dir is exists
// if exists return true, else return false
func Exists(dir string) bool {
	dir = strings.Replace(dir, "\\", "/", -1)
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// get current path
func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Errorf("%+v", err)
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// current path
var CurrentPath = GetCurrentPath()

// working dir (project dir)
var WorkingDir = getWorkingPath()

func getWorkingPath() string {
	wd, err := os.Getwd()
	if err == nil {
		workingDir := filepath.ToSlash(wd) + "/"
		return workingDir
	}
	return "/"
}

// mkdir
func Mkdir(dir string) bool {
	if Exists(dir) {
		return true
	}
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Errorf("mkdir with error: %+v", err)
		return false
	}
	return true
}

// get parent path
func GetParent(dir string) string {
	dir = strings.Replace(dir, "\\", "/", -1)
	//str := wingString.WString{dir}
	lastIndex := strings.LastIndex(dir[:len(dir)-1], "/")
	return dir[:lastIndex]
}

// path format, remove the last /
func GetPath(dir string) string {
	dir = strings.Replace(dir, "\\", "/", -1)
	if dir[len(dir)-1:] == "/" {
		return dir[:len(dir)-1]
	}
	return dir
}

// delete path
func Delete(dir string) bool {
	if !Exists(dir) {
		log.Warnf("delete dir %s is not exists", dir)
		return false
	}
	return nil == os.RemoveAll(dir)
}
