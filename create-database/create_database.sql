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
     VALUES ('El Hobbit', 'Una aventura épica en la Tierra Media', 15, 10, 1, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('Cien años de soledad', 'Una novela mágica sobre el amor, la familia y la historia', 12, 15, 2, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El Señor de los Anillos', 'La trilogía épica de Tolkien', 25, 20, 3, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('1984', 'Una novela distópica sobre un futuro totalitario', 10, 5, 4, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El principito', 'Una historia clásica sobre la amistad y el amor', 8, 12, 5, DATETIME('now'), DATETIME('now'));