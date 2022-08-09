CREATE DATABASE IF NOT EXISTS `storage` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `storage`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `firstName` varchar(60) NOT NULL,
  `lastName` varchar(60) NOT NULL,
  `email` varchar(60) NOT NULL,
  `age` int(11) NOT NULL,
  `height` float NOT NULL,
  `active` BOOLEAN,
  `creationDate` varchar(60),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*
type User struct {
	Id           int     `json:"-"`
	FirstName    string  `json:"firstName" binding:"required"`
	LastName     string  `json:"lastName" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Age          int     `json:"age" binding:"required"`
	Height       float64 `json:"height" binding:"required"`
	Active       bool    `json:"active" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}
*/