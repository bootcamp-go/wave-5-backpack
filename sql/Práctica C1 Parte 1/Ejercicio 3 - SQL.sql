-- Active: 1659473777000@@127.0.0.1@3306@melisprint
-- Crear una base de datos con el nombre "emple_dep"

CREATE DATABASE emple_dep;
-- Seleccionar la base de datos con el nombre "emple_dep"
USE emple_dep;

-- Tabla de empleados
CREATE TABLE empleados(
    legajo int not null primary key auto_increment,
    dni int,
    apellido varchar(50),
    nombre varchar(50),
    fecha_nacimiento date,
    fecha_incorporacion date,
    cargo varchar(50),
    sueldo_neto float,
    departamento_id int
);

-- Tabla de departamentos
CREATE TABLE departamentos(
    id int not null primary key auto_increment,
    nombre varchar(50),
    direccion varchar(50)
);

-- Foreign key para conectar empleados con departamentos (departamento_id)
ALTER TABLE empleados ADD FOREIGN KEY (departamento_id) REFER

-- 5 registros en la tabla departamentos
INSERT INTO departamentos(nombre, direccion) VALUES
    ('Sistemas', 'Av. Siempre Viva 123'),
    ('Ventas', 'Av. Siempre Viva 456'),
    ('Compras', 'Av. Siempre Viva 789'),
    ('RRHH', 'Av. Siempre Viva 012'),
    ('Finanzas', 'Av. Siempre Viva 345');

-- 15 en la tabla empleados
INSERT INTO empleados(dni, apellido, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto, departamento_id) VALUES
    (01234567, 'Perez',    'Juan',  '1980-01-01', '2000-01-01', 'Jefe',          5000000,  1),
    (12345678, 'Gonzalez', 'Juan',  '1981-02-02', '2001-02-02', 'Desarrollador', 30000000, 1),
    (23456789, 'Correa',   'Ana',   '1982-03-03', '2002-03-03', 'Vendedora',     20000000, 1),
    (34567890, 'Gomez',    'Juan',  '1983-04-04', '2003-04-04', 'Vendedora',     20000000, 2),
    (45678901, 'Perez',    'Ana',   '1984-05-05', '2004-05-05', 'Vendedora',     20000000, 2),
    (56789012, 'Gonzalez', 'Juan',  '1985-06-06', '2005-06-06', 'Desarrollador', 30000000, 2),
    (67890123, 'Correa',   'Ana',   '1986-07-07', '2006-07-07', 'Desarrollador', 30000000, 3),
    (78901234, 'Gomez',    'Juan',  '1987-08-08', '2007-08-08', 'Jefe',          5000000,  3),
    (89012345, 'Perez',    'Ana',   '1988-09-09', '2008-09-09', 'Jefe',          5000000,  3),
    (90123456, 'Ortiz',    'Pablo', '1989-10-10', '2009-10-10', 'Desarrollador', 30000000, 4),
    (01234567, 'Gonzalez', 'Juan',  '1990-11-11', '2010-11-11', 'Desarrollador', 30000000, 4),
    (12345678, 'Correa',   'Ana',   '1991-12-12', '2011-12-12', 'Jefe',          5000000,  4),
    (23456789, 'Gomez',    'Juan',  '1992-01-01', '2012-01-01', 'Jefe',          5000000,  5),
    (34567890, 'Perez',    'Ana',   '1993-02-02', '2013-02-02', 'Desarrollador', 30000000, 5),
    (45678901, 'Gonzalez', 'Juan',  '1994-03-03', '2014-03-03', 'Desarrollador', 30000000, 5);