package models

type TaskStatus string

const (
	StatusNew        TaskStatus = "New"
	StatusInProgress TaskStatus = "In Progress"
	StatusCompleted  TaskStatus = "Done"
)

func (s TaskStatus) IsValid() bool {
	switch s {
	case StatusNew, StatusInProgress, StatusCompleted:
		return true
	default:
		return false
	}
}

type Task struct { // можно было сделать ID и Подструктура Attributes
	ID          string     `json:"id,omitempty"`
	Title       string     `json:"title" validate:"required,min=1"`
	Description string     `json:"description" validate:"required,min=1"`
	Status      TaskStatus `json:"status" validate:"required"`
}
