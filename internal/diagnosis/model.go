package diagnosis

type Diagnosis struct {
	Code         string   `json:"code"`
	Severity     string   `json:"severity"`
	Summary      string   `json:"summary"`
	Explanation  string   `json:"explanation"`
	LikelyCauses []string `json:"likely_causes,omitempty"`
	NextSteps    []string `json:"next_steps,omitempty"`
}
