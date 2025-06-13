-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jun 13, 2025 at 09:19 AM
-- Server version: 11.7.2-MariaDB-deb12
-- PHP Version: 8.3.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `YNOV_B1_forum`
--
CREATE DATABASE IF NOT EXISTS `YNOV_B1_forum` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `YNOV_B1_forum`;

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `name` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`) VALUES
('fb614388-456c-11f0-96e0-bc2411e6ce9e', 'General Discussion'),
('fb61568f-456c-11f0-96e0-bc2411e6ce9e', 'Tips and Guides'),
('fb61888f-456c-11f0-96e0-bc2411e6ce9e', 'Warframes and Equipment'),
('fb618939-456c-11f0-96e0-bc2411e6ce9e', 'Primary Weapons'),
('fb61896a-456c-11f0-96e0-bc2411e6ce9e', 'Secondary Weapons'),
('fb618989-456c-11f0-96e0-bc2411e6ce9e', 'Melee Weapons'),
('fb6189a6-456c-11f0-96e0-bc2411e6ce9e', 'Mods and Builds'),
('fb6189c1-456c-11f0-96e0-bc2411e6ce9e', 'Missions and Factions'),
('fb6189e2-456c-11f0-96e0-bc2411e6ce9e', 'Quests and Lore'),
('fb6189fc-456c-11f0-96e0-bc2411e6ce9e', 'Dojo and Clans'),
('fb618a15-456c-11f0-96e0-bc2411e6ce9e', 'In-Game Events'),
('fb618a2f-456c-11f0-96e0-bc2411e6ce9e', 'Official Announcements'),
('fb618a49-456c-11f0-96e0-bc2411e6ce9e', 'Bugs and Technical Support'),
('fb618a64-456c-11f0-96e0-bc2411e6ce9e', 'Suggestions and Feedback'),
('fb618a80-456c-11f0-96e0-bc2411e6ce9e', 'Trading and Market');
INSERT INTO `categories` (`id`, `name`) VALUES
('fb618a9d-456c-11f0-96e0-bc2411e6ce9e', 'Screenshots and Fan Art'),
('fb618abc-456c-11f0-96e0-bc2411e6ce9e', 'Roleplay and Stories'),
('fb618adb-456c-11f0-96e0-bc2411e6ce9e', 'Tennocon and News');

-- --------------------------------------------------------

--
-- Table structure for table `comments`
--

CREATE TABLE `comments` (
  `id` int(11) NOT NULL,
  `content` varchar(200) NOT NULL,
  `createdAt` datetime NOT NULL,
  `post_id` int(11) NOT NULL,
  `user_id` uuid NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `notifications`
--

CREATE TABLE `notifications` (
  `id` int(11) NOT NULL,
  `title` varchar(35) NOT NULL,
  `description` varchar(150) NOT NULL,
  `createdAt` datetime NOT NULL,
  `from_user_id` uuid NOT NULL,
  `user_id` uuid NOT NULL,
  `post_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` int(11) NOT NULL,
  `title` varchar(150) NOT NULL,
  `content` text NOT NULL,
  `picture` longblob NOT NULL DEFAULT '',
  `validated` tinyint(1) NOT NULL,
  `createdAt` datetime NOT NULL,
  `user_id` uuid NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `posts_category`
--

CREATE TABLE `posts_category` (
  `category_id` uuid NOT NULL,
  `post_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `reactions`
--

CREATE TABLE `reactions` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `post_id` int(11) DEFAULT NULL,
  `comment_id` int(11) DEFAULT NULL,
  `user_id` uuid NOT NULL,
  `label` varchar(7) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `recent_activity`
--

CREATE TABLE `recent_activity` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `action` varchar(30) NOT NULL,
  `details` varchar(100) NOT NULL,
  `subTitle` varchar(200) DEFAULT NULL,
  `user_id` uuid NOT NULL,
  `post_id` int(11) NOT NULL,
  `createdAt` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `reports`
--

CREATE TABLE `reports` (
  `id` uuid NOT NULL,
  `post_id` int(11) NOT NULL,
  `user_id` uuid NOT NULL,
  `reportedAt` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `name` varchar(20) NOT NULL,
  `permission` bit(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `name`, `permission`) VALUES
('d3bf52aa-3b05-11f0-96e0-bc2411e6ce9e', 'user', b'001'),
('e862f1a2-3b05-11f0-96e0-bc2411e6ce9e', 'moderator', b'010'),
('e862f63f-3b05-11f0-96e0-bc2411e6ce9e', 'admin', b'100');

-- --------------------------------------------------------

--
-- Table structure for table `sessions`
--

CREATE TABLE `sessions` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `expireAt` timestamp NOT NULL,
  `user_id` uuid DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `pseudo` varchar(32) NOT NULL,
  `email` varchar(254) NOT NULL,
  `password` char(60) DEFAULT NULL,
  `bio` varchar(500) NOT NULL DEFAULT '',
  `avatar` mediumblob NOT NULL DEFAULT '\'\'',
  `createdAt` datetime NOT NULL DEFAULT current_timestamp(),
  `role_id` uuid NOT NULL,
  `google_id` varchar(64) DEFAULT NULL,
  `github_id` varchar(64) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `notifications`
--
ALTER TABLE `notifications`
  ADD PRIMARY KEY (`id`),
  ADD KEY `notification_user_FK` (`user_id`),
  ADD KEY `notification_post_FK` (`post_id`),
  ADD KEY `notification_from_user_FK` (`from_user_id`);

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_fk` (`user_id`);

--
-- Indexes for table `posts_category`
--
ALTER TABLE `posts_category`
  ADD KEY `post_category_fk` (`post_id`),
  ADD KEY `post_category_UNIQUE` (`category_id`,`post_id`) USING BTREE;

--
-- Indexes for table `reactions`
--
ALTER TABLE `reactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `reaction_type_fk` (`label`),
  ADD KEY `user_fk` (`user_id`) USING BTREE,
  ADD KEY `comment_fk` (`comment_id`),
  ADD KEY `post_fk` (`post_id`);

--
-- Indexes for table `recent_activity`
--
ALTER TABLE `recent_activity`
  ADD PRIMARY KEY (`id`),
  ADD KEY `recentActivity_user_FK` (`user_id`);

--
-- Indexes for table `reports`
--
ALTER TABLE `reports`
  ADD KEY `post_foreignKey` (`post_id`),
  ADD KEY `report_user_FK` (`user_id`);

--
-- Indexes for table `sessions`
--
ALTER TABLE `sessions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id_fk` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `pseudo` (`pseudo`),
  ADD UNIQUE KEY `idx_email` (`email`),
  ADD KEY `role_fk` (`role_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `comments`
--
ALTER TABLE `comments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `notifications`
--
ALTER TABLE `notifications`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `comments`
--
ALTER TABLE `comments`
  ADD CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `notifications`
--
ALTER TABLE `notifications`
  ADD CONSTRAINT `notification_from_user_FK` FOREIGN KEY (`from_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `notification_post_FK` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `notification_user_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `posts`
--
ALTER TABLE `posts`
  ADD CONSTRAINT `user_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `posts_category`
--
ALTER TABLE `posts_category`
  ADD CONSTRAINT `post_category_fk` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION;

--
-- Constraints for table `reactions`
--
ALTER TABLE `reactions`
  ADD CONSTRAINT `comment_fk` FOREIGN KEY (`comment_id`) REFERENCES `comments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `post_fk` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `user_foreignKey` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `recent_activity`
--
ALTER TABLE `recent_activity`
  ADD CONSTRAINT `recentActivity_user_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `reports`
--
ALTER TABLE `reports`
  ADD CONSTRAINT `post_foreignKey` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `report_user_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `sessions`
--
ALTER TABLE `sessions`
  ADD CONSTRAINT `user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
