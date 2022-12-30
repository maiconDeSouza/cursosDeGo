package enderecos

import (
	"fmt"
	"strings"
)

func TipoDeEndereco(end string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "alameda", "rodovia"}

	primeiraPalavraDoEndereco := strings.ToLower(strings.Split(end, " ")[0])

	if isValid := Some(tiposValidos, primeiraPalavraDoEndereco); isValid {
		return fmt.Sprintf("%v", strings.Title(primeiraPalavraDoEndereco))
	}
	return "Inválido!"
}

func Some(slc []string, str string) bool {
	for _, s := range slc {
		if s == str {
			return true
		}
	}
	return false
}
