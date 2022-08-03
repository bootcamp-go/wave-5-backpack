/*-----------------------------------------------------------------------*

     Assignment:	Ejercicio #3:  Base de Datos Relacionales
         Author:	Israel Fabela
	   Language:	mysql  Ver 8.0.29 for macos12.2 on arm64
		  Topic:	Base de Datos - SQL

	© Mercado Libre - IT Bootcamp 2022

-------------------------------------------------------------------------*/

-- Se crea base de datos para 'emple_dep'
CREATE DATABASE emple_dep;

-- Se usa la base de datos 'emple_dep'
USE emple_dep;

-- Creando tabla de 'departamento'
CREATE TABLE departamento (
	id_departamento INT NOT NULL AUTO_INCREMENT,
    nombre varchar(50) NOT NULL,
    direccion varchar(80) NULL,
    PRIMARY KEY (id_departamento)
);

-- Creando tabla de 'empleados'
CREATE TABLE empleados (
	no_legajo int NOT NULL AUTO_INCREMENT,
	id_departamento INT DEFAULT NULL,
    nombre varchar(50) NOT NULL,
    apellido varchar(50) NOT NULL,
    fecha_nacimiento DATETIME NOT NULL,
    fecha_incorporacion DATETIME NOT NULL,
    cargo varchar(70) DEFAULT NULL,
    sueldo_neto dec NOT NULL,
	dni varchar(40) DEFAULT NULL,
    PRIMARY KEY (no_legajo),
    FOREIGN KEY(id_departamento) REFERENCES departamento(id_departamento)
);

-- Insertando los datos a 'departamento'
INSERT INTO departamento (nombre,direccion) VALUES ('Data','Edificio A');
INSERT INTO departamento (nombre,direccion) VALUES ('ML','Edificio B');
INSERT INTO departamento (nombre,direccion) VALUES ('QA','Edificio C');
INSERT INTO departamento (nombre,direccion) VALUES ('SOFTWARE','Edificio D');
INSERT INTO departamento (nombre,direccion) VALUES ('SECURITY','Edificio E');

-- Insertando los datos a 'empleados'
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (1,'Michelle','Herring','1961-07-27','2014-12-03',55535,'ABC001');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (2,'Hammond','Fletcher','1975-06-13','2015-06-28',72119,'ABC002');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (3,'Hart','Silva','2002-11-30','2021-01-16',36260,'ABC003');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (4,'Rosie','Howard','1954-03-28','2012-03-19',50267,'ABC004');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (5,'Schwartz','Park','1962-04-11','2012-03-31',61693,'ABC005');

INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (1,'Herring','Butler','1970-07-27','2007-01-09',61080,'ABC006');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (2,'Elma','Sears','1987-10-08','2003-06-24',73263,'ABC007');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (3,'Penelope','Chavez','1971-06-15','2017-04-25',83104,'ABC008');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (4,'Lewis','Franco','1995-08-27','2021-02-16',40932,'ABC009');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (5,'Mcmillan','Head','1986-12-19','2012-02-06',61693,'ABC010');

INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (1,'Browning','Schneider','1969-10-09','1996-10-12',28302,'ABC011');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (2,'George','Hale','1972-07-22','2013-06-20',70222,'ABC012');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (3,'Warren','Phelps','1963-09-13','2009-09-2',54909,'ABC013');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (4,'Abbott','Pearson','1977-06-01','1994-04-06',29363,'ABC014');
INSERT INTO empleados (id_departamento,nombre,apellido,fecha_nacimiento,fecha_incorporacion,sueldo_neto,dni) 
	VALUES (5,'Tanisha','Fischer','2000-12-12','2004-05-17',59872,'ABC015');