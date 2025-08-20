-- phpMyAdmin SQL Dump
-- version 4.9.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 15, 2020 at 06:05 PM
-- Server version: 8.0.18
-- PHP Version: 7.3.11
SET
  SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

SET
  AUTOCOMMIT = 0;

START TRANSACTION;

SET
  time_zone = "+00:00";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;

/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;

/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;

/*!40101 SET NAMES utf8mb4 */
;

--
-- Database: `web_gudang`
--
-- --------------------------------------------------------
--
-- Table structure for table `incoming`
--
CREATE TABLE `incoming` (
  `id` int(11) NOT NULL,
  `transaksi_id` varchar(50) NOT NULL,
  `tanggal` varchar(20) NOT NULL,
  `lokasi` varchar(100) NOT NULL,
  `kode_barang` varchar(100) NOT NULL,
  `nama_barang` varchar(100) NOT NULL,
  `satuan` varchar(50) NOT NULL,
  `jumlah` int(255) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = latin1;

--
-- Dumping data for table `incoming`
--
INSERT INTO
  `incoming` (
    `id`,
    `transaksi_id`,
    `tanggal`,
    `lokasi`,
    `kode_barang`,
    `nama_barang`,
    `satuan`,
    `jumlah`
  )
VALUES
  (
    18,
    'TRX-157-616',
    '2020/05/16 01:00:57',
    'Jakarta Selatan',
    'JKS',
    'Genset',
    'Unit',
    124
  ),
  (
    19,
    'TRX-133-636',
    '2020/05/16 01:01:13',
    'Jakarta Selatan',
    'JKS',
    'Genset',
    'Unit',
    16
  );

-- --------------------------------------------------------
--
-- Table structure for table `items`
--
CREATE TABLE `items` (
  `id` int(11) NOT NULL,
  `item_name` varchar(100) NOT NULL,
  `item_code` varchar(100) NOT NULL,
  `unit_id` int(11) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- --------------------------------------------------------
--
-- Table structure for table `outcoming`
--
CREATE TABLE `outcoming` (
  `id` int(10) NOT NULL,
  `transaksi_id` varchar(50) NOT NULL,
  `tanggal_masuk` varchar(20) NOT NULL,
  `tanggal_keluar` varchar(20) NOT NULL,
  `lokasi` varchar(100) NOT NULL,
  `kode_barang` varchar(100) NOT NULL,
  `nama_barang` varchar(100) NOT NULL,
  `satuan` varchar(50) NOT NULL,
  `jumlah` varchar(10) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = latin1;

--
-- Dumping data for table `outcoming`
--
INSERT INTO
  `outcoming` (
    `id`,
    `transaksi_id`,
    `tanggal_masuk`,
    `tanggal_keluar`,
    `lokasi`,
    `kode_barang`,
    `nama_barang`,
    `satuan`,
    `jumlah`
  )
VALUES
  (
    23,
    'TRX-157-616',
    '2020/05/16 01:00:57',
    '2020/05/16 01:04:30',
    'Jakarta Selatan',
    'JKS',
    'Genset',
    'Unit',
    '13'
  );

-- --------------------------------------------------------
--
-- Table structure for table `test`
--
CREATE TABLE `test` (
  `id` int(11) NOT NULL,
  `name` int(11) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;

-- --------------------------------------------------------
--
-- Table structure for table `units`
--
CREATE TABLE `units` (
  `id` int(11) NOT NULL,
  `code` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;

--
-- Dumping data for table `units`
--
INSERT INTO
  `units` (`id`, `code`, `name`)
VALUES
  (1, 'KG', 'Kilogram'),
  (3, 'U', 'Unit'),
  (4, 'KL', 'Koli'),
  (5, 'PCS', 'Piece');

-- --------------------------------------------------------
--
-- Table structure for table `users`
--
CREATE TABLE `users` (
  `id` int(12) NOT NULL,
  `username` varchar(200) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(200) NOT NULL,
  `role` tinyint(4) NOT NULL DEFAULT '0',
  `last_login` varchar(20) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = latin1;

--
-- Dumping data for table `users`
--
INSERT INTO
  `users` (
    `id`,
    `username`,
    `email`,
    `password`,
    `role`,
    `last_login`
  )
VALUES
  (
    1,
    'admin',
    'admin@wms.com',
    '$2a$04$XzZVjxw99O9rNC1TnJ8s2.LILQA11n2Np8reVPIXV/KSRCVED3WLK',
    1,
    ''
  ),
  (
    2,
    'dany_arie',
    'dany_arie@yahoo.com',
    '$2a$04$TiV.22MinQqCegME2t5LU./GZf6VZkzlqBMRh.LZGCIyHigJn/LRq',
    0,
    ''
  ),
  (
    3,
    'rizal',
    'rizal@mail.com',
    '$2a$04$Z9yHSKxdU7qyr.xnW6CCqeV0GxQyFxse9vcLzWizkXIxKKabHReSa',
    1,
    ''
  );

-- --------------------------------------------------------
--
-- Stand-in structure for view `v_in_out_reporting`
-- (See below for the actual view)
--
CREATE TABLE `v_in_out_reporting` (
  `jumlah_keluar` double,
  `jumlah_masuk` decimal(65, 0),
  `kode_barang` varchar(100),
  `lokasi` varchar(100),
  `nama_barang` varchar(100),
  `satuan` varchar(50),
  `sisa` double
);

-- --------------------------------------------------------
--
-- Stand-in structure for view `v_total_in`
-- (See below for the actual view)
--
CREATE TABLE `v_total_in` (
  `jumlah_masuk` decimal(65, 0),
  `kode_barang` varchar(100),
  `lokasi` varchar(100),
  `nama_barang` varchar(100),
  `satuan` varchar(50)
);

-- --------------------------------------------------------
--
-- Stand-in structure for view `v_total_out`
-- (See below for the actual view)
--
CREATE TABLE `v_total_out` (`nama_barang` varchar(100), `total_out` double);

-- --------------------------------------------------------
--
-- Structure for view `v_in_out_reporting`
--
DROP TABLE IF EXISTS `v_in_out_reporting`;

CREATE ALGORITHM = UNDEFINED DEFINER = `root` @`localhost` SQL SECURITY DEFINER VIEW `v_in_out_reporting` AS
select
  `a`.`nama_barang` AS `nama_barang`,
  `a`.`lokasi` AS `lokasi`,
  `a`.`kode_barang` AS `kode_barang`,
  `a`.`satuan` AS `satuan`,
  `a`.`jumlah_masuk` AS `jumlah_masuk`,
  `b`.`total_out` AS `jumlah_keluar`,
(`a`.`jumlah_masuk` - `b`.`total_out`) AS `sisa`
from
  (
    `v_total_in` `a`
    left join `v_total_out` `b` on(
      (
        convert(`a`.`nama_barang` using utf8) = convert(`b`.`nama_barang` using utf8)
      )
    )
  );

-- --------------------------------------------------------
--
-- Structure for view `v_total_in`
--
DROP TABLE IF EXISTS `v_total_in`;

CREATE ALGORITHM = UNDEFINED DEFINER = `root` @`localhost` SQL SECURITY DEFINER VIEW `v_total_in` AS
select
  `incoming`.`nama_barang` AS `nama_barang`,
  `incoming`.`lokasi` AS `lokasi`,
  `incoming`.`kode_barang` AS `kode_barang`,
  `incoming`.`satuan` AS `satuan`,
  sum(`incoming`.`jumlah`) AS `jumlah_masuk`
from
  `incoming`
group by
  `incoming`.`nama_barang`,
  `incoming`.`lokasi`,
  `incoming`.`kode_barang`,
  `incoming`.`satuan`
having
  (sum(`incoming`.`jumlah`) >= 0);

-- --------------------------------------------------------
--
-- Structure for view `v_total_out`
--
DROP TABLE IF EXISTS `v_total_out`;

CREATE ALGORITHM = UNDEFINED DEFINER = `root` @`localhost` SQL SECURITY DEFINER VIEW `v_total_out` AS
select
  `outcoming`.`nama_barang` AS `nama_barang`,
  sum(`outcoming`.`jumlah`) AS `total_out`
from
  `outcoming`
group by
  `outcoming`.`nama_barang`
having
  (sum(`outcoming`.`jumlah`) >= 0);

--
-- Indexes for dumped tables
--
--
-- Indexes for table `incoming`
--
ALTER TABLE
  `incoming`
ADD
  PRIMARY KEY (`id`);

--
-- Indexes for table `outcoming`
--
ALTER TABLE
  `outcoming`
ADD
  PRIMARY KEY (`id`);

--
-- Indexes for table `test`
--
ALTER TABLE
  `test`
ADD
  PRIMARY KEY (`id`);

--
-- Indexes for table `units`
--
ALTER TABLE
  `units`
ADD
  PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE
  `users`
ADD
  PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--
--
-- AUTO_INCREMENT for table `incoming`
--
ALTER TABLE
  `incoming`
MODIFY
  `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 20;

--
-- AUTO_INCREMENT for table `outcoming`
--
ALTER TABLE
  `outcoming`
MODIFY
  `id` int(10) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 24;

--
-- AUTO_INCREMENT for table `test`
--
ALTER TABLE
  `test`
MODIFY
  `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `units`
--
ALTER TABLE
  `units`
MODIFY
  `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 6;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE
  `users`
MODIFY
  `id` int(12) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 4;

COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;

/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;

/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;