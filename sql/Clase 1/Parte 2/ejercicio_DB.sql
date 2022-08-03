CREATE DATABASE emple_dep;
USE emple_dep;

CREATE TABLE Departamento (
	id int primary key,
    nombre varchar(50),
    direccion varchar(50)
);


DROP TABLE Empleado;
CREATE TABLE Empleado (
    n_legajo INT PRIMARY KEY,
    dni BIGINT,
    nombre VARCHAR(20),
    apellido VARCHAR(20),
    fecha_nacimiento DATE,
    sueldo FLOAT,
    cargo VARCHAR(50),
    id_departamento INT
);

ALTER TABLE emple_dep.Empleado
ADD FOREIGN KEY (id_departamento) REFERENCES Departamento(id);

select * from Departamento;

INSERT INTO Departamento values (1, 'IT', 'Carrera 19 102-12');
INSERT INTO Departamento values (2, 'Marketing', 'Carrera 90 12-34');
INSERT INTO Departamento values (3, 'Finanzas', 'Calle 102 20-53');
INSERT INTO Departamento values (4, 'Crypto', 'Carrera 19 102-12');
INSERT INTO Departamento values (5, 'Seguridad', 'Carrera 19 102-12');

select * from Empleado;

INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1000, 10087690652, 'Ramos', 'Andres', STR_TO_DATE('2022-06-21', '%Y-%m-%d'), 1000, 'dev', 1);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1001, 90780567858, 'Tesone', 'Nicolas', STR_TO_DATE('2020-05-20', '%Y-%m-%d'), 2000, 'dev', 1);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1002, 38754690987, 'Perez', 'Juan', STR_TO_DATE('2019-04-12', '%Y-%m-%d'), 1500, 'market analyst', 2);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1003, 09856543433, 'Espinoza', 'Andrea', STR_TO_DATE('2022-01-21', '%Y-%m-%d'), 1200, 'social media analyst', 2);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1004, 09455246678, 'Zambrano', 'Viviana', STR_TO_DATE('2018-07-01', '%Y-%m-%d'), 8000, 'chief accountant', 3);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1005, 45785785007, 'Perdomo', 'Camila', STR_TO_DATE('2010-04-12', '%Y-%m-%d'), 900, 'junior accountant', 3);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1006, 97866457986, 'Avila', 'Camilo', STR_TO_DATE('2008-02-28', '%Y-%m-%d'), 1700, 'crypto accountant', 4);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1007, 65435678976, 'Viatela', 'Juan', STR_TO_DATE('2010-06-14', '%Y-%m-%d'), 1300, 'solidity dev', 4);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1008, 87656789765, 'Gonzalez', 'Raul', STR_TO_DATE('2008-07-20', '%Y-%m-%d'), 1100, 'cybersecurity analyst', 5);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1009, 04965899583, 'Castro', 'Pedro', STR_TO_DATE('2020-04-12', '%Y-%m-%d'), 800, 'junior cybersecurity', 5);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1010, 98765434587, 'Perdomo', 'Camila', STR_TO_DATE('2010-04-12', '%Y-%m-%d'), 900, 'junior accountant', 3);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1011, 97866457986, 'Avila', 'Camilo', STR_TO_DATE('2008-02-28', '%Y-%m-%d'), 1700, 'crypto accountant', 4);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1012, 65435678976, 'Viatela', 'Juan', STR_TO_DATE('2010-06-14', '%Y-%m-%d'), 1300, 'solidity dev', 4);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1013, 87656789765, 'Gonzalez', 'Raul', STR_TO_DATE('2008-07-20', '%Y-%m-%d'), 1100, 'cybersecurity analyst', 5);
INSERT INTO Empleado (n_legajo, dni, apellido, nombre, fecha_nacimiento, sueldo, cargo, id_departamento) values (1014, 04965899583, 'Castro', 'Pedro', STR_TO_DATE('2020-04-12', '%Y-%m-%d'), 800, 'junior cybersecurity', 5);