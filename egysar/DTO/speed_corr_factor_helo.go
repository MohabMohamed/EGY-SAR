package DTO

type SpeedCorrFactorHelo struct {
	Sixty            float64 `json:"60.0"`
	Ninety           float64 `json:"90.0"`
	OneHundredTwenty float64 `json:"120.0"`
	OneHundredforty  float64 `json:"140.0"`
}

func (s SpeedCorrFactorHelo) GetCorrectionOfSpeed(speed int32) float64 {
	switch speed {
	case 60:
		return s.Sixty
	case 90:
		return s.Ninety

	case 120:
		return s.OneHundredTwenty

	case 140:
		return s.OneHundredforty
	default:
		return -1
	}
}
