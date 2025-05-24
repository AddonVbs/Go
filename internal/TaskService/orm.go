package taskservers

type Task struct {
	ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Task1 string `json:"task"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
