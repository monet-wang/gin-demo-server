package domain

import (
	"database/sql"
)

type User struct {
	ID         int
	Name       string         `json:"name" binding:"required,name"`
	Age        sql.NullInt32  `json:"age"`
	Gender     sql.NullInt32  `json:"gender" binding:"oneof=1 2"`
	Phone      sql.NullString `json:"phone" binding:"e164"`
	CreateTime sql.NullTime   `json:"create_time"`
}

type UserList struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Gender     int    `json:"gender"`
	Phone      string `json:"phone"`
	CreateTime string `json:"create_time"`
}

type UpdateUser struct {
	ID     int
	Name   string `json:"name" binding:"required"`
	Age    int    `json:"age"`
	Gender int    `json:"gender" binding:"oneof=1 2"`
	Phone  string `json:"phone"`
}

type CreateUserResp struct {
	Id int
}
