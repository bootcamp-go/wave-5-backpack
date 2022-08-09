CREATE DATABASE IF NOT EXISTS `storage_clase`  DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `storage_clase`;
-- DROP TABLE `usuarios`;

CREATE TABLE `usuarios` (
  `id` int(11) NOT NULL,
  `nombre` varchar(60) NOT NULL,
  `apellido` varchar(60) NOT NULL,
  `email` varchar(60) NOT NULL,
  `edad` int(11) NOT NULL,
  `altura` int(11) NOT NULL,
  `activo` bool NOT NULL,
  `fecha_de_creacion` varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `usuarios`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `usuarios`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;
COMMIT;

SELECT * FROM usuarios;