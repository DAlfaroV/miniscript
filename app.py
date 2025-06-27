from flask import Flask, render_template, request
import subprocess
import tempfile
import os

app = Flask(__name__)

@app.route("/", methods=["GET", "POST"])
def index():
    ast_result = ""
    output_result = ""
    code_input = ""

    if request.method == "POST":
        code_input = request.form.get("code")

        # Guardar el código en un archivo temporal
        with tempfile.NamedTemporaryFile(suffix=".ms", delete=False, mode="w") as tmp:
            tmp.write(code_input)
            tmp_filename = tmp.name

        try:
            # Ejecutar el compilador/interprete
            result = subprocess.run(
                ["go", "run", "main.go", tmp_filename],
                capture_output=True,
                text=True
            )
            stdout = result.stdout

            # Parsear bloques delimitados
            if "__BEGIN_AST__" in stdout and "__BEGIN_OUTPUT__" in stdout:
                ast_part = stdout.split("__BEGIN_AST__")[1].split("__END_AST__")[0].strip()
                output_part = stdout.split("__BEGIN_OUTPUT__")[1].split("__END_OUTPUT__")[0].strip()
                ast_result = ast_part
                output_result = output_part
            else:
                output_result = "No se pudo ejecutar correctamente."

        except Exception as e:
            output_result = f"Error ejecutando: {e}"

        finally:
            os.remove(tmp_filename)

    return render_template(
        "index.html",
        ast=ast_result,
        output=output_result,
        code=code_input
    )

if __name__ == "__main__":
    app.run(debug=True)
