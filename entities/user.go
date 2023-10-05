package entities

type User struct {
	Id       uint32 `json:"Id"`
	Username string `json:"Username" binding:"min=1,max=31"`
	Email    string `json:"Email" binding:"min=1,max=255"`
}
