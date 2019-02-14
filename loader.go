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

	currentToken := "Bearer Q0NFREFOT0BDUkVTVFZJRVdDLkNPTX4yMDE5LTAyLTE0VDE4OjA0OjEx"

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

func GetGroupsLocation(existingGroupId int64, currentToken string) ([]*models.GroupsLocation, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/%d/groupslocation", existingGroupId), nil)
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
	items := make([]*models.GroupsLocation, 0)
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsPlanList(existingGroupId int64, currentToken string) ([]*models.GroupsPlanList, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/%d/groupsplanlist", existingGroupId), nil)
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
	items := make([]*models.GroupsPlanList, 0)
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsPriorAuth(existingGroupId int64, currentToken string) ([]*models.GroupsPriorAuth, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/%d/groupspriorauth", existingGroupId), nil)
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
	items := make([]*models.GroupsPriorAuth, 0)
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsDedCapMgmt(existingGroupId int64, currentToken string) ([]*models.GroupsDedCapMgmt, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/%d/groupsdedcapmgmt", existingGroupId), nil)
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
	items := make([]*models.GroupsDedCapMgmt, 0)
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsClaimAdminList(existingGroupId int64, currentToken string) ([]*models.GroupsClaimAdminList, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/%d/groupsclaimadminlist", existingGroupId), nil)
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
	items := make([]*models.GroupsClaimAdminList, 0)
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsClaimAdminFeeList(existingGroupId int64, currentToken string) ([]*models.GroupsClaimAdminFeeList, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/%d/groupsclaimadminfeelist", existingGroupId), nil)
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
	items := make([]*models.GroupsClaimAdminFeeList, 0)
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsSubPlan(existingGroupId int64, currentToken string) ([]*models.GroupsSubPlan, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/%d/groupssubplan", existingGroupId), nil)
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
	items := make([]*models.GroupsSubPlan, 0)
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}

func GetGroupsDynamicEnrollmentPharmacyDaysHours(existingGroupId int64, currentToken string) ([]*models.GroupsDynamicEnrollmentPharmacyDaysHours, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pmt-pharm-uat.tredium.com/api/plan-svc/groups/%d/groupsdynamicenrollmentpharmacydayshours", existingGroupId), nil)
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
	items := make([]*models.GroupsDynamicEnrollmentPharmacyDaysHours, 0)
	err = json.Unmarshal(bodyData, &items)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return items, nil
}
