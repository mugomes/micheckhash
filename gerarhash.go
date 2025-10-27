// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.mugomes.com.br

package main

import "fyne.io/fyne/v2"

func showGerarHash(a fyne.App) {
	w := a.NewWindow("Gerar Hash")
	w.Resize(fyne.NewSize(400, 300))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	w.Show()
}