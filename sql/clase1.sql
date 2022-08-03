##Mostrando DB, Creando DB, Usando DB.

show databases;
create database emple_dep;
use emple_dep;

##Creando la tabla empleado.

CREATE TABLE empleado (n_legajo int  NOT NULL, dni int NOT NULL, nombre varchar(25) NOT NULL, fecha_nacimiento date NOT NULL, fecha_incorporacion date NOT NULL, cargo varchar(60) NOT NULL, sueldo_neto float NOT NULL, id_departamento_fk int NOT NULL, PRIMARY KEY (n_legajo));
SELECT * FROM empleado;

DESCRIBE empleado;

#Anadiento un auto_increment a la columna n_legajo.

ALTER TABLE empleado MODIFY COLUMN n_legajo AUTO_INCREMENT;

##Creando tabla departamento.

CREATE TABLE departamento (id_departamento int NOT NULL AUTO_INCREMENT, nombre_departamento varchar(60) NOT NULL, direccion_departamento varchar(60) NOT NULL, PRIMARY KEY (id_departamento));

DESCRIBE departamento;

##Creando la FK en la tabla empleado, en el campo  id_dep.. con la llave primaria DEPARTAMENTO.

ALTER TABLE empleado ADD FOREIGN KEY (id_departamento_fk) REFERENCES departamento(id_departamento);

##INSERTANDO DATOS A LA TABLA departamento.

INSERT INTO departamento (nombre_departamento, direccion_departamento) VALUES ('People', 'Buenos Aires DOT');

INSERT INTO departamento (nombre_departamento, direccion_departamento) VALUES ('IT', 'Buenos Aires DOT');

INSERT INTO departamento (nombre_departamento, direccion_departamento) VALUES ('IT Aquisition', 'Buenos Aires DOT');

INSERT INTO departamento (nombre_departamento, direccion_departamento) VALUES ('Business', 'Buenos Aires DOT');

INSERT INTO departamento (nombre_departamento, direccion_departamento) VALUES ('CEO', 'Buenos Aires DOT');

##Insertando datos en la tabla empleados

INSERT INTO empleado (dni, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto, id_departamento_fk) VALUES (1005815706, 'Christian Daniel Ospina', '2001-04-16', '2021-06-21', 'Software developer',4500000, 2 );

INSERT INTO empleado (dni, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto, id_departamento_fk) VALUES (1005815701, 'Sara Jimena Ochoa', '2001-05-16', '2021-06-21', 'Software developer',4500000, 2 );

INSERT INTO empleado (dni, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto, id_departamento_fk) VALUES (1005815702, 'Jose Enrique Magallan', '2000-04-16', '2021-06-21', 'Software developer',4500000, 2 );

INSERT INTO empleado (dni, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto, id_departamento_fk) VALUES (1005815703, 'Mariana Juliana Velez', '1995-04-16', '2021-06-21', 'Software developer',4500000, 2);

INSERT INTO empleado (dni, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto, id_departamento_fk) VALUES (1005815704, 'Eilin  Juliana Gomez', '1999-04-16', '2021-06-21', 'Peopple Customer',4500000, 1);

INSERT INTO empleado (dni, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto, id_departamento_fk) VALUES (1005815705, 'Nahuel Daniel Gonzales', '1995-04-16', '2021-06-21', 'Software developer',4500000, 1);

INSERT INTO empleado (dni, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto, id_departamento_fk) VALUES (9994343, 'Agustin Emiliana Sala', '1996-04-16', '2015-06-21', 'IT Aquisition',7500000, 3);

INSERT INTO empleado (dni, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto, id_departamento_fk) VALUES (100345905, 'Andres Ramirez Guzman', '1994-04-16', '2012-06-21', 'Business Customer',10500000, 4);
