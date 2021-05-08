package deserialized_data

import (
	"egy-sar/egysar/DTO"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var data *DeserializedData

type DeserializedData struct {
	LeewayUnits                 []DTO.LeewayUnit
	LeewayErrors                []DTO.LeewayError
	FixedWingSW                 DTO.SweepWidthFixedWing
	HeliSW                      DTO.SweepWidthHeli
	MerchVesselSW               []DTO.SweepWidth
	SearchObjectIndices         []DTO.SearchObjectIndex
	SearchObjectIndexSpeedCorrs []DTO.SearchObjectIndexSpeedCorr
	SpeedCorrFactorsFixed       []DTO.SpeedCorrFactorFixed
	SpeedCorrFactorsHelo        []DTO.SpeedCorrFactorHelo
	SearchFactors               []DTO.SearchFactor
}

func ParseAllData() *DeserializedData {

	var leewayUnits []DTO.LeewayUnit
	byteValue := readJsonFile("leeway unit types.json")
	json.Unmarshal(byteValue, &leewayUnits)

	var leewayErrors []DTO.LeewayError
	byteValue = readJsonFile("leeway error.json")
	json.Unmarshal(byteValue, &leewayErrors)

	var fixedWingSW DTO.SweepWidthFixedWing
	byteValue = readJsonFile("sweep width fixed-wing.json")
	json.Unmarshal(byteValue, &fixedWingSW)

	var heliSW DTO.SweepWidthHeli
	byteValue = readJsonFile("Sweep width for heli.json")
	json.Unmarshal(byteValue, &heliSW)

	var merchVesselSW []DTO.SweepWidth
	byteValue = readJsonFile("Sweep width merchant vessel.json")
	json.Unmarshal(byteValue, &merchVesselSW)

	var searchObjectIndices []DTO.SearchObjectIndex
	byteValue = readJsonFile("Search Object Index(Search Eff).json")
	json.Unmarshal(byteValue, &searchObjectIndices)

	var searchObjectIndexSpeedCorrs []DTO.SearchObjectIndexSpeedCorr
	byteValue = readJsonFile("Search Object index(Speed Corr).json")
	json.Unmarshal(byteValue, &searchObjectIndexSpeedCorrs)

	var speedCorrFactorsFixed []DTO.SpeedCorrFactorFixed
	byteValue = readJsonFile("Speed correction factor fixed.json")
	json.Unmarshal(byteValue, &speedCorrFactorsFixed)

	var speedCorrFactorsHelo []DTO.SpeedCorrFactorHelo
	byteValue = readJsonFile("Speed correction factor helo.json")
	json.Unmarshal(byteValue, &speedCorrFactorsHelo)

	var searchFactors []DTO.SearchFactor
	byteValue = readJsonFile("Search Factors.json")
	json.Unmarshal(byteValue, &searchFactors)
	ResultData := DeserializedData{
		LeewayUnits:                 leewayUnits,
		LeewayErrors:                leewayErrors,
		FixedWingSW:                 fixedWingSW,
		HeliSW:                      heliSW,
		MerchVesselSW:               merchVesselSW,
		SearchObjectIndices:         searchObjectIndices,
		SearchObjectIndexSpeedCorrs: searchObjectIndexSpeedCorrs,
		SpeedCorrFactorsFixed:       speedCorrFactorsFixed,
		SpeedCorrFactorsHelo:        speedCorrFactorsHelo,
		SearchFactors:               searchFactors,
	}
	return &ResultData
}

func readJsonFile(fileName string) []byte {
	jsonFile, err := os.Open("data/" + fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	return byteValue
}

func GetData() *DeserializedData {

	if data == nil {
		data = ParseAllData()
	}
	return data
}
