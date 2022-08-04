INSERT INTO
departamentos (id_departamento, nombre_departamento, direccion_departamento)
VALUES
(1, 'Finanzas', 'Bogotá'),
(2, 'Recursos Humanos', 'Cali'),
(3, 'IT', 'Buenos Aires'),
(4, 'IT', 'Medellín'),
(5, 'Shipping', 'San Andrés');

INSERT INTO
empleados (legajo, dni, id_departamento, apellido, nombre, fecha_de_nacimiento, fecha_de_incorporacion, cargo, sueldo_neto)
VALUES
(1, 000, 4,'Lucumi','Luberley','1995-04-20','2016-04-20', 'Developer' ,2000000),
(2, 001, 5,'Pinzon','Stiven','1998-11-10','2016-05-18', 'Manager' ,38000000),
(3, 101, 3,'Zapata','Carlos','1997-12-23','2018-04-13', 'Software Developer' ,2300500),
(4, 123, 1, 'Salas', 'Camilo', '2022-06-21', '2022-06-21', 'Asistente', 3456789.00),
(5, 345, 2, 'Solis', 'Maria', '2022-06-21', '2022-06-21', 'Secretaria', 3454489.00),
(6, 234, 3, 'Arenas', 'Luz', '2022-06-21', '2022-06-21', 'Mensajera', 9956789.00),
(7, 456, 4,'Bastidas', 'Natalia', '2022-06-21', '2022-06-21', 'Cajera', 2156789.00),
(8, 789, 5,'Ante', 'Linda', '2022-06-21', '2022-06-21', 'Subgerente', 6756789.00),
(9, 100, 5,'Pinzon','Sandra','1998-11-10','2016-05-18', 'Manager' ,38000000),
(10, 489, 3,'Hurtado','Carla','1997-12-23','2018-04-13', 'Software Developer' ,2300500),
(11, 098, 1, 'Alvarez', 'Carlota', '2022-06-21', '2022-06-21', 'Asistente', 3456789.00),
(12, 667, 2, 'Toro', 'Luisa', '2022-06-21', '2022-06-21', 'Secretaria', 3454489.00),
(13, 876, 3, 'Henao', 'Luz', '2022-06-21', '2022-06-21', 'Mensajera', 9956789.00),
(14, 406, 4,'Bueno', 'Nairobi', '2022-06-21', '2022-06-21', 'Cajera', 2156789.00),
(15, 709, 5,'Moncada', 'Jaime', '2022-06-21', '2022-06-21', 'Subgerente', 6756789.00);