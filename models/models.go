package models

type User struct {
	ID       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

type Result struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	WPM       int    `json:"wpm"`       // Слов в минуту
	Accuracy  int    `json:"accuracy"`  // Точность в процентах
	Timestamp string `json:"timestamp"` // Время выполнения теста
}
