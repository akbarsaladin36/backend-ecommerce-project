<h1 align="center">REST API Ecommerce - Backend</h1>

This API is created by me for build my personal project that will be helpful in future. This backend api using Golang with Fiber and GORM package. [More about Golang](https://go.dev/)

## Built With

[![Go/Golang](https://img.shields.io/badge/Golang-1.23.2-cyan.svg?style=rounded-square)](https://go.dev/)
[![Fiber](https://img.shields.io/badge/Fiber-v.2.52.8-blue.svg?style=rounded-square)](https://gin-gonic.com/docs/)
[![GORM](https://img.shields.io/badge/Gorm-v.1.30.0-blue.svg?style=rounded-square)](https://gorm.io/)


## Requirements

1. <a href="https://go.dev/">Go</a>
2. <a href="https://gofiber.io/">Gin</a>
3. <a href="https://gorm.io/">GORM</a>
4. <a href="https://www.mysql.com/">MySQL</a>

## How to run the app ?

1. Clone this project
2. Open app's directory in CMD or Terminal
3. Turn on Web Server and MySQL can using Third-party tool like xampp, etc.
4. Create a database with the name #nama_database, and Import file sql to **phpmyadmin**
5. Open Postman desktop application or Chrome web app extension that has installed before
6. Choose HTTP Method and enter request url.(ex. localhost:3600/)
7. You can see all the end point [here](https://documenter.getpostman.com/view/14780095/2sAXjM3XAX)
8. Type `go run main.go` to activated the server.
9. If you have "air" package installed, then you can type `air server --port [YOUR_APP_PORT]`.(ex. "air server --port 3600")

## Set up project

1. Create a file ".env" then setup environment variable like code below:

```
APP_PORT=<FIBER_LISTENING_PORT>
DATABASE_HOSTNAME=<YOUR_DB_HOSTNAME>
DATABASE_PORT=<YOUR_DB_PORT>
DATABASE_USERNAME=<YOUR_DB_USERNAME>
DATABASE_PASSWORD=<YOUR_DB_PASSWORD>
DATABASE_NAME=<YOUR_DATABASE_NAME>
JWT_KEY=<YOUR_SECRET_KEY_JWT>
```

## Feature

1. Authentication (Login & Register User)
2. CRUD :
    1. Admin :
        1. Users
        2. Products
        3. Carts
        4. Invoices
        5. Profile
    2. User :
        1. Products (Get Products & Product Detail)
        2. Carts (Get Carts By User, Cart Detail, Create Cart)
        3. Invoices (Get Invoices By User, Invoice Detail, Create Invoice)
        4. Profile

## License

© [Muhammad Akbar Saladin Siregar](https://github.com/akbarsaladin36/)
