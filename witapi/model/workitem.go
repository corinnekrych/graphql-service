package model

//TODO explore how to automate map goa mapping
type WorkItemAttributes struct {
	Description string `json:"system.description"`
	Title       string `json:"system.title"`
	State       string `json:"system.state"`
	Number      int    `json:"system.number"`
}

//
type WorkItem struct {
	Id                 string             `json:"id"`
	WorkItemAttributes WorkItemAttributes `json:"attributes"`
	Type               string             `json:"type"`
}
