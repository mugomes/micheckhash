// MiCheckHash
// Copyright (C) 2025 Murilo Gomes Julio. Todos os direitos reservados.

// Este software e código-fonte é distribuído sob os termos do Contrato de Licença de Usuário Final do MiCheckHash.

package controls

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

func LoadTranslations() error {
	// getPath()
	var sPath string = "langs/"
	if runtime.GOOS == "linux" {
		if IsDevMode() {
			sPath, _ = os.Getwd()
		} else {
			exe, _ := os.Executable()
			sPath = filepath.Dir(exe)
		}

		sPath += "/langs/"
	}
	file, err := os.Open(sPath + GetSystemLanguage() + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&tr)
}

// T retorna o texto traduzido com formatação opcional.
func T(key string, args ...interface{}) string {
	msg, ok := tr[key]
	if !ok {
		msg = key // fallback se não achar
	}
	return fmt.Sprintf(msg, args...)
}
