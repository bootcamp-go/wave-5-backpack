CREATE DATABASE IF NOT EXISTS `storage` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `storage`;

CREATE TABLE `transactions` (
                            `id` int(11) NOT NULL,
                            `transaction_code` varchar(60) NOT NULL,
                            `type_currency` varchar(60) NOT NULL,
                            `amount` int(11) NOT NULL,
                            `transmitter` float NOT NULL,
                            `receiver` varchar(60) NOT NULL,
                            `date` varchar(60) NOT NULL,
                            `completed` bool NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `transactions`
    ADD PRIMARY KEY (`id`);

ALTER TABLE `transactions`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;
