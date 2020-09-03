package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	// ...
	t.Fail()
}

// para rodar os testes da biblioteca padrão : go test std (não recomendavel)
// para rodar todos os testes de uma pasta: go test ./...
