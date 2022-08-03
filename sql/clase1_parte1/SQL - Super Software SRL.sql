CREATE DATABASE EMPLOYEE_X_DEPARTAMENT;

USE EMPLOYEE_X_DEPARTAMENT;

CREATE TABLE DEPARTAMENTS(
	ID int NOT NULL PRIMARY KEY auto_increment,
    name varchar(60),
    address varchar(100)
);

CREATE TABLE EMPLOYEES(
	file_number int NOT NULL PRIMARY KEY auto_increment,
    DNI int,
    last_name varchar(60),
    name varchar(60),
    date_birthdate date,
    date_incorporation date,
    position varchar(100),
    net_salary decimal(10, 2),
    departament_id int NOT NULL
);

ALTER TABLE EMPLOYEES ADD FOREIGN KEY (departament_id) REFERENCES DEPARTAMENTS(id);

INSERT INTO 
	DEPARTAMENTS (id, name, address) 
VALUES 
	(1, "Marketing", ""),
	(2, "Develpment", ""),
	(3, "Security", ""),
	(4, "People", ""),
    (5, "Logistic", "")
;

INSERT INTO 
	EMPLOYEES (DNI, last_name, name, date_birthdate, date_incorporation, position, net_salary, departament_id) 
VALUES 
	(123, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 1),
    (124, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 2),
    (125, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 3),
    (126, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 3),
    (127, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 5),
    (128, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 4),
    (129, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 4),
    (130, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 4),
    (131, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 1),
    (132, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 2),
    (133, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 3),
    (134, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 4),
    (135, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 1),
    (136, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 5),
    (137, "torres", "michael", "1998-08-03", "2022-06-21", "Software developer", 123, 2);

    
SELECT * FROM DEPARTAMENTS;
SELECT * FROM EMPLOYEES;

DROP DATABASE EMPLOYEE_X_DEPARTAMENT;

DROP table EMPLOYEES;