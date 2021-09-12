package entity

type UserInfo struct {
	UserId          int64  `json:"user_id,omitempty"`
	Role            string `json:"role,omitempty"`
	IsAuthenticated bool   `json:"is_authenticated,omitempty"`
}
