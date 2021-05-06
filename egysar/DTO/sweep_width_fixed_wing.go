package DTO

type SweepWidthFixedWing struct {
	Alt150m []SweepWidth `json:"alt150m"`
	Alt300m []SweepWidth `json:"alt300m"`
	Alt600m []SweepWidth `json:"alt600m"`
}
