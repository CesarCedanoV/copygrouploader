package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"git.tredium.com/scm/copygrouploader/models"
	"github.com/tealeg/xlsx"
)

func main() {

	response := LoaderResponse{}

	currentToken := "Bearer Q0NFREFOT0BDUkVTVFZJRVdDLkNPTX4yMDE5LTAyLTE0VDE2OjI3OjEx"

	excelFileName := "groups_to_copy_cesar.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalln(err)
	}

	for _, sheet := range xlFile.Sheets {

		existing_group_num := strings.Trim(sheet.Rows[1].Cells[0].Value, " ")

		existingGroup, err := GetGroupsByQuery(existing_group_num, currentToken)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(existingGroup)

		for key, row := range sheet.Rows {
			// Omit first row about the Table Headers
			if key == 0 {
				continue
			}

			newGroup := existingGroup
			newGroup.GroupNum = strings.Trim(row.Cells[1].Value, " ")
			newGroup.GroupName = strings.Trim(row.Cells[2].Value, " ")
			start_date, err := time.Parse("20060102", strings.Trim(row.Cells[3].Value, " "))
			if err != nil {
				response.RowsFailed = append(response.RowsFailed, int64(key+1))
				response.TotalFailed++
				fmt.Println("In row:", (key + 1), "Error", err)
				continue
			}
			newGroup.StartDate = start_date

			err = AddGroup(*newGroup, currentToken)
			if err != nil {
				response.RowsFailed = append(response.RowsFailed, int64(key+1))
				response.TotalFailed++
				fmt.Println("In row:", (key + 1), "Error:", err)
				continue
			}

			response.RowsAdded = append(response.RowsAdded, int64(key+1))
			response.TotalAdded++

		}
	}
}

func GetGroupsByQuery(existing_group_num string, currentToken string) (*models.Groups, error) {

	client := &http.Client{}
	dto, _ := json.Marshal(map[string]interface{}{"groups_group_num": existing_group_num})
	req, err := http.NewRequest("POST", "https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/query", bytes.NewReader(dto))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", currentToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Error: {\"status_code\": %d}", resp.StatusCode))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	existingGroups := make([]models.Groups, 0)
	err = json.Unmarshal(bodyData, &existingGroups)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &existingGroups[0], nil
}

func AddGroup(group models.Groups, currentToken string) error {

	client := &http.Client{}
	dto, _ := json.Marshal(group)
	fmt.Println(string(dto))
	req, err := http.NewRequest("POST", "https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/", bytes.NewReader(dto))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", currentToken)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return errors.New(fmt.Sprintf("Error: {\"status_code\": %d}", resp.StatusCode))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	fmt.Println(string(bodyData))

	return nil
}

type LoaderResponse struct {
	RowsAdded   []int64
	TotalAdded  int64
	RowsFailed  []int64
	TotalFailed int64
}
