// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package controls

import "github.com/mugomes/mglang"

func LoadTranslations() {
	lang := mglang.GetLang()
	if lang == "pt" {
		mglang.Set("File", "Arquivo")
		mglang.Set("Save", "Salvar")
		mglang.Set("Hash Type", "Tipo de Hash")
		mglang.Set("Select file", "Selecione o arquivo")
		mglang.Set("Generate Hash", "Gerar Hash")
		mglang.Set("Success!", "Sucesso!")
		mglang.Set("Different!", "Diferente")
		mglang.Set("Tools", "Ferramentas")
		mglang.Set("About", "Sobre")
		mglang.Set("Check Update", "Verificar Atualização")
		mglang.Set("Support MiCheckHash", "Apoie MiCheckHash")
		mglang.Set("About MiCheckHash", "Sobre MiCheckHash")
		mglang.Set("Type/Paste the Hash", "Digite/Cole o Hash")
		mglang.Set("Check Now", "Verificar Agora")
		mglang.Set("Verifying Hash... Please wait!", "Verificando Hash... Aguarde!")
		mglang.Set("Generating Hash... Please wait!", "Gerando Hash... Aguarde!")
		mglang.Set("File created successfully!", "Arquivo criado com sucesso!")
	}
}

func T(key string, args ...interface{}) string {
	return mglang.T(key, args...)
}
