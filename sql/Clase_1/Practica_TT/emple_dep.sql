-- MySQL dump 10.13  Distrib 8.0.30, for macos12 (x86_64)
--
-- Host: 127.0.0.1    Database: emple_dep
-- ------------------------------------------------------
-- Server version	5.7.34

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
-- Table structure for table `departamento`
--

DROP TABLE IF EXISTS `departamento`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `departamento` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Nombre` varchar(45) DEFAULT NULL,
  `Direccion` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `departamento`
--

LOCK TABLES `departamento` WRITE;
/*!40000 ALTER TABLE `departamento` DISABLE KEYS */;
INSERT INTO `departamento` VALUES (1,'IT','Isidora 435'),(2,'IT','Isidora 435'),(3,'IT','Isidora 435'),(4,'IT','Isidora 435'),(5,'Marketing','Alameda 3100');
/*!40000 ALTER TABLE `departamento` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `empleados`
--

DROP TABLE IF EXISTS `empleados`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `empleados` (
  `DNI` int(11) NOT NULL AUTO_INCREMENT,
  `Nombre` varchar(50) DEFAULT NULL,
  `Apellido` varchar(50) DEFAULT NULL,
  `Fecha_Nacimiento` varchar(50) DEFAULT NULL,
  `N_legajo` float DEFAULT NULL,
  `Fecha_Incorporacion` varchar(50) DEFAULT NULL,
  `Sueldo_neto` float DEFAULT NULL,
  `Id_dpto` int(11) DEFAULT NULL,
  PRIMARY KEY (`DNI`)
) ENGINE=InnoDB AUTO_INCREMENT=7898357 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `empleados`
--

LOCK TABLES `empleados` WRITE;
/*!40000 ALTER TABLE `empleados` DISABLE KEYS */;
INSERT INTO `empleados` VALUES (1134,'Pedro','Gonzalez','05-11-1995',40599,'01-03-2022',1900000,1),(1234,'Juan','Perez','02-08-1991',40598,'01-03-2022',1200000,1),(4561,'Pedro','Pascal','30-03-1988',40602,'01-03-2022',1700000,3),(4891,'Juan','Rebolledo','05-10-1992',40603,'01-03-2022',1500000,3),(9639,'Ester','Allendes','09-09-1985',40604,'01-03-2022',2000000,1),(12345,'Karina','Manriquez','26-09-1989',40601,'01-03-2022',1800000,2),(15645,'Bernarda','Cordero','13-01-1965',40607,'01-03-2022',2200000,5),(19241,'Julio','Nuñez','08-12-1962',40608,'01-03-2022',1900000,2),(78915,'Fernando','Flores','26-08-1991',40606,'01-03-2022',1900000,3),(98793,'Elias','Gutierrez','07-08-1995',40608,'01-03-2022',1900000,5),(111234,'Ana','Rodriguez','15-11-1990',40600,'01-03-2022',2000000,2),(156489,'Javiera','Nuñez','07-02-1965',40606,'01-03-2022',2200000,4),(974265,'Nicolas','Rios','08-01-2000',40609,'01-03-2022',2000000,5),(987135,'Karol','Rios','10-10-1991',40605,'01-03-2022',2900000,1),(7898356,'Nataly','Herrera','05-11-1985',40607,'01-03-2022',1900000,4);
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

-- Dump completed on 2022-08-04 20:41:30
