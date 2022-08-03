CREATE DATABASE emple_dep;

USE emple_dep;

CREATE TABLE departamento(
    id int not null auto_increment,
    nombre varchar(50),
    direccion varchar(100),
    PRIMARY KEY (id)
);

CREATE TABLE empleado(
	legajo int not null auto_increment,
    dni int,
    nombre varchar(50),
    apellido varchar(50),
    fecha_nacimiento datetime(0),
    fecha_incorporacion datetime(0),
    cargo varchar(50),
    sueldo_neto float8,
    id_departamento int,
    PRIMARY KEY (legajo),
    FOREIGN KEY (id_departamento) REFERENCES departamento (id)
);