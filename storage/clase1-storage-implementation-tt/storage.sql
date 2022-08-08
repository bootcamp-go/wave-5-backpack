CREATE DATABASE IF NOT EXISTS `storage` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `storage`;

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `codeTransaction` varchar(60) NOT NULL,
  `currency` varchar(60) NOT NULL,
  `amount` float NOT NULL,
  `transmitter` varchar(60) NOT NULL,
  `receiver`varchar(60) NOT NULL,
  `date`varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;
