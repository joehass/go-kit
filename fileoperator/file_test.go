package fileoperator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type StoneSummary struct {
	ActID       string  `file_title:"活动id"`
	GameUid     int64   `file_title:"-"`
	Statues     int     `file_title:"statues（禁止状态）"`
	FinalReward float64 `file_title:"final-reward"`
	IsFlag      bool    `file_title:"标识"`
}

var datas = []interface{}{
	&StoneSummary{
		ActID:       "a",
		GameUid:     1,
		Statues:     2,
		FinalReward: 3.1,
		IsFlag:      false,
	},
	&StoneSummary{
		ActID:       "b",
		GameUid:     1,
		Statues:     2,
		FinalReward: 3.2,
		IsFlag:      true,
	},
}

func TestR(t *testing.T) {
	val := getContentList(datas)
	fmt.Println(val)
	s := &StoneSummary{}
	titles := getTitle(s)
	fmt.Println(titles)
}

func TestWriteFile(t *testing.T) {
	fileUrl,err := WriteFile(&FileOperator{
		FilePath: "",
		FileName: "",
		Model:    StoneSummary{},
		Data:     datas,
		FileType: FileTypeCSV,
	})
	assert.Nil(t,err)
	fmt.Println(fileUrl)
}

func TestWriteFileWithZip(t *testing.T) {
	fileUrl,err := WriteFile(&FileOperator{
		FilePath: "",
		FileName: "111",
		Model:    StoneSummary{},
		Data:     datas,
		FileType: FileTypeCSV,
	},WithZip())
	assert.Nil(t,err)
	fmt.Println(fileUrl)
}