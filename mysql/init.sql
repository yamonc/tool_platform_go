-- MySQL dump 10.13  Distrib 5.7.17, for Win64 (x86_64)
--
-- Host: mysql.yy    Database: biligo
-- ------------------------------------------------------
-- Server version	5.7.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `display_name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `user_status` int(11) DEFAULT '1',
  `dept_id` int(11) DEFAULT NULL,
  `login_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;
INSERT INTO `sys_user` VALUES (1,'admin','admin','$2a$10$hntRCHBkQ9.9M2O5Mq/Uo.n/zdKFjzFYGT3AjgMW5DB110o4Cbo6m',1,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_token`
--

DROP TABLE IF EXISTS `sys_user_token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_token` (
  `token` varchar(255) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `expired_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_token`
--

LOCK TABLES `sys_user_token` WRITE;
/*!40000 ALTER TABLE `sys_user_token` DISABLE KEYS */;
INSERT INTO `sys_user_token` VALUES ('1859f35b-56d2-4199-8e23-d1bc4c9c8583',1,'2019-08-25 22:06:27'),('2a67f231-e00a-4d51-9df2-a06d4529fc39',1,'2019-08-25 21:33:26'),('318387e0-cd2f-41ae-80b5-dc72ea2b8fb8',1,'2019-08-25 22:03:14'),('5094b788-ab5d-4017-b22b-635ac79c92c1',1,'2019-08-25 21:57:35'),('50d96af4-e046-4f95-88e0-88d7553955c5',1,'2019-08-25 21:33:30'),('a2f6ef9b-3f80-4fd0-b068-fc1130871728',1,'2019-08-25 21:48:31'),('cb37b3bd-e4f9-48c5-9936-7bdddbd8e88e',1,'2019-08-25 21:33:28'),('d880db41-fdb4-486d-baff-b2e89a668600',1,'2019-08-25 22:09:00'),('ed0cdc82-49f2-44fb-99bc-63162cb1afe6',1,'2019-08-25 21:23:33');
/*!40000 ALTER TABLE `sys_user_token` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-08-18 22:13:33
# 初始化car_record表
insert into `car_record` values('1',null, '更换发动机机油 *1', '5000', '6');
insert into `car_record` values('2',null, '更换发动机机油滤清器 *1', '5000', '12');
insert into `car_record` values('3',null, '更换空气滤清器滤芯', '20000', null);
insert into `car_record` values('4',null, '检查气门间隙*2', '120000', null);
insert into `car_record` values('5', null,'更换火花塞', '100000', null);
insert into `car_record` values('6',null, '检查传动皮带', '400000', null);
insert into `car_record` values('7',null, '更换发动机冷却液', '200000', null);
insert into `car_record` values('8',null, '添加纯正喷油嘴清洗剂', '5000', '6');
insert into `car_record` values('9',null, '更换变速器油', '400000', '12');
insert into `car_record` values('10',null, '更换粉尘滤清器', '20000', '12');
insert into `car_record` values('11',null, '检查前后制动器', '10000', '12');
insert into `car_record` values('12',null, '更换制动液', null, '36');
insert into `car_record` values('13',null, '更换油箱中的燃油滤清器', '180000', '36');
insert into `car_record` values('14',null, '轮胎换位（每个月至少检查一次轮胎充气和状况）', '10000', null);
insert into `car_record` values('15',null, '传动轴防尘罩横拉杆接头、转向齿轮箱和防尘套悬架部件', '10000', '12');
insert into `car_record` values('16',null, '制动软管和管路（包括 ABS/VSA）排气系统燃油管路和连接', '20000', '12');

create table car
(
    uuid        varchar(255) null,
    name        varchar(255) null comment '车辆名称',
    remark      text         null comment '备注',
    daily_km    varchar(255) null comment '每日行驶距离',
    is_alarm    tinyint(1)   null comment '是否开启提醒',
    create_time timestamp    null comment '创建时间'
);


