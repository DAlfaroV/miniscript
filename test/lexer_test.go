package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DAlfaroV/miniscript/internal/lexer"
)

func TestLexerOnExampleFiles(t *testing.T) {
	exampleDir := "examples/"

	pattern := filepath.Join(exampleDir, "*.ms")
	files, err := filepath.Glob(pattern)

	if err != nil {
		t.Fatalf("Error buscando archivos de ejemplo: %v", err)
	}

	if len(files) == 0 {
		t.Fatalf("No se encontraron archivos .ms en %s", exampleDir)
	}

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			contentBytes, err := os.ReadFile(file)

			if err != nil {
				t.Fatalf("No se pudo leer %s: %v", file, err)
			}

			src := string(contentBytes)
			lex := lexer.NewLexer(src)
			tokens, err := lex.ScanTokens()
			if err != nil {
				t.Fatalf("Error léxico en %s: %v", file, err)
			}

			if len(tokens) == 0 {
				t.Errorf("No se produjeron tokens para %s", file)
			}

			last := tokens[len(tokens)-1]
			if last.Type != lexer.TOKEN_EOF {
				t.Errorf("El último token de %s no es EOF, es %v", file, last.Type)
			}
		})
	}
}
