package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	currentToken := "Bearer Q0NFREFOT0BDUkVTVFZJRVdDLkNPTX4yMDE5LTAyLTE0VDEzOjU0OjUw"

	excelFileName := "groups_to_copy.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	existing_group_num := ""

	for _, sheet := range xlFile.Sheets {

		existing_group_num = strings.Trim(sheet.Rows[1].Cells[0].Value, " ")
		for _, row := range sheet.Rows {
			dto := map[string]interface{}{
				"group_num":  strings.Trim(row.Cells[1].Value, " "),
				"group_name": strings.Trim(row.Cells[2].Value, " "),
				"start_date": strings.Trim(row.Cells[3].Value, " "),
				"udf":        strings.Trim(row.Cells[4].Value, " "),
			}

			group, _ := json.Marshal(dto)

			fmt.Printf("%s\n", group)

		}
	}

	GetGroupsByQuery(existing_group_num, currentToken)

	query_groups_endpoint := "https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/query"
	fmt.Println(query_groups_endpoint)

	copy_groups_endpoint := "https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/copy"
	fmt.Println(copy_groups_endpoint)

}

func GetGroupsByQuery(existing_group_num string, currentToken string) {

	client := &http.Client{}
	dto, _ := json.Marshal(map[string]interface{}{"groups_group_num": existing_group_num})
	fmt.Println(string(dto))
	req, err := http.NewRequest("POST", "https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/query", bytes.NewReader(dto))
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", currentToken)
	resp, err := client.Do(req)
	// defer resp.Body.Close()

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			log.Fatalln(err2)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	// group, _ := json.Marshal(map[string]interface{}{"groups_group_num": existing_group_num})
	// resp, err := http.Post("https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/query", "application/json", bytes.NewBuffer(group))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println(string(body))
}
	modifyDTO := new(models.ModifyGroupsPlanListsDTO)
	err = json.Unmarshal(bodyData, modifyDTO)
	if err != nil {
		errResponse := liberror.NewAppError("Error parsing json.", applicationContext+err.Error())
		libhttp.EncodeErrorResponse(w, errResponse, http.StatusBadRequest)
		return
	}
