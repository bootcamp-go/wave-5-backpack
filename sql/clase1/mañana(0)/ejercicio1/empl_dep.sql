CREATE DATABASE  IF NOT EXISTS `empl_dep` /*!40100 DEFAULT CHARACTER SET utf8mb3 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `empl_dep`;
-- MySQL dump 10.13  Distrib 8.0.29, for Linux (x86_64)
--
-- Host: localhost    Database: empl_dep
-- ------------------------------------------------------
-- Server version	8.0.29-0ubuntu0.22.04.2

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
-- Table structure for table `departamentos`
--

DROP TABLE IF EXISTS `departamentos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `departamentos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(255) DEFAULT NULL,
  `direccion` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `departamentos`
--

LOCK TABLES `departamentos` WRITE;
/*!40000 ALTER TABLE `departamentos` DISABLE KEYS */;
INSERT INTO `departamentos` VALUES (1,'Recursos Humanos','Monroe 860'),(2,'Bootcamp Go','Monroe 860'),(3,'Investigación y Tecnología','Monroe 860'),(4,'Global Red Team','Monroe 860'),(5,'Administración','Monroe 860');
/*!40000 ALTER TABLE `departamentos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `empleados`
--

DROP TABLE IF EXISTS `empleados`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `empleados` (
  `dni` int NOT NULL,
  `num_legajo` varchar(45) CHARACTER SET utf8mb3 DEFAULT NULL,
  `nombre` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `apellido` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `fecha_nac` date DEFAULT NULL,
  `fecha_ingreso` date DEFAULT NULL,
  `cargo` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `sueldo` float(10,2) DEFAULT NULL,
  `neto` float(10,2) DEFAULT NULL,
  `departamentos_id` int DEFAULT NULL,
  PRIMARY KEY (`dni`),
  UNIQUE KEY `dni_UNIQUE` (`dni`),
  KEY `fk_empleados_1_idx` (`departamentos_id`),
  CONSTRAINT `fk_empleados_dep` FOREIGN KEY (`departamentos_id`) REFERENCES `departamentos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=armscii8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `empleados`
--

LOCK TABLES `empleados` WRITE;
/*!40000 ALTER TABLE `empleados` DISABLE KEYS */;
INSERT INTO `empleados` VALUES (50398102,'ABC123','Marcos','Beltran','1990-04-01','2022-07-01','Backend Developer GO',90000.00,80000.00,2),(50398103,'ABC124','Franco','Perez','1990-04-02','2022-07-01','Fronted Developer ReactJs',90000.00,80000.00,2),(50398104,'ABC125','Luciana','Rodriguez','1990-04-03','2022-07-01','Backend Developer PHP',90000.00,80000.00,2),(50398105,'ABC126','Pricilla','Sosa','1990-04-04','2022-07-01','Backend Developer GO',90000.00,80000.00,2),(50398106,'ABC127','Esteban','Blanco','1990-04-05','2022-07-01','Frontend Developer JS',90000.00,80000.00,2),(50398107,'ABC128','Patricio','Velez','1990-04-06','2022-07-01','Recruiter',90000.00,80000.00,1),(50398108,'ABC129','Nahuel','Pardo','1990-04-07','2022-07-01','Analista de Seguridad',90000.00,80000.00,4),(50398109,'ABC110','Francisco','Valle','1990-04-08','2022-07-01','Investigador IT',90000.00,80000.00,3),(50398110,'ABC111','Gisel','Pintos','1990-04-09','2022-07-01','Recruiter',90000.00,80000.00,1),(50398111,'ABC112','Iam','Charls','1990-04-10','2022-07-01','Recruiter',90000.00,80000.00,1),(50398112,'ABC113','Cristian','Valdes','1990-04-11','2022-07-01','Administrador',90000.00,80000.00,5),(50398113,'ABC114','Mora','Moreno','1990-04-12','2022-07-01','Backend Developer GO',90000.00,80000.00,2),(50398114,'ABC115','Maria','Soles','1990-04-13','2022-07-01','Backend Developer GO',90000.00,80000.00,2),(50398115,'ABC116','Karim','Benzema','1990-04-14','2022-07-01','Backend Developer GO',90000.00,80000.00,2),(50398116,'ABC117','Santiago','Lopez','1990-04-15','2022-07-01','Backend Developer GO',90000.00,80000.00,2),(50398117,'ABC118','Federico','Aimar','1990-04-16','2022-07-01','Backend Developer GO',90000.00,80000.00,2);
/*!40000 ALTER TABLE `empleados` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-09 10:58:08
