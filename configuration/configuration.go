package configuration

import uuid "github.com/satori/go.uuid"

type McdcConfiguration struct {
	ApiKeys  []ApiKey `json:"apiKeys"`
	Database Database `json:"database"`
}

type ApiKey struct {
	Key     uuid.UUID `json:"key"`
	Name    string    `json:"name"`
	Enabled bool      `json:"enabled"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}
