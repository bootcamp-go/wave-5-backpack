CREATE DATABASE emple_dep;

USE emple_dep;

CREATE TABLE departamentos (
	id_departamento INT,
    nombre_departamento VARCHAR(50),
    direccion_departamento VARCHAR(50),
    PRIMARY KEY (id_departamento)
);

CREATE TABLE empleados (
	legajo INT,
    dni INT,
    apellido VARCHAR(50),
    nombre VARCHAR(50),
    fecha_de_nacimiento DATE,
    fecha_de_incorporacion DATE,
    cargo VARCHAR(50),
    sueldo_neto FLOAT,
    id_departamento INT,
    PRIMARY KEY (legajo),
    FOREIGN KEY (id_departamento) REFERENCES departamentos(id_departamento)
);
