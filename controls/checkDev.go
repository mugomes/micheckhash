// MiCheckHash
// Copyright (C) 2025 Murilo Gomes Julio. Todos os direitos reservados.

// Este software e código-fonte é distribuído sob os termos do Contrato de Licença de Usuário Final do MiCheckHash.

package controls

import (
	"os"
	"strconv"
)

func IsDevMode() bool {
	b, _ := strconv.ParseBool(os.Getenv("MIDEV"))
	return b
}