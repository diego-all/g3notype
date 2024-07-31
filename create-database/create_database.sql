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
     VALUES ('El Hobbit', 'Aventura épica de Tolkien', 15, 10, 1234, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('Cien años de soledad', 'Novela clásica de García Márquez', 12, 20, 5678, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('1984', 'Distopía de Orwell', 10, 15, 9012, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El señor de los anillos', 'Trilogía épica de Tolkien', 25, 5, 3456, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El principito', 'Cuento clásico de Saint-Exupéry', 8, 30, 7890, DATETIME('now'), DATETIME('now'));