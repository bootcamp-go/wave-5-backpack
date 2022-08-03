DROP DATABASE IF EXISTS EMPLE_DEP;
CREATE DATABASE EMPLE_DEP;

USE EMPLE_DEP;

CREATE TABLE EMPLEADO(
    nro_legajo int primary key auto_increment,
    dni varchar(50) unique,
    nombre varchar(50),
    apellido varchar(50),
    cargo varchar(50),
    sueldo_neto float,
    fecha_nacimiento datetime,
    fecha_incorporacion datetime,
    departamento_id int 
);

CREATE TABLE DEPARTAMENTO(
    id int primary key auto_increment,
    nombre varchar(50),
    direccion varchar(100)
);

ALTER TABLE EMPLEADO
ADD CONSTRAINT fk_empleado_departamento
FOREIGN KEY(departamento_id) REFERENCES DEPARTAMENTO(id);

INSERT INTO DEPARTAMENTO(nombre, direccion)
VALUES ('TI', 'Av Apoquindo 5230, Santiago, Chile'),
('VENTAS', 'Av Apoquindo 5000, Santiago, Chile'),
('MARKETING', 'Av Apoquindo 4600, Santiago, Chile'),
('LOGISTICA', 'Av Apoquindo 4000, Santiago, Chile'),
('CALL CENTER', 'Av Apoquindo 3200, Santiago, Chile');


INSERT INTO EMPLEADO (dni, nombre, apellido, cargo, sueldo_neto, fecha_nacimiento, fecha_incorporacion, departamento_id)
VALUES ('195711-7', 'claudio', 'figueroa', 'software developer', 1900000, '1997-02-22', '2022-06-22', 1),
 ('1957315-7', 'andres', 'sepulveda', 'vendedor',800000, '1997-02-22', '2022-07-10', 2),
 ('19572316-7', 'claudia', 'miraflores', 'publicador',500000, '1997-02-22', '2022-06-25', 3),
 ('152312-7', 'flor', 'aleman', 'empaquetador',1000000, '1997-02-22', '2022-05-21', 4),
 ('19572311-7', 'jaime', 'ortiz', 'operadora',700000, '1997-02-22', '2022-03-13', 5),
 ('19572318-7', 'florencia', 'lagos', 'software developer',950000, '1997-02-22', '2021-06-15', 1),
 ('1972319-7', 'mary', 'hurtado', 'software developer',1250000, '1997-02-22', '2022-02-03', 1),
 ('19572314-7', 'nicolas', 'tesone', 'software developer',1350000, '1997-02-22', '2022-08-25', 1),
 ('1972317-7', 'horacion', 'lopez', 'publicador',1100000, '1997-02-22', '2021-09-15', 3),
 ('1957231-7', 'juan', 'ortiz', 'empaquetador',1500000, '1997-02-22', '2022-06-25', 4),
 ('1952312-7', 'jimena', 'rodriguez', 'operadora',1400000, '1997-02-22', '2022-06-22', 5),
 ('19572331-7', 'daniela', 'sotomayor', 'vendedor',800000, '1997-02-22', '2022-06-22', 2),
 ('1957261-7', 'javier', 'del prado', 'publicador',1900000, '1997-02-22', '2022-06-22', 3),
 ('19571-7', 'abelardo', 'guimenez', 'empaquetador',1950000, '1997-02-22', '2022-06-22', 4),
 ('195721-7', 'rodrigo', 'cifuentes', 'software developer',1500000, '1997-02-22', '2022-06-22', 1);