CREATE DATABASE emple_dep;
USE emple_dep;

CREATE TABLE departamentos (
	ID int NOT NULL auto_increment,
	nombre varchar(50),
    direccion varchar(50),
    PRIMARY KEY(ID)
);

CREATE TABLE empleados (
	legajo int NOT NULL auto_increment,
	dni int NOT NULL,
    apellido varchar(50),
    nombre varchar(50),
    fecha_de_nacimiento date,
    fecha_de_incorporacion date,
    cargo varchar(50),
    sueldo_neto float,
    PRIMARY KEY(legajo),
	departamentoID int,
    FOREIGN KEY(departamentoID) REFERENCES departamentos(ID)
);

INSERT INTO 
departamentos (ID, nombre, direccion)
VALUES 
(1, 'finanzas', 'Cordoba'),
(2, 'recursos humanos', 'Mendoza'),
(3, 'IT', 'Buenos Aires'),
(4, 'IT', 'Rafaela'),
(5, 'Shipping', 'Tucuman');

INSERT INTO
empleados(dni, departamentoID, apellido, nombre, fecha_de_nacimiento, fecha_de_incorporacion, cargo, sueldo_neto)
VALUES
(234, 4,'Darin','Ricardo','1995-04-20','2016-04-20', 'Developer' ,2000000),
(345, 5,'Lopez','Stiven','1998-11-10','2016-05-18', 'Manager' ,38000000),
(489, 3,'Zapata','Carlos','1997-12-23','2018-04-13', 'Software Developer' ,2300500),
(123, 1, 'salas', 'Carlos', '2022-06-21', '2022-06-21', 'asistente', 3456789.00),
(345, 2, 'solis', 'Maria', '2022-06-21', '2022-06-21', 'secretaria', 3454489.00),
(234, 3, 'arenas', 'Maria', '2022-06-21', '2022-06-21', 'mensajera', 9956789.00),
(456, 4,'bastidas', 'Fernando', '2022-06-21', '2022-06-21', 'cajera', 2156789.00),
(789, 5,'zapata', 'Lina', '2022-06-21', '2022-06-21', 'subgerente', 6756789.00);
