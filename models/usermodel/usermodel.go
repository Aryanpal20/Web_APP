package models

import seller "gin/models/sellermodel"

type User struct {
	ID       int                 `json:"id"`
	Email    string              `json:"email"`
	Password string              `json:"password"`
	Role     string              `json:"role"`
	Seller   []seller.Collection `json:"ForeignKey:User_ID"`
}
