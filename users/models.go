package users

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id                uuid.UUID `gorm:"primaryKey" json:"id"`
	MinecraftNickname string    `gorm:"unique" json:"minecraftNickname"`
	DiscordId         string    `gorm:"unique" json:"discordId"`
	Verified          bool      `json:"verified"`
}

type Status struct {
	Verified bool `json:"verified"`
}

type McdcError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HttpStatus int    `json:"-"`
}
