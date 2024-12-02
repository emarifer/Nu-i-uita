package db

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/emarifer/Nu-i-uita/internal/models"

	"github.com/ostafen/clover/v2/query"

	c "github.com/ostafen/clover/v2"
	d "github.com/ostafen/clover/v2/document"
)

const (
	db_files                   = "nuiuita_db"
	language_collection        = "LANGUAGE"
	master_password_collection = "MASTER_PASSWORD"
	password_entry_collection  = "PASSWORD_ENTRIES"
	dump_file_extension        = "nuitabak"
)

type Db struct {
	clover *c.DB
}

type DbDump struct {
	Mp           string
	Pwds         []models.PasswordEntryDto
	LanguageCode string
}

func setupCollections(db *c.DB) {
	db.CreateCollection(language_collection)
	db.CreateCollection(master_password_collection)
	db.CreateCollection(password_entry_collection)
}

func NewDb() *Db {
	// It is necessary to create the folder that
	// will contain the DB storage files
	err := os.MkdirAll(db_files, os.ModePerm)
	if err != nil {
		panic(err)
	}
	db, err := c.Open(db_files)
	if err != nil {
		panic(err)
	}

	setupCollections(db)

	return &Db{clover: db}
}

func (db *Db) Close() error {
	return db.clover.Close()
}

func (db *Db) GetLanguageCode() string {
	doc, err := db.clover.FindFirst(query.NewQuery(language_collection))
	if err != nil {
		panic(err)
	}
	if doc == nil {
		return ""
	}

	if code, ok := doc.Get("code").(string); ok {
		return code
	}

	return ""
}

func (db *Db) SaveLanguageCode(code string) {
	db.clover.Delete(query.NewQuery(language_collection))
	doc := d.NewDocument()
	doc.Set("code", code)
	_, err := db.clover.InsertOne(language_collection, doc)
	if err != nil {
		panic(err)
	}
}

func (db *Db) InsertMasterPassword(mpStr string) string {
	mp := models.NewMasterPassword(mpStr)
	doc := d.NewDocumentOf(mp)
	id, err := db.clover.InsertOne(master_password_collection, doc)
	if err != nil {
		panic(err)
	}
	// mp.Id = id
	// return mp, nil

	return id
}

// RecoverMasterPassword returns the master password struct
// from the database or an empty struct if it has not yet been set
func (db *Db) RecoverMasterPassword() models.MasterPassword {
	doc, err := db.clover.FindFirst(query.NewQuery(master_password_collection))
	if err != nil {
		panic(err)
	}
	if doc == nil {
		return models.MasterPassword{}
	}

	var mp models.MasterPassword
	err = doc.Unmarshal(&mp)
	if err != nil {
		panic(err)
	}

	return mp
}

func (db *Db) GetCryptoInstance() *models.Crypto {
	mp := db.RecoverMasterPassword()
	instance := mp.GetCrypto()

	return &instance
}

func (db *Db) InsertPasswordEntry(password models.PasswordEntry) string {
	crypto := db.GetCryptoInstance()

	passDto := password.DTO(crypto)

	doc := d.NewDocumentOf(passDto)
	id, err := db.clover.InsertOne(password_entry_collection, doc)
	if err != nil {
		log.Println(err)
	}

	return id
}

func (db *Db) PasswordCount() int {
	c, _ := db.clover.Count(query.NewQuery(password_entry_collection))

	return c
}

func (db *Db) GetAllPasswords() []models.PasswordEntry {
	//TODO: maybe I need to check that the collection exists
	pwDtosDocs, _ := db.clover.FindAll(
		query.
			NewQuery(password_entry_collection).
			Sort(query.SortOption{Field: "timestamp", Direction: -1}),
	)

	return db.loadManyPasswordEntry(pwDtosDocs)
}

/*
// FilterPasswords filter the collection by search criteria
func (db *Db) FilterPasswords(search string) []models.PasswordEntry {
	searchPattern := fmt.Sprintf("(?i)%s", search)
	pwDtosDocs, _ := db.clover.FindAll(
		query.
			NewQuery(password_entry_collection).
			Where(
				query.Field("website").Like(searchPattern).Or(
					query.Field("username").Like(searchPattern),
				)))

	return db.loadManyPasswordEntry(pwDtosDocs)
}
*/

func (db *Db) GetPasswordById(id string) models.PasswordEntry {
	crypto := db.GetCryptoInstance()
	doc, _ := db.clover.FindById(password_entry_collection, id)
	dto := loadPasswordEntryDto(doc)

	return dto.ToPasswordEntry(crypto)
}

func (db *Db) loadManyPasswordEntry(docs []*d.Document) []models.PasswordEntry {
	crypto := db.GetCryptoInstance()
	result := make([]models.PasswordEntry, len(docs))
	for i, doc := range docs {
		dto := loadPasswordEntryDto(doc)
		result[i] = dto.ToPasswordEntry(crypto)
	}

	return result
}

func loadPasswordEntryDto(doc *d.Document) *models.PasswordEntryDto {
	var dto models.PasswordEntryDto
	doc.Unmarshal(&dto)
	dto.Id = doc.ObjectId()

	return &dto
}

func loadManyPasswordEntryDto(docs []*d.Document) []models.PasswordEntryDto {
	result := make([]models.PasswordEntryDto, len(docs))
	for i, doc := range docs {
		result[i] = *loadPasswordEntryDto(doc)
	}

	return result
}

func (db *Db) DeletePasswordEntry(id string) {
	db.clover.DeleteById(password_entry_collection, id)
}

func (db *Db) UpdatePasswordEntry(pe models.PasswordEntry) bool {
	dto := pe.DTO(db.GetCryptoInstance())
	updates := dto.ToMap()
	// doc := d.NewDocumentOf(dto)
	// creating the query when the Id field (whose default name is "_id")
	// has the value of the todo that we pass to it
	q := query.
		NewQuery(password_entry_collection).
		Where(query.Field("_id").Eq(pe.Id))

	err := db.clover.Update(q, updates)

	return err == nil
}

// Drop cleans the database, deleting collections and creating empty ones
func (db *Db) DropCollections() {
	db.clover.DropCollection(language_collection)
	db.clover.DropCollection(password_entry_collection)
	db.clover.DropCollection(master_password_collection)

	// this is to avoid fatal errors in case of drop then continuing
	setupCollections(db.clover)
}

func (db *Db) GenerateDump(baseFolder string) (string, error) {
	dumpDate := time.Now().Format("200601021504")
	fileName := filepath.Join(baseFolder, fmt.Sprintf("pwds_%s.%s", dumpDate, dump_file_extension))
	f, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	mp := db.RecoverMasterPassword()
	pwDtosDocs, _ := db.clover.FindAll(
		query.NewQuery(password_entry_collection),
	)
	pwds := loadManyPasswordEntryDto(pwDtosDocs)
	lc := db.GetLanguageCode()
	data := DbDump{
		mp.Value,
		pwds,
		lc,
	}

	encoder := gob.NewEncoder(f)
	errEncoding := encoder.Encode(data)

	return fileName, errEncoding
}

func (db *Db) ImportDump(password string, dumpFileLocation string) error {
	file, err := os.Open(dumpFileLocation)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var importedDump DbDump
	err = decoder.Decode(&importedDump)
	if err != nil {
		return errors.New("cannot read data dump")
	}

	mp := models.NewMasterPasswordFromB64(importedDump.Mp)
	if !mp.Check(password) {
		return errors.New("invalid dump password")
	}

	newMp := models.NewMasterPassword(password)
	cryptoImport := newMp.GetCrypto()

	for _, dto := range importedDump.Pwds {
		pe := dto.ToPasswordEntry(&cryptoImport)
		db.InsertPasswordEntry(pe)
	}

	db.SaveLanguageCode(importedDump.LanguageCode)

	return nil
}
