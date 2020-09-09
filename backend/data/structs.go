package structs

type Transaction struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Registered string `json:"registered"`
	Role       string `json:"role"`
	Status     string `json:"status"`
}
