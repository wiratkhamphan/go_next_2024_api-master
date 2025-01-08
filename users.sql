-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 08, 2025 at 10:13 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.0.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `shoplek`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `username` varchar(255) NOT NULL,
  `level` int(11) NOT NULL,
  `section_id` int(11) NOT NULL,
  `status` varchar(50) NOT NULL DEFAULT 'active'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `username`, `level`, `section_id`, `status`) VALUES
(5, '123', '', '$2a$14$9IA3Sx8I1L0fEtcrCERZ6eQKODkb1CZXFXpZI/2xW7OG2FoaVkxc6', '', 0, 0, 'active'),
(6, '123', 'wirat_@hotmail.com', '$2a$14$nw2Q1QoBi41IiDNrV9g0Re69figC/SLbGh9.WUj4zh6.7/JmVvw7a', 'Lek01', 0, 0, 'active'),
(10, '', 'wirat_@@gmail.com', '$2a$10$UvZQAfLIWXxms3mptARrmuXck9BZrIl/ZsAxLpItXtTLMnMCXWQey', '', 0, 0, 'active'),
(11, 'LEK', 'wirat1_@@gmail.com', '$2a$10$Pkz0uO2JAhgNiuVdlw6iw.MoODlgFx9jgFPpFktujE.0LBWRjaAZW', '', 0, 0, 'active'),
(14, 'LEK1', 'wirat1_@@gmail.com1', '$2a$10$ORPIXCjRdzFhnuMT3H8otOBUfMhGFydsYgoYXU/MB6F63wUHnvqG.', '', 0, 0, 'active'),
(15, 'LEK1', 'wirat1_@@gmail.com11', '$2a$10$xyE.CSukhxTqVC.KfJT1LO8mvb30kvD/ypfSwbrPJqvkmhBSe/Awy', '', 0, 0, 'active'),
(16, 'LEK1', 'wirat1_@@gmail.com111', '$2a$10$96O8XEuy.LBNRVlt4Q47QukTY4X1eqHkwRmvgo.lCQuGnephfJO32', '', 0, 0, 'active');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
