#### En con_dependencias:
<br>

``` $ go run main.go test/examples/hello_world.ms ```
<br>

#### En con_gui
<br>

Output interprete y AST en terminal:
<br>

``` $ go run main.go test/examples/hello_world.ms ```

Iniciar webapp para usar gui:
<br>

``` $ python app.py ```
<br>

Abrir en navegador: http://127.0.0.1:5000/

####  En compila_c:

En esta rama existe 'main_compilador.go'
Se traduce codigo en archivo .ms a Clang, luego el archivo .c se compila usando gcc

Traduce a C con go:
<br>
``` $ go run main_compilador.go test/examples/hello_world.ms ```

Compila .c:
<br>
``` $ gcc hello_world.c runtime.c -o hello ```

Ejecuta compilado:
<br>
``` ./hello ```
