package gui

import (
	"log"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

func New(port string, isDebug bool) {
	if AppName == "" {
		AppName = "GO Sysview"
	}

	astilectronOptions := astilectron.Options{
		AppName:            AppName,
		AppIconDarwinPath:  "resources/icon.icns",
		AppIconDefaultPath: "resources/icon.png",
		SingleInstance:     true,
		VersionAstilectron: VersionAstilectron,
		VersionElectron:    VersionElectron,
	}

	astilectronWindow := []*bootstrap.Window{{
		Options: &astilectron.WindowOptions{
			Center: astikit.BoolPtr(true),
			Height: astikit.IntPtr(900),
			Width:  astikit.IntPtr(1440),
		},
		Homepage: "http://localhost:" + port,
	}}

	options := bootstrap.Options{
		Asset:              Asset,
		AssetDir:           AssetDir,
		AstilectronOptions: astilectronOptions,
		Debug:              isDebug,
		RestoreAssets:      RestoreAssets,
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			go func() {
				if err := ws[0].Session.ClearCache(); err != nil {
					log.Fatal("Failed to clear cache")
				}
			}()
			return nil
		},
		Windows: astilectronWindow,
	}

	if err := bootstrap.Run(options); err != nil {
		log.Fatal("Bootstraping failed")
	}
}
