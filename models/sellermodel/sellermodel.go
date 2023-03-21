package sellermodel

import "time"

type Collection struct {
	ID       int       `json:"id"`
	Image    string    `json:"img"`
	Rupees   string    `json:"rupees"`
	User_ID  int       `json:"user_id"`
	Createat time.Time `json:"created_at"`
}
