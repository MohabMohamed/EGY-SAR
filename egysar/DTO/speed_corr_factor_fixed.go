package DTO

type SpeedCorrFactorFixed struct {
	OneHundredFifty  float64 `json:"150.0"`
	OneHundredEighty float64 `json:"180.0"`
	TwoHundredTen    float64 `json:"210.0"`
}

func (s SpeedCorrFactorFixed) GetCorrectionOfSpeed(speed int32) float64 {
	switch speed {
	case 150:
		return s.OneHundredFifty
	case 180:
		return s.OneHundredEighty
	case 210:
		return s.TwoHundredTen
	default:
		return -1
	}
}
