CREATE DATABASE IF NOT EXISTS `storage` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE storage;

CREATE TABLE `TRANSACTIONS` (
	`id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `cod_transaction` varchar(20) NOT NULL,
	`currency` varchar(4),
    `amount` int(20),
    `sender` varchar(100),
    `receiver` varchar(100),
    `date_order` varchar(100)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

COMMIT;

DROP DATABASE storage;