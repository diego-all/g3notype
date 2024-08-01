-- DDL statements
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


-- DML statements [Dummy data] by Gemini

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('The Hitchhikers Guide to the Galaxy', 'A humorous science fiction novel by Douglas Adams', 10, 100, 12345, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('Pride and Prejudice', 'A romantic novel by Jane Austen', 8, 50, 67890, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('1984', 'A dystopian novel by George Orwell', 12, 75, 24680, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('To Kill a Mockingbird', 'A novel by Harper Lee', 9, 120, 13579, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('The Lord of the Rings', 'An epic high fantasy novel by J. R. R. Tolkien', 15, 150, 56789, DATETIME('now'), DATETIME('now'));