// https://github.com/qax-os/excelize

package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func Demo() {
	f := excelize.NewFile()
	sh := "Sheet1" //S小写会出错
	index := f.NewSheet(sh)
	f.SetActiveSheet(index)
	f.SetCellStr(sh, "A1", "int")
	f.SetCellStr(sh, "B1", "float")
	f.SetCellStr(sh, "C1", "bool")
	for i := 0; i < 10; i++ {
		line := strconv.Itoa(i + 2)
		f.SetCellInt(sh, "A"+line, i)
		f.SetCellFloat(sh, "B"+line, 3.14159265, 2, 64)
		f.SetCellBool(sh, "C"+line, i&1 == 0)
	}

	sh2 := "Sheet2"
	index2 := f.NewSheet(sh2)
	f.SetActiveSheet(index2)
	f.SetColWidth(sh2, "A", "A", 20) //单列设置宽度
	f.SetColWidth(sh2, "B", "C", 30) //两列都设为相同宽度
	f.SetCellStr(sh2, "A1", "One")
	f.SetCellStr(sh2, "B1", "Two")
	f.SetCellStr(sh2, "C1", "Three")

	//if err := f.SaveAs("demo.xlsx"); err != nil {
	//	fmt.Println(err) //保存到本地
	//}

	buf, err := f.WriteToBuffer()
	if err != nil {
		fmt.Println(err)
		return
	}
	// buf直接从第三方sdk上传
	fmt.Println(buf.Len())
}
