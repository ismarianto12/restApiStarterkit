-- MySQL dump 10.13  Distrib 9.3.0, for Linux (aarch64)
--
-- Host: localhost    Database: golang_rest
-- ------------------------------------------------------
-- Server version	9.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `barang`
--

DROP TABLE IF EXISTS `barang`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `barang` (
  `id` int NOT NULL AUTO_INCREMENT,
  `kode` varchar(30) DEFAULT NULL,
  `nama` varchar(50) DEFAULT NULL,
  `kategory_id` int DEFAULT NULL,
  `deskripsi` text,
  `gambar` varchar(14) DEFAULT NULL,
  `created_at` varchar(34) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `updated_at` varchar(34) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_by` varchar(34) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `updated_by` varchar(34) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `deleted_at` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `barang`
--

LOCK TABLES `barang` WRITE;
/*!40000 ALTER TABLE `barang` DISABLE KEYS */;
INSERT INTO `barang` VALUES (5,'CROT 86','sabun coli adalah',0,'adadasdsa adad','crotdalam.png','2025-07-22 16:34:31.158','2025-07-22 16:34:31.158',NULL,NULL,NULL),(6,'1231','Sabun Coli',0,'ada','ada.png','2025-07-22 16:34:31.549','2025-07-22 16:34:31.549',NULL,NULL,NULL),(7,'1231','Sabun Coli',0,'ada','ada.png','2025-07-22 16:34:31.99','2025-07-22 16:34:31.99',NULL,NULL,NULL),(8,'1231','Sabun Coli',0,'ada','ada.png','2025-07-22 16:34:33.453','2025-07-22 16:34:33.453',NULL,NULL,NULL),(9,'1231','Sabun Coli',1,'ada','ada.png','2025-07-22 16:34:58.889','2025-07-22 16:34:58.889',NULL,NULL,NULL),(10,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:18.673','2025-07-22 16:35:18.673',NULL,NULL,NULL),(11,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:19.172','2025-07-22 16:35:19.172',NULL,NULL,NULL),(12,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:19.692','2025-07-22 16:35:19.692',NULL,NULL,NULL),(13,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:20.015','2025-07-22 16:35:20.015',NULL,NULL,NULL),(14,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:20.321','2025-07-22 16:35:20.321',NULL,NULL,NULL),(15,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:21.736','2025-07-22 16:35:21.736',NULL,NULL,NULL),(16,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:22.142','2025-07-22 16:35:22.142',NULL,NULL,NULL),(17,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:22.577','2025-07-22 16:35:22.577',NULL,NULL,NULL),(18,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:23.005','2025-07-22 16:35:23.005',NULL,NULL,NULL),(19,'1231','Sabun Coli',2,'ada','ada.png','2025-07-22 16:35:23.398','2025-07-22 16:35:23.398',NULL,NULL,NULL);
/*!40000 ALTER TABLE `barang` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `purchasing`
--

DROP TABLE IF EXISTS `purchasing`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `purchasing` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int DEFAULT NULL,
  `facture_code` varchar(30) DEFAULT NULL,
  `pemeriksa` varchar(30) DEFAULT NULL,
  `created_by` varchar(30) DEFAULT NULL,
  `updated_by` varchar(30) DEFAULT NULL,
  `deleted_at` varchar(14) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `purchasing`
--

LOCK TABLES `purchasing` WRITE;
/*!40000 ALTER TABLE `purchasing` DISABLE KEYS */;
INSERT INTO `purchasing` VALUES (1,1,'adsad','adada','','',NULL),(2,1,'dad','adada','','',NULL);
/*!40000 ALTER TABLE `purchasing` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stock`
--

DROP TABLE IF EXISTS `stock`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `stock` (
  `id` int NOT NULL AUTO_INCREMENT,
  `barang_id` int DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `deskripsi` text,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stock`
--

LOCK TABLES `stock` WRITE;
/*!40000 ALTER TABLE `stock` DISABLE KEYS */;
/*!40000 ALTER TABLE `stock` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `suplier`
--

DROP TABLE IF EXISTS `suplier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `suplier` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama` varchar(14) DEFAULT NULL,
  `kode` varchar(15) DEFAULT NULL,
  `contact` varchar(15) DEFAULT NULL,
  `address` text,
  `created_at` varchar(14) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `suplier`
--

LOCK TABLES `suplier` WRITE;
/*!40000 ALTER TABLE `suplier` DISABLE KEYS */;
INSERT INTO `suplier` VALUES (26,'','','','',NULL),(27,'','','','',NULL),(28,'','','','',NULL),(29,'','','','',NULL),(30,'','','','',NULL),(31,'','','','',NULL),(32,'','','','',NULL),(33,'','','','',NULL),(34,'','','','',NULL),(35,'','','','',NULL),(36,'','','','',NULL),(37,'','','','',NULL),(38,'','','','',NULL),(39,'','','','',NULL),(40,'','','','',NULL),(41,'','','','',NULL),(42,'','','','',NULL);
/*!40000 ALTER TABLE `suplier` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(15) DEFAULT NULL,
  `image` varchar(30) DEFAULT NULL,
  `level_id` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (23,'rian','ismarianto@gmail.com','123','2025-07-23 15:42:31','2025-07-23 15:42:31',NULL,'admin','','1');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-07-25 13:57:08
