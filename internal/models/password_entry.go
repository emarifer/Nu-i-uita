package models

import (
	"fmt"
	"time"
)

type PasswordEntry struct {
	Id       string
	Website  string
	Username string
	Password string
}

func NewPasswordEntry(website, username, password string) PasswordEntry {

	return PasswordEntry{
		"",
		website,
		username,
		password,
	}
}

func NewPasswordEntryWithId(
	id, website, username, password string,
) PasswordEntry {

	return PasswordEntry{
		id,
		website,
		username,
		password,
	}
}

func (p *PasswordEntry) DTO(crypto *Crypto) PasswordEntryDto {
	encrypted, err := crypto.encryptB64(p.Password)
	if err != nil {
		fmt.Println("Could not decrypt password entry")
	}

	return PasswordEntryDto{
		p.Id,
		p.Website,
		p.Username,
		time.Now().Unix(),
		encrypted,
	}
}

type PasswordEntryDto struct {
	Id        string
	Website   string `clover:"website"`
	Username  string `clover:"username"`
	Timestamp int64  `clover:"timestamp"`
	Password  string `clover:"password"`
}

func (p *PasswordEntryDto) ToPasswordEntry(crypto *Crypto) PasswordEntry {
	clear, err := crypto.decryptB64(p.Password)
	if err != nil {
		return PasswordEntry{}
	}

	return NewPasswordEntryWithId(p.Id, p.Website, p.Username, clear)
}

func (p *PasswordEntryDto) ToMap() map[string]interface{} {
	result := map[string]interface{}{}

	result["website"] = p.Website
	result["username"] = p.Username
	result["timestamp"] = p.Timestamp
	result["password"] = p.Password

	return result
}
