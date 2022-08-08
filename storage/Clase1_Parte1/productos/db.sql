CREATE DATABASE IF NOT EXISTS `storage` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `storage`;

CREATE TABLE `products`(
	`id` int(11) NOT NULL,
    `nombre` varchar(60) NOT NULL,
    `color` varchar(60) NOT NULL,
    `precio` int(11) NOT NULL,
    `stock` int(11) NOT NULL,
    `codigo` varchar(60) NOT NULL,
    `publicado` bool NOT NULL,
    `fecha_creacion` varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `products`
	ADD PRIMARY KEY (`id`);
    
ALTER TABLE `products`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;