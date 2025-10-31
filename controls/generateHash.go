// MiCheckHash
// Copyright (C) 2025 Murilo Gomes Julio. Todos os direitos reservados.

// Este software e código-fonte é distribuído sob os termos do Contrato de Licença de Usuário Final do MiCheckHash.

package controls

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
)
	

func GetHash(sTipoHash string, file io.Reader) string {
	var hashsum hash.Hash

	switch sTipoHash {
	case "md5":
		hashsum = md5.New()
	case "sha1":
		hashsum = sha1.New()
	case "sha256":
		hashsum = sha256.New()
	case "sha512":
		hashsum = sha512.New()
	default:
		hashsum = md5.New()
	}

	if _, err := io.Copy(hashsum, file); err != nil {
		return ""
	}

	hashInBytes := hashsum.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}