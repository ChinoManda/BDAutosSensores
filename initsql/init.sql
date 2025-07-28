CREATE DATABASE IF NOT EXISTS carsdb;
use carsdb;

CREATE TABLE brands(
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50)
);

CREATE TABLE models(
  id INT AUTO_INCREMENT PRIMARY KEY,
  brand_id INT,
  FOREIGN KEY(brand_id) REFERENCES brands(id),
  name VARCHAR(50)
);

CREATE TABLE engines(
  id INT AUTO_INCREMENT PRIMARY KEY, 
  name VARCHAR(50),
  cilinders INT
);

CREATE TABLE displacements(
  id INT AUTO_INCREMENT PRIMARY KEY,
  CC float
);

CREATE TABLE cars(
  id INT AUTO_INCREMENT PRIMARY KEY,
  year INT,
  brand_id INT,
  model_id INT,
  engine_id INT,
  displacement_id INT,
  FOREIGN KEY(brand_id) REFERENCES brands(id),
  FOREIGN KEY(model_id) REFERENCES models(id),
  FOREIGN KEY(engine_id) REFERENCES engines(id),
  FOREIGN KEY(displacement_id) REFERENCES displacements(id)
);

INSERT INTO engines(name, cilinders) VALUES ('AP', 4), ('Tipo', 4), ('TSI', 4), ('4AGE', 4), ('O7K', 5);

INSERT INTO displacements(CC) VALUES (1.0), (1.1), (1.2), (1.3), (1.4), (1.5), (1.6), (1.7), (1.8), (1.9), (2.0), (2.5);


INSERT INTO brands(name) VALUES 
  ('Volkswagen'),  -- id = 1
  ('Fiat'),        -- id = 2
  ('Toyota'),      -- id = 3
  ('Nissan'),      -- id = 4
  ('Peugeot'),     -- id = 5
  ('Honda');       -- id = 6

INSERT INTO models(name, brand_id) VALUES
-- Volkswagen
('Gol', 1),
('Golf', 1),
('Bora', 1),
('Vento', 1),
('Nivus', 1),
('Polo', 1),
('Passat', 1),

-- Fiat
('Uno', 2),
('147', 2),
('128', 2),

-- Toyota
('Corolla', 3),
('Supra', 3),
('Starlet', 3),

-- Nissan
('Sentra', 4),
('Silvia', 4),

-- Peugeot
('206', 5),
('207', 5),
('307', 5),
('205', 5),
('306', 5),
('106', 5),

-- Honda
('Civic', 6),
('Accord', 6),
('Prelude', 6),
('CR-V', 6);


INSERT INTO cars(year, brand_id, model_id, engine_id ,displacement_id) VALUES ('1993', 1, 1, 1, 9);

INSERT INTO cars(year, brand_id, model_id, engine_id ,displacement_id) VALUES ('1991', 2, 9, 2, 5);
INSERT INTO cars(year, brand_id, model_id, engine_id ,displacement_id) VALUES ('2005', 6, 21, 3 ,5);
INSERT INTO cars(year, brand_id, model_id, engine_id ,displacement_id) VALUES ('1986', 3, 11, 4 ,7);
INSERT INTO cars(year, brand_id, model_id, engine_id ,displacement_id) VALUES ('2018', 1, 3, 5 ,12);
