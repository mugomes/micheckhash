// MiCheckHash
// Copyright (C) 2025 Murilo Gomes Julio. Todos os direitos reservados.

// Este software e código-fonte é distribuído sob os termos do Contrato de Licença de Usuário Final do MiCheckHash.

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func showGerarHash(a fyne.App) {
	w := a.NewWindow("Gerar Hash")
	w.Resize(fyne.NewSize(500, 300))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	lblTipoHash := widget.NewLabel("Tipo de Hash")
	lblTipoHash.Move(fyne.NewPos(5, 0))
	lblTipoHash.TextStyle = fyne.TextStyle{Bold: true}
	sOptions := []string{"MD5", "SHA1", "SHA256", "SHA512"}
	cboTipoHash := widget.NewSelect(sOptions, func(string) {})
	cboTipoHash.PlaceHolder = "MD5"
	cboTipoHash.Resize(fyne.NewSize(w.Canvas().Size().Width-17, 30))
	cboTipoHash.Move(fyne.NewPos(5, lblTipoHash.Position().Y+37))

	lblArquivo := widget.NewLabel("Arquivo")
	lblArquivo.TextStyle = fyne.TextStyle{Bold: true}
	lblArquivo.Move(fyne.NewPos(5, cboTipoHash.Position().Y+37))
	txtArquivo := widget.NewEntry()
	txtArquivo.Resize(fyne.NewSize(w.Canvas().Size().Width-52, 38.4))
	txtArquivo.Move(fyne.NewPos(5, lblArquivo.Position().Y+37))
	txtArquivo.Disable()

	btnArquivo := widget.NewButton("...", func() {
		dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
			if r != nil {
				sPath := r.URI().Path()
				txtArquivo.SetText(sPath)
			} else {
				return
			}
		}, w)
	})

	btnArquivo.Resize(fyne.NewSize(30, 38.4))
	btnArquivo.Move(fyne.NewPos(txtArquivo.Size().Width+10, txtArquivo.Position().Y))

	var btnGerar *widget.Button
	var txtInfo *widget.Entry
	var btnSave *widget.Button
	var sTipoHash = ""
	btnGerar = widget.NewButton("Generate Hash", func() {
		go func() {
			fyne.Do(func() {
				txtInfo.SetText("Generate Hash... Please wait!")
				btnGerar.Disable()
			})

			sFilename := txtArquivo.Text

			sTipoHash = cboTipoHash.Selected
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
			fyne.Do(func() {
				txtInfo.SetText(hex.EncodeToString(hashInBytes))
				btnGerar.Enable()
			})
		}()
	})

	ctnGerar := container.NewHBox(
		layout.NewSpacer(),
		btnGerar,
		layout.NewSpacer(),
	)

	ctnGerar.Resize(fyne.NewSize(w.Canvas().Size().Width-17, 30))
	ctnGerar.Move(fyne.NewPos(5, txtArquivo.Position().Y+60))

	txtInfo = widget.NewEntry()
	txtInfo.Disable()
	txtInfo.Resize(fyne.NewSize(w.Canvas().Size().Width-87, 38.04))
	txtInfo.Move(fyne.NewPos(5, ctnGerar.Position().Y+52))
	btnSave = widget.NewButton("Save", func() {
		sFilename := filepath.Base(txtArquivo.Text)
		sConteudo := []byte(fmt.Sprintf("%s %s", txtInfo.Text, sFilename))
		if err := os.WriteFile(txtArquivo.Text+"."+sTipoHash, sConteudo, 0644); err != nil {
			dialog.NewError(err, w)
		}
	})
	btnSave.Resize(fyne.NewSize(67, 38.4))
	btnSave.Move(fyne.NewPos(txtInfo.Size().Width+10, txtInfo.Position().Y))

	layout := container.NewWithoutLayout(
		lblTipoHash,
		cboTipoHash,
		lblArquivo,
		txtArquivo,
		btnArquivo,
		ctnGerar,
		txtInfo,
		btnSave,
	)
	w.SetContent(layout)
	w.Show()
}
