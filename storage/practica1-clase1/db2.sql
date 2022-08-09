-- Active: 1659473777000@@127.0.0.1@3306@storage
CREATE TABLE `warehouses` (
  `id` int(11) NOT NULL,
  `name` varchar(40) NOT NULL,
  `adress` varchar(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Volcado de datos para la tabla `warehouses`
INSERT INTO `warehouses` (`id`, `name`, `adress`) VALUES
(1, 'Main Warehouse', '221b Baker Street');

ALTER TABLE `warehouses`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `warehouses`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

ALTER TABLE `products` ADD `id_warehouse` INT NOT NULL AFTER `price`;

UPDATE `products` SET `id_warehouse` = '1';