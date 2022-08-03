create database emple_dep;
use emple_dep;
create table departamentos(
id int primary key auto_increment,
nombre varchar(50),
direccion varchar(80)
);
create table empleados(
dni int primary key,
nombre varchar(50),
apellido varchar(50),
fecha_nacimiento date,
nro_legajo int,
cargo varchar(50),
fecha_incorporacion date,
sueldo_neto int,
id_departamento int, 
foreign key (id_departamento) references departamentos(id) 
);

insert into departamentos(nombre, direccion)
values
('TI','Orgrimmar'),
('Contabilidad','Goblinar'),
('RRHH','StormWind'),
('Logistica','Gnomeraggan'),
('Gerencia','Dalaran');

insert into empleados(dni, nombre, apellido, fecha_nacimiento, nro_legajo, cargo, fecha_incorporacion, sueldo_neto, id_departamento)
values
('123456789','Matias','Carrasco','01-01-1984','666','Desarrollador','01-01-2022','5555555','1'),
('987654321','Matias','Carrasco','01-01-1984','666','Teleco','01-01-2022','5555555','2'),
('678123785','Matias','Carrasco','01-01-1984','666','Contador','01-01-2022','5555555','3'),
('784356712','Matias','Carrasco','01-01-1984','666','Genrente','01-01-2022','5555555','4'),
('761234675','Matias','Carrasco','01-01-1984','666','Diseñador','01-01-2022','5555555','5'),
('781235681','Matias','Carrasco','01-01-1984','666','ingeniero','01-01-2022','5555555','5'),
('534532231','Matias','Carrasco','01-01-1984','666','Abogado','01-01-2022','5555555','4'),
('121212212','Matias','Carrasco','01-01-1984','666','Veterinario','01-01-2022','5555555','3'),
('323232323','Matias','Carrasco','01-01-1984','666','Laboratorista','01-01-2022','5555555','2'),
('434343434','Matias','Carrasco','01-01-1984','666','Actuario','01-01-2022','5555555','1'),
('545454325','Matias','Carrasco','01-01-1984','666','Aseo','01-01-2022','5555555','2'),
('345345255','Matias','Carrasco','01-01-1984','666','Doctor','01-01-2022','5555555','3'),
('345634534','Matias','Carrasco','01-01-1984','666','Enefermero','01-01-2022','5555555','3'),
('634634324','Matias','Carrasco','01-01-1984','666','Desarrollador','01-01-2022','5555555','4'),
('657885653','Matias','Carrasco','01-01-1984','666','Asistente Social','01-01-2022','5555555','5')
;
show tables;
select * from departamentos;
select * from empleados;




