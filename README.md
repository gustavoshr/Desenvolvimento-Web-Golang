# Gopher Store 🛒

A web-based product management application built with Go, developed as a study project from the [Go: create a web application] track on Alura.

---

## About

Gopher Store is a full CRUD application that runs in the browser. You can list, create, edit, and delete products, with all data persisted in a PostgreSQL database.

---

## Features

- View all registered products
- Add new products with name, description, price, and quantity
- Edit existing products
- Delete products

---

## Tech Stack

| Technology | Usage |
|---|---|
| Go | Backend and HTTP server |
| PostgreSQL | Database |
| HTML Templates | Native Go templating |
| Bootstrap 4 | Frontend styling |
| lib/pq | PostgreSQL driver |
| godotenv | Environment variable loading |

---

## Project Structure

```
Gopher Store/
  Base/
    database.go         # PostgreSQL connection via .env
  Front-end/
    index.html          # Product listing
    new-products.html   # Create product form
    edit.html           # Edit product form
    _head.html          # Reusable head component
    _nav.html           # Reusable navbar component
  models/
    products.go         # Struct and database functions
  main.go               # HTTP routes and handlers
  .env                  # Environment variables (not versioned)
```

---

## Getting Started

Create the database and table in PostgreSQL:

```sql
CREATE DATABASE gopher_store;

\c gopher_store

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    quantity INTEGER NOT NULL
);
```

Set up the `.env` file at the project root:

```env
DB_USER=postgres
DB_NAME=
DB_PASSWORD=your_password
DB_HOST=localhost
DB_SSLMODE=disable
```

Install dependencies and run:

```bash
go mod tidy
go run main.go
```

Open in your browser.

---

## What I Learned

This was my first hands-on experience with web development in Go. Throughout the project I learned how to build an HTTP server using the `net/http` package, work with Go's native HTML templates, connect an application to a PostgreSQL database, organize code into packages, and manage environment variables with `godotenv`.

---

Developed by Gustavo Bispo.
