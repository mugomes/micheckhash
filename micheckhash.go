// MiCheckHash
// Copyright (C) 2025 Murilo Gomes Julio. Todos os direitos reservados.

// Este software e código-fonte é distribuído sob os termos do Contrato de Licença de Usuário Final do MiCheckHash.

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	c "mugomes/micheckhash/controls"
	"net/url"
	"os"
	"strings"
)

const VERSION_APP string = "5.0.1"

type myTheme struct {
	fyne.Theme
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 16
	}
	return m.Theme.Size(name)
}

func main() {
	c.LoadTranslations()

	sIcon := fyne.NewStaticResource("micheckhash.png", resourceMicheckhashPngData)

	a := app.NewWithID("br.com.mugomes.micheckhash")
	a.SetIcon(sIcon)
	w := a.NewWindow("MiCheckHash")
	w.Resize(fyne.NewSize(500, 379))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	a.Settings().SetTheme(&myTheme{theme.DarkTheme()})

	mnuTools := fyne.NewMenu(c.T("Tools"),
		fyne.NewMenuItem(
			c.T("Generate Hash"), func() {
				showGerarHash(a)
			}),
	)

	mnuAbout := fyne.NewMenu(c.T("About"),
		fyne.NewMenuItem(c.T("Check Update"), func() {
			url, _ := url.Parse("https://www.mugomes.com.br/2025/07/micheckhash.html")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("Support MiCheckHash"), func() {
			url, _ := url.Parse("https://www.mugomes.com.br/p/apoie.html")
			a.OpenURL(url)
		}),
		fyne.NewMenuItem(c.T("Technical Support"), func() {
			url, _ := url.Parse("https://www.mugomes.com.br/2025/07/micheckhash.html#support")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("About MiCheckHash"), func() {
			showAbout(a)
		}),
	)

	w.SetMainMenu(fyne.NewMainMenu(mnuTools, mnuAbout))

	lblTipo := widget.NewLabel(c.T("Hash Type"))
	lblTipo.TextStyle = fyne.TextStyle{Bold: true}

	sOptions := []string{"MD5", "SHA1", "SHA256", "SHA512"}
	txtTipo := widget.NewSelect(sOptions, func(string) {})
	txtTipo.PlaceHolder = "MD5"
	txtTipo.Resize(fyne.NewSize(w.Canvas().Size().Width-7, 40))
	txtTipo.Move(fyne.NewPos(0, lblTipo.Position().Y+37))

	lblArquivo := widget.NewLabel(c.T("File"))
	lblArquivo.TextStyle = fyne.TextStyle{Bold: true}
	lblArquivo.Move(fyne.NewPos(0, txtTipo.Position().Y+37))
	txtArquivo := widget.NewEntry()
	txtArquivo.SetPlaceHolder(c.T("Select file"))
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

	lblHash := widget.NewLabel(c.T("Type/Paste the Hash"))
	lblHash.Move(fyne.NewPos(0, txtArquivo.Position().Y+37))
	txtHash := widget.NewEntry()
	txtHash.Resize(fyne.NewSize(w.Canvas().Size().Width-7, 37))
	txtHash.Move(fyne.NewPos(0, lblHash.Position().Y+37))

	var lblInfo *widget.Label
	var btnCheck *widget.Button

	btnCheck = widget.NewButton(c.T("Check Now"), func() {
		go func() {
			fyne.Do(func() {
				lblInfo.SetText(c.T("Verifying Hash... Please wait!"))
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

			fileHash := c.GetHash(sTipoHash, file)

			fyne.Do(func() {
				lblInfo.SetText("")
			})
			if fileHash == sHash {
				fyne.Do(func() {
					dialog.NewInformation("MiCheckHash", c.T("Success!"), w).Show()
				})
			} else {
				fyne.Do(func() {
					dialog.NewInformation("MiCheckHash", c.T("Different!"), w).Show()
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
