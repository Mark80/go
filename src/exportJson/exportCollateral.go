package main

import (
  "fmt"

  "io/ioutil"
  "strings"
  "encoding/json"
  "strconv"
  "os"
)

type CollateralGroup struct {
  ClientId           string  `json:"clientId"`
  CollateralCategory string  `json:"collateralCategory"`
  Types              []Types `json:"types"`
}

type Types struct {
  CollateralType                string `json:"collateralType"`
  AramisCollateralCategory      string `json:"aramisCollateralCategory"`
  AramisCollateralTypes         string `json:"aramisCollateralTypes"`
  CollateralFactorAirbAndRci    float64    `json:"collateralFactorAirbAndRci"`
  CollateralFactorFirb          float64    `json:"collateralFactorFirb"`
  CollateralRwSta               float64 `json:"collateralRwSta"`
  CollateralFactorAirbAndRciRef string `json:"collateralFactorAirbAndRciRef"`
  CollateralFactorFirbRef       string `json:"collateralFactorFirbRef"`
  CollateralRwStaRef            string `json:"collateralRwStaRef"`
  MaxCollateralFactor           float64    `json:"maxCollateralFactor"`
  ImpactOnlyPillar1             string `json:"impactOnlyPillar1"`
  ImpactOnlyPillar2             string `json:"impactOnlyPillar2,null"`
}

func main() {


  clientId := os.Args[1:][0]
  path := os.Args[1:][1]

  //"/Users/marco/go/src/exportJson/collateralSerbia.csv"
  rowData, err := ioutil.ReadFile(path)

  if err != nil {
	fmt.Println(err.Error())
	return
  }

  data := string(rowData)

  lines := strings.Split(data, "\n")

  var mapping = make(map[string][]string)

  for _, line := range lines {

	columns := strings.Split(line, ";")
	collateralCategory := columns[0]

	if mapping[collateralCategory] != nil {
	  mapping[collateralCategory] = append(mapping[collateralCategory], line)
	} else {
	  mapping[collateralCategory] = []string{line}
	}

  }

  var collaterlaGroups = []CollateralGroup{}
  for collateralCategory, lines := range mapping {

	var collateralGroup CollateralGroup
	collateralGroup.ClientId = clientId
	collateralGroup.CollateralCategory = collateralCategory

	var listTypes = []Types{
	}
	for _, line := range lines {
	  var types = Types{
		MaxCollateralFactor : 1,
		CollateralFactorAirbAndRci: 1,
		CollateralFactorFirb : 1,
		CollateralRwSta :1,
	  }
	  columns := strings.Split(line, ";")
	  types.CollateralType = columns[2]
	  types.AramisCollateralCategory = columns[3]
	  types.AramisCollateralTypes = columns[4]

	  var rangeStr string =columns[5]



	  if strings.Contains(rangeStr, "-") {
		rangeStr = strings.Split(rangeStr, "-")[1]
	  }

	  var newRangeStr string = strings.Replace(rangeStr, "%", "", -1)

	  maxCollateralFactor, _ := strconv.ParseFloat(strings.TrimSpace(newRangeStr), 64)

	  types.MaxCollateralFactor = maxCollateralFactor / 100
	  listTypes = append(listTypes, types)
	}

	collateralGroup.Types = listTypes
	collaterlaGroups = append(collaterlaGroups, collateralGroup)

  }

  jsonString, er :=  json.MarshalIndent(collaterlaGroups ,"" ," ")

  if er != nil {
	fmt.Println(er.Error())
  }

  ioutil.WriteFile("collateral01287.json", jsonString, 777)

}
