USE emple_dep

INSERT INTO departamento (nombre,direccion) VALUES ('Software','Piso 1');
INSERT INTO departamento (nombre,direccion) VALUES ('QA','Piso 2');
INSERT INTO departamento (nombre,direccion) VALUES ('Documentación','Piso 3');
INSERT INTO departamento (nombre,direccion) VALUES ('Análisis','Piso 4');
INSERT INTO departamento (nombre,direccion) VALUES ('Deployment','Piso 5');

INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (1,'DNI0001','Juan','SW',1300,'2000-02-14');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (2,'DNI0002','Oscar','QA',1600,'1998-08-20');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (3,'DNI0003','Karina','DOC',1000,'2002-12-07');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (4,'DNI0004','Omar','AN',1700,'2005-04-19');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (5,'DNI0005','Sandra','DEP',1200,'1997-07-12');

INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (1,'DNI0006','Karen','SW',1300,'2001-02-14');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (2,'DNI0007','Jose','QA',1600,'1999-08-20');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (3,'DNI0008','Pedro','DOC',1000,'1995-12-07');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (4,'DNI0009','Erik','AN',1700,'2006-04-19');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (5,'DNI0010','Melisa','DEP',1200,'2005-07-12');

INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (1,'DNI0011','Irving','SW',1300,'2003-02-14');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (2,'DNI0012','Melina','QA',1600,'1998-08-20');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (3,'DNI0013','Edgar','DOC',1000,'1999-12-07');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (4,'DNI0014','Ingrid','AN',1700,'2006-04-19');
INSERT INTO empleado (id_departamento,dni,nombre,cargo,sueldo_neto,fecha_incorporacion) VALUES (5,'DNI0015','Angel','DEP',1200,'2009-07-12');
