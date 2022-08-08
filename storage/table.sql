CREATE DATABASE IF NOT EXISTS transactions;
USE transactions;

CREATE TABLE transactions(
    id INT PRIMARY KEY AUTO_INCREMENT,
    transaction_code varchar(50) UNIQUE,
    currency varchar(5),
    amount DECIMAL,
    sender varchar(100),
    reciever varchar(100),
    transaction_date DATETIME
)