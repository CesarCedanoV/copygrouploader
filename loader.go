package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "groups_to_copy.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {

			dto := map[string]interface{}{
				"existing_group_num": strings.Trim(row.Cells[0].Value, " "),
				"group_num":          strings.Trim(row.Cells[1].Value, " "),
				"group_name":         strings.Trim(row.Cells[2].Value, " "),
				"start_date":         strings.Trim(row.Cells[3].Value, " "),
				"udf":                strings.Trim(row.Cells[4].Value, " "),
			}

			group, _ := json.Marshal(dto)

			fmt.Println(string(group))

			// fmt.Printf("%s\n", group)

			// for _, cell := range row.Cells {
			// 	text := cell.String()
			// 	fmt.Printf("%s\n", text)
			// }

		}
	}
}
