use storage;

drop table if exists products;
drop table if exists warehouses;

CREATE TABLE products (
	id         INT PRIMARY KEY AUTO_INCREMENT,
	name       VARCHAR(50),
	color      VARCHAR(50),
	price      DOUBLE,
	stock      INT,
	code       VARCHAR(50),
	published  TINYINT(1),
	created_at DATE
);

INSERT INTO `products` (`name`, `color`, price, stock, code, published, created_at) VALUES
('product 1', 'red', 10.99, 100, 'HJ988BH', 1, CURDATE());

CREATE TABLE `warehouses` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(40) NOT NULL,
  `adress` varchar(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Volcado de datos para la tabla `warehouses`
INSERT INTO `warehouses` (`id`, `name`, `adress`) VALUES
(1, 'Main Warehouse', '221b Baker Street');

ALTER TABLE `products` ADD `warehouse_id` INT NOT NULL;

UPDATE `products` SET `warehouse_id` = '1';


SELECT * FROM products;
SELECT * FROM warehouses;