## URL SHORTENER
Tech Stack: Golang(Gin, Cobra), MongoDB, Docker

### API
API, written in Gin Webframework, has 1 endpoint "/:id".
After opening http://localhost:PORT/:id, server looks for original url in MongoDB database
and redirect to this site or display "Error": "Url not found" if given shortened url does not exist.
When there is "!" before shortened url server returns html page with "Go to [page]" text,
where [page] is original or "URL not found" if given shortened url does not exist.

## To run:

## LOCALLY:

(GOLANG VERSION 1.18 IS REQUIRED)
- go to "url-shortener" directory
- start mongodb
- run "go mod download"
- set enviorment variables like SHORT_URL_MAX_LEN in .env file, example: SHORT_URL_MAX_LEN=6
- run "go build -o url-shortener-api" ("go build -o url-shortener-api.exe" for Windows)
- start api by running "./url-shortener-api" (".\url-shortener-api.exe" for Windows)
- now mongoDB is working and api is listening at given PORT

## USING DOCKER:

- go to "url-shortener" directory
- set enviorment variables in Dockerfile (lines starting with "ENV", set values after "=" symbol, example: ENV SHORT_URL_MAX_LEN=6)
- run "docker build -t url-api ."
- run "cd .."
- run "docker-compose up" in websensa directory, it starts api and mongodb, add " -d" at the end if you want to run app in background
- now api is working and listening at given PORT
- run "docker-compose down" to make api and database stop working

## CLI

Written in cobra. There is 4 commands: add, update, delete, list.
You can run cli app with -h or --help flag to see all commands and their description
Usage:

- **add**
  command is used to adding new url to database  
  First method:  
  program will ask user if he want custom url or random one  
  Example:  
  url-shortener add http://example.com  
  Second method:  
  Flags:  
  -r, --random: for random shortened url  
  -c, --custom [value]: for value given by user  
  Example:  
  url-shortener add http://example.com -c xyz123  
  url-shortener add http://example1.com --random  
- **list**
command shows all created urls  
Example:  
url-shortener list  
- **delete**
command deletes url from database, user type shortened url to specify which one is gonna be deleted.  
User can give more than one shortened url with spaces  
Example:  
url-shortener delete xyz123  
url-shortener delete xyz123 example  
- **update**
command updates already created url, user need to type shortened url which one is gonna be updated.  
At least 1 flag is required.  
Flags:  
-s, --short-url [value] : new shortened url  
-u, --url [value] : new url  
Example:  
url-shortener update xyz123 -s 123abc  
url-shortener update xyz123 -s 111aaa -u http://allegro.pl  

To run:

(GOLANG VERSION 1.18 IS REQUIRED)
- do API part first (we need running database to use cli)
- go to "cli" directory
- run "go mod download"
- set enviorment variables like API_PORT or MONGO_URI in .env file, example: SHORT_URL_MAX_LEN=6

on Linux or MacOS:
- run "go build -o url-shortener"
- cli app is ready to use in "cli" directory, to do it run "./url-shortener [command] [flags]"

on Windows:
- run "go build -o url-shortener.exe"
- cli app is ready to use in "cli" directory, to do it run ".\url-shortener.exe [command] [flags]"
