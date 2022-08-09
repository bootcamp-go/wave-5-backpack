CREATE DATABASE IF NOT EXISTS Storage DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE Storage;

CREATE TABLE transactions(
    id int(11) NOT NULL,
    code varchar(60) NOT NULL,
    coin varchar(60) NOT NULL,
    amount float NOT NULL,
    emisor varchar(60) NOT NULL,
    receptor varchar(60) NOT NULL,
    date varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE transactions
    ADD PRIMARY KEY (id);

ALTER TABLE transactions
    MODIFY id int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;