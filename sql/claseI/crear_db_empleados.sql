CREATE DATABASE emple_dep;

USE emple_dep;

CREATE TABLE empleado (
    n_legajo INT NOT NULL PRIMARY KEY,
    dni INT,
    apellido VARCHAR(32),
    nombre VARCHAR(32),
    fecha_nacimiento DATE,
    fecha_ingreso DATE,
    cargo VARCHAR(32),
    sueldo_neto DOUBLE
);
-- se establecio el n_legajo como PK por ser un identificador unico en la compania

CREATE TABLE departamento (
	id INT NOT NULL PRIMARY KEY,
	nombre VARCHAR(32),
    direccion VARCHAR(64)
);
-- se establecio el id como identificador unico ya que facilita su identificacion

ALTER TABLE empleado ADD COLUMN id_departamento INT;
ALTER TABLE empleado ADD FOREIGN KEY (id_departamento) REFERENCES departamento(id);
-- se establecio una relacion de (empleado)1 a N(departamento)
-- y la FK se adiciono a la tabla empleado con relacion a la tabla departamento

INSERT INTO departamento (id, nombre, direccion)
VALUES
	(1000, 'compras', 'calle 8'),
    (1001, 'people', 'avenida central'),
    (1002, 'soporte', 'calle 38'),
    (1003, 'it', 'calle 8'),
    (1004, 'reparto', 'calle 1');

INSERT INTO empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, fecha_ingreso, cargo, sueldo_neto, id_departamento)
VALUES
	(1200, 12345, 'Perez', 'Carlos', '1980-11-11','2000-11-11','jefe proceso',20000,1002),
    (1204, 123456, 'Martinez','Olga','1975-01-11','2010-01-01','coordinadora', 10000,1000),
    (1205, 12346, 'Caicedo', 'Ferney','1990-02-03','1999-04-05','desarrollador',100000,1003),
    (1206, 28272,'Angel','Luis','1995-01-01','2015-06-30','repartidor',2000,1004),
    (1207,1264,'Velez','Camilo','1991-01-12','2022-06-21','desarrollador',9000,1003),
    (1208,123448,'Moreno','Hector','1992-01-12','2015-01-12','auxiliar',1200,1001),
    (1209,87684,'Paez','Fabio','1993-01-12','2016-01-12','auxiliar',1150,1004),
    (1210,7327,'Rodriguez','Miguel','1994-01-12','2017-01-12','jefe proceso',20000,1001),
    (1211,893,'Aguilar','Joshua','1995-01-12','2018-01-12','jefe proceso',21000,1003),
    (1212,89384,'Prieto','Roberto','1991-01-12','2019-01-12','jefe proceso',22000,1004),
    (1213,83439,'Sanchez','Juan','1991-01-13','2018-01-12','auxiliar',1150,1004),
    (1214,9139,'Duque','Pablo','1992-02-12','2017-01-12','auxiliar',1200,1002),
    (1215,93309,'Ariza','Pablo','1991-01-17','2016-01-12','coordinadora',10000,1003),
    (1216,903920,'Romero','Andres','1994-03-12','2015-01-12','coordinadora',10000,1001),
    (1217,90390,'Castillo','Laura','1992-01-12','2015-01-12','coordinadora',1000,1004);
