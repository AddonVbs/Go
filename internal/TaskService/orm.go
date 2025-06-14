package taskservice

type Task struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Task string `json:"task"`
	//ALTER
	UserID int `json:"user_id"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
