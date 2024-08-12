-- DDL statements
CREATE TABLE IF NOT EXISTS products (
 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(100) NOT NULL,
	description VARCHAR(100) NOT NULL,
	price INTEGER NOT NULL,
	quantity INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT DATETIME NOT NULL,
 	updated_at TIMESTAMP NOT NULL
 	);


-- DML statements [Dummy data] by Gemini
-- Tener cuidado con caracteres especiales que no acepta SQL lite como apostrofe por ejemplo para posesivo sajon

INSERT INTO products (name, description, price, quantity, created_at, updated_at)
VALUES ('Teléfono móvil', 'Smartphone de última generación', 799, 10, DATETIME('now'), DATETIME('now'));
INSERT INTO products (name, description, price, quantity, created_at, updated_at)
VALUES ('Camiseta', 'Camiseta de algodón', 20, 50, DATETIME('now'), DATETIME('now'));
INSERT INTO products (name, description, price, quantity, created_at, updated_at)
VALUES ('Sartén antiadherente', 'Sartén para cocinar', 35, 20, DATETIME('now'), DATETIME('now'));
INSERT INTO products (name, description, price, quantity, created_at, updated_at)
VALUES ('Balón de fútbol', 'Balón oficial de la FIFA', 50, 15, DATETIME('now'), DATETIME('now'));
INSERT INTO products (name, description, price, quantity, created_at, updated_at)
VALUES ('Muñeca', 'Muñeca de peluche para niños', 15, 30, DATETIME('now'), DATETIME('now'));