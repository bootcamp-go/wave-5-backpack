-- Creación de base de datos
CREATE DATABASE emple_dep;

-- Colocando la base de datos a disposición para su uso
USE emple_dep;

-- Tabla de departamentos
CREATE TABLE departamentos (
    id int NOT NULL auto_increment,
    nombre varchar(50),
    direccion varchar(50),
    PRIMARY KEY(id)
);

-- Tabla de empleados
CREATE TABLE empleados (
    legajo int NOT NULL auto_increment,
    dni int NOT NULL,
    depto_id int,
    apellido varchar(50),
    nombre varchar(50),
    fecha_nacimiento date,
    fecha_incorporacion date,
    cargo varchar(50),
    sueldo_neto float,
    PRIMARY KEY(legajo),
    FOREIGN KEY(depto_id) REFERENCES departamentos(id)
);

-- Insertando información a tabla departamentos
INSERT INTO departamentos(nombre, direccion) VALUES
    ('Finanzas', 'Bogotá'),
    ('Recursos Humanos', 'Cali'),
    ('IT', 'Buenos Aires'),
    ('IT', 'Medellín'),
    ('Shipping', 'Bogotá');

-- Insertando información a tabla empleados
INSERT INTO empleados(dni, depto_id, apellido, nombre, fecha_nacimiento, fecha_incorporacion, cargo, sueldo_neto) VALUES 
    (1,3,'Lucumi','Luz','1960-06-22','2010-05-23', 'Developer' ,2200000),
    (2,2,'Perez','Maria','1961-06-23','2012-06-30', 'Manager' ,3500000),
    (3,3,'Martinez','Pablo','1962-07-28','2016-09-26', 'Senior Developer' ,5500000),
    (4,4,'Hernandez','Martha','1963-09-30','2014-05-28', 'Endpoint Designer' ,23000000),
    (5,4,'Lucumi','Luberley','1995-04-20','2016-04-20', 'Developer' ,2000000),
    (6,5,'Pinzon','Stiven','1998-11-10','2016-05-18', 'Manager' ,38000000),
    (7,3,'Lopez','Candelaria','1995-06-22','2010-05-23', 'Software Developer' ,2200000),
    (8,2,'Riascos','Mariana','1998-06-23','2012-06-30', 'IT Manager' ,3500000),
    (9,3,'Perez','Carla','1999-07-28','2016-09-26', 'Senior Developer' ,5500000),
    (10,5,'Cascada','Sofia','1999-09-30','2014-05-28', 'Shipment Analyst' ,23000000),
    (11,3,'Martinez','Pablo','1995-04-20','2016-04-20', 'Developer' ,2000000),
    (12,5,'Pinzon','Santiago','1998-11-10','2016-05-18', 'Manager' ,38000000),
    (13,5,'Triana','Joshua','1990-09-30','2014-05-28', 'Shipment Analyst' ,23000000),
    (14,3,'Delgago','Carlos','1992-04-20','2020-04-20', 'Developer' ,2000000),
    (15,3,'Velandia','Diana','1991-11-10','2022-05-18', 'Software Developer' ,38000000);
