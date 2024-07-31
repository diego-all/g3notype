Tengo una API construida en golang, y me esta presentando un error:

{
    "error": true,
    "message": "sql: Scan error on column index 2, name \"descripcion\": converting driver.Value type string (\"Limpiador random\") to a int: invalid syntax"
}

Podrias explicarme el error y ayudarme a solucionarlo dandome la respuesta en español.

Estoy intentando utiliza el endpoint http://localhost:9090/books/update/7

Aca el archivo handlers-Book.go con: 

func (app *application) UpdateBook(w http.ResponseWriter, r *http.Request) {

	var bookReq bookRequest
	var payload jsonResponse

	err := app.readJSON(w, r, &bookReq)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	bookID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_, err = app.models.Book.GetOneById(bookID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var book = models.Book{
	Nombre:	bookReq.Nombre,
	Descripcion:	bookReq.Descripcion,
	Precio:	bookReq.Precio,
	Cantidad:	bookReq.Cantidad,
	Random:	bookReq.Random,
	UpdatedAt:   time.Now(),
 	Id:          bookID,
}

	_, err = app.models.Book.Update(book)
	if err != nil {
		app.errorJSON(w, err)
		return
	}


	payload = jsonResponse{
	    Error:   false,
	    Message: "Book successfully updated",
	    Data:    envelope{"book": book.Nombre},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

{
"nombre": "oelo",
"descripcion": "oelo",
"precio": 10000000,
"cantidad":500,
"random": "random saludos" 
}


Books.go

package models

import (
	"context"
	"time"
)

type Book struct {
 	Id	int	`json:"id"`
	Nombre	string	`json:"nombre"`
	Descripcion	string	`json:"descripcion"`
	Precio	int	`json:"precio"`
	Cantidad	int	`json:"cantidad"`
	Random	string	`json:"random"`
	CreatedAt   time.Time `json:"created_at"`
 	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Book) Update(book Book) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel() // resource leaks

	stmt := `update book set
 	Nombre = $1,
	Descripcion = $2,
	Precio = $3,
	Cantidad = $4,
	Random = $5,
	updated_at = $7
 	where id = $8`

	_, err := db.ExecContext(ctx, stmt,
	book.Nombre,
	book.Descripcion,
	book.Precio,
	book.Cantidad,
	book.Random,
	time.Now(),
 	book.Id,
)

	if err != nil {
		return 0, err
	}

	return 0, nil
}
}



routes.go

package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/health", app.Health)

	// Book
	mux.Post("/books", app.CreateBook)
	mux.Get("/books/get/{id}", app.GetBook)
	mux.Put("/books/update/{id}", app.UpdateBook)
	mux.Get("/books/all", app.AllBooks)
	mux.Delete("/books/delete/{id}", app.DeleteBook)

	return mux
}


Aca estan lso archivos para crear la base de datos:

CREATE TABLE IF NOT EXISTS books (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
   nombre VARCHAR(100) NOT NULL,
   descripcion VARCHAR(100) NOT NULL,
   precio INTEGER NOT NULL,
   cantidad INTEGER NOT NULL,
   random VARCHAR(100) NOT NULL,
   created_at TIMESTAMP DEFAULT DATETIME NOT NULL,
	updated_at TIMESTAMP NOT NULL
	);

-- GENERADO CON SCRIPT 

-- DML statements [Dummy data]
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El Hobbit', 'Una aventura épica de Tolkien', 15, 10, 10, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('Cien años de soledad', 'Una novela mágica de García Márquez', 20, 15, 20, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('cualquier cosa', 'Una distopía clásica de Orwell', 12, 8, 15, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El principito', 'Una historia conmovedora de Saint-Exupéry', 10, 12, 5, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('Matar a un ruiseñor', 'Un clásico de la literatura americana', 18, 10, 25, DATETIME('now'), DATETIME('now'));