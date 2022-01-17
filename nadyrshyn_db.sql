-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: localhost    Database: nadyrshyn_db
-- ------------------------------------------------------
-- Server version	8.0.27

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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
  UNIQUE KEY `ingredients_name_uindex` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1954 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ingredients`
--

LOCK TABLES `ingredients` WRITE;
/*!40000 ALTER TABLE `ingredients` DISABLE KEYS */;
INSERT INTO `ingredients` VALUES (1867,'Mozarella','2022-01-17 19:42:08',NULL),(1868,'Peperoni','2022-01-17 19:42:08',NULL),(1869,'Tomatoes','2022-01-17 19:42:08',NULL),(1870,'BBQ sauce','2022-01-17 19:42:08',NULL),(1871,'Corn','2022-01-17 19:42:08',NULL),(1872,'Onion','2022-01-17 19:42:08',NULL),(1873,'Mushrooms','2022-01-17 19:42:08',NULL),(1874,'Bavarian sausages','2022-01-17 19:42:08',NULL),(1877,'Chicken','2022-01-17 19:42:08',NULL),(1878,'Pineapple','2022-01-17 19:42:08',NULL),(1880,'Cheese','2022-01-17 19:42:08',NULL),(1881,'Caramelized onions','2022-01-17 19:42:08',NULL),(1883,'Original sauce','2022-01-17 19:42:08',NULL),(1889,'Cheese cheddar','2022-01-17 19:42:08',NULL),(1890,'Bacon','2022-01-17 19:42:08',NULL),(1892,'Onion rings','2022-01-17 19:42:08',NULL),(1906,'Salmon','2022-01-17 19:42:08',NULL),(1907,'Nori','2022-01-17 19:42:08',NULL),(1908,'Rice','2022-01-17 19:42:08',NULL),(1909,'Cucumber','2022-01-17 19:42:08',NULL),(1910,'Cream cheese','2022-01-17 19:42:08',NULL),(1916,'Unagi sauce','2022-01-17 19:42:08',NULL),(1923,'Japanese tamago','2022-01-17 19:42:08',NULL),(1924,'Pasta','2022-01-17 19:42:09',NULL),(1925,'Meat','2022-01-17 19:42:09',NULL),(1926,'Mayo','2022-01-17 19:42:09',NULL),(1930,'Garlic','2022-01-17 19:42:09',NULL),(1931,'Basil','2022-01-17 19:42:09',NULL),(1934,'Flour','2022-01-17 19:42:09',NULL),(1935,'Peas','2022-01-17 19:42:09',NULL),(1936,'Carrots','2022-01-17 19:42:09',NULL),(1937,'Chili Pepper','2022-01-17 19:42:09',NULL),(1942,'Peanuts','2022-01-17 19:42:09',NULL),(1943,'Grapes','2022-01-17 19:42:09',NULL),(1945,'Strawberry','2022-01-17 19:42:09',NULL),(1946,'Butter','2022-01-17 19:42:09',NULL),(1947,'Cottage Cheese','2022-01-17 19:42:09',NULL),(1948,'Sugar','2022-01-17 19:42:09',NULL),(1950,'Potatoes','2022-01-17 19:42:09',NULL),(1953,'Milk','2022-01-17 19:42:09',NULL);
/*!40000 ALTER TABLE `ingredients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menus`
--

DROP TABLE IF EXISTS `menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menus` (
  `id` int NOT NULL,
  `supplier_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `menu_id_uindex` (`id`),
  KEY `menus_suppliers_id_fk` (`supplier_id`),
  CONSTRAINT `menus_suppliers_id_fk` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menus`
--

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;
INSERT INTO `menus` VALUES (1,1,'2022-01-17 19:42:08',NULL),(2,2,'2022-01-17 19:42:08',NULL),(3,3,'2022-01-17 19:42:08',NULL),(4,4,'2022-01-17 19:42:08',NULL),(5,5,'2022-01-17 19:42:09',NULL),(6,6,'2022-01-17 19:42:09',NULL),(7,7,'2022-01-17 19:42:09',NULL);
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
  `price` float DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `orders_products_id_uindex` (`id`),
  KEY `orders_products_orders_id_fk` (`order_id`),
  KEY `order_product_products_id_fk` (`product_id`),
  CONSTRAINT `order_product_products_id_fk` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `orders_products_orders_id_fk` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_product`
--

LOCK TABLES `order_product` WRITE;
/*!40000 ALTER TABLE `order_product` DISABLE KEYS */;
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
  `Address` varchar(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `orders_id_uindex` (`id`),
  KEY `orders_users_id_fk` (`user_id`),
  CONSTRAINT `orders_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
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
  CONSTRAINT `product_ingredients_ingredients_id_fk` FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients` (`id`),
  CONSTRAINT `product_ingredients_products_id_fk` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1707 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_ingredients`
--

LOCK TABLES `product_ingredients` WRITE;
/*!40000 ALTER TABLE `product_ingredients` DISABLE KEYS */;
INSERT INTO `product_ingredients` VALUES (1620,1,1867,'2022-01-17 19:42:08',NULL),(1621,1,1868,'2022-01-17 19:42:08',NULL),(1622,1,1869,'2022-01-17 19:42:08',NULL),(1623,1,1870,'2022-01-17 19:42:08',NULL),(1624,2,1871,'2022-01-17 19:42:08',NULL),(1625,2,1872,'2022-01-17 19:42:08',NULL),(1626,2,1873,'2022-01-17 19:42:08',NULL),(1627,2,1874,'2022-01-17 19:42:08',NULL),(1628,2,1867,'2022-01-17 19:42:08',NULL),(1629,2,1870,'2022-01-17 19:42:08',NULL),(1630,3,1877,'2022-01-17 19:42:08',NULL),(1631,3,1878,'2022-01-17 19:42:08',NULL),(1632,3,1867,'2022-01-17 19:42:08',NULL),(1633,4,1880,'2022-01-17 19:42:08',NULL),(1634,4,1881,'2022-01-17 19:42:08',NULL),(1635,4,1869,'2022-01-17 19:42:08',NULL),(1636,4,1883,'2022-01-17 19:42:08',NULL),(1637,5,1877,'2022-01-17 19:42:08',NULL),(1638,5,1869,'2022-01-17 19:42:08',NULL),(1639,5,1873,'2022-01-17 19:42:08',NULL),(1640,5,1872,'2022-01-17 19:42:08',NULL),(1641,5,1870,'2022-01-17 19:42:08',NULL),(1642,6,1889,'2022-01-17 19:42:08',NULL),(1643,6,1890,'2022-01-17 19:42:08',NULL),(1644,6,1883,'2022-01-17 19:42:08',NULL),(1645,6,1892,'2022-01-17 19:42:08',NULL),(1646,7,1867,'2022-01-17 19:42:08',NULL),(1647,7,1868,'2022-01-17 19:42:08',NULL),(1648,7,1869,'2022-01-17 19:42:08',NULL),(1649,7,1870,'2022-01-17 19:42:08',NULL),(1650,8,1871,'2022-01-17 19:42:08',NULL),(1651,8,1872,'2022-01-17 19:42:08',NULL),(1652,8,1873,'2022-01-17 19:42:08',NULL),(1653,8,1874,'2022-01-17 19:42:08',NULL),(1654,8,1867,'2022-01-17 19:42:08',NULL),(1655,8,1870,'2022-01-17 19:42:08',NULL),(1656,9,1877,'2022-01-17 19:42:08',NULL),(1657,9,1878,'2022-01-17 19:42:08',NULL),(1658,9,1867,'2022-01-17 19:42:08',NULL),(1659,10,1906,'2022-01-17 19:42:08',NULL),(1660,10,1907,'2022-01-17 19:42:08',NULL),(1661,10,1908,'2022-01-17 19:42:08',NULL),(1662,10,1909,'2022-01-17 19:42:08',NULL),(1663,10,1910,'2022-01-17 19:42:08',NULL),(1664,11,1906,'2022-01-17 19:42:08',NULL),(1665,11,1907,'2022-01-17 19:42:08',NULL),(1666,11,1908,'2022-01-17 19:42:08',NULL),(1667,11,1909,'2022-01-17 19:42:08',NULL),(1668,11,1910,'2022-01-17 19:42:08',NULL),(1669,11,1916,'2022-01-17 19:42:08',NULL),(1670,12,1906,'2022-01-17 19:42:08',NULL),(1671,12,1907,'2022-01-17 19:42:08',NULL),(1672,12,1908,'2022-01-17 19:42:08',NULL),(1673,12,1909,'2022-01-17 19:42:08',NULL),(1674,12,1910,'2022-01-17 19:42:08',NULL),(1675,12,1916,'2022-01-17 19:42:08',NULL),(1676,12,1923,'2022-01-17 19:42:08',NULL),(1677,13,1924,'2022-01-17 19:42:09',NULL),(1678,13,1925,'2022-01-17 19:42:09',NULL),(1679,13,1926,'2022-01-17 19:42:09',NULL),(1680,14,1871,'2022-01-17 19:42:09',NULL),(1681,14,1872,'2022-01-17 19:42:09',NULL),(1682,14,1908,'2022-01-17 19:42:09',NULL),(1683,15,1930,'2022-01-17 19:42:09',NULL),(1684,15,1931,'2022-01-17 19:42:09',NULL),(1685,15,1867,'2022-01-17 19:42:09',NULL),(1686,16,1877,'2022-01-17 19:42:09',NULL),(1687,16,1934,'2022-01-17 19:42:09',NULL),(1688,16,1935,'2022-01-17 19:42:09',NULL),(1689,16,1936,'2022-01-17 19:42:09',NULL),(1690,17,1937,'2022-01-17 19:42:09',NULL),(1691,17,1872,'2022-01-17 19:42:09',NULL),(1692,17,1908,'2022-01-17 19:42:09',NULL),(1693,17,1926,'2022-01-17 19:42:09',NULL),(1694,18,1934,'2022-01-17 19:42:09',NULL),(1695,18,1942,'2022-01-17 19:42:09',NULL),(1696,18,1943,'2022-01-17 19:42:09',NULL),(1697,19,1934,'2022-01-17 19:42:09',NULL),(1698,19,1945,'2022-01-17 19:42:09',NULL),(1699,19,1946,'2022-01-17 19:42:09',NULL),(1700,20,1947,'2022-01-17 19:42:09',NULL),(1701,20,1948,'2022-01-17 19:42:09',NULL),(1702,20,1946,'2022-01-17 19:42:09',NULL),(1703,21,1950,'2022-01-17 19:42:09',NULL),(1704,21,1934,'2022-01-17 19:42:09',NULL),(1705,21,1948,'2022-01-17 19:42:09',NULL),(1706,21,1953,'2022-01-17 19:42:09',NULL);
/*!40000 ALTER TABLE `product_ingredients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` int NOT NULL,
  `menu_id` int NOT NULL,
  `name` varchar(50) NOT NULL,
  `price` float NOT NULL,
  `image` varchar(500) DEFAULT NULL,
  `type` varchar(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `products_id_uindex` (`id`),
  KEY `products_menus_id_fk` (`menu_id`),
  CONSTRAINT `products_menus_id_fk` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,1,'Pizza Pepperoni with tomatoes',6.07,'https://roll-club.kh.ua/wp-content/uploads/2021/04/okean-1.jpg.webp','pizza','2022-01-17 19:42:08',NULL),(2,1,'Pizza Texas',2.5,'https://roll-club.kh.ua/wp-content/uploads/2021/04/4-mjasa-1.jpg.webp','pizza','2022-01-17 19:42:08',NULL),(3,1,'Pizza Hawaiian',7.85,'https://roll-club.kh.ua/wp-content/uploads/2021/04/rostbif-v-tunce-1.jpg.webp','pizza','2022-01-17 19:42:08',NULL),(4,2,'Original burger',2.65,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/49e75c35-6fff-4a0f-8d6a-959cf4721c74_Medium_.jpg','burger','2022-01-17 19:42:08',NULL),(5,2,'Crispy Chicken burger',3.83,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/93bce037-709e-41a0-9beb-ab3670663b40_Medium_.jpg','burger','2022-01-17 19:42:08',NULL),(6,2,'Bacon Cheese Burger',6.84,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/ba3a1ca2-2ec6-4cf4-a843-f7f0fc0f175a_Medium_.jpg','burger','2022-01-17 19:42:08',NULL),(7,3,'Pizza Ocean',7.21,'https://roll-club.kh.ua/wp-content/uploads/2019/03/kapricheza.jpg.webp','pizza','2022-01-17 19:42:08',NULL),(8,3,'Pizza Florida',7.97,'https://roll-club.kh.ua/wp-content/uploads/2015/09/4-syra.jpg.webp','pizza','2022-01-17 19:42:08',NULL),(9,3,'Pizza Italiano',5.96,'https://roll-club.kh.ua/wp-content/uploads/2014/08/ukrainskaja.jpg.webp','pizza','2022-01-17 19:42:08',NULL),(10,4,'Philadelphia with salmon',7.08,'https://ninjasushi.com.ua/img/452.png','sushi','2022-01-17 19:42:08',NULL),(11,4,'Unagi Philadelphia',7.25,'https://ninjasushi.com.ua/img/102-600x600.png','sushi','2022-01-17 19:42:08',NULL),(12,4,'Himawari',7.09,'https://roll-club.kh.ua/wp-content/uploads/2014/08/ukrainskaja.jpg.webp','sushi','2022-01-17 19:42:08',NULL),(13,5,'Swedish Meatballs',3.87,'https://target.scene7.com/is/image/Target/GUEST_9066181d-e570-4eb1-b85d-1c8a547857fe?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-17 19:42:09',NULL),(14,5,'Unwrapped Burrito Bowl',4.9,'https://target.scene7.com/is/image/Target/GUEST_7c640617-96cd-4236-96f0-38a995828861?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-17 19:42:09',NULL),(15,5,'Tortellini Bowls',3.84,'https://target.scene7.com/is/image/Target/GUEST_dd195641-5a34-4be1-9e14-c4632ee5942e?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-17 19:42:09',NULL),(16,6,'Chicken Pot Pie',5.63,'https://i5.walmartimages.com/asr/64dc7c05-7d2c-4e29-854d-bada644f019e_1.f4de24032b1a5f48063903f488213f99.jpeg','frozen_meal','2022-01-17 19:42:09',NULL),(17,6,'Beef & Bean Green Chili Burritos',7.63,'https://i5.walmartimages.com/asr/94fb73ff-59ce-48c2-acd3-42bbb3c52e67.56499b84aa3d79c7a9c51f29da5c3626.png','appetizer','2022-01-17 19:42:09',NULL),(18,6,'Peanut Butter & Grape Jelly Sandwich',2.59,'https://i5.walmartimages.com/asr/f490a06c-4ea2-41f5-9151-a05449e93e43.64a407ec5ca286d97dd37cdede9f51fc.jpeg','appetizer','2022-01-17 19:42:09',NULL),(19,7,'Angel Food Cake',7.45,'https://i5.walmartimages.com/asr/7be23ae2-0733-4fe1-a13c-13c1121db61a.37700bee4e0c25b6a3f93496ae17c7c7.jpeg','dessert','2022-01-17 19:42:09',NULL),(20,7,'New York Style Cheesecake',5.08,'https://i5.walmartimages.com/asr/7df68f47-78dc-4308-8a46-2ca7e1fa50e6.b2973bca56a71b9dc3e773ee08aeb1a6.jpeg','dessert','2022-01-17 19:42:09',NULL),(21,7,'Sweet Potato Pie',7.24,'https://i5.walmartimages.com/asr/7bd145de-7975-4c04-842e-188a5a0be79d_2.57a9e072cd4d529d16ccd9125f6a035d.jpeg','pastry','2022-01-17 19:42:09',NULL);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `suppliers`
--

DROP TABLE IF EXISTS `suppliers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `suppliers` (
  `id` int NOT NULL,
  `name` varchar(20) NOT NULL,
  `type` varchar(20) DEFAULT NULL,
  `address` varchar(20) DEFAULT NULL,
  `image` varchar(500) DEFAULT NULL,
  `opening` varchar(10) DEFAULT NULL,
  `closing` varchar(10) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `suppliers`
--

LOCK TABLES `suppliers` WRITE;
/*!40000 ALTER TABLE `suppliers` DISABLE KEYS */;
INSERT INTO `suppliers` VALUES (1,'Pizza Club','restaurant','','https://play-lh.googleusercontent.com/qMewibe3u5Wvq3fBf3Ca3_QItjHCOKeGrOAzVXWxqzgRpMwxYlD5CA6M2M5L78SwNA','09:00','20:00','2022-01-17 19:42:08',NULL),(2,'Burger Club','restaurant','','https://eda.ua/images/506509-195-195-burger_club_harkov.jpg','10:00','20:00','2022-01-17 19:42:08',NULL),(3,'Saint Frank Coffee','coffee_shop','','http://cdn.shopify.com/s/files/1/1578/1589/files/colorf-01_198x200.png','12:00','18:00','2022-01-17 19:42:08',NULL),(4,'Sushi Space','restaurant','','https://image.freepik.com/free-vector/sushi-restaurant-logo_8169-12.jpg','08:00','20:00','2022-01-17 19:42:08',NULL),(5,'Target','supermarket','','https://1000logos.net/wp-content/uploads/2021/04/Target-logo.png','00:00','24:00','2022-01-17 19:42:09',NULL),(6,'Wallmart','supermarket','','https://cdn.corporate.walmart.com/dims4/WMT/ea03f5e/2147483647/strip/true/crop/855x305+0+0/resize/1578x563!/quality/90/?url=https%3A%2F%2Fcdn.corporate.walmart.com%2F98%2F28%2F342ccbff478ab025592645fafcfc%2Fwalmart-logo.png','00:00','24:00','2022-01-17 19:42:09',NULL),(7,'Linea Caffe','coffee_shop','','https://lineacaffe.com/wp-content/themes/lineacaffe/images/linea-logo.svg','07:00','19:00','2022-01-17 19:42:09',NULL);
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
  `name` varchar(10) DEFAULT NULL,
  `email` varchar(20) DEFAULT NULL,
  `password_hash` varchar(30) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_uindex` (`id`),
  UNIQUE KEY `user_email_uindex` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=259 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (49,'Name','Email2','Password','2022-01-15 21:30:14',NULL);
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

-- Dump completed on 2022-01-17 21:43:49
