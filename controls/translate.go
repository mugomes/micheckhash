// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package controls

import (
	"fmt"
	"os"
	"strings"
)

type Translations map[string]string

var tr Translations

// Detecta o idioma do sistema e retorna apenas o código base (pt, en, es, etc.)
func GetSystemLanguage() string {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if lang == "" {
		lang = os.Getenv("LC_MESSAGES")
	}
	if lang == "" {
		return "en" // fallback padrão
	}

	// Exemplo: "pt_BR.UTF-8" → "pt"
	parts := strings.Split(lang, ".")
	base := parts[0]
	baseParts := strings.Split(base, "_")
	return strings.ToLower(baseParts[0])
}

func LoadTranslations() {
	if GetSystemLanguage() == "pt" {
		valor := make(map[string]string)
		valor["File"] = "Arquivo"
		valor["Save"] = "Salvar"
		valor["Hash Type"] = "Tipo de Hash"
		valor["Select file"] = "Selecione o arquivo"
		valor["Generate Hash"] = "Gerar Hash"
		valor["Success!"] = "Sucesso!"
		valor["Different!"] = "Diferente"
		valor["Tools"] = "Ferramentas"
		valor["About"] = "Sobre"
		valor["Check Update"] = "Verificar Atualização"
		valor["Support MiCheckHash"] = "Apoie MiCheckHash"
		valor["Technical Support"] = "Suporte Técnico"
		valor["About MiCheckHash"] = "Sobre MiCheckHash"
		valor["Type/Paste the Hash"] = "Digite/Cole o Hash"
		valor["Check Now"] = "Verificar Agora"
		valor["Verifying Hash... Please wait!"] = "Verificando Hash... Aguarde!"
		valor["Generating Hash... Please wait!"] = "Gerando Hash... Aguarde!"
		tr = valor
	}
}

// T retorna o texto traduzido com formatação opcional.
func T(key string, args ...interface{}) string {
	msg, ok := tr[key]
	if !ok {
		msg = key // fallback se não achar
	}
	return fmt.Sprintf(msg, args...)
}
