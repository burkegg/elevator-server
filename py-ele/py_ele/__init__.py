# -*- coding: utf-8 -*-
from flask import Flask

# Handle Python3 import issues
try:
    import v1
except ModuleNotFoundError:
    from . import v1


def create_app():
    app = Flask(__name__, static_folder='static')
    app.register_blueprint(
        v1.bp,
        url_prefix='/v1')
    return app

if __name__ == '__main__':
    create_app().run(debug=True)
