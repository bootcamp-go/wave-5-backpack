CREATE DATABASE emple_dep;

CREATE TABLE emple_dep.Departments (
    ID INT PRIMARY KEY NOT NULL,
    Nombre VARCHAR (50) NOT NULL,
    Direccion VARCHAR (100) NOT NULL
    
);
CREATE TABLE emple_dep.Employees (
    Nro_Legajo INT PRIMARY KEY,
    DNI VARCHAR (50) NOT NULL,
    Apellido VARCHAR (50) NOT NULL,
    Nombre VARCHAR (50) NOT NULL,
    Fecha_Nacimiento DATETIME,
    Fecha_Incorporacion DATETIME,
    Cargo VARCHAR (50) NOT NULL,
    Sueldo_Neto INT NOT NULL ,
    ID_Departamento INT NOT NULL,
    FOREIGN KEY (ID_Departamento) REFERENCES emple_dep.Departments (ID)
);

SELECT * FROM emple_dep.Departments;
INSERT INTO `emple_dep`.`Departments` (`ID`, `Nombre`, `Direccion`) VALUES ('1', 'RRHH', 'direccion1');
INSERT INTO `emple_dep`.`Departments` (`ID`, `Nombre`, `Direccion`) VALUES ('2', 'IT', 'direccion2');
INSERT INTO `emple_dep`.`Departments` (`ID`, `Nombre`, `Direccion`) VALUES ('3', 'Marketing', 'direccion3');
INSERT INTO `emple_dep`.`Departments` (`ID`, `Nombre`, `Direccion`) VALUES ('4', 'Logistica', 'direccion4');
INSERT INTO `emple_dep`.`Departments` (`ID`, `Nombre`, `Direccion`) VALUES ('5', 'Almacen', 'direccion5');


SELECT * FROM emple_dep.Employees;
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('1', '123456', 'Maradona', 'Diego', '19601030', '20220802', 'Software Developer', '10000', '1');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('2', '354354', 'Perez', 'Pedro', '19601030', '20220802', 'Cargo', '5000', '1');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('3', '5345345', 'Castro', 'Julian', '19601030', '20220802', 'Cargo', '5555', '1');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('4', '5365767', 'Rodriguez', 'Fernanda', '19601030', '20220802', 'Cargo', '9999', '1');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('5', '876867', 'Mendez', 'Tatiana', '19601030', '20220802', 'Cargo', '60000', '2');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('6', '9778', 'Izaguirre', 'Martin', '19601030', '20220802', 'Cargo', '40000', '2');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('7', '53423432', 'Fernandez', 'Jose', '19601030', '20220802', 'Cargo', '700', '3');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('8', '675689', 'Vaz', 'Valentina', '19601030', '20220802', 'Cargo', '300000', '3');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('9', '7567587', 'Fabra', 'Claudio', '19601030', '20220802', 'Cargo', '25000', '3');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('10', '4562347', 'Valdez', 'Alvaro', '19601030', '20220802', 'Cargo', '320000', '3');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('11', '9897845', 'Carreño', 'Nicolas', '19601030', '20220802', 'Cargo', '45800', '4');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('12', '2343248', 'Carrasco', 'Camila', '19601030', '20220802', 'Cargo', '9870', '4');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('13', '1854345', 'Vivas', 'Tatiana', '19601030', '20220802', 'Cargo', '70456', '4');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Nacimiento`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('14', '823157567', 'Gallardo', 'Roberto', '19601030', '20220802', 'Cargo', '5343', '4');
INSERT INTO `emple_dep`.`Employees` (`Nro_Legajo`, `DNI`, `Apellido`, `Nombre`, `Fecha_Incorporacion`, `Cargo`, `Sueldo_Neto`, `ID_Departamento`) VALUES ('15', '2465758', 'Alonso', 'Juana', '20220802', 'Cargo', '96432', '5');


