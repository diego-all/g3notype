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
