create database emple_dep

USE emple_dep;
CREATE TABLE empleado(
    no_legajo int not null primary key,
    dni varchar(255),
    apellido varchar(255),
    nombre varchar(255),
    fecha_ingreso date,
    fecha_corporacion date,
    sueldo double 
);

CREATE TABLE departamento(
    id int not null primary key,
    nombre varchar(255),
    direccion varchar(255)
);

ALTER TABLE empleado ADD COLUMN id_depto int;
ALTER TABLE empleado ADD FOREIGN KEY (id_depto) REFERENCES departamento(id);
INSERT INTO departamento 
VALUES (1, "Credits", "Mercado Libre México");

INSERT INTO departamento 
VALUES (2, "Fintech", "Mercado Libre México");

INSERT INTO departamento 
VALUES (3, "Marketplace", "Mercado Libre México");

INSERT INTO departamento 
VALUES (4, "UX/UI", "Mercado Libre México");

INSERT INTO departamento 
VALUES (5, "Shops", "Mercado Libre México");
INSERT INTO empleado
VALUES (1, "EMP010101", "Esquivel","Andrea", "2022-06-22", "2022-06-22", 20000, 1);

INSERT INTO empleado
VALUES (2, "EMP020202", "Rueda","Gabriela", "2022-06-22", "2022-06-22", 26000, 1);

INSERT INTO empleado
VALUES (3, "EMP030303", "Velez","Cristian", "2022-06-21", "2022-06-22", 25000, 2);

INSERT INTO empleado
VALUES (4, "EMP04044", "Santiago","Oliver", "2022-06-22", "2022-06-22", 30000, 4);

INSERT INTO empleado
VALUES (5, "EMP050505", "Lopez","Victor", "2022-06-23", "2022-06-22", 18000, 3);

INSERT INTO empleado
VALUES (6, "EMP060606", "Velez","Cristian", "2022-06-19", "2022-06-22", 25000, 2);

INSERT INTO empleado
VALUES (7, "EMP070707", "Mercado","Tania", "2022-06-20", "2022-06-22", 35000, 4);

INSERT INTO empleado
VALUES (8, "EMP080808", "Marquez","Carlos Obed", "2022-06-22", "2022-06-22", 25000, 4);

INSERT INTO empleado
VALUES (9, "EMP090909", "Rosendo","Miguel Angel", "2022-06-19", "2022-06-22", 15000, 4);

INSERT INTO empleado
VALUES (10, "EMP101010", "Brooke","Antonio", "2022-06-19", "2022-06-22", 45000, 5);

INSERT INTO empleado
VALUES (11, "EMP11111", "Hood","Miguel", "2022-06-18", "2022-06-22", 31000, 1);

INSERT INTO empleado
VALUES (12, "EMP121212", "Bronach","Lively", "2022-06-20", "2022-06-20", 33000, 3);

INSERT INTO empleado
VALUES (13, "EMP131313", "Evans","Apolo", "2022-06-20", "2022-06-20", 32000, 3);

INSERT INTO empleado
VALUES (14, "EMP141414", "Roberts","Luke", "2022-06-20", "2022-06-20", 39000,3);

INSERT INTO empleado
VALUES (15, "EMP151515", "Cassiev","Alexandra", "2022-06-22", "2022-06-20", 35000, 5);