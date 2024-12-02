package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		// Title:         "Nu-i uita • minimalist password manager",
		Width:         450,
		Height:        300,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnBeforeClose:    app.beforeClose,
		Bind: []interface{}{
			app,
		},
		// Linux platform specific options
		Linux: &linux.Options{
			Icon: icon,
			// WindowIsTranslucent: true,
			WebviewGpuPolicy: linux.WebviewGpuPolicyNever,
			// ProgramName:         "wails",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
