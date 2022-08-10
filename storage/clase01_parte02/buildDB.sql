-- Creo la base de datos, si no existe, y la seleciono.
CREATE DATABASE IF NOT EXISTS storage;
USE storage;
-- Si existe una tabla "warehouse" la elimino y creo una nueva.
DROP TABLE IF EXISTS `warehouses`;
CREATE TABLE `warehouses` (
  `id` INT NOT NULL PRIMARY KEY,
  `name` VARCHAR(40) NOT NULL,
  `address` VARCHAR(150) NOT NULL
);
-- Creo un warehouse principal.
INSERT INTO `warehouses` (`id`, `name`, `address`) VALUES (1, 'Main Warehouse', '221 Baker Street');
-- Si existe una tabla "products" la elimino y creo una nueva.
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(60) NOT NULL,
  `lastname` VARCHAR(60) NOT NULL,
  `email` VARCHAR(60) NOT NULL,
  `age` INT NOT NULL,
  `height` FLOAT NOT NULL,
  `active` BOOL NOT NULL,
  `createdat` DATE NOT NULL
);
-- Creo unos productos de prueba
INSERT INTO `users` VALUES 
    (1,'Claudio','Bieler',"taca@gmail.com",39,1.78, true, "2005-07-04 00:00:00"),
    (2,'Lionel','Messi',"lio@gmail.com",35,1.70, true, "2015-08-04 00:00:00"),
    (3,'Emiliano','Martinez',"dibu@gmail.com",29,1.87, true, "2020-01-10 00:00:00"),
    (4,'Angel','Di Maria',"fideo@gmail.com",32,1.81, true, "2019-06-09 00:00:00"),
    (5,'Julian','Alvarez',"araña@gmail.com",20,1.75, true, "2012-03-07 00:00:00");
-- Desactivo el modo seguro de updates para poder actualizar toda la tabla de products
SET SQL_SAFE_UPDATES = 0;
-- Agrego la clave foranea de "warehouses" a la tabla "products"
ALTER TABLE `users` ADD `warehouse_id` INT NOT NULL AFTER `createdat`;
-- Asigno todos los productos al warehouse con id 1
UPDATE `users` SET `warehouse_id` = '1';