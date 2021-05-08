package egysar

import (
	"egy-sar/egysar/DTO"
	"egy-sar/egysar/deserialized_data"
	"log"
	"math"

	"github.com/mitchellh/mapstructure"
)

type Leeway struct {
	LeewayValue float64 `json:"leeway"`
	DatumRight  float64 `json:"datumRight"`
	DatumLeft   float64 `json:"datumLeft"`
	LeewayError float64 `json:"leewayError"`
	DivAngle    float64 `json:"divAngle"`
}

type LeewayInputs struct {
	DriftTime     float64
	WindDirection float64
	WindSpeed     float64
	UnitName      string
}

func NewLeeway() *Leeway {
	return &Leeway{}
}

func (leeway *Leeway) Calculate(data map[string]interface{}) *Leeway {
	var inputs LeewayInputs
	err := mapstructure.Decode(data, &inputs)
	if err != nil {
		log.Println(err)
	}
	leewayUnits := deserialized_data.GetData().LeewayUnits

	var matchedUnit DTO.LeewayUnit
	for _, unit := range leewayUnits {
		if unit.Name == inputs.UnitName {
			matchedUnit = unit
			break
		}
	}
	leeway.DivAngle = matchedUnit.DivAngle
	leeway.LeewayError = matchedUnit.Error

	direction := math.Mod((inputs.WindDirection + 180.0), 360.0)
	leeway.DatumRight = math.Mod((direction + matchedUnit.DivAngle), 360.0)
	if leeway.DatumRight < 0 {
		leeway.DatumRight += 360
	}

	leeway.DatumLeft = math.Mod((direction - matchedUnit.DivAngle), 360.0)
	if leeway.DatumLeft < 0 {
		leeway.DatumLeft += 360
	}

	idx := matchedUnit.Series - 1
	leewayErr := deserialized_data.GetData().LeewayErrors[idx]

	leeway.LeewayValue = inputs.DriftTime * leewayErr.GetErrValue(inputs.WindSpeed)

	return leeway
}
