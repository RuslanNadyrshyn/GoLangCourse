-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: localhost    Database: nadyrshyn_db
-- ------------------------------------------------------
-- Server version	8.0.28-0ubuntu0.20.04.3

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
-- Table structure for table `basket_product`
--

DROP TABLE IF EXISTS `basket_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `basket_product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `basket_id` int DEFAULT NULL,
  `product_id` int DEFAULT NULL,
  `price` float DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `baskets_products_id_uindex` (`id`),
  KEY `baskets_products_baskets_id_fk` (`basket_id`),
  KEY `basket_product_products_id_fk` (`product_id`),
  CONSTRAINT `basket_product_products_id_fk` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `baskets_products_baskets_id_fk` FOREIGN KEY (`basket_id`) REFERENCES `baskets` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `basket_product`
--

LOCK TABLES `basket_product` WRITE;
/*!40000 ALTER TABLE `basket_product` DISABLE KEYS */;
/*!40000 ALTER TABLE `basket_product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `baskets`
--

DROP TABLE IF EXISTS `baskets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `baskets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `price` float DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `baskets_id_uindex` (`id`),
  KEY `baskets_users_id_fk` (`user_id`),
  CONSTRAINT `baskets_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `baskets`
--

LOCK TABLES `baskets` WRITE;
/*!40000 ALTER TABLE `baskets` DISABLE KEYS */;
/*!40000 ALTER TABLE `baskets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ingredients`
--

DROP TABLE IF EXISTS `ingredients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ingredients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ingredients_id_uindex` (`id`),
  UNIQUE KEY `ingredients_name_uindex` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=16137 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ingredients`
--

LOCK TABLES `ingredients` WRITE;
/*!40000 ALTER TABLE `ingredients` DISABLE KEYS */;
INSERT INTO `ingredients` VALUES (15528,'Salmon','2022-01-30 22:13:11',NULL),(15529,'Cheese','2022-01-30 22:13:11',NULL),(15530,'Mozarella','2022-01-30 22:13:11',NULL),(15532,'Nori','2022-01-30 22:13:11',NULL),(15533,'Caramelized onions','2022-01-30 22:13:11',NULL),(15534,'Peperoni','2022-01-30 22:13:11',NULL),(15536,'Rice','2022-01-30 22:13:11',NULL),(15537,'Tomatoes','2022-01-30 22:13:11',NULL),(15540,'Cucumber','2022-01-30 22:13:11',NULL),(15541,'Original sauce','2022-01-30 22:13:11',NULL),(15542,'BBQ sauce','2022-01-30 22:13:11',NULL),(15544,'Cream cheese','2022-01-30 22:13:11',NULL),(15545,'Chicken','2022-01-30 22:13:11',NULL),(15547,'Corn','2022-01-30 22:13:11',NULL),(15550,'Onion','2022-01-30 22:13:11',NULL),(15553,'Mushrooms','2022-01-30 22:13:11',NULL),(15558,'Bavarian sausages','2022-01-30 22:13:11',NULL),(15565,'Unagi sauce','2022-01-30 22:13:11',NULL),(15568,'Cheese cheddar','2022-01-30 22:13:11',NULL),(15569,'Bacon','2022-01-30 22:13:11',NULL),(15573,'Pineapple','2022-01-30 22:13:11',NULL),(15579,'Onion rings','2022-01-30 22:13:11',NULL),(15582,'Pasta','2022-01-30 22:13:11',NULL),(15583,'Flour','2022-01-30 22:13:11',NULL),(15586,'Meat','2022-01-30 22:13:11',NULL),(15589,'Strawberry','2022-01-30 22:13:11',NULL),(15590,'Peas','2022-01-30 22:13:11',NULL),(15591,'Japanese tamago','2022-01-30 22:13:11',NULL),(15592,'Mayo','2022-01-30 22:13:11',NULL),(15593,'Butter','2022-01-30 22:13:11',NULL),(15594,'Carrots','2022-01-30 22:13:11',NULL),(15596,'Cottage Cheese','2022-01-30 22:13:11',NULL),(15598,'Chili Pepper','2022-01-30 22:13:11',NULL),(15599,'Sugar','2022-01-30 22:13:11',NULL),(15604,'Garlic','2022-01-30 22:13:11',NULL),(15606,'Potatoes','2022-01-30 22:13:11',NULL),(15607,'Basil','2022-01-30 22:13:11',NULL),(15612,'Peanuts','2022-01-30 22:13:11',NULL),(15613,'Milk','2022-01-30 22:13:11',NULL),(15614,'Grapes','2022-01-30 22:13:11',NULL);
/*!40000 ALTER TABLE `ingredients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menus`
--

DROP TABLE IF EXISTS `menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menus` (
  `id` int NOT NULL AUTO_INCREMENT,
  `supplier_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `menu_id_uindex` (`id`),
  KEY `menus_suppliers_id_fk` (`supplier_id`),
  CONSTRAINT `menus_suppliers_id_fk` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1265 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menus`
--

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;
INSERT INTO `menus` VALUES (1258,1,'2022-01-30 22:44:53',NULL),(1259,2,'2022-01-30 22:44:53',NULL),(1260,3,'2022-01-30 22:44:53',NULL),(1261,4,'2022-01-30 22:44:53',NULL),(1262,5,'2022-01-30 22:44:53',NULL),(1263,6,'2022-01-30 22:44:53',NULL),(1264,7,'2022-01-30 22:44:53',NULL);
/*!40000 ALTER TABLE `menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_product`
--

DROP TABLE IF EXISTS `order_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `count` int DEFAULT NULL,
  `price` float DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `orders_products_id_uindex` (`id`),
  KEY `orders_products_orders_id_fk` (`order_id`),
  KEY `order_product_products_id_fk` (`product_id`),
  CONSTRAINT `order_product_products_id_fk` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `orders_products_orders_id_fk` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_product`
--

LOCK TABLES `order_product` WRITE;
/*!40000 ALTER TABLE `order_product` DISABLE KEYS */;
INSERT INTO `order_product` VALUES (1,1,3,5,7.85,'2022-04-22 21:27:26',NULL),(2,1,2,1,2.5,'2022-04-22 21:27:26',NULL),(3,1,4,3,2.65,'2022-04-22 21:27:27',NULL),(4,1,15,1,3.84,'2022-04-22 21:27:27',NULL),(5,1,11,1,7.25,'2022-04-22 21:27:27',NULL),(6,1,1,7,6.07,'2022-04-22 21:27:27',NULL),(7,1,5,6,3.83,'2022-04-22 21:27:27',NULL),(8,2,1,1,6.07,'2022-04-22 21:29:14',NULL),(9,2,2,1,2.5,'2022-04-22 21:29:15',NULL),(10,3,1,1,6.07,'2022-04-22 21:33:57',NULL),(11,3,2,1,2.5,'2022-04-22 21:33:58',NULL),(12,4,1,1,6.07,'2022-04-22 22:13:44',NULL),(13,4,2,1,2.5,'2022-04-22 22:13:44',NULL),(14,5,2,1,2.5,'2022-04-22 22:30:13',NULL),(15,5,1,3,6.07,'2022-04-22 22:30:13',NULL),(16,6,2,1,2.5,'2022-04-22 22:33:10',NULL),(17,6,1,3,6.07,'2022-04-22 22:33:10',NULL),(18,7,2,1,2.5,'2022-04-22 22:34:37',NULL),(19,7,1,3,6.07,'2022-04-22 22:34:37',NULL),(20,8,4,1,2.65,'2022-04-22 22:35:16',NULL),(21,9,3,1,7.85,'2022-04-22 22:36:05',NULL),(22,10,2,1,2.5,'2022-04-22 22:40:03',NULL),(23,11,4,1,2.65,'2022-04-22 22:53:21',NULL),(24,12,18,15,2.59,'2022-04-23 19:35:26',NULL);
/*!40000 ALTER TABLE `order_product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `price` float DEFAULT NULL,
  `user_id` int NOT NULL,
  `address` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `orders_id_uindex` (`id`),
  KEY `orders_users_id_fk` (`user_id`),
  CONSTRAINT `orders_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,126.26,266,'addr2','2022-04-22 21:27:25',NULL),(2,8.57,267,'addr3','2022-04-22 21:29:14',NULL),(3,8.57,268,'sad','2022-04-22 21:33:57',NULL),(4,8.57,270,'fsfd','2022-04-22 22:13:43',NULL),(5,20.71,271,'adad','2022-04-22 22:30:13',NULL),(6,20.71,272,'adad','2022-04-22 22:33:10',NULL),(7,20.71,273,'adad','2022-04-22 22:34:37',NULL),(8,2.65,274,'asd','2022-04-22 22:35:16',NULL),(9,7.85,275,'dfsd','2022-04-22 22:36:04',NULL),(10,2.5,276,'sdas','2022-04-22 22:40:03',NULL),(11,2.65,277,'dsfsd','2022-04-22 22:53:21',NULL),(12,38.85,283,'adasd','2022-04-23 19:35:25',NULL);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_ingredients`
--

DROP TABLE IF EXISTS `product_ingredients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_ingredients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int DEFAULT NULL,
  `ingredient_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `product_ingredients_id_uindex` (`id`),
  KEY `product_ingredients_ingredients_id_fk` (`ingredient_id`),
  KEY `product_ingredients_products_id_fk` (`product_id`),
  CONSTRAINT `product_ingredients_ingredients_id_fk` FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `product_ingredients_products_id_fk` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14341 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_ingredients`
--

LOCK TABLES `product_ingredients` WRITE;
/*!40000 ALTER TABLE `product_ingredients` DISABLE KEYS */;
INSERT INTO `product_ingredients` VALUES (14254,4,15529,'2022-01-30 22:44:53',NULL),(14255,1,15530,'2022-01-30 22:44:53',NULL),(14256,7,15530,'2022-01-30 22:44:53',NULL),(14257,4,15533,'2022-01-30 22:44:53',NULL),(14258,10,15528,'2022-01-30 22:44:53',NULL),(14259,1,15534,'2022-01-30 22:44:53',NULL),(14260,7,15534,'2022-01-30 22:44:53',NULL),(14261,4,15537,'2022-01-30 22:44:53',NULL),(14262,10,15532,'2022-01-30 22:44:53',NULL),(14263,1,15537,'2022-01-30 22:44:53',NULL),(14264,7,15537,'2022-01-30 22:44:53',NULL),(14265,10,15536,'2022-01-30 22:44:53',NULL),(14266,1,15542,'2022-01-30 22:44:53',NULL),(14267,7,15542,'2022-01-30 22:44:53',NULL),(14268,4,15541,'2022-01-30 22:44:53',NULL),(14269,10,15540,'2022-01-30 22:44:53',NULL),(14270,10,15544,'2022-01-30 22:44:53',NULL),(14271,8,15547,'2022-01-30 22:44:53',NULL),(14272,2,15547,'2022-01-30 22:44:53',NULL),(14273,5,15545,'2022-01-30 22:44:53',NULL),(14274,11,15528,'2022-01-30 22:44:53',NULL),(14275,2,15550,'2022-01-30 22:44:53',NULL),(14276,8,15550,'2022-01-30 22:44:53',NULL),(14277,5,15537,'2022-01-30 22:44:53',NULL),(14278,8,15553,'2022-01-30 22:44:53',NULL),(14279,11,15532,'2022-01-30 22:44:53',NULL),(14280,5,15553,'2022-01-30 22:44:53',NULL),(14281,2,15553,'2022-01-30 22:44:53',NULL),(14282,8,15558,'2022-01-30 22:44:53',NULL),(14283,5,15550,'2022-01-30 22:44:53',NULL),(14284,11,15536,'2022-01-30 22:44:53',NULL),(14285,8,15530,'2022-01-30 22:44:53',NULL),(14286,2,15558,'2022-01-30 22:44:53',NULL),(14287,5,15542,'2022-01-30 22:44:53',NULL),(14288,11,15540,'2022-01-30 22:44:53',NULL),(14289,8,15542,'2022-01-30 22:44:53',NULL),(14290,2,15530,'2022-01-30 22:44:53',NULL),(14291,6,15568,'2022-01-30 22:44:53',NULL),(14292,11,15544,'2022-01-30 22:44:53',NULL),(14293,2,15542,'2022-01-30 22:44:53',NULL),(14294,9,15545,'2022-01-30 22:44:53',NULL),(14295,6,15569,'2022-01-30 22:44:53',NULL),(14296,11,15565,'2022-01-30 22:44:53',NULL),(14297,9,15573,'2022-01-30 22:44:53',NULL),(14298,3,15545,'2022-01-30 22:44:53',NULL),(14299,6,15541,'2022-01-30 22:44:53',NULL),(14300,9,15530,'2022-01-30 22:44:53',NULL),(14301,12,15528,'2022-01-30 22:44:53',NULL),(14302,6,15579,'2022-01-30 22:44:53',NULL),(14303,3,15573,'2022-01-30 22:44:53',NULL),(14304,12,15532,'2022-01-30 22:44:53',NULL),(14305,3,15530,'2022-01-30 22:44:53',NULL),(14306,12,15536,'2022-01-30 22:44:53',NULL),(14307,12,15540,'2022-01-30 22:44:53',NULL),(14308,13,15582,'2022-01-30 22:44:53',NULL),(14309,16,15545,'2022-01-30 22:44:53',NULL),(14310,13,15586,'2022-01-30 22:44:53',NULL),(14311,19,15583,'2022-01-30 22:44:53',NULL),(14312,12,15544,'2022-01-30 22:44:53',NULL),(14313,16,15583,'2022-01-30 22:44:53',NULL),(14314,13,15592,'2022-01-30 22:44:53',NULL),(14315,19,15589,'2022-01-30 22:44:53',NULL),(14316,16,15590,'2022-01-30 22:44:53',NULL),(14317,12,15565,'2022-01-30 22:44:53',NULL),(14318,19,15593,'2022-01-30 22:44:53',NULL),(14319,16,15594,'2022-01-30 22:44:53',NULL),(14320,12,15591,'2022-01-30 22:44:53',NULL),(14321,14,15547,'2022-01-30 22:44:53',NULL),(14322,20,15596,'2022-01-30 22:44:53',NULL),(14323,14,15550,'2022-01-30 22:44:53',NULL),(14324,17,15598,'2022-01-30 22:44:53',NULL),(14325,20,15599,'2022-01-30 22:44:53',NULL),(14326,14,15536,'2022-01-30 22:44:53',NULL),(14327,17,15550,'2022-01-30 22:44:53',NULL),(14328,20,15593,'2022-01-30 22:44:53',NULL),(14329,17,15536,'2022-01-30 22:44:54',NULL),(14330,15,15604,'2022-01-30 22:44:54',NULL),(14331,17,15592,'2022-01-30 22:44:54',NULL),(14332,15,15607,'2022-01-30 22:44:54',NULL),(14333,21,15606,'2022-01-30 22:44:54',NULL),(14334,21,15583,'2022-01-30 22:44:54',NULL),(14335,15,15530,'2022-01-30 22:44:54',NULL),(14336,18,15583,'2022-01-30 22:44:54',NULL),(14337,21,15599,'2022-01-30 22:44:54',NULL),(14338,18,15612,'2022-01-30 22:44:54',NULL),(14339,18,15614,'2022-01-30 22:44:54',NULL),(14340,21,15613,'2022-01-30 22:44:54',NULL);
/*!40000 ALTER TABLE `product_ingredients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menu_id` int NOT NULL,
  `name` varchar(50) NOT NULL,
  `price` float NOT NULL,
  `image` varchar(1000) DEFAULT NULL,
  `type` varchar(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `products_id_uindex` (`id`),
  KEY `products_menus_id_fk` (`menu_id`),
  CONSTRAINT `products_menus_id_fk` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3406 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,1258,'Pizza Pepperoni with tomatoes',6.07,'https://roll-club.kh.ua/wp-content/uploads/2021/04/okean-1.jpg.webp','pizza','2022-01-30 22:44:53',NULL),(2,1258,'Pizza Texas',2.5,'https://roll-club.kh.ua/wp-content/uploads/2021/04/4-mjasa-1.jpg.webp','pizza','2022-01-30 22:44:53',NULL),(3,1258,'Pizza Hawaiian',7.85,'https://roll-club.kh.ua/wp-content/uploads/2021/04/rostbif-v-tunce-1.jpg.webp','pizza','2022-01-30 22:44:53',NULL),(4,1259,'Original burger',2.65,'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTvzg-HoIDaAED4WcwA6F-AX_tFxYMsNMzMjQ&usqp=CAU','burger','2022-01-30 22:44:53','2022-03-30 01:39:12'),(5,1259,'Crispy Chicken burger',3.83,'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTUe86lXj9CqmvO-eqVIdlKIDwBVORPBS48gA&usqp=CAU','burger','2022-01-30 22:44:53','2022-03-30 01:25:45'),(6,1259,'Bacon Cheese Burger',6.84,'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRXFU-fBgF98gLOJ-HC1NvHn9w5VkA7B8i2jQ&usqp=CAU','burger','2022-01-30 22:44:53','2022-03-30 01:33:59'),(7,1260,'Pizza Ocean',7.21,'https://roll-club.kh.ua/wp-content/uploads/2019/03/kapricheza.jpg.webp','pizza','2022-01-30 22:44:53',NULL),(8,1260,'Pizza Florida',7.97,'https://roll-club.kh.ua/wp-content/uploads/2015/09/4-syra.jpg.webp','pizza','2022-01-30 22:44:53',NULL),(9,1260,'Pizza Italiano',5.96,'https://roll-club.kh.ua/wp-content/uploads/2014/08/ukrainskaja.jpg.webp','pizza','2022-01-30 22:44:53',NULL),(10,1261,'Philadelphia with salmon',7.08,'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT1VuGxfGkAwZfR0CoVmvQvoZrl7fH3kJDC3w&usqp=CAU','sushi','2022-01-30 22:44:53','2022-03-30 01:42:36'),(11,1261,'Unagi Philadelphia',7.25,'https://roll-club.kh.ua/wp-content/uploads/2021/12/unagi-lajt.jpg','sushi','2022-01-30 22:44:53','2022-03-30 01:43:55'),(12,1261,'Himawari',7.09,'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRTL5TMUqNbYx4EKb-byk2fFAIYnlZxQ1dDBw&usqp=CAU','sushi','2022-01-30 22:44:53','2022-03-30 01:42:36'),(13,1262,'Swedish Meatballs',3.87,'https://target.scene7.com/is/image/Target/GUEST_9066181d-e570-4eb1-b85d-1c8a547857fe?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-30 22:44:53',NULL),(14,1262,'Unwrapped Burrito Bowl',4.9,'https://target.scene7.com/is/image/Target/GUEST_7c640617-96cd-4236-96f0-38a995828861?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-30 22:44:53',NULL),(15,1262,'Tortellini Bowls',3.84,'https://target.scene7.com/is/image/Target/GUEST_dd195641-5a34-4be1-9e14-c4632ee5942e?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-30 22:44:53',NULL),(16,1263,'Chicken Pot Pie',5.63,'https://i5.walmartimages.com/asr/64dc7c05-7d2c-4e29-854d-bada644f019e_1.f4de24032b1a5f48063903f488213f99.jpeg','frozen_meal','2022-01-30 22:44:53',NULL),(17,1263,'Beef & Bean Green Chili Burritos',7.63,'https://i5.walmartimages.com/asr/94fb73ff-59ce-48c2-acd3-42bbb3c52e67.56499b84aa3d79c7a9c51f29da5c3626.png','appetizer','2022-01-30 22:44:53',NULL),(18,1263,'Peanut Butter & Grape Jelly Sandwich',2.59,'https://i5.walmartimages.com/asr/f490a06c-4ea2-41f5-9151-a05449e93e43.64a407ec5ca286d97dd37cdede9f51fc.jpeg','appetizer','2022-01-30 22:44:54',NULL),(19,1264,'Angel Food Cake',7.45,'https://i5.walmartimages.com/asr/7be23ae2-0733-4fe1-a13c-13c1121db61a.37700bee4e0c25b6a3f93496ae17c7c7.jpeg','dessert','2022-01-30 22:44:53',NULL),(20,1264,'New York Style Cheesecake',5.08,'https://i5.walmartimages.com/asr/7df68f47-78dc-4308-8a46-2ca7e1fa50e6.b2973bca56a71b9dc3e773ee08aeb1a6.jpeg','dessert','2022-01-30 22:44:53',NULL),(21,1264,'Sweet Potato Pie',7.24,'https://i5.walmartimages.com/asr/7bd145de-7975-4c04-842e-188a5a0be79d_2.57a9e072cd4d529d16ccd9125f6a035d.jpeg','pastry','2022-01-30 22:44:54',NULL);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `suppliers`
--

DROP TABLE IF EXISTS `suppliers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `suppliers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `type` varchar(20) DEFAULT NULL,
  `image` varchar(500) DEFAULT NULL,
  `opening` varchar(10) DEFAULT NULL,
  `closing` varchar(10) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `suppliers_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1332 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `suppliers`
--

LOCK TABLES `suppliers` WRITE;
/*!40000 ALTER TABLE `suppliers` DISABLE KEYS */;
INSERT INTO `suppliers` VALUES (1,'Pizza Club','restaurant','https://play-lh.googleusercontent.com/qMewibe3u5Wvq3fBf3Ca3_QItjHCOKeGrOAzVXWxqzgRpMwxYlD5CA6M2M5L78SwNA','09:00','20:00','2022-01-30 22:44:53',NULL),(2,'Burger Club','restaurant','https://static.tildacdn.com/tild3931-3935-4639-a564-383462316536/3.png','10:00','20:00','2022-01-30 22:44:53','2022-03-30 01:17:43'),(3,'Saint Frank Coffee','coffee_shop','http://cdn.shopify.com/s/files/1/1578/1589/files/colorf-01_198x200.png','12:00','18:00','2022-01-30 22:44:53',NULL),(4,'Sushi Space','restaurant','https://image.freepik.com/free-vector/sushi-restaurant-logo_8169-12.jpg','08:00','20:00','2022-01-30 22:44:53',NULL),(5,'Target','supermarket','https://1000logos.net/wp-content/uploads/2021/04/Target-logo.png','00:00','24:00','2022-01-30 22:44:53',NULL),(6,'Wallmart','supermarket','https://cdn.corporate.walmart.com/dims4/WMT/ea03f5e/2147483647/strip/true/crop/855x305+0+0/resize/1578x563!/quality/90/?url=https%3A%2F%2Fcdn.corporate.walmart.com%2F98%2F28%2F342ccbff478ab025592645fafcfc%2Fwalmart-logo.png','00:00','24:00','2022-01-30 22:44:53',NULL),(7,'Linea Caffe','coffee_shop','https://lineacaffe.com/wp-content/themes/lineacaffe/images/linea-logo.svg','07:00','19:00','2022-01-30 22:44:53',NULL);
/*!40000 ALTER TABLE `suppliers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `email` varchar(20) DEFAULT NULL,
  `password_hash` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_uindex` (`id`),
  UNIQUE KEY `user_email_uindex` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=290 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (49,'Name','Email2','Password','2022-01-15 21:30:14',NULL),(261,'as',NULL,NULL,'2022-04-22 20:30:38',NULL),(262,'ab',NULL,NULL,'2022-04-22 20:59:13',NULL),(263,'ruslan',NULL,NULL,'2022-04-22 21:16:40',NULL),(264,'ruslan',NULL,NULL,'2022-04-22 21:20:41',NULL),(265,'name1',NULL,NULL,'2022-04-22 21:23:33',NULL),(266,'name2',NULL,NULL,'2022-04-22 21:27:25',NULL),(267,'name3',NULL,NULL,'2022-04-22 21:29:14',NULL),(268,'asd',NULL,NULL,'2022-04-22 21:33:57',NULL),(269,'test name','testeamail','$2a$10$AYIB0nXxsEKZ3DlvLD5H4eG2ohPaB3jVGxUB4wD0OSyiJ5s8fPjha','2022-04-22 22:07:38',NULL),(270,'ads',NULL,NULL,'2022-04-22 22:13:43',NULL),(271,'adad',NULL,NULL,'2022-04-22 22:30:13',NULL),(272,'adad',NULL,NULL,'2022-04-22 22:33:10',NULL),(273,'adad',NULL,NULL,'2022-04-22 22:34:37',NULL),(274,'asd',NULL,NULL,'2022-04-22 22:35:15',NULL),(275,'dfdsf',NULL,NULL,'2022-04-22 22:36:04',NULL),(276,'namd',NULL,NULL,'2022-04-22 22:40:03',NULL),(277,'fdsfd',NULL,NULL,'2022-04-22 22:53:21',NULL),(278,'ruslan','email','$2a$10$v.4kGlIiWpTynnaH0/xTfeTtyDD0.JLoDymH1eIkPGveK8bBqZXLe','2022-04-22 22:54:03',NULL),(279,'a','a','$2a$10$LvJEAxZPuzuEE.hoF/BpweAR4bz5yGVVnRTU1p0QCVBKWVYRGXU.y','2022-04-22 23:12:58',NULL),(280,'testname','testemail','$2a$10$z4QoV/hOFDDvvxjiUx9SH.cxtQ5kR6fP6oIdkugKyZwd2WWGbMmaO','2022-04-23 18:52:35',NULL),(281,'name','mail','$2a$10$IjmqAeJGglK2D9Lnv0oUQukmsvXZUlmSWdxu8QZpr/hxwC3FrJxfC','2022-04-23 19:15:38',NULL),(282,'name','asdasd','$2a$10$8f7hVteyGY8Zh4bEgmhgIuBhFbyEr3kLYR82d0X.s9kNAf2A37.BS','2022-04-23 19:28:23',NULL),(283,'fsdf',NULL,NULL,'2022-04-23 19:35:24',NULL),(284,'name','mail1','$2a$10$r9HhtMEB.XORSi0RMqJscuqC/7gqyhqkZNGZ2oQgy5FHoHCZe.VmS','2022-04-26 21:55:22',NULL),(285,'test','test','$2a$10$dA5HEjOyoT3NuhiU2sKAeuhWaAoKHlI5BYQNFKnqUSsF.pf42IG1a','2022-04-26 22:01:08',NULL),(286,'test','test1','$2a$10$RuUoGV7LhNoaSN9VuoxqXOC19n9098E7TnEj3vRIldIvDUn2tB4uu','2022-04-26 22:05:54',NULL),(287,'test','test2','$2a$10$XsM8jbr7AqaOn/fyc3L9yO6fqLcFncswka/pVeq78jqdZ2izS5dS6','2022-04-26 22:09:39',NULL),(288,'test','test3','$2a$10$zVoo5PDO9E8YhpxlOXqmYORhb3ISi29A4c0KCJigoqOXuUwLOzv/u','2022-04-26 22:10:46',NULL),(289,'test','test4','$2a$10$ObRy786X8G3gBbGUB8mJ4OT.YwVls166h6YDhqvAMMBN9xwCFvsBa','2022-04-26 22:17:35',NULL);
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

-- Dump completed on 2022-04-28  3:59:45
