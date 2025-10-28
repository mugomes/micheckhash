// MiCheckHash
// Copyright (C) 2025 Murilo Gomes Julio. Todos os direitos reservados.

// Este software e código-fonte é distribuído sob os termos do Contrato de Licença de Usuário Final do MiCheckHash.

package main

import (
	"encoding/hex"
	"hash"
	"io"
	lang "mugomes/micheckhash/modules"
	"net/url"
	"os"
	"strings"

	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const VERSION_APP = "4.0.0"

type myTheme struct {
	fyne.Theme
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 16 // Set your desired font size here
	}
	return m.Theme.Size(name)
}

func main() {
	_ = lang.LoadTranslations()

	sIcon, err := fyne.LoadResourceFromPath("micheckhash.png")
	if err != nil {
		panic(err)
	}

	a := app.NewWithID("br.com.mugomes.micheckhash")
	a.SetIcon(sIcon)
	w := a.NewWindow("MiCheckHash")
	w.Resize(fyne.NewSize(500, 379))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	a.Settings().SetTheme(&myTheme{theme.DarkTheme()})

	mnuTools := fyne.NewMenu(lang.T("Tools"),
		fyne.NewMenuItem(
			lang.T("Generate Hash"), func() {
				showGerarHash(a)
			}),
	)

	mnuAbout := fyne.NewMenu(lang.T("About"),
		fyne.NewMenuItem(lang.T("Check Update"), func() {
			url, _ := url.Parse("https://www.mugomes.com.br/2025/07/micheckhash.html")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(lang.T("Technical Support"), func() {
			url, _ := url.Parse("https://www.mugomes.com.br/2025/07/micheckhash.html#support")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(lang.T("About MiCheckHash"), func() {
			showAbout(a)
		}),
	)

	w.SetMainMenu(fyne.NewMainMenu(mnuTools, mnuAbout))

	lblTipo := widget.NewLabel(lang.T("Hash Type"))
	lblTipo.TextStyle = fyne.TextStyle{Bold: true}

	sOptions := []string{"MD5", "SHA1", "SHA256", "SHA512"}
	txtTipo := widget.NewSelect(sOptions, func(string) {})
	txtTipo.PlaceHolder = "MD5"
	txtTipo.Resize(fyne.NewSize(w.Canvas().Size().Width-7, 40))
	txtTipo.Move(fyne.NewPos(0, lblTipo.Position().Y+37))

	lblArquivo := widget.NewLabel(lang.T("File"))
	lblArquivo.TextStyle = fyne.TextStyle{Bold: true}
	lblArquivo.Move(fyne.NewPos(0, txtTipo.Position().Y+37))
	txtArquivo := widget.NewEntry()
	txtArquivo.SetPlaceHolder(lang.T("Select file"))
	txtArquivo.Move(fyne.NewPos(0, lblArquivo.Position().Y+37))
	txtArquivo.Resize(fyne.NewSize(w.Canvas().Size().Width-52, 38.4))
	txtArquivo.Disable()
	btnArquivo := widget.NewButton("...", func() {
		dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
			if r != nil {
				sPath := r.URI().Path()
				txtArquivo.SetText(sPath)
			}
		}, w)
	})
	btnArquivo.Resize(fyne.NewSize(30, 38.4))
	btnArquivo.Move(fyne.NewPos(txtArquivo.Size().Width+10, txtArquivo.Position().Y))

	lblHash := widget.NewLabel(lang.T("Type/Paste the Hash"))
	lblHash.Move(fyne.NewPos(0, txtArquivo.Position().Y+37))
	txtHash := widget.NewEntry()
	txtHash.Resize(fyne.NewSize(w.Canvas().Size().Width-7, 37))
	txtHash.Move(fyne.NewPos(0, lblHash.Position().Y+37))

	var lblInfo *widget.Label
	var btnCheck *widget.Button

	btnCheck = widget.NewButton(lang.T("Check Now"), func() {
		go func() {
			fyne.Do(func() {
				lblInfo.SetText(lang.T("Verifying Hash... Please wait!"))
				btnCheck.Disable()
			})

			sFilename := txtArquivo.Text
			sHash := txtHash.Text
			sTipoHash := txtTipo.Selected
			if sTipoHash == "" {
				sTipoHash = "md5"
			} else {
				sTipoHash = strings.ToLower(sTipoHash)
			}

			file, _ := os.Open(sFilename)
			defer file.Close()

			var hashsum hash.Hash

			if sTipoHash == "md5" {
				hashsum = md5.New()
			} else if sTipoHash == "sha1" {
				hashsum = sha1.New()
			} else if sTipoHash == "sha256" {
				hashsum = sha256.New()
			} else if sTipoHash == "sha512" {
				hashsum = sha512.New()
			}

			if _, err := io.Copy(hashsum, file); err != nil {
				return
			}

			hashInBytes := hashsum.Sum(nil)
			fileHash := hex.EncodeToString(hashInBytes)

			fyne.Do(func() {
				lblInfo.SetText("")
			})
			if fileHash == sHash {
				fyne.Do(func() {
					dialog.NewInformation("MiCheckHash", lang.T("Success!"), w).Show()
				})
			} else {
				fyne.Do(func() {
					dialog.NewInformation("MiCheckHash", lang.T("Different!"), w).Show()
				})
			}

			fyne.Do(func() {
				btnCheck.Enable()
			})
		}()
	})

	ctn := container.NewHBox(
		layout.NewSpacer(),
		btnCheck,
		layout.NewSpacer(),
	)
	ctn.Resize(fyne.NewSize(w.Canvas().Size().Width, 30))
	ctn.Move(fyne.NewPos(0, txtHash.Position().Y+57))

	lblInfo = widget.NewLabel("")
	lblInfo.Move(fyne.NewPos(0, ctn.Position().Y+37))
	//btnCheck.Move(fyne.NewPos(0, ))

	layout := container.NewWithoutLayout(
		lblTipo,
		txtTipo,
		lblArquivo,
		txtArquivo,
		btnArquivo,
		lblHash,
		txtHash,
		ctn,
		lblInfo,
	)
	w.SetContent(layout)
	w.ShowAndRun()
}
