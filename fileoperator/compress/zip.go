package compress

import (
	"archive/zip"
	"fmt"
	"github.com/go-kit/fileoperator/util"
	"os"
)


// Write2TmpZip 将数据写入压缩文件
func Write2TmpZip(path string,fileName string) (string, error) {

	if err:= util.CreateTmpDir(path);err != nil {
		return "", err
	}

	////结束后立刻删除
	//defer func() {
	//	err := os.RemoveAll(path)
	//	if err != nil {
	//		log.Error(ctx, "clear tmp zip fail",
	//			zap.String("path", path),
	//			zap.Error(err))
	//	}
	//}()

	tmpZip := fmt.Sprintf("%s%s.zip", path,fileName)

	if err := dir2zip(path, tmpZip);err != nil {
		return "", err
	}

	return tmpZip, nil

}

// dir2zip zip压缩
//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func dir2zip(src string, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()

	file, err := os.Open(src)
	if err != nil {
		return err
	}
	err = compress(file, w)
	if err != nil {
		return err
	}

	return nil
}
