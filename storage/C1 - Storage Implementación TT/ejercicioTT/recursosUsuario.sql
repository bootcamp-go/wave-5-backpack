CREATE DATABASE IF NOT EXISTS `storage` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `storage`;

CREATE TABLE `usuarios` (
  `id` int(11) NOT NULL,
  `nombre` varchar(60) NOT NULL,
  `apellido` varchar(60) NOT NULL,
  `email` varchar(60) NOT NULL,
  `edad` int(3) NOT NULL,
  `altura` float NOT NULL,
  `activo` boolean NOT NULL DEFAULT 0,
  `fecha` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `usuarios`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `usuarios`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;
