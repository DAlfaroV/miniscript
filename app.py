# app.py
from flask import Flask, request, render_template_string
import subprocess
import tempfile

app = Flask(__name__)

HTML = """
<h1>MiniScript online</h1>
<form method="post">
<textarea name="code" rows="20" cols="80">{{ code }}</textarea><br>
<input type="submit" value="Run">
</form>
<h2>Output:</h2>
<pre>{{ output }}</pre>
"""

@app.route("/", methods=["GET", "POST"])
def index():
    code = ""
    output = ""
    if request.method == "POST":
        code = request.form["code"]
        # guardar en archivo temporal
        with tempfile.NamedTemporaryFile(suffix=".ms", delete=False) as tf:
            tf.write(code.encode())
            tf.flush()
            # correr go run main.go archivo
            result = subprocess.run(
                ["go", "run", "main.go", tf.name],
                capture_output=True, text=True
            )
            output = result.stdout + result.stderr
    return render_template_string(HTML, code=code, output=output)

if __name__ == "__main__":
    app.run(debug=True)
