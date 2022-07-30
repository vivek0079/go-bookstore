# Go Bookstore

### Introduction

This is a book management project build using golang and gorm.

GORM is an ORM (Object–relational mapping) library for Golang. Read about GORM [here](https://gorm.io/)

### Setup

In this project we have used mysql as the database. 

**Step 1: [Checkout project]**

Checkout the project to your local using

`git checkout https://github.com/vivek0079/go-bookstore.git`

**Step 2: [Setup mysql in docker]**

* Pull an image of mysql from docker hub using 

`docker pull mysql/mysql-server:latest`

* Run the mysql image using 

`docker run -p 13306:3306 --name=[container_name] -d [image_tag_name]`

* Get the mysql password from the docker logs using 

`docker logs [container_name] | grep GENERATED ROOT PASSWORD`

* Save this password and replace the same in `app.go`. 


* If needed you can set allow empty password by passing `MYSQL_ALLOW_EMPTY_PASSWORD` as env var. 
**Note: However this is recommended for local development.** 

**Step 3: [Start Golang server]**

Start the golang server using 

`go build` followed by `go run main.go`

**Step 4: [Hack]**

Now you can use postman to play around with the APIs :)
