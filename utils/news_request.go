package utils

type UpdateArgs struct {
	UID       string `json:"uid" form:"uid"` //binding:"required"
	UNX       string `json:"unx" form:"unx"`
	Ver       string `json:"ver" form:"ver" `
	Frm       string `json:"frm" form:"frm" `
	SoftID    int64  `json:"softid" form:"softid"`
	OS        string `json:"os" form:"os"`
	Start     string `json:"start" form:"start"`
	EventType string `json:"type" form:"type"`
	City      string `json:"city" form:"city"`
}
