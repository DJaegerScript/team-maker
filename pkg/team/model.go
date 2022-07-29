package team

type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
	Image   string   `json:"image"`
}
