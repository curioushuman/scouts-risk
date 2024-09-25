package model

// Poor mans string enum
// ! NOTE: it does not prevent h.Severity = HazardSeverity("garbage")
type HazardSeverity string
const (
	HazardSeverityInsignificant HazardSeverity = "insignificant"
	HazardSeverityMinor   			HazardSeverity = "minor"
	HazardSeverityModerate  		HazardSeverity = "moderate"
	HazardSeveritySignificant 	HazardSeverity = "significant"
	HazardSeveritySevere 		  	HazardSeverity = "severe"
)

type HazardLikelihood string
const (
	HazardLikelihoodRare 					HazardLikelihood = "rare"
	HazardLikelihoodUnlikely   		HazardLikelihood = "unlikely"
	HazardLikelihoodPossible  		HazardLikelihood = "possible"
	HazardLikelihoodLikely 				HazardLikelihood = "likely"
	HazardLikelihoodAlmostCertain HazardLikelihood = "almost_certain"
)

type Hazard struct {
	Label 	 		 string
	Id    	 		 string
	Controls 		 []Control
	Consequences []Consequence
	Severity	 	 HazardSeverity
	Likelihood	 HazardLikelihood
}

type Consequence struct {
	Label string
	Id    string
}
