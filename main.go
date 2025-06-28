package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/DAlfaroV/miniscript/internal/ast"
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

	program, err := parser.Parser.ParseString("", string(data))
	if err != nil {
		fmt.Printf("Error de parsing: %v\n", err)
		os.Exit(1)
	}

	// Generar el .c
	output := compileToC(program)

	outfile := strings.TrimSuffix(filepath.Base(filename), ".ms") + ".c"
	err = os.WriteFile(outfile, []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error escribiendo archivo C: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Código C generado en: %s\n", outfile)

	// Llamar a gcc para compilar
	cmdline := exec.Command("gcc", outfile, "runtime.c", "-o", strings.TrimSuffix(outfile, ".c"))
	cmdline.Stdout = os.Stdout
	cmdline.Stderr = os.Stderr
	if err := cmdline.Run(); err != nil {
		fmt.Println("Fallo compilación:", err)
		os.Exit(1)
	}

	fmt.Println("Listo. Ejecuta ./" + strings.TrimSuffix(outfile, ".c"))
}

// compileToC genera el código C completo
func compileToC(p *ast.Program) string {
	var b strings.Builder

	b.WriteString(`#include "runtime.h"` + "\n\n")
	b.WriteString("int main() {\n")

	// primero declarar todas las variables encontradas
	declared := map[string]bool{}
	for _, stmt := range p.Statements {
		if stmt.Assignment != nil {
			if !declared[stmt.Assignment.Name] {
				b.WriteString(fmt.Sprintf("  Value %s;\n", stmt.Assignment.Name))
				declared[stmt.Assignment.Name] = true
			}
		}
	}

	// generar sentencias
	for _, stmt := range p.Statements {
		emitStatement(stmt, &b)
	}

	b.WriteString("  return 0;\n}\n")
	return b.String()
}

// emitStatement genera C para una Statement
func emitStatement(s *ast.Statement, b *strings.Builder) {
	if s.Print != nil {
		expr := emitExpr(s.Print.Expr)
		b.WriteString(fmt.Sprintf("  value_print(%s);\n", expr))
	}
	if s.Assignment != nil {
		emitAssignment(s.Assignment, b)
	}
	if s.If != nil {
		cond := emitExpr(s.If.Condition)
		b.WriteString(fmt.Sprintf("  if (%s.b) {\n", cond))
		for _, inner := range s.If.Body {
			emitStatement(inner, b)
		}
		b.WriteString("  }\n")
	}
	if s.While != nil {
		cond := emitExpr(s.While.Condition)
		b.WriteString(fmt.Sprintf("  while (%s.b) {\n", cond))
		for _, inner := range s.While.Body {
			emitStatement(inner, b)
		}
		b.WriteString("  }\n")
	}
}

// emitAssignment genera C para una asignación
func emitAssignment(a *ast.Assignment, b *strings.Builder) {
	expr := emitExpr(a.Value)
	b.WriteString(fmt.Sprintf("  %s = %s;\n", a.Name, expr))
}

// emitExpr devuelve el string en C que representa la expresión
func emitExpr(e *ast.Expr) string {
	// solo soportaremos Left y Rest simples
	base := emitTerm(e.Left)
	for _, opTerm := range e.Rest {
		right := emitTerm(opTerm.Term)
		// solo operaciones de enteros por ahora
		// podemos expandirlo con una mini-runtime más adelante
		switch opTerm.Op {
		case "+", "-", "*", "/":
			base = fmt.Sprintf("({ Value tmp; tmp.type=VAL_INT; tmp.i=%s.i %s %s.i; tmp; })", base, opTerm.Op, right)
		case "==":
			base = fmt.Sprintf("({ Value tmp; tmp.type=VAL_BOOL; tmp.b=(%s.i == %s.i); tmp; })", base, right)
		case "!=":
			base = fmt.Sprintf("({ Value tmp; tmp.type=VAL_BOOL; tmp.b=(%s.i != %s.i); tmp; })", base, right)
		case ">", "<", ">=", "<=":
			base = fmt.Sprintf("({ Value tmp; tmp.type=VAL_BOOL; tmp.b=(%s.i %s %s.i); tmp; })", base, opTerm.Op, right)
		default:
			base = "({ Value tmp; tmp.type=VAL_INT; tmp.i=0; tmp; }) /* op no soportado */"
		}
	}
	return base
}

// emitTerm devuelve el string en C que representa un término
func emitTerm(t *ast.Term) string {
	if t.Int != nil {
		return fmt.Sprintf("({ Value tmp; tmp.type=VAL_INT; tmp.i=%d; tmp; })", *t.Int)
	}
	if t.Float != nil {
		return fmt.Sprintf("({ Value tmp; tmp.type=VAL_FLOAT; tmp.f=%f; tmp; })", *t.Float)
	}
	if t.String != nil {
		return fmt.Sprintf("({ Value tmp; tmp.type=VAL_STRING; tmp.s=\"%s\"; tmp; })", *t.String)
	}
	if t.Ident != nil {
		return *t.Ident
	}
	return "({ Value tmp; tmp.type=VAL_INT; tmp.i=0; tmp; })"
}
