package main

import (
	"context"

	"github.com/emarifer/Nu-i-uita/internal/db"
	"github.com/emarifer/Nu-i-uita/internal/models"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx               context.Context
	db                *db.Db
	selectedDirectory string
	selectedFile      string
}

// NewApp creates a new App application struct
func NewApp() *App {
	db := db.NewDb()

	return &App{db: db}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	var fileLocation string
	a.ctx = ctx

	runtime.EventsOn(a.ctx, "change_titles", func(optionalData ...interface{}) {
		if appTitle, ok := optionalData[0].(string); ok {
			runtime.WindowSetTitle(a.ctx, appTitle)
		}
		if selectedDirectory, ok := optionalData[1].(string); ok {
			a.selectedDirectory = selectedDirectory
		}
		if selectedFile, ok := optionalData[2].(string); ok {
			a.selectedFile = selectedFile
		}
	})

	runtime.EventsOn(a.ctx, "quit", func(optionalData ...interface{}) {
		runtime.Quit(a.ctx)
	})

	runtime.EventsOn(a.ctx, "export_data", func(optionalData ...interface{}) {
		d, _ := runtime.OpenDirectoryDialog(a.ctx, runtime.
			OpenDialogOptions{
			Title: a.selectedDirectory,
		})

		if d != "" {
			f, err := a.db.GenerateDump(d)
			if err != nil {
				runtime.EventsEmit(a.ctx, "saved_as", err.Error())
				return
			}
			runtime.EventsEmit(a.ctx, "saved_as", f)
		}
	})

	runtime.EventsOn(a.ctx, "import_data", func(optionalData ...interface{}) {
		fileLocation, _ = runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
			Title: a.selectedFile,
		})

		// fmt.Println("SELECTED FILE:", fileLocation)
		if fileLocation != "" {
			runtime.EventsEmit(a.ctx, "enter_password")
		}
	})

	runtime.EventsOn(a.ctx, "password", func(optionalData ...interface{}) {
		// fmt.Printf("MY PASS: %v", optionalData...)
		if pass, ok := optionalData[0].(string); ok {
			if len(fileLocation) != 0 {
				err := a.db.ImportDump(pass, fileLocation)
				if err != nil {
					runtime.EventsEmit(a.ctx, "imported_data", err.Error())
					return
				}
				runtime.EventsEmit(a.ctx, "imported_data", "success")
			}
		}
	})
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	defer a.db.Close()

	return false
}

// Retrieves the configured language code
func (a *App) GetLanguage() string {

	return a.db.GetLanguageCode()
}

func (a *App) SaveLanguage(code string) {

	a.db.SaveLanguageCode(code)
}

// Save the master password in the DB
func (a *App) SaveMasterPassword(mpStr string) string {
	return a.db.InsertMasterPassword(mpStr)
}

// GetMasterPassword returns true if the master password exists
// in the database or false if it has not yet been set
func (a *App) GetMasterPassword() bool {

	return a.db.RecoverMasterPassword().Value != ""
}

func (a *App) CheckMasterPassword(passFromUI string) bool {
	mp := a.db.RecoverMasterPassword()

	return mp.Check(passFromUI, a.db.SetMasterPassword)
}

func (a *App) AddPasswordEntry(website, username, password string) string {
	pe := models.NewPasswordEntry(website, username, password)

	return a.db.InsertPasswordEntry(pe)
}

func (a *App) GetPasswordCount() int {

	return a.db.PasswordCount()
}

func (a *App) GetAllEntries() []models.PasswordEntry {

	return a.db.GetAllPasswords()
}

func (a *App) GetEntryById(id string) models.PasswordEntry {

	return a.db.GetPasswordById(id)
}

func (a *App) UpdateEntry(pe models.PasswordEntry) bool {

	return a.db.UpdatePasswordEntry(pe)
}

func (a *App) DeleteEntry(id string) {
	a.db.DeletePasswordEntry(id)
}

func (a *App) Drop() {
	a.db.DropCollections()
}
