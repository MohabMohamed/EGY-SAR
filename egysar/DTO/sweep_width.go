package DTO

type SweepWidth struct {
	Index        int32   `json:"index"`
	SearchObject string  `json:"searchObject"`
	One          float64 `json:"1.0"`
	Three        float64 `json:"3.0"`
	Five         float64 `json:"5.0"`
	Ten          float64 `json:"10.0"`
	Fifteen      float64 `json:"15.0"`
	Twenty       float64 `json:"20.0"`
}

func (s SweepWidth) getWuForExactVisibility(visibility int32) float64 {
	switch visibility {
	case 1:
		return s.One
	case 3:
		return s.Three
	case 5:
		return s.Five
	case 10:
		return s.Ten
	case 15:
		return s.Fifteen
	case 20:
		return s.Twenty
	default:
		return -1
	}
}

func (s SweepWidth) getVisibilityBefore(num float64) int32 {
	if num < 3 {
		return 1
	}
	if num < 5 {
		return 3
	}
	if num < 10 {
		return 5
	}
	if num < 15 {
		return 10
	}
	if num < 20 {
		return 15
	}
	return 20
}

func (s SweepWidth) getLowerWuValue(visibility float64) float64 {
	return s.getWuForExactVisibility(s.getVisibilityBefore(visibility))
}
func (s SweepWidth) getVisibilityAfter(num float64) int32 {
	if num >= 20 {
		return 20
	}
	if num > 15 {
		return 20
	}
	if num > 10 {
		return 15
	}
	if num > 5 {
		return 10
	}
	if num > 3 {
		return 5
	}
	if num > 1 {
		return 3
	}
	return 1

}

func (s SweepWidth) getUpperWuValue(visibility float64) float64 {
	return s.getWuForExactVisibility(s.getVisibilityAfter(visibility))
}

func (s SweepWidth) interpolate(visibility float64) float64 {
	visibilityBefore := s.getVisibilityBefore(visibility)
	visibilityAfter := s.getVisibilityAfter(visibility)
	visibilityDiff := visibility - float64(visibilityBefore)
	lowerWuValue := s.getLowerWuValue(visibility)
	WuDiff := s.getUpperWuValue(visibility) - lowerWuValue

	return lowerWuValue + ((visibilityDiff * WuDiff) / float64((visibilityAfter - visibilityBefore)))
}

func (s SweepWidth) GetWuValue(visibility float64) float64 {
	vis := int32(visibility)
	if vis == 1 || vis == 3 || vis == 5 ||
		vis == 10 || vis == 15 || vis == 20 {
		return s.getWuForExactVisibility(vis)
	}
	return s.interpolate(visibility)

}
