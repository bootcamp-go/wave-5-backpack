/*-----------------------------------------------------------------------*

     Assignment:	C2 - TT | Practica #2
         Author:	Israel Fabela
	   Language:	mysql  Ver 8.0.29 for macos12.2 on arm64
		  Topic:	Storage Implementation

	© Mercado Libre - IT Bootcamp 2022

-------------------------------------------------------------------------*/


-- Se crea la base de datos, sino existe, la selecciono.
CREATE DATABASE IF NOT EXISTS `storage`;
USE `storage`;

-- Si existe una tabla `warehouse` la elimino y creo una nueva.
DROP TABLE IF EXISTS `warehouses`;
CREATE TABLE `warehouses` (
`id` INT NOT NULL PRIMARY KEY,
`name` VARCHAR(50) NOT NULL,
`address` VARCHAR(150) NOT NULL
);

--  Se crea un `warehouses` principal.
INSERT INTO `warehouses` (`id`,`name`,address) VALUES (1, 'Main Warehouse', '221 Baker Street');

-- Si existe una table `transactions` la elimino y creo una nueva.
DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL PRIMARY KEY,
  `codeTransaction` varchar(60) NOT NULL,
  `currency` varchar(60) NOT NULL,
  `amount` float NOT NULL,
  `transmitter` varchar(60) NOT NULL,
  `receiver`varchar(60) NOT NULL,
  `date`varchar(60) NOT NULL
);

-- Inserto unas 'transacciones' de prueba.
INSERT INTO transactions VALUES
	(1, 'abc', 'JPY', 1012.76, 'SMFG', 'Mitsubishi UFJ','2018-06-24'),
	(2, 'cde', 'EUR', 983.07, 'Lloyds Banking', 'Deutsche Bank-Rg','2019-05-11'),
	(3, 'efg', 'MXN', 2302.75, 'BBVA', 'Banorte','2017-12-24'),
	(4, 'hij', 'USD', 1012.76, 'Bank Of America', 'Morgan Stanley','2016-04-07'),
	(5, 'klm', 'KPW', 1012.76, 'Shang Pudong', 'Kb Financial Gro','2015-07-23'),
	(6, 'nop', 'CHF', 1012.76, 'Bankcomm-H', 'Bank Of Ningbo','2018-06-24');

-- Desactivo el modo seguro de updates para poder actualizar toda la tabla de products
SET SQL_SAFE_UPDATES = 0;

-- Agrego la clave foranea de "warehouses" a la tabla "products"
ALTER TABLE transactions ADD `warehouse_id` INT NOT NULL AFTER `amount`;

-- Asigno todos los productos al warehouse con id 1
UPDATE `transactions` SET `warehouse_id`='1';

-- Asignar al 'ID' y 'warehouse_ID' un valor de AutoIncrementar o Default, respectivamente.
ALTER TABLE `transactions`
	MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;
    
ALTER TABLE `transactions`
	MODIFY  `warehouse_id` INT NOT NULL DEFAULT  '1';