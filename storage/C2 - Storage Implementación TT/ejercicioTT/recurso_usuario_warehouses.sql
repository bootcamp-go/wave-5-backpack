-- creación de base de datos, si no existe, se selecciona
CREATE DATABASE IF NOT EXISTS `storage` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `storage`;
-- Si existe una tabla "warehouse" se elimina y se crea una nueva
DROP TABLE IF EXISTS `WAREHOUSES`;

CREATE TABLE `warehouses` (
  `id` INT NOT NULL PRIMARY KEY,
  `name` VARCHAR(40) NOT NULL,
  `address` VARCHAR(150) NOT NULL
);

-- Creación del warehouse principal
INSERT INTO `warehouses` (`id`,`name`,`address`) VALUES (1,'Main Warehouse','221 Baker Street');

-- Si existe una tabla "usuarios" la elimina y se crea una nueva
DROP TABLE IF EXISTS `usuarios`;

CREATE TABLE `usuarios` (
  `id` int(11) NOT NULL,
  `nombre` varchar(60) NOT NULL,
  `apellido` varchar(60) NOT NULL,
  `email` varchar(60) NOT NULL,
  `edad` int(3) NOT NULL,
  `altura` float NOT NULL,
  `activo` boolean DEFAULT 1,
  `fecha` datetime
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Creo algunos usuarios
INSERT INTO `usuarios` VALUES
(1,"Martha","Hernandez","martha@hotmail.es",61,1.64,1,'2020-01-19 03:15:07'),
(2,"Luber","Lucumi","luber@hotmail.es",62,1.80,1,'2020-01-22 03:14:07'),
(3,"Luz","Lucumi Hernandez","luz@hotmail.es",26,1.65,1,'2020-02-19 03:26:07'),
(4,"Paquita","Lucumi Hernandez","paquis@guau.dog",1,0.40,1,'2022-05-19 03:06:07');

SET SQL_SAFE_UPDATES=0;
-- Agrego clave foránea de warehouse
ALTER TABLE `usuarios` ADD `warehouse_id` INT NOT NULL AFTER `fecha`;
-- Asigno todos los usuarios al warehouse con id 1
UPDATE `usuarios` SET `warehouse_id` = 1;

ALTER TABLE `usuarios`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `usuarios`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;