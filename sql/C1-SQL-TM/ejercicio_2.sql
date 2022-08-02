-- Una vez modelada y planteada la base de datos, responder a las siguientes preguntas:
-- 
-- a. ¿Cuál es la primary key para la tabla empleados?
--       el no de legajo
--
-- b. ¿Cuál es la primary key para la tabla departamentos?
--       el ID departamento
--
-- c. ¿Qué relación/es existirían entre las tablas?
--    ¿En qué tabla debería haber foreign key?
--    ¿A qué campo de qué tabla hace referencia dicha foreign key?
--       La relación es de un empleado pertenece a un departamento,
--       En la tabla empleado debe tener una foreign key,
--       dicha foreign key hace referencia a la llave primaria de departamento

CREATE DATABASE emple_dep;

USE emple_dep

CREATE TABLE departamento (
  id_departamento INT NOT NULL AUTO_INCREMENT,
  nombre varchar(50) NOT NULL,
  direccion varchar(50) NULL,
  PRIMARY KEY (id_departamento)
);

create table empleado (
  no_legajo INT NOT NULL AUTO_INCREMENT,
  id_departamento INT NOT NULL,
  dni VARCHAR(20) NOT NULL,
  nombre VARCHAR(50) NOT NULL,
  apellido VARCHAR(50) NULL,
  cargo VARCHAR(50) NOT NULL,
  sueldo_neto DECIMAL(15,2) NOT NULL,
  fecha_nacimiento DATETIME NULL,
  fecha_incorporacion DATETIME NOT NULL,
  PRIMARY KEY (no_legajo),
  FOREIGN KEY (id_departamento) REFERENCES departamento(id_departamento)
);
