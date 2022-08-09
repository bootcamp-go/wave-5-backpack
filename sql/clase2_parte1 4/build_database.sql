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
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` INT NOT NULL PRIMARY KEY,
  `name` VARCHAR(60) NOT NULL,
  `type` VARCHAR(60) NOT NULL,
  `count` INT NOT NULL,
  `price` DOUBLE NOT NULL
);
-- Creo unos productos de prueba
INSERT INTO `products` VALUES 
    (1,'Campera','Indumentaria',20,1000),
    (2,'Pantalon','Indumentaria',20,2500.50),
    (3,'Coca Cola','Gaseosa',10,300),
    (4,'Fanta','Gaseosa',10,200.99),
    (5,'Heladera','Electrodomesticos',56,80499.99),
    (6,'Horno','Electrodomesticos',56,55899.99);
-- Desactivo el modo seguro de updates para poder actualizar toda la tabla de products
SET SQL_SAFE_UPDATES = 0;
-- Agrego la clave foranea de "warehouses" a la tabla "products"
ALTER TABLE `products` ADD `warehouse_id` INT NOT NULL AFTER `price`;
-- Asigno todos los productos al warehouse con id 1
UPDATE `products` SET `warehouse_id` = '1';
