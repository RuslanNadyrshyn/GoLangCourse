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
  UNIQUE KEY `ingredients_id_uindex` (`id`),
  UNIQUE KEY `ingredients_name_uindex` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=9224 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ingredients`
--

LOCK TABLES `ingredients` WRITE;
/*!40000 ALTER TABLE `ingredients` DISABLE KEYS */;
INSERT INTO `ingredients` VALUES (9137,'Salmon','2022-01-28 20:26:11',NULL),(9138,'Mozarella','2022-01-28 20:26:11',NULL),(9139,'Cheese','2022-01-28 20:26:11',NULL),(9141,'Nori','2022-01-28 20:26:11',NULL),(9142,'Peperoni','2022-01-28 20:26:11',NULL),(9143,'Caramelized onions','2022-01-28 20:26:11',NULL),(9145,'Rice','2022-01-28 20:26:11',NULL),(9146,'Tomatoes','2022-01-28 20:26:11',NULL),(9149,'Cucumber','2022-01-28 20:26:11',NULL),(9150,'BBQ sauce','2022-01-28 20:26:11',NULL),(9151,'Original sauce','2022-01-28 20:26:11',NULL),(9153,'Cream cheese','2022-01-28 20:26:11',NULL),(9154,'Corn','2022-01-28 20:26:11',NULL),(9156,'Chicken','2022-01-28 20:26:11',NULL),(9158,'Onion','2022-01-28 20:26:11',NULL),(9162,'Mushrooms','2022-01-28 20:26:11',NULL),(9167,'Bavarian sausages','2022-01-28 20:26:11',NULL),(9173,'Unagi sauce','2022-01-28 20:26:11',NULL),(9176,'Cheese cheddar','2022-01-28 20:26:11',NULL),(9179,'Bacon','2022-01-28 20:26:11',NULL),(9181,'Pineapple','2022-01-28 20:26:11',NULL),(9187,'Onion rings','2022-01-28 20:26:11',NULL),(9193,'Pasta','2022-01-28 20:26:11',NULL),(9195,'Flour','2022-01-28 20:26:11',NULL),(9196,'Japanese tamago','2022-01-28 20:26:11',NULL),(9198,'Meat','2022-01-28 20:26:11',NULL),(9199,'Strawberry','2022-01-28 20:26:11',NULL),(9200,'Peas','2022-01-28 20:26:11',NULL),(9201,'Mayo','2022-01-28 20:26:11',NULL),(9202,'Carrots','2022-01-28 20:26:11',NULL),(9203,'Butter','2022-01-28 20:26:11',NULL),(9205,'Chili Pepper','2022-01-28 20:26:11',NULL),(9207,'Cottage Cheese','2022-01-28 20:26:11',NULL),(9210,'Sugar','2022-01-28 20:26:11',NULL),(9212,'Garlic','2022-01-28 20:26:11',NULL),(9215,'Basil','2022-01-28 20:26:11',NULL),(9216,'Potatoes','2022-01-28 20:26:11',NULL),(9220,'Peanuts','2022-01-28 20:26:11',NULL),(9222,'Milk','2022-01-28 20:26:11',NULL),(9223,'Grapes','2022-01-28 20:26:11',NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=694 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menus`
--

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;
INSERT INTO `menus` VALUES (687,817,'2022-01-28 20:26:11',NULL),(688,818,'2022-01-28 20:26:11',NULL),(689,819,'2022-01-28 20:26:11',NULL),(690,820,'2022-01-28 20:26:11',NULL),(691,821,'2022-01-28 20:26:11',NULL),(692,822,'2022-01-28 20:26:11',NULL),(693,823,'2022-01-28 20:26:11',NULL);
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
  CONSTRAINT `product_ingredients_ingredients_id_fk` FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `product_ingredients_products_id_fk` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7434 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_ingredients`
--

LOCK TABLES `product_ingredients` WRITE;
/*!40000 ALTER TABLE `product_ingredients` DISABLE KEYS */;
INSERT INTO `product_ingredients` VALUES (7347,1862,9137,'2022-01-28 20:26:11',NULL),(7348,1863,9138,'2022-01-28 20:26:11',NULL),(7349,1865,9139,'2022-01-28 20:26:11',NULL),(7350,1864,9138,'2022-01-28 20:26:11',NULL),(7351,1862,9141,'2022-01-28 20:26:11',NULL),(7352,1863,9142,'2022-01-28 20:26:11',NULL),(7353,1865,9143,'2022-01-28 20:26:11',NULL),(7354,1864,9142,'2022-01-28 20:26:11',NULL),(7355,1862,9145,'2022-01-28 20:26:11',NULL),(7356,1863,9146,'2022-01-28 20:26:11',NULL),(7357,1864,9146,'2022-01-28 20:26:11',NULL),(7358,1865,9146,'2022-01-28 20:26:11',NULL),(7359,1862,9149,'2022-01-28 20:26:11',NULL),(7360,1863,9150,'2022-01-28 20:26:11',NULL),(7361,1864,9150,'2022-01-28 20:26:11',NULL),(7362,1865,9151,'2022-01-28 20:26:11',NULL),(7363,1862,9153,'2022-01-28 20:26:11',NULL),(7364,1866,9154,'2022-01-28 20:26:11',NULL),(7365,1868,9154,'2022-01-28 20:26:11',NULL),(7366,1869,9137,'2022-01-28 20:26:11',NULL),(7367,1867,9156,'2022-01-28 20:26:11',NULL),(7368,1866,9158,'2022-01-28 20:26:11',NULL),(7369,1869,9141,'2022-01-28 20:26:11',NULL),(7370,1868,9158,'2022-01-28 20:26:11',NULL),(7371,1867,9146,'2022-01-28 20:26:11',NULL),(7372,1869,9145,'2022-01-28 20:26:11',NULL),(7373,1866,9162,'2022-01-28 20:26:11',NULL),(7374,1867,9162,'2022-01-28 20:26:11',NULL),(7375,1868,9162,'2022-01-28 20:26:11',NULL),(7376,1869,9149,'2022-01-28 20:26:11',NULL),(7377,1866,9167,'2022-01-28 20:26:11',NULL),(7378,1867,9158,'2022-01-28 20:26:11',NULL),(7379,1869,9153,'2022-01-28 20:26:11',NULL),(7380,1868,9167,'2022-01-28 20:26:11',NULL),(7381,1866,9138,'2022-01-28 20:26:11',NULL),(7382,1867,9150,'2022-01-28 20:26:11',NULL),(7383,1869,9173,'2022-01-28 20:26:11',NULL),(7384,1866,9150,'2022-01-28 20:26:11',NULL),(7385,1868,9138,'2022-01-28 20:26:11',NULL),(7386,1870,9176,'2022-01-28 20:26:11',NULL),(7387,1868,9150,'2022-01-28 20:26:11',NULL),(7388,1871,9156,'2022-01-28 20:26:11',NULL),(7389,1872,9137,'2022-01-28 20:26:11',NULL),(7390,1870,9179,'2022-01-28 20:26:11',NULL),(7391,1873,9156,'2022-01-28 20:26:11',NULL),(7392,1872,9141,'2022-01-28 20:26:11',NULL),(7393,1871,9181,'2022-01-28 20:26:11',NULL),(7394,1870,9151,'2022-01-28 20:26:11',NULL),(7395,1873,9181,'2022-01-28 20:26:11',NULL),(7396,1872,9145,'2022-01-28 20:26:11',NULL),(7397,1871,9138,'2022-01-28 20:26:11',NULL),(7398,1870,9187,'2022-01-28 20:26:11',NULL),(7399,1873,9138,'2022-01-28 20:26:11',NULL),(7400,1872,9149,'2022-01-28 20:26:11',NULL),(7401,1872,9153,'2022-01-28 20:26:11',NULL),(7402,1872,9173,'2022-01-28 20:26:11',NULL),(7403,1874,9193,'2022-01-28 20:26:11',NULL),(7404,1875,9156,'2022-01-28 20:26:11',NULL),(7405,1876,9195,'2022-01-28 20:26:11',NULL),(7406,1875,9195,'2022-01-28 20:26:11',NULL),(7407,1872,9196,'2022-01-28 20:26:11',NULL),(7408,1874,9198,'2022-01-28 20:26:11',NULL),(7409,1875,9200,'2022-01-28 20:26:11',NULL),(7410,1876,9199,'2022-01-28 20:26:11',NULL),(7411,1874,9201,'2022-01-28 20:26:11',NULL),(7412,1875,9202,'2022-01-28 20:26:11',NULL),(7413,1876,9203,'2022-01-28 20:26:11',NULL),(7414,1877,9154,'2022-01-28 20:26:11',NULL),(7415,1877,9158,'2022-01-28 20:26:11',NULL),(7416,1878,9205,'2022-01-28 20:26:11',NULL),(7417,1879,9207,'2022-01-28 20:26:11',NULL),(7418,1878,9158,'2022-01-28 20:26:11',NULL),(7419,1877,9145,'2022-01-28 20:26:11',NULL),(7420,1879,9210,'2022-01-28 20:26:11',NULL),(7421,1878,9145,'2022-01-28 20:26:11',NULL),(7422,1879,9203,'2022-01-28 20:26:11',NULL),(7423,1880,9212,'2022-01-28 20:26:11',NULL),(7424,1878,9201,'2022-01-28 20:26:11',NULL),(7425,1880,9215,'2022-01-28 20:26:11',NULL),(7426,1881,9216,'2022-01-28 20:26:11',NULL),(7427,1882,9195,'2022-01-28 20:26:11',NULL),(7428,1881,9195,'2022-01-28 20:26:11',NULL),(7429,1880,9138,'2022-01-28 20:26:11',NULL),(7430,1881,9210,'2022-01-28 20:26:11',NULL),(7431,1882,9220,'2022-01-28 20:26:11',NULL),(7432,1881,9222,'2022-01-28 20:26:11',NULL),(7433,1882,9223,'2022-01-28 20:26:11',NULL);
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
  `image` varchar(500) DEFAULT NULL,
  `type` varchar(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `products_id_uindex` (`id`),
  KEY `products_menus_id_fk` (`menu_id`),
  CONSTRAINT `products_menus_id_fk` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1883 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1862,687,'Philadelphia with salmon',7.08,'https://ninjasushi.com.ua/img/452.png','sushi','2022-01-28 20:26:11',NULL),(1863,688,'Pizza Pepperoni with tomatoes',6.07,'https://roll-club.kh.ua/wp-content/uploads/2021/04/okean-1.jpg.webp','pizza','2022-01-28 20:26:11',NULL),(1864,689,'Pizza Ocean',7.21,'https://roll-club.kh.ua/wp-content/uploads/2019/03/kapricheza.jpg.webp','pizza','2022-01-28 20:26:11',NULL),(1865,690,'Original burger',2.65,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/49e75c35-6fff-4a0f-8d6a-959cf4721c74_Medium_.jpg','burger','2022-01-28 20:26:11',NULL),(1866,688,'Pizza Texas',2.5,'https://roll-club.kh.ua/wp-content/uploads/2021/04/4-mjasa-1.jpg.webp','pizza','2022-01-28 20:26:11',NULL),(1867,690,'Crispy Chicken burger',3.83,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/93bce037-709e-41a0-9beb-ab3670663b40_Medium_.jpg','burger','2022-01-28 20:26:11',NULL),(1868,689,'Pizza Florida',7.97,'https://roll-club.kh.ua/wp-content/uploads/2015/09/4-syra.jpg.webp','pizza','2022-01-28 20:26:11',NULL),(1869,687,'Unagi Philadelphia',7.25,'https://ninjasushi.com.ua/img/102-600x600.png','sushi','2022-01-28 20:26:11',NULL),(1870,690,'Bacon Cheese Burger',6.84,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/ba3a1ca2-2ec6-4cf4-a843-f7f0fc0f175a_Medium_.jpg','burger','2022-01-28 20:26:11',NULL),(1871,688,'Pizza Hawaiian',7.85,'https://roll-club.kh.ua/wp-content/uploads/2021/04/rostbif-v-tunce-1.jpg.webp','pizza','2022-01-28 20:26:11',NULL),(1872,687,'Himawari',7.09,'https://roll-club.kh.ua/wp-content/uploads/2014/08/ukrainskaja.jpg.webp','sushi','2022-01-28 20:26:11',NULL),(1873,689,'Pizza Italiano',5.96,'https://roll-club.kh.ua/wp-content/uploads/2014/08/ukrainskaja.jpg.webp','pizza','2022-01-28 20:26:11',NULL),(1874,691,'Swedish Meatballs',3.87,'https://target.scene7.com/is/image/Target/GUEST_9066181d-e570-4eb1-b85d-1c8a547857fe?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-28 20:26:11',NULL),(1875,692,'Chicken Pot Pie',5.63,'https://i5.walmartimages.com/asr/64dc7c05-7d2c-4e29-854d-bada644f019e_1.f4de24032b1a5f48063903f488213f99.jpeg','frozen_meal','2022-01-28 20:26:11',NULL),(1876,693,'Angel Food Cake',7.45,'https://i5.walmartimages.com/asr/7be23ae2-0733-4fe1-a13c-13c1121db61a.37700bee4e0c25b6a3f93496ae17c7c7.jpeg','dessert','2022-01-28 20:26:11',NULL),(1877,691,'Unwrapped Burrito Bowl',4.9,'https://target.scene7.com/is/image/Target/GUEST_7c640617-96cd-4236-96f0-38a995828861?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-28 20:26:11',NULL),(1878,692,'Beef & Bean Green Chili Burritos',7.63,'https://i5.walmartimages.com/asr/94fb73ff-59ce-48c2-acd3-42bbb3c52e67.56499b84aa3d79c7a9c51f29da5c3626.png','appetizer','2022-01-28 20:26:11',NULL),(1879,693,'New York Style Cheesecake',5.08,'https://i5.walmartimages.com/asr/7df68f47-78dc-4308-8a46-2ca7e1fa50e6.b2973bca56a71b9dc3e773ee08aeb1a6.jpeg','dessert','2022-01-28 20:26:11',NULL),(1880,691,'Tortellini Bowls',3.84,'https://target.scene7.com/is/image/Target/GUEST_dd195641-5a34-4be1-9e14-c4632ee5942e?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-28 20:26:11',NULL),(1881,693,'Sweet Potato Pie',7.24,'https://i5.walmartimages.com/asr/7bd145de-7975-4c04-842e-188a5a0be79d_2.57a9e072cd4d529d16ccd9125f6a035d.jpeg','pastry','2022-01-28 20:26:11',NULL),(1882,692,'Peanut Butter & Grape Jelly Sandwich',2.59,'https://i5.walmartimages.com/asr/f490a06c-4ea2-41f5-9151-a05449e93e43.64a407ec5ca286d97dd37cdede9f51fc.jpeg','appetizer','2022-01-28 20:26:11',NULL);
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
  `address` varchar(20) DEFAULT NULL,
  `image` varchar(500) DEFAULT NULL,
  `opening` varchar(10) DEFAULT NULL,
  `closing` varchar(10) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `suppliers_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=824 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `suppliers`
--

LOCK TABLES `suppliers` WRITE;
/*!40000 ALTER TABLE `suppliers` DISABLE KEYS */;
INSERT INTO `suppliers` VALUES (817,'Sushi Space','restaurant','','https://image.freepik.com/free-vector/sushi-restaurant-logo_8169-12.jpg','08:00','20:00','2022-01-28 20:26:11',NULL),(818,'Pizza Club','restaurant','','https://play-lh.googleusercontent.com/qMewibe3u5Wvq3fBf3Ca3_QItjHCOKeGrOAzVXWxqzgRpMwxYlD5CA6M2M5L78SwNA','09:00','20:00','2022-01-28 20:26:11',NULL),(819,'Saint Frank Coffee','coffee_shop','','http://cdn.shopify.com/s/files/1/1578/1589/files/colorf-01_198x200.png','12:00','18:00','2022-01-28 20:26:11',NULL),(820,'Burger Club','restaurant','','https://eda.ua/images/506509-195-195-burger_club_harkov.jpg','10:00','20:00','2022-01-28 20:26:11',NULL),(821,'Target','supermarket','','https://1000logos.net/wp-content/uploads/2021/04/Target-logo.png','00:00','24:00','2022-01-28 20:26:11',NULL),(822,'Wallmart','supermarket','','https://cdn.corporate.walmart.com/dims4/WMT/ea03f5e/2147483647/strip/true/crop/855x305+0+0/resize/1578x563!/quality/90/?url=https%3A%2F%2Fcdn.corporate.walmart.com%2F98%2F28%2F342ccbff478ab025592645fafcfc%2Fwalmart-logo.png','00:00','24:00','2022-01-28 20:26:11',NULL),(823,'Linea Caffe','coffee_shop','','https://lineacaffe.com/wp-content/themes/lineacaffe/images/linea-logo.svg','07:00','19:00','2022-01-28 20:26:11',NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=261 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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

-- Dump completed on 2022-01-28 23:32:08
