package models

type Supplier struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Address      string `json:"address"`
	Image        string `json:"image"`
	WorkingHours struct {
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"workingHours"`
	Menu []struct {
		Id          int      `json:"id"`
		Name        string   `json:"name"`
		Price       float64  `json:"price"`
		Image       string   `json:"image"`
		Type        string   `json:"type"`
		Ingredients []string `json:"ingredients"`
	} `json:"menu"`
	CreatedAt string
	UpdatedAt string
}
