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
-- Tener cuidado con caracteres especiales que no acepta SQL lite como apostrofe por ejemplo para posesivo sajon

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El Hobbit', 'Una aventura épica de fantasía', 15, 5, 10, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('Cien años de soledad', 'Una obra maestra de la literatura latinoamericana', 12, 8, 15, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('1984', 'Una distopía clásica sobre el control totalitario', 18, 3, 20, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El Principito', 'Una historia conmovedora sobre la amistad y la imaginación', 10, 12, 25, DATETIME('now'), DATETIME('now'));
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El Señor de los Anillos', 'Una trilogía épica de fantasía', 25, 7, 30, DATETIME('now'), DATETIME('now'));