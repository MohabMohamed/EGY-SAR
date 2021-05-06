package DTO

type LeewayError struct {
	Series       int32   `json:"Series"`
	Zero         float64 `json:"zero"`
	Two          float64 `json:"two"`
	Four         float64 `json:"four"`
	Six          float64 `json:"six"`
	Eight        float64 `json:"eight"`
	Ten          float64 `json:"ten"`
	Twelve       float64 `json:"twelve"`
	Fourteen     float64 `json:"fourteen"`
	Sixteen      float64 `json:"sixteen"`
	Eighteen     float64 `json:"eighteen"`
	Twenty       float64 `json:"twenty"`
	Twenty_two   float64 `json:"twenty_two"`
	Twenty_four  float64 `json:"twenty_four"`
	Twenty_six   float64 `json:"twenty_six"`
	Twenty_eight float64 `json:"twenty_eight"`
	Thirty       float64 `json:"thirty"`
	Thirty_two   float64 `json:"thirty_two"`
	Thirty_four  float64 `json:"thirty_four"`
	Thirty_six   float64 `json:"thirty_six"`
	Thirty_eight float64 `json:"thirty_eight"`
	Forty        float64 `json:"forty"`
	Forty_two    float64 `json:"forty_two"`
	Forty_four   float64 `json:"forty_four"`
}

func (l LeewayError) getErrValueForExactWindSpeed(windSpeed int32) float64 {
	switch windSpeed {
	case 0:
		return l.Zero
	case 2:
		return l.Two
	case 4:
		return l.Four
	case 6:
		return l.Six
	case 8:
		return l.Eight
	case 10:
		return l.Ten
	case 12:
		return l.Twelve
	case 14:
		return l.Fourteen
	case 16:
		return l.Sixteen
	case 18:
		return l.Eighteen
	case 20:
		return l.Twenty
	case 22:
		return l.Twenty_two
	case 24:
		return l.Twenty_four
	case 26:
		return l.Twenty_six
	case 28:
		return l.Twenty_eight
	case 30:
		return l.Thirty
	case 32:
		return l.Thirty_two
	case 34:
		return l.Thirty_four
	case 36:
		return l.Thirty_six
	case 38:
		return l.Thirty_eight
	case 40:
		return l.Forty
	case 42:
		return l.Forty_two
	case 44:
		return l.Forty_four
	default:
		return -1

	}

}

func (l LeewayError) getTheEvenBefore(num float64) int32 {
	return (((int32)(num)) >> 1) << 1
}

func (l LeewayError) getLowerErrValue(windSpeed float64) float64 {
	return l.getErrValueForExactWindSpeed(l.getTheEvenBefore(windSpeed))
}

func (l LeewayError) getUpperErrValue(windSpeed float64) float64 {
	return l.getErrValueForExactWindSpeed(l.getTheEvenBefore(windSpeed) + 2)
}
func (l LeewayError) interpolate(windSpeed float64) float64 {
	theEvenBefore := l.getTheEvenBefore(windSpeed)

	windSpeedDiff := windSpeed - float64(theEvenBefore)
	lowerErrValue := l.getLowerErrValue(windSpeed)
	errDiff := l.getUpperErrValue(windSpeed) - lowerErrValue

	return lowerErrValue + ((windSpeedDiff * errDiff) / 2)
}

func (l LeewayError) GetErrValue(windSpeed float64) float64 {
	return l.interpolate(windSpeed)

}
