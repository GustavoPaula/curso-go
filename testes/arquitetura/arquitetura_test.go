package arquiteturas

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // Executando teste de forma paralela
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}

	t.Fail()
}
