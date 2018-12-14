package entity

type Resource struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Layer string `json:"layer"`
}

type Verb struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Method string `json:"method"`
}
