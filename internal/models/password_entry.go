package models

import (
	"log"
	"time"
)

type PasswordEntry struct {
	Id        string
	Website   string
	Username  string
	Password  string
	Timestamp int64
}

func NewPasswordEntry(website, username, password string) PasswordEntry {

	return PasswordEntry{
		"",
		website,
		username,
		password,
		time.Now().Unix(),
	}
}

func newPasswordEntryWithIdAndTimestamp(
	id, website, username, password string, timestamp int64,
) PasswordEntry {

	return PasswordEntry{
		id,
		website,
		username,
		password,
		timestamp,
	}
}

// ToDTO turns PasswordEntry into a DTO that stores
// the encrypted password using the master password as key.
func (p *PasswordEntry) ToDTO(crypto *Crypto) PasswordEntryDTO {
	encryptedPass, err := crypto.encryptB64(p.Password)
	if err != nil {
		log.Println("could not encrypt password entry")
	}

	return PasswordEntryDTO{
		p.Id,
		p.Website,
		p.Username,
		encryptedPass,
		p.Timestamp,
	}
}

// PasswordEntryDTO represents the object that will be saved in the database.
// This struct is used so that the change of state that means
// the passage from an unencrypted password to an encrypted one
// is "detected" by the type system, making it impossible
// to save objects with an unencrypted password in the database.
type PasswordEntryDTO struct {
	Id        string
	Website   string `clover:"website"`
	Username  string `clover:"username"`
	Password  string `clover:"password"`
	Timestamp int64  `clover:"timestamp"`
}

// ToPasswordEntry converts the DTO into a PasswordEntry
// that stores the decrypted password so that it is visible to the user,
// using the master password as the key to perform the decryption.
func (p *PasswordEntryDTO) ToPasswordEntry(crypto *Crypto) PasswordEntry {
	decryptedPass, err := crypto.decryptB64(p.Password)
	if err != nil {
		return PasswordEntry{}
	}

	return newPasswordEntryWithIdAndTimestamp(
		p.Id,
		p.Website,
		p.Username,
		decryptedPass,
		p.Timestamp,
	)
}

// ToMap converts the DTO to a map, which is necessary for clover-db
// to update the DTO. It adds a new timestamp to the object to be updated.
func (p *PasswordEntryDTO) ToMap() map[string]interface{} {
	result := map[string]interface{}{}

	result["website"] = p.Website
	result["username"] = p.Username
	result["password"] = p.Password
	result["timestamp"] = time.Now().Unix() // update timestamp

	return result
}
