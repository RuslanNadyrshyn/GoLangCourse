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
	Menu      []Product `json:"menu"`
	CreatedAt string
	UpdatedAt string
}
