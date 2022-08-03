CREATE DATABASE IF NOT EXISTS emple_dep;
USE emple_dep;

CREATE TABLE  IF NOT EXISTS Departamento(
    Id int NOT NULL AUTO_INCREMENT,
    Nombre varchar(50) NOT NULL,
    Direccion varchar(50) NOT NULL,
    PRIMARY KEY (Id)
);

CREATE TABLE  IF NOT EXISTS Empleado (
    Legajo int NOT NULL,
    DNI int NOT NULL,
    Apellido VARCHAR(50) NOT NULL,
    Nombre VARCHAR(50) NOT NULL,
    Fecha_nacimiento DATE NOT NULL,
    Fecha_incorporacion DATE NOT NULL,
    Cargo VARCHAR(50) NOT NULL,
    Sueldo_neto INT NOT NULL,
    Fk_id_departamento INT NOT NULL,
    PRIMARY KEY (Legajo),
    FOREIGN KEY (Fk_id_departamento) REFERENCES Departamento(Id)
);

INSERT INTO departamento (nombre, direccion)
VALUES ('Departamento 1', 'Carrera 1 calle 1'),
 ('Departamento 2', 'Carrera 2 calle 2'),
 ('Departamento 3', 'Carrera 3 calle 3'),
 ('Departamento 4', 'Carrera 4 calle 4'),
 ('Departamento 5', 'Carrera 5 calle 5');
 
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('321', '111', 'Issacov', 'Deana', '2022-03-28', '2021-06-06', 'Account Representative I', 1361, 5);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('36', '9487', 'Pretsel', 'Bartholemy', '2021-04-14', '2022-10-08', 'Staff Scientist', 2922, 5);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('74', '6021', 'Hurt', 'Johnnie', '2022-01-19', '2022-02-08', 'Staff Scientist', 4034, 2);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('0168', '36', 'Winspurr', 'Sanson', '2022-11-21', '2022-09-12', 'Sales Representative', 1775, 3);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('436', '2045', 'Jandel', 'Blakeley', '2022-01-17', '2021-04-30', 'Recruiting Manager', 4057, 1);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('690', '10', 'Critoph', 'Marta', '2022-10-09', '2022-05-07', 'Information Systems Manager', 1756, 3);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('84146', '49028', 'Czapla', 'Monty', '2022-09-10', '2021-06-12', 'Research Associate', 1949, 1);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('882', '8', 'Tellett', 'Gunner', '2022-12-04', '2022-05-09', 'Chemical Engineer', 3983, 4);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('31', '36939', 'Bredbury', 'Keely', '2022-06-17', '2022-08-09', 'Environmental Specialist', 3245, 4);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('01', '422', 'Pedrocco', 'Abbey', '2022-07-11', '2022-03-26', 'Health Coach II', 2737, 5);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('03724', '49', 'Sonner', 'Alameda', '2022-01-05', '2022-10-17', 'Quality Control Specialist', 4200, 1);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('492', '6317', 'Corton', 'Roddie', '2022-07-03', '2022-05-05', 'Budget/Accounting Analyst IV', 3467, 3);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('6811', '72', 'Verrillo', 'Lusa', '2022-08-19', '2021-11-09', 'Analyst Programmer', 3911, 1);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('0073', '3', 'Losbie', 'Mallory', '2022-11-18', '2022-04-16', 'Administrative Officer', 4588, 3);
insert into empleado (Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, Fk_id_departamento) values ('8', '16205', 'Penburton', 'Glenine', '2022-06-16', '2021-07-25', 'Associate Professor', 2904, 1);

SELECT * FROM departamento;
SELECT * FROM empleado;

SELECT e.Legajo, e.DNI, e.Apellido, e.Nombre, e.Fecha_nacimiento, e.Fecha_incorporacion, e.Cargo, e.Sueldo_neto, d.nombre as "Nombre dpto" , d.direccion
FROM empleado as e
INNER JOIN departamento as d
ON e.fk_id_departamento = d.id;