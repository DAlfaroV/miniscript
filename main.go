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

	program, err := parser.Parser.ParseString(filename, string(data)) // Esta por el print de punteros original.
	println("\n ============== AST ============ \n")
	// Usnado json para mostrar arbol en vez de punteros hexa
	output, err := json.MarshalIndent(program, "", "  ")
	if err != nil {
		fmt.Println("Error al serializar AST:", err)
		os.Exit(1)
	}
	fmt.Println(string(output))

	println("\n\n\n ============== OUTPUT de interprete ============ \n")
	env := interpreter.Env{Variables: make(map[string]interface{})}
	err = env.Run(program)
	if err != nil {
		fmt.Println("Error al ejecutar:", err)
		os.Exit(1)
	}

}
