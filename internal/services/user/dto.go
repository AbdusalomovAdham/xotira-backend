package user

type Create struct {
	FullName string
	Password *string
	Region   string
	City     string
	Email    string
	Avatar   string
}

type Update struct {
	Id       int     `json:"id" form:"id"`
	FullName *string `json:"full_name" form:"full_name"`
	Password *string `json:"password" form:"password"`
	Region   *string `json:"region" form:"region"`
	City     *string `json:"city" form:"city"`
	Avatar   *string `json:"avatar" form:"avatar"`
	Email    *string `json:"email" form:"email"`
}
