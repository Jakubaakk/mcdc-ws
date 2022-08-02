package users

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id                uuid.UUID      `gorm:"primaryKey" json:"id"`
	MinecraftNickname sql.NullString `gorm:"unique" json:"minecraftNickname"`
	DiscordId         sql.NullString `gorm:"unique" json:"discordId"`
	Verified          bool
	Subscribed        bool
}

type UserJson struct {
	Id                uuid.UUID `json:"id"`
	MinecraftNickname string    `json:"minecraftNickname"`
	DiscordId         string    `json:"discordId"`
	Verified          bool      `json:"verified"`
	Subscribed        bool      `json:"subscribed"`
}

type Status struct {
	Verified bool `json:"verified"`
}

type McdcError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HttpStatus int    `json:"-"`
}

func (self *User) toUserJson() *UserJson {
	return &UserJson{
		Id:                self.Id,
		MinecraftNickname: self.MinecraftNickname.String,
		DiscordId:         self.DiscordId.String,
		Verified:          self.Verified,
		Subscribed:        self.Subscribed,
	}
}

func (self *UserJson) toUser() *User {
	return &User{
		Id:                self.Id,
		MinecraftNickname: sql.NullString{String: self.MinecraftNickname, Valid: self.MinecraftNickname != ""},
		DiscordId:         sql.NullString{String: self.DiscordId, Valid: self.DiscordId != ""},
		Verified:          self.Verified,
		Subscribed:        self.Subscribed,
	}
}
