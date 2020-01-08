package datamodels

// User is our User example model.
// Keep note that the tags for public-use (for our web app)
// should be kept in other file like "web/viewmodels/user.go"
// which could wrap by embedding the datamodels.User or
// define completely new fields instead but for the shake
// of the example, we will use this datamodel
// as the only one User model in our application.
type User struct {
	ID             int64     `json:"id" form:"id"`
	Username      string    `json:"username" form:"username"`
	Email					string 		`json:"email" form:"email"`
	MobilePhone		string 		`json:"mobile_phone" form:"mobile_phone"`
	CreateTime		int64			`json:"create_time" form:"create_time"`
	UpdateTime		int64 		`json:"update_time" form:"update_time"`
	Status				bool			`json:"staus" form:"staus"`
}

// Post : Post
type Post struct {
	ID             int64     `json:"id" form:"id"`
	Title	string `json:"title" form:"title"`
	Content	string `json:"content" form:"content"`
	CreateTime	int64 `json:"create_time" form:"create_time"`
	Status	string `json:"status" form:"status"`
	TagIds	string `json:"tag_ids" form:"tag_ids"`
	UserID	int64 `json:"user_id" form:"user_id"`
}
