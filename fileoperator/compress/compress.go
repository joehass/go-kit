package compress

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func compress(file *os.File, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			filesInfo:=strings.Split(fi.Name(),".")
			if len(filesInfo)>=2 && filesInfo[1] == "zip"{
				continue
			}
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, zw)
			if err != nil {
				f.Close()
				return err
			}
			f.Close()
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

