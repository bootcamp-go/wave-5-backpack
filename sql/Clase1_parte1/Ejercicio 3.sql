CREATE DATABASE IF NOT EXISTS emple_dep;
USE emple_dep;

CREATE TABLE IF NOT EXISTS Departamento(
	Id int NOT NULL AUTO_INCREMENT,
    Nombre varchar(50) NOT NULL,
    Direccion varchar(50) NOT NULL,
    PRIMARY KEY (Id)
);

CREATE TABLE IF NOT EXISTS Empleado (
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

INSERT INTO Departamento(Nombre, Direccion)
VALUES 
('Departamento 1', 'Carrera 35 # 23'), 
('Departamento 2', 'Carrera 45 # 76'),
('Departamento 3', 'Carrera 67 # 21'), 
('Departamento 4', 'Carrera 87 # 43'),
('Departamento 5', 'Carrera 30 # 56');

INSERT INTO Empleado(Legajo, DNI, Apellido, Nombre, Fecha_nacimiento, Fecha_incorporacion, Cargo, Sueldo_neto, FK_id_departamento)
VALUES 
('06', '1082', 'Deighan', 'Belle', '2022-08-31', '2021-01-13', 'Environmental Specialist', 4183, 3),
('1742', '9832', 'Fedorchenko', 'Florence', '2022-09-07', '2022-02-22', 'Research Assistant III', 1721, 1),
('8676', '8573', 'Palser', 'Drucill', '2022-08-12', '2021-09-30', 'Professor', 4133, 2),
('895', '93332', 'Anersen', 'Ilysa', '2022-05-27', '2022-02-23', 'Speech Pathologist', 2079, 5),
('2', '396', 'Birds', 'Lilias', '2021-02-24', '2022-02-20', 'Sales Associate', 2899, 2),
('37688', '63', 'Guage', 'Ali', '2021-05-19', '2021-07-11', 'Health Coach II', 4001, 5),
('03', '691', 'Denyakin', 'Itch', '2021-04-29', '2021-04-13', 'Systems Administrator III', 4846, 2),
('9066', '0843', 'Harbron', 'Ray', '2022-12-22', '2021-04-12', 'Product Engineer', 1619, 4),
('88715', '931', 'Daniau', 'Lancelot', '2022-08-05', '2022-09-17', 'Nuclear Power Engineer', 2825, 2),
('26532', '52986', 'Meenan', 'Catherina', '2022-02-07', '2021-06-22', 'Clinical Specialist', 2794, 4),
('34', '03', 'Brigden', 'Petronille', '2021-02-05', '2022-02-25', 'Dental Hygienist', 3425, 2),
('68388', '4', 'Hargess', 'Lionello', '2022-01-13', '2021-08-19', 'Structural Analysis Engineer', 2031, 2),
('328', '05042', 'Mutimer', 'Anthia', '2021-08-03', '2021-06-04', 'Internal Auditor', 1492, 2),
('5', '789', 'Dinesen', 'Cyb', '2021-02-08', '2021-02-21', 'Food Chemist', 4243, 2),
('80', '00996', 'Davies', 'Inez', '2021-07-31', '2022-04-12', 'VP Marketing', 3534, 3);

SELECT * FROM Departamento;
SELECT * FROM Empleado;

SELECT e.Legajo, e.DNI, e.Apellido, e.Nombre, e.Fecha_nacimiento, e.Fecha_incorporacion, e.Cargo, e.Sueldo_neto, d.nombre as "Nombre dpto" , d.direccion
FROM empleado as e
INNER JOIN departamento as d
ON e.fk_id_departamento = d.id;