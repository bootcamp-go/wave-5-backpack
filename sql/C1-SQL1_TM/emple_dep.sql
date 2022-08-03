-- MySQL dump 10.13  Distrib 8.0.30, for macos12 (x86_64)
--
-- Host: localhost    Database: emple_dep
-- ------------------------------------------------------
-- Server version	8.0.29

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Departamentos`
--

DROP TABLE IF EXISTS `Departamentos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Departamentos` (
  `id_departamento` int NOT NULL AUTO_INCREMENT,
  `nombre_departamento` varchar(45) DEFAULT NULL,
  `direccion_departamento` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_departamento`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Departamentos`
--

LOCK TABLES `Departamentos` WRITE;
/*!40000 ALTER TABLE `Departamentos` DISABLE KEYS */;
INSERT INTO `Departamentos` VALUES (1,'Santander','Cra 3'),(2,'Cundinamarca','Cra15'),(3,'Antioquia','Cra 22'),(4,'Arauca','Cra 34'),(5,'Medellin','Cra 33'),(6,'Cesar','Cra 68');
/*!40000 ALTER TABLE `Departamentos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Empleados`
--

DROP TABLE IF EXISTS `Empleados`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Empleados` (
  `DNI` int NOT NULL AUTO_INCREMENT,
  `numero_legado` int DEFAULT NULL,
  `apellido` varchar(45) DEFAULT NULL,
  `nombre` varchar(45) DEFAULT NULL,
  `fecha_nacimiento` datetime DEFAULT NULL,
  `fecha_incorporacion` datetime DEFAULT NULL,
  `cargo` varchar(45) DEFAULT NULL,
  `sueldo_neto` int DEFAULT NULL,
  `departamento_ID` int DEFAULT NULL,
  PRIMARY KEY (`DNI`),
  KEY `tdepartamento_templeado_idx` (`departamento_ID`),
  CONSTRAINT `tdepartamento_templeado` FOREIGN KEY (`departamento_ID`) REFERENCES `Departamentos` (`id_departamento`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Empleados`
--

LOCK TABLES `Empleados` WRITE;
/*!40000 ALTER TABLE `Empleados` DISABLE KEYS */;
INSERT INTO `Empleados` VALUES (1,2132,'Carvajal','Sebastian','1997-10-22 00:00:00','2022-01-10 00:00:00','Software developer',4500000,3),(2,3212,'Ulloa','Carlos','1997-09-22 00:00:00','2022-02-22 00:00:00','Software developer',4500000,2),(3,1132,'Martinez','Jose','1998-10-30 00:00:00','2022-04-23 00:00:00','Software developer',4500000,4),(4,1122,'Peña','Armando','1998-10-30 00:00:00','2022-04-23 00:00:00','Software developer',4500000,2),(5,1134,'Rodriguez','Michael','1998-10-30 00:00:00','2022-04-23 00:00:00','Tester',4500000,3),(6,4321,'Gomez','Jesus','1998-10-30 00:00:00','2022-04-23 00:00:00','Tester',3500000,1),(7,4621,'Miranda','Juan','1998-10-30 00:00:00','2022-04-23 00:00:00','Tester',3500000,4),(8,4221,'Amaya','Lucas','1998-10-30 00:00:00','2022-04-23 00:00:00','Tester',3500000,3),(9,4221,'Carrascal','Luis','1998-10-30 00:00:00','2022-04-23 00:00:00','FrontEnd developer',4200000,2),(10,4221,'Montero','Luisa','1998-10-30 00:00:00','2022-04-23 00:00:00','BackEnd developer',5200000,1);
/*!40000 ALTER TABLE `Empleados` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-08-03 10:27:08
