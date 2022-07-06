package model

type Service struct {
	Id      int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Service string `json:"service"`
	Serial  string `json:"serial"`
	Source  string `json:"source"`
	Price   string `json:"price"`
	Date    string `json:"date"`
	Status  bool   `json:"status"`
}

type ServiceProcess struct {
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Service string `json:"service"`
	Serial  string `json:"serial"`
	Source  string `json:"source"`
	Price   string `json:"price"`
	Date    string `json:"date"`
}
