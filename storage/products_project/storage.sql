CREATE DATABASE IF NOT EXISTS Storage DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE Storage;

CREATE TABLE products(
    id int(11) NOT NULL,
    name varchar(60) NOT NULL,
    type varchar(60) NOT NULL,
    count int(11) NOT NULL,
    price float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE products
    ADD PRIMARY KEY (id);

ALTER TABLE products
    MODIFY id int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;