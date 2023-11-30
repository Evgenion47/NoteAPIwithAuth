package models

type AuthInfo struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type Note struct {
	ID          int64  `json:"id"`
	OwnerID     string `json:"ownerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completion  bool   `json:"completion"`
	CreatedAt   string `json:"createdAt"`
}

type UpdateNote struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completion  bool   `json:"completion"`
}

type CreateNote struct {
	OwnerId     string `json:"ownerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Notes struct {
	Notes []Note
}

type NoteId struct {
	ID int64 `json:"id"`
}

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
