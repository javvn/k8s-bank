# CREATE TABLE `employees` (
#      `emp_no` int NOT NULL,
#      `birth_date` date NOT NULL,
#      `first_name` varchar(14) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
#      `last_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
#      `gender` enum('M','F') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
#      `hire_date` date NOT NULL,
#      PRIMARY KEY (`emp_no`),
#      KEY `ix_hiredate` (`hire_date`),
#      KEY `ix_gender_birthdate` (`gender`,`birth_date`),
#      KEY `ix_firstname` (`first_name`)
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci STATS_PERSISTENT=0;


CREATE TABLE `user` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `birthday` date NOT NULL,
                        `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        `gender` enum('M','F') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        `email` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        `leave` tinyint(1) NOT NULL DEFAULT 0,
                        `createdAt` date NOT NULL,
                        `updatedAt` date NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


INSERT INTO `user` VALUES (10001,'1953-09-02','Georgi','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10002,'1964-06-02','Bezalel','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10003,'1959-12-03','Parto','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10004,'1954-05-01','Chirstian','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10005,'1955-01-21','Kyoichi','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10006,'1953-04-20','Anneke','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10007,'1957-05-23','Tzvetan','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10008,'1958-02-19','Saniya','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10009,'1952-04-19','Sumant','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10010,'1963-06-01','Duangkaew','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10011,'1953-11-07','Mary','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10012,'1960-10-04','Patricio','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10013,'1963-06-07','Eberhardt','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10014,'1956-02-12','Berni','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10015,'1959-08-19','Guoxiang','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10016,'1961-05-02','Kazuhito','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10017,'1958-07-06','Cristinel','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10018,'1954-06-19','Kazuhide','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10019,'1953-01-23','Lillian','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10020,'1952-12-24','Mayuko','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10021,'1960-02-20','Ramzi','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10022,'1952-07-08','Shahaf','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10023,'1953-09-29','Bojan','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10024,'1958-09-05','Suzette','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10025,'1958-10-31','Prasadram','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10026,'1953-04-03','Yongqiao','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10027,'1962-07-10','Divier','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10028,'1963-11-26','Domenick','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10029,'1956-12-13','Otmar','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10030,'1958-07-14','Elvis','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10031,'1959-01-27','Karsten','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10032,'1960-08-09','Jeong','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10033,'1956-11-14','Arif','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10034,'1962-12-29','Bader','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10035,'1953-02-08','Alain','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10036,'1959-08-10','Adamantios','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10037,'1963-07-22','Pradeep','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10038,'1960-07-20','Huan','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10039,'1959-10-01','Alejandro','M','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00'),
                          (10040,'1959-09-13','Weiyi','F','01013411233','test@test.com',0,'2023-04-26 00:00:00','2023-04-26 00:00:00');



# CREATE TABLE `employee_name` (
                                   #                                  `emp_no` int NOT NULL,
                                   #                                  `first_name` varchar(14) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                   #                                  `last_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                   #                                  PRIMARY KEY (`emp_no`),
                                   #                                  FULLTEXT KEY `fx_name` (`first_name`,`last_name`) /*!50100 WITH PARSER `ngram` */
                                       # ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
# /*!40101 SET character_set_client = @saved_cs_client */;
