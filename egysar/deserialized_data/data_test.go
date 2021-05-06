package deserialized_data_test

import (
	"egy-sar/egysar/DTO"
	"egy-sar/egysar/deserialized_data"
	"os"
	"testing"
)

var data deserialized_data.DeserializedData

func init() {
	os.Chdir("../..")
	data = deserialized_data.ParseAllData()
}

func TestLeewayUnits(t *testing.T) {
	got := data.LeewayUnits[0]
	want := DTO.LeewayUnit{
		Name:     "Bait/wharf box lightly loaded ( 90 kg )",
		DivAngle: 15.0,
		Error:    0.1,
		Series:   13}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	got = data.LeewayUnits[1]
	want = DTO.LeewayUnit{Name: "Bait/wharf box, full ( 365kg )",
		DivAngle: 35.0,
		Error:    0.1,
		Series:   16}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestLeewayErrors(t *testing.T) {

	got := data.LeewayErrors[0]
	want := DTO.LeewayError{Series: 1,
		Zero:         0.0,
		Two:          0.183,
		Four:         0.36667,
		Six:          0.55,
		Eight:        0.665385,
		Ten:          0.7807,
		Twelve:       0.8961,
		Fourteen:     1.0115,
		Sixteen:      1.1269,
		Eighteen:     1.2423,
		Twenty:       1.3576,
		Twenty_two:   1.473077,
		Twenty_four:  1.588462,
		Twenty_six:   1.70384,
		Twenty_eight: 1.819231,
		Thirty:       1.934615,
		Thirty_two:   2.05,
		Thirty_four:  2.165385,
		Thirty_six:   2.280775,
		Thirty_eight: 2.39616,
		Forty:        2.511545,
		Forty_two:    2.62693,
		Forty_four:   2.742315,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	got = data.LeewayErrors[1]
	want = DTO.LeewayError{Series: 2.0,
		Zero:         0.0,
		Two:          0.095,
		Four:         0.19,
		Six:          0.28500000000000003,
		Eight:        0.38000000000000006,
		Ten:          0.47500000000000003,
		Twelve:       0.5700000000000001,
		Fourteen:     0.6650000000000001,
		Sixteen:      0.7600000000000002,
		Eighteen:     0.8550000000000003,
		Twenty:       0.9500000000000004,
		Twenty_two:   1.0450000000000004,
		Twenty_four:  1.1400000000000003,
		Twenty_six:   1.235,
		Twenty_eight: 1.3299999999999998,
		Thirty:       1.4249999999999996,
		Thirty_two:   1.5199999999999994,
		Thirty_four:  1.614999999999999,
		Thirty_six:   1.71,
		Thirty_eight: 1.805,
		Forty:        1.9,
		Forty_two:    1.995,
		Forty_four:   2.09}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestFixedWingSW(t *testing.T) {

	got := data.FixedWingSW.Alt150m[0]
	want := DTO.SweepWidth{
		Index:        1,
		SearchObject: "person in water",
		One:          0.0,
		Three:        0.1,
		Five:         0.1,
		Ten:          0.1,
		Fifteen:      0.1,
		Twenty:       0.1,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	got = data.FixedWingSW.Alt600m[1]
	want = DTO.SweepWidth{
		Index:        2,
		SearchObject: "power boat 10(33)",
		One:          0.3,
		Three:        2.2,
		Five:         3.4,
		Ten:          5.5,
		Fifteen:      6.9,
		Twenty:       8.0,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestHeliSW(t *testing.T) {

	got := data.HeliSW.Alt300m[1]
	want := DTO.SweepWidth{
		Index:        2,
		SearchObject: "power boat 10(33)",
		One:          0.7,
		Three:        2.6,
		Five:         3.9,
		Ten:          6.3,
		Fifteen:      7.9,
		Twenty:       9.1,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestMerchVesselSW(t *testing.T) {

	got := data.MerchVesselSW[1]
	want := DTO.SweepWidth{
		Index:        2,
		SearchObject: "Boat 12m (40ft)",
		One:          0.0,
		Three:        2.8,
		Five:         4.5,
		Ten:          7.6,
		Fifteen:      9.4,
		Twenty:       11.6,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestSearchObjectIndices(t *testing.T) {

	got := data.SearchObjectIndices[1]
	want := DTO.SearchObjectIndex{
		SearchObject: "Aircraft over 5700 KG",
		Index:        2,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestSearchObjectIndexSpeedCorrs(t *testing.T) {

	got := data.SearchObjectIndexSpeedCorrs[1]
	want := DTO.SearchObjectIndexSpeedCorr{
		SearchObject: "power boat 10(33)",
		Index:        5,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestSpeedCorrFactorsFixed(t *testing.T) {

	got := data.SpeedCorrFactorsFixed[1]
	want := DTO.SpeedCorrFactorFixed{
		OneHundredFifty:  1.1,
		OneHundredEighty: 1.0,
		TwoHundredTen:    0.9,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestSpeedCorrFactorsHelo(t *testing.T) {

	got := data.SpeedCorrFactorsHelo[1]
	want := DTO.SpeedCorrFactorHelo{
		Sixty:            1.3,
		Ninety:           1.0,
		OneHundredTwenty: 0.9,
		OneHundredforty:  0.8,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestSearchFactors(t *testing.T) {

	got := data.SearchFactors[3]
	want := DTO.SearchFactor{
		SearchFactorValue:      0.3,
		RelativeEffortIdealCo:  0.2,
		RelativeEffortNormalCo: 0.2,
	}

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}
