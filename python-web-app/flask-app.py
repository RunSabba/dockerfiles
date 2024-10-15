from flask import Flask, render_template, request

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def home():
    greeting = ""
    if request.method == 'POST':
        name = request.form['name']
        greeting = f"Hello, {name}! Welcome to RUNSABBA CLOUD!"
    return render_template('index.html', greeting=greeting)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)

