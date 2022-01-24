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
) ENGINE=InnoDB AUTO_INCREMENT=3306 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ingredients`
--

LOCK TABLES `ingredients` WRITE;
/*!40000 ALTER TABLE `ingredients` DISABLE KEYS */;
INSERT INTO `ingredients` VALUES (3132,'Mozarella','2022-01-24 15:58:53',NULL),(3133,'Salmon','2022-01-24 15:58:53',NULL),(3135,'Cheese','2022-01-24 15:58:53',NULL),(3136,'Peperoni','2022-01-24 15:58:53',NULL),(3137,'Nori','2022-01-24 15:58:53',NULL),(3138,'Caramelized onions','2022-01-24 15:58:53',NULL),(3140,'Tomatoes','2022-01-24 15:58:53',NULL),(3141,'Rice','2022-01-24 15:58:53',NULL),(3144,'BBQ sauce','2022-01-24 15:58:53',NULL),(3145,'Original sauce','2022-01-24 15:58:53',NULL),(3147,'Cucumber','2022-01-24 15:58:53',NULL),(3148,'Corn','2022-01-24 15:58:53',NULL),(3149,'Chicken','2022-01-24 15:58:53',NULL),(3151,'Onion','2022-01-24 15:58:53',NULL),(3152,'Cream cheese','2022-01-24 15:58:53',NULL),(3155,'Mushrooms','2022-01-24 15:58:53',NULL),(3161,'Bavarian sausages','2022-01-24 15:58:53',NULL),(3166,'Cheese cheddar','2022-01-24 15:58:53',NULL),(3169,'Unagi sauce','2022-01-24 15:58:53',NULL),(3175,'Bacon','2022-01-24 15:58:53',NULL),(3176,'Pineapple','2022-01-24 15:58:53',NULL),(3184,'Onion rings','2022-01-24 15:58:53',NULL),(3187,'Pasta','2022-01-24 15:58:53',NULL),(3190,'Flour','2022-01-24 15:58:53',NULL),(3191,'Japanese tamago','2022-01-24 15:58:53',NULL),(3192,'Meat','2022-01-24 15:58:53',NULL),(3193,'Peas','2022-01-24 15:58:53',NULL),(3195,'Strawberry','2022-01-24 15:58:53',NULL),(3196,'Mayo','2022-01-24 15:58:53',NULL),(3197,'Carrots','2022-01-24 15:58:53',NULL),(3198,'Butter','2022-01-24 15:58:53',NULL),(3200,'Chili Pepper','2022-01-24 15:58:53',NULL),(3202,'Cottage Cheese','2022-01-24 15:58:53',NULL),(3204,'Garlic','2022-01-24 15:58:53',NULL),(3207,'Sugar','2022-01-24 15:58:53',NULL),(3210,'Basil','2022-01-24 15:58:53',NULL),(3211,'Peanuts','2022-01-24 15:58:53',NULL),(3213,'Potatoes','2022-01-24 15:58:53',NULL),(3215,'Grapes','2022-01-24 15:58:53',NULL),(3218,'Milk','2022-01-24 15:58:53',NULL);
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
  CONSTRAINT `menus_suppliers_id_fk` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=187 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menus`
--

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;
INSERT INTO `menus` VALUES (173,229,'2022-01-24 15:58:53',NULL),(174,230,'2022-01-24 15:58:53',NULL),(175,231,'2022-01-24 15:58:53',NULL),(176,232,'2022-01-24 15:58:53',NULL),(177,233,'2022-01-24 15:58:53',NULL),(178,234,'2022-01-24 15:58:53',NULL),(179,235,'2022-01-24 15:58:53',NULL),(180,236,'2022-01-24 16:03:58',NULL),(181,237,'2022-01-24 16:03:58',NULL),(182,238,'2022-01-24 16:03:58',NULL),(183,239,'2022-01-24 16:03:58',NULL),(184,240,'2022-01-24 16:03:58',NULL),(185,241,'2022-01-24 16:03:58',NULL),(186,242,'2022-01-24 16:03:58',NULL);
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
  KEY `product_ingredients_products_id_fk` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2433 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_ingredients`
--

LOCK TABLES `product_ingredients` WRITE;
/*!40000 ALTER TABLE `product_ingredients` DISABLE KEYS */;
INSERT INTO `product_ingredients` VALUES (2393,421,3132,'2022-01-24 15:58:53',NULL),(2394,422,3133,'2022-01-24 15:58:53',NULL),(2395,424,3135,'2022-01-24 15:58:53',NULL),(2396,423,3136,'2022-01-24 15:58:53',NULL),(2397,422,3137,'2022-01-24 15:58:53',NULL),(2398,424,3138,'2022-01-24 15:58:53',NULL),(2399,423,3140,'2022-01-24 15:58:53',NULL),(2400,422,3141,'2022-01-24 15:58:53',NULL),(2401,424,3145,'2022-01-24 15:58:53',NULL),(2402,423,3144,'2022-01-24 15:58:53',NULL),(2403,422,3147,'2022-01-24 15:58:53',NULL),(2404,425,3148,'2022-01-24 15:58:53',NULL),(2405,427,3149,'2022-01-24 15:58:53',NULL),(2406,422,3152,'2022-01-24 15:58:53',NULL),(2407,426,3151,'2022-01-24 15:58:53',NULL),(2408,425,3155,'2022-01-24 15:58:53',NULL),(2409,426,3161,'2022-01-24 15:58:53',NULL),(2410,429,3166,'2022-01-24 15:58:53',NULL),(2411,428,3169,'2022-01-24 15:58:53',NULL),(2412,430,3176,'2022-01-24 15:58:53',NULL),(2413,429,3175,'2022-01-24 15:58:53',NULL),(2414,429,3184,'2022-01-24 15:58:53',NULL),(2415,433,3187,'2022-01-24 15:58:53',NULL),(2416,434,3190,'2022-01-24 15:58:53',NULL),(2417,432,3191,'2022-01-24 15:58:53',NULL),(2418,433,3192,'2022-01-24 15:58:53',NULL),(2419,434,3193,'2022-01-24 15:58:53',NULL),(2420,435,3195,'2022-01-24 15:58:53',NULL),(2421,433,3196,'2022-01-24 15:58:53',NULL),(2422,434,3197,'2022-01-24 15:58:53',NULL),(2423,435,3198,'2022-01-24 15:58:53',NULL),(2424,437,3200,'2022-01-24 15:58:53',NULL),(2425,438,3202,'2022-01-24 15:58:53',NULL),(2426,439,3204,'2022-01-24 15:58:53',NULL),(2427,438,3207,'2022-01-24 15:58:53',NULL),(2428,439,3210,'2022-01-24 15:58:53',NULL),(2429,440,3211,'2022-01-24 15:58:53',NULL),(2430,441,3213,'2022-01-24 15:58:53',NULL),(2431,440,3215,'2022-01-24 15:58:53',NULL),(2432,441,3218,'2022-01-24 15:58:53',NULL);
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
  CONSTRAINT `products_menus_id_fk` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=463 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (421,173,'Pizza Pepperoni with tomatoes',6.07,'https://roll-club.kh.ua/wp-content/uploads/2021/04/okean-1.jpg.webp','pizza','2022-01-24 15:58:53',NULL),(422,174,'Philadelphia with salmon',7.08,'https://ninjasushi.com.ua/img/452.png','sushi','2022-01-24 15:58:53',NULL),(423,175,'Pizza Ocean',7.21,'https://roll-club.kh.ua/wp-content/uploads/2019/03/kapricheza.jpg.webp','pizza','2022-01-24 15:58:53',NULL),(424,176,'Original burger',2.65,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/49e75c35-6fff-4a0f-8d6a-959cf4721c74_Medium_.jpg','burger','2022-01-24 15:58:53',NULL),(425,173,'Pizza Texas',2.5,'https://roll-club.kh.ua/wp-content/uploads/2021/04/4-mjasa-1.jpg.webp','pizza','2022-01-24 15:58:53',NULL),(426,175,'Pizza Florida',7.97,'https://roll-club.kh.ua/wp-content/uploads/2015/09/4-syra.jpg.webp','pizza','2022-01-24 15:58:53',NULL),(427,176,'Crispy Chicken burger',3.83,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/93bce037-709e-41a0-9beb-ab3670663b40_Medium_.jpg','burger','2022-01-24 15:58:53',NULL),(428,174,'Unagi Philadelphia',7.25,'https://ninjasushi.com.ua/img/102-600x600.png','sushi','2022-01-24 15:58:53',NULL),(429,176,'Bacon Cheese Burger',6.84,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/ba3a1ca2-2ec6-4cf4-a843-f7f0fc0f175a_Medium_.jpg','burger','2022-01-24 15:58:53',NULL),(430,173,'Pizza Hawaiian',7.85,'https://roll-club.kh.ua/wp-content/uploads/2021/04/rostbif-v-tunce-1.jpg.webp','pizza','2022-01-24 15:58:53',NULL),(431,175,'Pizza Italiano',5.96,'https://roll-club.kh.ua/wp-content/uploads/2014/08/ukrainskaja.jpg.webp','pizza','2022-01-24 15:58:53',NULL),(432,174,'Himawari',7.09,'https://roll-club.kh.ua/wp-content/uploads/2014/08/ukrainskaja.jpg.webp','sushi','2022-01-24 15:58:53',NULL),(433,177,'Swedish Meatballs',3.87,'https://target.scene7.com/is/image/Target/GUEST_9066181d-e570-4eb1-b85d-1c8a547857fe?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-24 15:58:53',NULL),(434,178,'Chicken Pot Pie',5.63,'https://i5.walmartimages.com/asr/64dc7c05-7d2c-4e29-854d-bada644f019e_1.f4de24032b1a5f48063903f488213f99.jpeg','frozen_meal','2022-01-24 15:58:53',NULL),(435,179,'Angel Food Cake',7.45,'https://i5.walmartimages.com/asr/7be23ae2-0733-4fe1-a13c-13c1121db61a.37700bee4e0c25b6a3f93496ae17c7c7.jpeg','dessert','2022-01-24 15:58:53',NULL),(436,177,'Unwrapped Burrito Bowl',4.9,'https://target.scene7.com/is/image/Target/GUEST_7c640617-96cd-4236-96f0-38a995828861?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-24 15:58:53',NULL),(437,178,'Beef & Bean Green Chili Burritos',7.63,'https://i5.walmartimages.com/asr/94fb73ff-59ce-48c2-acd3-42bbb3c52e67.56499b84aa3d79c7a9c51f29da5c3626.png','appetizer','2022-01-24 15:58:53',NULL),(438,179,'New York Style Cheesecake',5.08,'https://i5.walmartimages.com/asr/7df68f47-78dc-4308-8a46-2ca7e1fa50e6.b2973bca56a71b9dc3e773ee08aeb1a6.jpeg','dessert','2022-01-24 15:58:53',NULL),(439,177,'Tortellini Bowls',3.84,'https://target.scene7.com/is/image/Target/GUEST_dd195641-5a34-4be1-9e14-c4632ee5942e?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-24 15:58:53',NULL),(440,178,'Peanut Butter & Grape Jelly Sandwich',2.59,'https://i5.walmartimages.com/asr/f490a06c-4ea2-41f5-9151-a05449e93e43.64a407ec5ca286d97dd37cdede9f51fc.jpeg','appetizer','2022-01-24 15:58:53',NULL),(441,179,'Sweet Potato Pie',7.24,'https://i5.walmartimages.com/asr/7bd145de-7975-4c04-842e-188a5a0be79d_2.57a9e072cd4d529d16ccd9125f6a035d.jpeg','pastry','2022-01-24 15:58:53',NULL),(442,180,'Pizza Ocean',7.21,'https://roll-club.kh.ua/wp-content/uploads/2019/03/kapricheza.jpg.webp','pizza','2022-01-24 16:03:58',NULL),(443,181,'Philadelphia with salmon',7.08,'https://ninjasushi.com.ua/img/452.png','sushi','2022-01-24 16:03:58',NULL),(444,183,'Original burger',2.65,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/49e75c35-6fff-4a0f-8d6a-959cf4721c74_Medium_.jpg','burger','2022-01-24 16:03:58',NULL),(445,182,'Pizza Pepperoni with tomatoes',6.07,'https://roll-club.kh.ua/wp-content/uploads/2021/04/okean-1.jpg.webp','pizza','2022-01-24 16:03:58',NULL),(446,183,'Crispy Chicken burger',3.83,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/93bce037-709e-41a0-9beb-ab3670663b40_Medium_.jpg','burger','2022-01-24 16:03:58',NULL),(447,180,'Pizza Florida',7.97,'https://roll-club.kh.ua/wp-content/uploads/2015/09/4-syra.jpg.webp','pizza','2022-01-24 16:03:58',NULL),(448,182,'Pizza Texas',2.5,'https://roll-club.kh.ua/wp-content/uploads/2021/04/4-mjasa-1.jpg.webp','pizza','2022-01-24 16:03:58',NULL),(449,181,'Unagi Philadelphia',7.25,'https://ninjasushi.com.ua/img/102-600x600.png','sushi','2022-01-24 16:03:58',NULL),(450,183,'Bacon Cheese Burger',6.84,'https://316024.selcdn.ru/wiget/4d4c871a-107f-11e7-80df-d8d38565926f/ba3a1ca2-2ec6-4cf4-a843-f7f0fc0f175a_Medium_.jpg','burger','2022-01-24 16:03:58',NULL),(451,181,'Himawari',7.09,'https://roll-club.kh.ua/wp-content/uploads/2014/08/ukrainskaja.jpg.webp','sushi','2022-01-24 16:03:58',NULL),(452,182,'Pizza Hawaiian',7.85,'https://roll-club.kh.ua/wp-content/uploads/2021/04/rostbif-v-tunce-1.jpg.webp','pizza','2022-01-24 16:03:58',NULL),(453,180,'Pizza Italiano',5.96,'https://roll-club.kh.ua/wp-content/uploads/2014/08/ukrainskaja.jpg.webp','pizza','2022-01-24 16:03:58',NULL),(454,184,'Swedish Meatballs',3.87,'https://target.scene7.com/is/image/Target/GUEST_9066181d-e570-4eb1-b85d-1c8a547857fe?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-24 16:03:58',NULL),(455,185,'Angel Food Cake',7.45,'https://i5.walmartimages.com/asr/7be23ae2-0733-4fe1-a13c-13c1121db61a.37700bee4e0c25b6a3f93496ae17c7c7.jpeg','dessert','2022-01-24 16:03:58',NULL),(456,186,'Chicken Pot Pie',5.63,'https://i5.walmartimages.com/asr/64dc7c05-7d2c-4e29-854d-bada644f019e_1.f4de24032b1a5f48063903f488213f99.jpeg','frozen_meal','2022-01-24 16:03:58',NULL),(457,184,'Unwrapped Burrito Bowl',4.9,'https://target.scene7.com/is/image/Target/GUEST_7c640617-96cd-4236-96f0-38a995828861?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-24 16:03:58',NULL),(458,185,'New York Style Cheesecake',5.08,'https://i5.walmartimages.com/asr/7df68f47-78dc-4308-8a46-2ca7e1fa50e6.b2973bca56a71b9dc3e773ee08aeb1a6.jpeg','dessert','2022-01-24 16:03:58',NULL),(459,186,'Beef & Bean Green Chili Burritos',7.63,'https://i5.walmartimages.com/asr/94fb73ff-59ce-48c2-acd3-42bbb3c52e67.56499b84aa3d79c7a9c51f29da5c3626.png','appetizer','2022-01-24 16:03:58',NULL),(460,184,'Tortellini Bowls',3.84,'https://target.scene7.com/is/image/Target/GUEST_dd195641-5a34-4be1-9e14-c4632ee5942e?wid=1416&hei=1416&fmt=webp','frozen_meal','2022-01-24 16:03:58',NULL),(461,185,'Sweet Potato Pie',7.24,'https://i5.walmartimages.com/asr/7bd145de-7975-4c04-842e-188a5a0be79d_2.57a9e072cd4d529d16ccd9125f6a035d.jpeg','pastry','2022-01-24 16:03:58',NULL),(462,186,'Peanut Butter & Grape Jelly Sandwich',2.59,'https://i5.walmartimages.com/asr/f490a06c-4ea2-41f5-9151-a05449e93e43.64a407ec5ca286d97dd37cdede9f51fc.jpeg','appetizer','2022-01-24 16:03:58',NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=243 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `suppliers`
--

LOCK TABLES `suppliers` WRITE;
/*!40000 ALTER TABLE `suppliers` DISABLE KEYS */;
INSERT INTO `suppliers` VALUES (229,'Pizza Club','restaurant','','https://play-lh.googleusercontent.com/qMewibe3u5Wvq3fBf3Ca3_QItjHCOKeGrOAzVXWxqzgRpMwxYlD5CA6M2M5L78SwNA','09:00','20:00','2022-01-24 15:58:53',NULL),(230,'Sushi Space','restaurant','','https://image.freepik.com/free-vector/sushi-restaurant-logo_8169-12.jpg','08:00','20:00','2022-01-24 15:58:53',NULL),(231,'Saint Frank Coffee','coffee_shop','','http://cdn.shopify.com/s/files/1/1578/1589/files/colorf-01_198x200.png','12:00','18:00','2022-01-24 15:58:53',NULL),(232,'Burger Club','restaurant','','https://eda.ua/images/506509-195-195-burger_club_harkov.jpg','10:00','20:00','2022-01-24 15:58:53',NULL),(233,'Target','supermarket','','https://1000logos.net/wp-content/uploads/2021/04/Target-logo.png','00:00','24:00','2022-01-24 15:58:53',NULL),(234,'Wallmart','supermarket','','https://cdn.corporate.walmart.com/dims4/WMT/ea03f5e/2147483647/strip/true/crop/855x305+0+0/resize/1578x563!/quality/90/?url=https%3A%2F%2Fcdn.corporate.walmart.com%2F98%2F28%2F342ccbff478ab025592645fafcfc%2Fwalmart-logo.png','00:00','24:00','2022-01-24 15:58:53',NULL),(235,'Linea Caffe','coffee_shop','','https://lineacaffe.com/wp-content/themes/lineacaffe/images/linea-logo.svg','07:00','19:00','2022-01-24 15:58:53',NULL),(236,'Saint Frank Coffee','coffee_shop','','http://cdn.shopify.com/s/files/1/1578/1589/files/colorf-01_198x200.png','12:00','18:00','2022-01-24 16:03:58',NULL),(237,'Sushi Space','restaurant','','https://image.freepik.com/free-vector/sushi-restaurant-logo_8169-12.jpg','08:00','20:00','2022-01-24 16:03:58',NULL),(238,'Pizza Club','restaurant','','https://play-lh.googleusercontent.com/qMewibe3u5Wvq3fBf3Ca3_QItjHCOKeGrOAzVXWxqzgRpMwxYlD5CA6M2M5L78SwNA','09:00','20:00','2022-01-24 16:03:58',NULL),(239,'Burger Club','restaurant','','https://eda.ua/images/506509-195-195-burger_club_harkov.jpg','10:00','20:00','2022-01-24 16:03:58',NULL),(240,'Target','supermarket','','https://1000logos.net/wp-content/uploads/2021/04/Target-logo.png','00:00','24:00','2022-01-24 16:03:58',NULL),(241,'Linea Caffe','coffee_shop','','https://lineacaffe.com/wp-content/themes/lineacaffe/images/linea-logo.svg','07:00','19:00','2022-01-24 16:03:58',NULL),(242,'Wallmart','supermarket','','https://cdn.corporate.walmart.com/dims4/WMT/ea03f5e/2147483647/strip/true/crop/855x305+0+0/resize/1578x563!/quality/90/?url=https%3A%2F%2Fcdn.corporate.walmart.com%2F98%2F28%2F342ccbff478ab025592645fafcfc%2Fwalmart-logo.png','00:00','24:00','2022-01-24 16:03:58',NULL);
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

-- Dump completed on 2022-01-24 18:10:46
