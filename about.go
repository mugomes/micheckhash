// MiCheckHash
// Copyright (C) 2025 Murilo Gomes Julio. Todos os direitos reservados.

// Este software e código-fonte é distribuído sob os termos do Contrato de Licença de Usuário Final do MiCheckHash.

package main

import (
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showAbout(a fyne.App) {
	w := a.NewWindow("Sobre")
	w.Resize(fyne.NewSize(597, 470))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	lblSoftware := canvas.NewText("MiCheckHash - Version:"+VERSION_APP, color.Opaque)
	lblSoftware.TextSize = 18
	lblSoftware.TextStyle.Bold = true
	lblSoftware.Move(fyne.NewPos(9, 7))

	lblDesenvolvedor1 := widget.NewLabel("Desenvolvido por:")
	lblDesenvolvedor1.TextStyle = fyne.TextStyle{Bold: true}
	lblDesenvolvedor1.Move(fyne.NewPos(0, lblSoftware.MinSize().Height+10))

	lblDesenvolvedor2 := widget.NewLabel("Murilo Gomes Julio")
	lblDesenvolvedor2.Move(fyne.NewPos(lblDesenvolvedor1.MinSize().Width-10, lblDesenvolvedor1.Position().Y))

	lblSite1 := widget.NewLabel("Site:")
	lblSite1.TextStyle = fyne.TextStyle{Bold: true}
	lblSite1.Move(fyne.NewPos(0, lblDesenvolvedor1.Position().Y+37))

	sURL, _ := url.Parse("https://www.mugomes.com.br")
	lblSite2 := widget.NewHyperlink("https://www.mugomes.com.br", sURL)
	lblSite2.Move(fyne.NewPos(lblSite1.MinSize().Width-10, lblDesenvolvedor2.Position().Y+37))

	lblCopyright1 := widget.NewLabel("Copyright (C) 2024-2025 Murilo Gomes Julio. Todos os direitos reservados.")
	lblCopyright1.TextStyle = fyne.TextStyle{Bold: true}
	lblCopyright1.Move(fyne.NewPos(0, lblSite1.Position().Y+37))

	lblLicense1 := widget.NewLabel("License:")
	lblLicense1.TextStyle = fyne.TextStyle{Bold: true}
	lblLicense1.Move(fyne.NewPos(0, lblCopyright1.Position().Y+37))

	lblLicense2 := widget.NewLabel("Gratuito")
	lblLicense2.Move(fyne.NewPos(lblLicense1.MinSize().Width-10, lblCopyright1.Position().Y+37))

	txtLicense := widget.NewRichTextFromMarkdown(`
	Todos os direitos deste software e código-fonte foram reservados pelo Desenvolvedor. Qualquer reprodução, distribuição não autorizada ou uso deste software ou código-fonte sem a devida autorização do desenvolvedor é estritamente proibido e sujeito a ações legais.

	Ao baixar, acessar, instalar, copiar ou utilizar o Software ou código-fonte fornecido pelo Desenvolvedor, o Usuário reconhece que leu, compreendeu e concordou com os termos e condições do Contrato de Licença do Usuário Final (EULA) do MiCheckHash.
	`)
	txtLicense.Wrapping = fyne.TextWrapWord
	
	vBoxLicense := container.NewVScroll(txtLicense)
	vBoxLicense.Move(fyne.NewPos(0, lblLicense1.Position().Y+37))
	vBoxLicense.Resize(fyne.NewSize(597, 257))
	// sURL, _ := url.Parse("https://github.com/mugomes/micheckhash")
	// lblSite2 := &clickText{
	// 	Text: canvas.NewText("https://github.com/mugomes/micheckhash", color.Black),
	// 	URL: sURL,
	// }
	// lblSite2.TextSize = 15
	// lblSite2.Move(fyne.NewPos(lblSite1.MinSize().Width+5, lblSite1.Position().Y))

	// hBox1 := container.NewHBox(lblDesenvolvedor1,lblDesenvolvedor2)
	// hBox2 := container.NewHBox(lblSite1,lblSite2)
	// vBox := container.NewVBox(hBox1, hBox2)

	layout := container.NewWithoutLayout(
		lblSoftware,
		lblDesenvolvedor1,
		lblDesenvolvedor2,
		lblSite1,
		lblSite2,
		lblCopyright1,
		lblLicense1,
		lblLicense2,
		vBoxLicense)

	w.SetContent(layout)
	w.Show()
}
