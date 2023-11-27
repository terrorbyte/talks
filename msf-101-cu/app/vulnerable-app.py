#!/usr/bin/env python

from flask import Flask, request, send_from_directory, send_file
import yaml

app = Flask(__name__)

@app.route('/')
def index_html():
    return send_file('static/index.html')

@app.route('/<path:path>')
def index(path):
    return send_from_directory('static', path)

@app.route('/check', methods=['POST'])
def yaml_post():
    if request.method == 'POST':
        f = request.files['yaml']
        try:
            yaml.load(f, Loader=yaml.UnsafeLoader)
            return 'That content was very valid, congrats'
        except yaml.YAMLError:
            return 'That content was not valid, wtf did you just send me?'

def run():
    app.run(host="192.168.100.2", port=8000, debug=True, passthrough_errors=True, use_debugger=False, use_reloader=False)

if __name__ == "__main__":
    run()
