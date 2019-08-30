package allchown

import (
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/pkg/errors"
)

// Change numeric uid and gid of the named filedir recursively
func Change(dir string, uid, gid int) error {
	f := genChownFunc(uid, gid)
	return process(dir, f)
}

// ChangeAs specified file
func ChangeAs(dir, file string) error {
	fileInfo, _ := os.Stat(file)
	fileSys := fileInfo.Sys()
	uid := int(fileSys.(*syscall.Stat_t).Uid)
	gid := int(fileSys.(*syscall.Stat_t).Gid)

	return Change(dir, uid, gid)
}

func process(dir string, chownF func(path string, info os.FileInfo, err error) error) error {
	return filepath.Walk(dir, chownF)
}

func genChownFunc(uid, gid int) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if err := os.Chown(path, uid, gid); err != nil {
			return errors.Wrap(err, path+" -> "+strconv.Itoa(uid)+":"+strconv.Itoa(gid))
		}
		return nil
	}
}
