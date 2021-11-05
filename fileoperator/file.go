package fileoperator

import (
	"encoding/csv"
	"fmt"
	"github.com/go-kit/fileoperator/compress"
	"github.com/go-kit/fileoperator/util"
	"log"
	"os"
	"reflect"
)

type FileOperator struct {
	FilePath string
	FileName string
	Model interface{}
	Data []interface{}
	FileType FileType
}

const (
	excelTag = "file_title"
	skipTag  = "-"
)

type FileType int

const (
	FileTypeCSV   = 0
	FileTypeExcel = 1
)

// WriteFile 写入文件
func WriteFile(fileOperator *FileOperator,opts ...Option) (string,error) {
	var (
		filePath = fileOperator.FilePath
		fileName = fileOperator.FileName
		fileType = fileOperator.FileType
		model = fileOperator.Model
		data = fileOperator.Data
		opt = newOption(opts...)

	)
	if filePath == "" {
		filePath = util.GetDefaultFilePath()
	}

	if fileName == "" {
		fileName = util.GetDefaultFileName()
	}

	if err := util.CreateTmpDir(filePath);err!=nil{
		return "",err
	}

	tmpUrl := getFilePath(filePath, fileName, fileType)

	f, err := os.OpenFile(tmpUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return "",err
	}
	defer f.Close()
	_, err = f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	if err != nil {
		return "",err
	}
	w := csv.NewWriter(f)
	title := getTitle(model)
	err = w.Write(title)
	if err != nil {
		return "",err
	}
	w.Flush()

	rowDatas := getContentList(data)
	err = w.WriteAll(rowDatas)
	if err != nil {
		return "",err
	}
	w.Flush()

	if opt.IsZip{
		//结束后立刻删除
		defer func() {
			err := os.RemoveAll(tmpUrl)
			if err != nil {
				log.Printf("clear tmp zip fail,path:%s \n",tmpUrl)
			}
		}()
		zipUrl,err := compress.Write2TmpZip(filePath,fileName)
		if err != nil {
			return "",err
		}

		return zipUrl,nil
	}

	return tmpUrl,nil
}

func getFilePath(filePath string, fileName string, fileType FileType) string {
	switch fileType {
	case FileTypeCSV:
		return fmt.Sprintf("%s/%s.csv", filePath, fileName)
	case FileTypeExcel:
		return fmt.Sprintf("%s/%s.xlsx", filePath, fileName)
	default:
		return fmt.Sprintf("%s/%s.csv", filePath, fileName)
	}
}

func getTitle(data interface{}) []string {
	_, titles := getRowAndTagData(data)
	return titles
}

func getContentList(data []interface{}) [][]string {
	list := make([][]string, 0, len(data))
	for _, v := range data {
		rowData, _ := getRowAndTagData(v)
		list = append(list, rowData)
	}

	return list
}

func getRowAndTagData(val interface{}) ([]string, []string) {
	rowData := make([]string, 0)
	tagData := make([]string, 0)
	valV := reflect.ValueOf(val)
	if valV.Kind() == reflect.Ptr {
		valV = valV.Elem()
	}
	valT := reflect.TypeOf(val)
	if valT.Kind() == reflect.Ptr {
		valT = valT.Elem()
	}
	for i := 0; i < valT.NumField(); i++ {
		tag := valT.Field(i).Tag.Get(excelTag)
		if tag == skipTag {
			continue
		}
		fieldVal := valV.Field(i).Interface()
		fVal := fmt.Sprintf("%v", fieldVal)
		rowData = append(rowData, fVal)
		tagData = append(tagData, tag)
	}

	return rowData, tagData
}
