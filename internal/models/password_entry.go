package models

import (
	"fmt"
	"time"
)

type PasswordEntry struct {
	Id        string
	Website   string `clover:"website"`
	Username  string `clover:"username"`
	Timestamp int64  `clover:"timestamp"`
	Password  string `clover:"password"`
}

func NewPasswordEntry(website, username, password string) PasswordEntry {

	return PasswordEntry{
		Id:        "",
		Website:   website,
		Username:  username,
		Timestamp: time.Now().Unix(),
		Password:  password,
	}
}

func newPassEntryWithIdAndTimestamp(
	id, website, username, password string, timestamp int64,
) PasswordEntry {

	return PasswordEntry{
		Id:        id,
		Website:   website,
		Username:  username,
		Timestamp: timestamp,
		Password:  password,
	}
}

func (p *PasswordEntry) EncryptPassEntry(crypto *Crypto) PasswordEntry {
	encryptedPass, err := crypto.encryptB64(p.Password)
	if err != nil {
		fmt.Println("could not encrypt password entry")
	}

	return newPassEntryWithIdAndTimestamp(
		p.Id, p.Website, p.Username, encryptedPass, p.Timestamp,
	)
}

func (p *PasswordEntry) DecryptPassEntry(crypto *Crypto) PasswordEntry {
	decryptedPass, err := crypto.decryptB64(p.Password)
	if err != nil {
		return PasswordEntry{}
	}

	return newPassEntryWithIdAndTimestamp(
		p.Id, p.Website, p.Username, decryptedPass, p.Timestamp,
	)
}

func (p *PasswordEntry) ToMap() map[string]interface{} {
	result := map[string]interface{}{}

	result["website"] = p.Website
	result["username"] = p.Username
	result["timestamp"] = time.Now().Unix() // update timestamp
	result["password"] = p.Password

	return result
}
