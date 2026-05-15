package diagnosis

type Diagnosis struct {
	Code        string   `json:"code"`
	Title       string   `json:"title"`
	Confidence  int      `json:"confidence"`
	Explanation []string `json:"explanation"`
	NextSteps   []string `json:"next_steps"`
}
