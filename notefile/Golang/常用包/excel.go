// https://github.com/qax-os/excelize

package excel

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

func Write() {
	f := excelize.NewFile()
	defer f.Close()

	sh := "Sheet1" //默认的首个sheet，如果设为其它，则首个sheet为空的Sheet1
	index, err := f.NewSheet(sh)
	if err != nil {
		log.Fatal(err)
	}
	f.SetActiveSheet(index)

	f.SetColWidth(sh, "A", "A", 30) //单列设置宽度180px
	f.SetColWidth(sh, "B", "C", 20) //两列设为相同宽度120px
	f.SetRowHeight(sh, 1, 20)       //首行高度设为20px
	f.SetCellStr(sh, "A1", "aaa")
	f.SetCellStr(sh, "B1", "bbb")
	f.SetCellStr(sh, "C1", "ccc")

	for i := int64(2); i < 12; i++ {
		line := strconv.FormatInt(i, 10)
		f.SetCellFloat(sh, "A"+line, 3.141592653, 5, 64) //保留5位小数
		f.SetCellInt(sh, "B"+line, i)
		f.SetCellBool(sh, "C"+line, i&1 == 0)
	}

	f.SetSheetName(sh, "工作表") //重命名sheet

	if err = f.SaveAs("local.xlsx"); err != nil {
		log.Fatal(err) //保存到本地
	}
	buf, err := f.WriteToBuffer()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(buf.Len()) //buf可上传到storage
}

func Read() {
	f, err := excelize.OpenFile("local.xlsx") //也可以通过OpenReader从内存中读取
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sh := f.GetSheetList()[0]
	a2, _ := f.GetCellValue(sh, "A2") //读取指定单元格
	log.Println(a2)

	rows, _ := f.GetRows(sh) //读取全部行列
	for i, row := range rows {
		log.Println(i, row)
	}
}
