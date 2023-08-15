package structs

type Session struct {
	Id int64 `bson:"_id"`
	UserId int64 `bson:"userid"`
	Key string `bson:"key"`
}

type User struct {
	Id int64 `bson:"_id"`
	Username string `bson:"username"`
	Email string `bson:"email"`
	Password string `bson:"password"`
	Profile string `bson:"profile"`
	CreationDate int64 `bson:"creationdate"`
}
