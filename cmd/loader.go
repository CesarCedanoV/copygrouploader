package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	"git.tredium.com/scm/copygrouploader/models"
	"github.com/tealeg/xlsx"
)

func main() {

	response := LoaderResponse{}

	currentToken := "Bearer Q0NFREFOT0BDUkVTVFZJRVdDLkNPTX4yMDE5LTAyLTE1VDE1OjU5OjMz"

	excelFileName := "../groups_to_copy.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	poolSize := 50
	if err != nil {
		log.Fatalln(err)
	}

	for _, sheet := range xlFile.Sheets {

		existingGroupNum := strings.Trim(sheet.Rows[1].Cells[0].Value, " ")

		existingGroup, err := GetGroupsByQuery(existingGroupNum, currentToken)
		if err != nil {
			log.Fatalln(err)
		}
		existingGroup.GroupsLocation, err = GetGroupsLocation(existingGroup.Id, currentToken)
		if err != nil {
			log.Fatalln(err)
		}
		existingGroup.GroupsPlanList, err = GetGroupsPlanList(existingGroup.Id, currentToken)
		if err != nil {
			log.Fatalln(err)
		}
		existingGroup.GroupsPriorAuth, err = GetGroupsPriorAuth(existingGroup.Id, currentToken)
		if err != nil {
			log.Fatalln(err)
		}
		existingGroup.GroupsDedCapMgmt, err = GetGroupsDedCapMgmt(existingGroup.Id, currentToken)
		if err != nil {
			log.Fatalln(err)
		}
		existingGroup.GroupsClaimAdminList, err = GetGroupsClaimAdminList(existingGroup.Id, currentToken)
		if err != nil {
			log.Fatalln(err)
		}
		existingGroup.GroupsClaimAdminFeeList, err = GetGroupsClaimAdminFeeList(existingGroup.Id, currentToken)
		if err != nil {
			log.Fatalln(err)
		}
		existingGroup.GroupsSubPlan, err = GetGroupsSubPlan(existingGroup.Id, currentToken)
		if err != nil {
			log.Fatalln(err)
		}
		existingGroup.GroupsDynamicEnrollmentPharmacyDaysHours, err = GetGroupsDynamicEnrollmentPharmacyDaysHours(existingGroup.Id, currentToken)
		if err != nil {
			log.Fatalln(err)
		}

		min := func(a, b int) int {
			if a > b {
				return b
			}

			return a
		}
		ch := make(chan bool)
		sheet.Rows = sheet.Rows[1:100]
		l := int(math.Ceil(float64(len(sheet.Rows)) / float64(poolSize)))
		for i := 0; i < l; i++ {
			start := i * poolSize
			end := min((i+1)*poolSize, len(sheet.Rows))
			responses := 0
			chunk := sheet.Rows[start:end]

			for index, row := range chunk {
				go func(index int, row *xlsx.Row, response *LoaderResponse) {
					key := start + index
					fmt.Println("Start Process in Row ", key+1, ".")
					newGroup := existingGroup
					newGroup.GroupNum = strings.Trim(row.Cells[1].Value, " ")
					newGroup.GroupName = strings.Trim(row.Cells[2].Value, " ")
					startDate, err := time.Parse("20060102", strings.Trim(row.Cells[3].Value, " "))
					if err != nil {
						response.RowsFailed = append(response.RowsFailed, int64(key+1))
						response.TotalFailed++
						fmt.Println("In row:", (key + 1), "TimeParse Error:", err)
						ch <- true
					}
					newGroup.StartDate = startDate

					err = AddGroup(*newGroup, currentToken)
					if err != nil {
						response.RowsFailed = append(response.RowsFailed, int64(key+1))
						response.TotalFailed++
						fmt.Println("In row:", (key + 1), err)
						ch <- true
					}

					response.RowsAdded = append(response.RowsAdded, int64(key+1))
					response.TotalAdded++

					ch <- true
				}(index, row, &response)
			}
			isBreak := false
			for {
				if isBreak {
					break
				}
				select {
				case <-ch:
					responses++
					if responses >= end-start {
						isBreak = true
					}
				}
			}

		}

		// for key, row := range sheet.Rows {
		// 	// Omit first row about the Table Headers
		// 	if key == 0 {
		// 		continue
		// 	}

		// 	newGroup := existingGroup
		// 	newGroup.GroupNum = strings.Trim(row.Cells[1].Value, " ")
		// 	newGroup.GroupName = strings.Trim(row.Cells[2].Value, " ")
		// 	startDate, err := time.Parse("20060102", strings.Trim(row.Cells[3].Value, " "))
		// 	if err != nil {
		// 		response.RowsFailed = append(response.RowsFailed, int64(key+1))
		// 		response.TotalFailed++
		// 		fmt.Println("In row:", (key + 1), "TimeParse Error:", err)
		// 		continue
		// 	}
		// 	newGroup.StartDate = startDate

		// 	err = AddGroup(*newGroup, currentToken)
		// 	if err != nil {
		// 		response.RowsFailed = append(response.RowsFailed, int64(key+1))
		// 		response.TotalFailed++
		// 		fmt.Println("In row:", (key + 1), err)
		// 		continue
		// 	}

		// 	response.RowsAdded = append(response.RowsAdded, int64(key+1))
		// 	response.TotalAdded++

		// }
	}
	log.Println("Response:", response)
}

func GetGroupsByQuery(existingGroupNum string, currentToken string) (*models.Groups, error) {

	client := &http.Client{}
	dto, _ := json.Marshal(map[string]interface{}{"groups_group_num": existingGroupNum})
	req, err := http.NewRequest("POST", "https://scl-pharm-dev.tredium.com/api/plan-svc/groups/query", bytes.NewReader(dto))
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
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("GetGroupByQuery Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
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

	if len(existingGroups) > 0 {
		for _, group := range existingGroups {
			if strings.ToUpper(group.GroupNum) == strings.ToUpper(existingGroupNum) {
				return &group, nil
			}
		}
	}
	return nil, nil
}

func AddGroup(group models.Groups, currentToken string) error {

	client := &http.Client{}
	dto, _ := json.Marshal(group)
	req, err := http.NewRequest("POST", "https://scl-pharm-dev.tredium.com/api/plan-svc/groups/", bytes.NewReader(dto))
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
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("AddGroup Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

type LoaderResponse struct {
	RowsAdded   []int64
	TotalAdded  int64
	RowsFailed  []int64
	TotalFailed int64
}

func GetGroupsLocation(existingGroupId int64, currentToken string) (items []*models.GroupsLocation, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scl-pharm-dev.tredium.com/api/plan-svc/groups/%d/groupslocation", existingGroupId), nil)
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
		if resp.StatusCode == http.StatusNotFound {
			return items, nil
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("GetGroupsLocation Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsPlanList(existingGroupId int64, currentToken string) (items []*models.GroupsPlanList, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scl-pharm-dev.tredium.com/api/plan-svc/groups/%d/groupsplanlist", existingGroupId), nil)
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
		if resp.StatusCode == http.StatusNotFound {
			return items, nil
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("GetGroupsPlanList Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsPriorAuth(existingGroupId int64, currentToken string) (items []*models.GroupsPriorAuth, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scl-pharm-dev.tredium.com/api/plan-svc/groups/%d/groupspriorauth", existingGroupId), nil)
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
		if resp.StatusCode == http.StatusNotFound {
			return items, nil
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("GetGroupsPriorAuth Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsDedCapMgmt(existingGroupId int64, currentToken string) (items []*models.GroupsDedCapMgmt, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scl-pharm-dev.tredium.com/api/plan-svc/groups/%d/groupsdedcapmgmt", existingGroupId), nil)
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
		if resp.StatusCode == http.StatusNotFound {
			return items, nil
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("GetGroupsDedCapMgmt Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsClaimAdminList(existingGroupId int64, currentToken string) (items []*models.GroupsClaimAdminList, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scl-pharm-dev.tredium.com/api/plan-svc/groups/%d/groupsclaimadminlist", existingGroupId), nil)
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
		if resp.StatusCode == http.StatusNotFound {
			return items, nil
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("GetGroupsClaimAdminList Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsClaimAdminFeeList(existingGroupId int64, currentToken string) (items []*models.GroupsClaimAdminFeeList, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scl-pharm-dev.tredium.com/api/plan-svc/groups/%d/groupsclaimadminfeelist", existingGroupId), nil)
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
		if resp.StatusCode == http.StatusNotFound {
			return items, nil
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("GetGroupsClaimAdminFeeList Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsSubPlan(existingGroupId int64, currentToken string) (items []*models.GroupsSubPlan, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scl-pharm-dev.tredium.com/api/plan-svc/groups/%d/groupssubplan", existingGroupId), nil)
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
		if resp.StatusCode == http.StatusNotFound {
			return items, nil
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("GetGroupsSubPlan Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsDynamicEnrollmentPharmacyDaysHours(existingGroupId int64, currentToken string) (items []*models.GroupsDynamicEnrollmentPharmacyDaysHours, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scl-pharm-dev.tredium.com/api/plan-svc/groups/%d/groupsdynamicenrollmentpharmacydayshours", existingGroupId), nil)
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
		if resp.StatusCode == http.StatusNotFound {
			return items, nil
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("GetGroupsDynamicEnrollmentPharmacyDaysHours Error: \n{\n\"status_code\": %d,\n\"body\": %s\n}\n", resp.StatusCode, string(bodyBytes)))
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}
