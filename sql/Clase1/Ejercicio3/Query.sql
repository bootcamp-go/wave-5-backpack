CREATE DATABASE emple_dep;

USE emple_dep;

CREATE TABLE empleado(
	NoLegajo INT,
    DNI VARCHAR(30),
    Apellido VARCHAR(25),
    Nombre VARCHAR(25),
    FechaNacimiento DATE,
    FechaIncorporacion DATE,
    Cargo VARCHAR(30),
    SueldoNeto INT
);

CREATE TABLE departamento (
	ID INT NOT NULL PRIMARY KEY,
	Nombre VARCHAR(32),
    Direccion VARCHAR(64)
);

ALTER TABLE empleado ADD COLUMN IDDepartamento INT;
ALTER TABLE empleado ADD FOREIGN KEY (IDDepartamento) REFERENCES departamento(ID);

INSERT INTO departamento (ID, Nombre, Direccion)
VALUES
	(1, 'Ventas', 'calle 1'),
    (2, 'Contabilidad', 'calle 1'),
    (3, 'Limpieza', 'calle 1'),
    (4, 'Mantenimiento', 'calle 2'),
    (5, 'Soporte', 'calle 2');

INSERT INTO empleado (NoLegajo, DNI, Apellido, Nombre, FechaNacimiento, FechaIncorporacion, Cargo, SueldoNeto, IdDepartamento)
VALUES
	(1,15246,'Perez', 'Carlos', '1980-11-11','2000-11-11','jefe proceso',20000,1),
    (2,17696,'Martinez','Olga','1975-01-11','2010-01-01','coordinadora', 10000,3),
    (3,87425,'Caicedo', 'Ferney','1990-02-03','1999-04-05','desarrollador',100000,3),
    (4,78462,'Angel','Luis','1995-01-01','2015-06-30','repartidor',2000,4),
    (5,54547,'Velez','Camilo','1991-01-12','2022-06-21','desarrollador',4,5),
    (6,58945,'Moreno','Hector','1992-01-12','2015-01-12','auxiliar',1200,1),
    (7,47001,'Paez','Fabio','1993-01-12','2016-01-12','auxiliar',1150,2),
    (8,65489,'Rodriguez','Miguel','1994-01-12','2017-01-12','jefe proceso',20000,4),
    (9,15481,'Aguilar','Joshua','1995-01-12','2018-01-12','jefe proceso',21000,5),
    (10,74511,'Prieto','Roberto','1991-01-12','2019-01-12','jefe proceso',22000,5),
    (11,63256,'Sanchez','Juan','1991-01-13','2018-01-12','auxiliar',1150,5),
    (12,98547,'Duque','Pablo','1992-02-12','2017-01-12','auxiliar',1200,2),
    (13,21578,'Ariza','Pablo','1991-01-17','2016-01-12','coordinadora',10000,3),
    (14,97822,'Romero','Andres','1994-03-12','2015-01-12','coordinadora',10000,4),
    (15,12024,'Castillo','Laura','1992-01-12','2015-01-12','coordinadora',1000,1);