package diagnosis

type PossibleCause struct {
	Name       string `json:"name"`
	Confidence int    `json:"confidence"`
}

type Diagnosis struct {
	Code         string          `json:"code"`
	Severity     string          `json:"severity"`
	Summary      string          `json:"summary"`
	Explanation  string          `json:"explanation"`
	Confidence   int             `json:"confidence"`
	LikelyCauses []string        `json:"likely_causes,omitempty"`
	Possible     []PossibleCause `json:"possible_causes,omitempty"`
	NextSteps    []string        `json:"next_steps,omitempty"`
}
