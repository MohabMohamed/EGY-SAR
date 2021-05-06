package DTO

type SearchFactor struct {
	SearchFactorValue      float64 `json:"Search factor (fs)"`
	RelativeEffortIdealCo  float64 `json:"RelativeEffortIdealCo"`
	RelativeEffortNormalCo float64 ` json:"RelativeEffortNormalCo"`
}

func (s SearchFactor) GetRelativeEffByCond(condition string) float64 {
	switch condition {
	case "RelativeEffortIdealCo":
		return s.RelativeEffortIdealCo
	case "RelativeEffortNormalCo":
		return s.RelativeEffortNormalCo
	default:
		return -1
	}
}
