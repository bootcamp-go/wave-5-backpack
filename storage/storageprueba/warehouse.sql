CREATE TABLE warehouses (
    id int(11) NOT NULL, 
    name varchar(30) NOT NULL,
    address varchar(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO warehouses (`id`, `name`, `address`) VALUES (1, 'Main Warehouse', '221b Baker Street');

ALTER TABLE warehouses ADD PRIMARY KEY (`id`);

ALTER TABLE warehouses MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

ALTER TABLE products ADD `id_warehouse` INT NOT NULL AFTER `fechaCreacion`;


UPDATE products SET `id_warehouse` = '1';
