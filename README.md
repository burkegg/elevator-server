# elevator-server
Example code for interviews

# Welcome to the Elevator Server Project
In this repo we are developing a standardized API for the control system of an
elevator installation with 2-N floors and 1-M cars. Consider the consumers
of this API to be some kind of tablet that has HTTP(S) access to the server
on a protected network, both from inside the elevator car, and
on the floors where the elevator can get called.

# Swagger
For this API, we are using Swagger (newly aka Open API Initiative) as a type
of contract language for defining the API (see `elevator.yml`). Here are some
links that might be helpful:
* [Swagger Homepage](http://swagger.io/)
* [Swagger Spec](http://swagger.io/specification/)
* [JSON Schema Homepage](http://json-schema.org/)
* [JSON Schema Validation](http://json-schema.org/latest/json-schema-validation.html)
* [HTTP Status Codes](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html)

# Generate
There are several benefits to using a formalized contract language like Swagger.
One is that it is programming-language agnostic, meaning that it can be a good
meeting place for API producers and consumers to discuss pending changes.
Another is that you can use code to automagically generate what would be
boilerplate code. This is what we will play with in these exercises. Drop into
either the `py-ele` or `go-ele` directories for instructions on running the
generation.

# Design & Implementation Tasks
We chose to model an elevator in the hopes that the interactions are fairly well
understood by just about everyone. There will be some secondary/support
functionalities already
provided, but you'll need to design and implement some of the main functionalities.
For this exercise, we've provided unit tests for functionality to determine which car is
closest to a given floor and to call the closest car to a given floor. Your task will be to
design and implement endpoints for this functionality while ensuring your implementation
passes the unit tests. Please feel free to complete this challenge in either Python or Go,
whichever is more comfortable for you.
Here are a few questions to get you thinking about what might be needed:
* When calling a car to a floor, what information does the controller need to know?
* What 'read-only' types of information should we plan for?
* What interaction patterns should be expected from the endpoints?
