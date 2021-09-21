![tests](https://unal.edu.co/golang-rest-example/workflows/Unit%20Testing/badge.svg)

# Golang REST API Example

This is a complete example of a production-ready REST API built with **Golang**, **Gin**, **Zap** and another technologies such as **JWT**

## Config the environment
The program parses a single `.yalm` and set the proper variables into the program.

Copy the `config/example.config.yaml` file and paste it in the same folder, but with one of the following two nams:

- `production.config.yaml`
- `dev.config.yaml`

and update the content in it according your settings. 

## Build and Run
Use the Makefile as reference, e.g:
   
    $ make dev
    > Run the app in development mode
or
    
    $ make production
    $ make production-run
    > The first one builds the project into the dist/ foder.
    > The second one run it.