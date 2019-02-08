# Python Version of the Elevator Server
This doc should give you links to everything you need to know to develop
the elevator-server code in Python (currently only tested in Python 3.6).

# Generate
We are using the [`swagger_py_codegen`](https://github.com/guokr/swagger-py-codegen)
tool to build the generic [Flask](http://flask.pocoo.org/) and
[Flask-RESTful](http://flask-restful-cn.readthedocs.org/en/0.3.5/) scaffolding
for our project.

Basically, you need to install the tool:
`$ pip install swagger-py-codegen`

First, let's get the repo:
```
$ git clone git@github.com:OnBeep/elevator-server.git
```
Then, to run the generation:

```
$ cd $REPO_ROOT
$ swagger_py_codegen -s ./elevator.yml py-ele
$ cd py-ele
$ pip install -r requirements.txt
```

FYI: The force-overridden files have already been added to the `.gitignore` file.

To run the server:
```
$ cd $REPO_ROOT/py-ele/py_ele
$ python __init__.py
```

To see the welcome message from a running server:
```
$ curl http://127.0.0.1:5000/v1/welcome
```

To run the unit tests:
```
$ cd $REPO_ROOT/py-ele/py_ele
$ pytest v1/api/tests/
```
