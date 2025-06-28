package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DAlfaroV/miniscript/internal/interpreter"
	"github.com/DAlfaroV/miniscript/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: miniscript archivo.ms")
		os.Exit(1)
	}

	filename := os.Args[1]

	data, err := os.ReadFile(filepath.Clean(filename))
	if err != nil {
		fmt.Printf("Error leyendo el archivo: %v\n", err)
		os.Exit(1)
	}

	// Parsear el programa
	program, err := parser.Parser.ParseString(filename, string(data))
	if err != nil {
		fmt.Printf("Error al parsear: %v\n", err)
		os.Exit(1)
	}

	// Mostrar AST con delimitadores
	fmt.Println("__BEGIN_AST__")
	astJSON, err := json.MarshalIndent(program, "", "  ")
	if err != nil {
		fmt.Println("Error al serializar AST:", err)
		os.Exit(1)
	}
	fmt.Println(string(astJSON))
	fmt.Println("__END_AST__")

	// Ejecutar el interprete con delimitadores
	fmt.Println("__BEGIN_OUTPUT__")
	env := interpreter.Env{Variables: make(map[string]interface{})}
	err = env.Run(program)
	if err != nil {
		fmt.Println("Error al ejecutar:", err)
		os.Exit(1)
	}
	fmt.Println("__END_OUTPUT__")
}
