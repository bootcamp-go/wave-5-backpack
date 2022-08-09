DROP DATABASE IF EXISTS `storage_bootcamp`;

CREATE DATABASE IF NOT EXISTS `storage_bootcamp` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `storage_bootcamp`;

CREATE TABLE `transactions` (
	`id` int(11) NOT NULL,
    `monto` float NOT NULL,
    `cod` varchar(60) NOT NULL,
    `moneda` varchar(20) NOT NULL,
    `emisor` varchar(60) NOT NULL,
    `receptor` varchar(60) NOT NULL,
    `fecha` datetime
);

ALTER TABLE `transactions` 
	ADD PRIMARY KEY (`id`);

ALTER TABLE `transactions`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;
