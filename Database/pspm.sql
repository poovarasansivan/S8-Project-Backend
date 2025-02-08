-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 08, 2025 at 02:37 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pspm`
--

-- --------------------------------------------------------

--
-- Table structure for table `full_stack_projects`
--

CREATE TABLE `full_stack_projects` (
  `id` int(10) NOT NULL,
  `project_id` varchar(10) NOT NULL,
  `project_title` varchar(50) NOT NULL,
  `project_description` text NOT NULL,
  `assigend_by` varchar(25) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `full_stack_projects`
--

INSERT INTO `full_stack_projects` (`id`, `project_id`, `project_title`, `project_description`, `assigend_by`, `created_at`, `updated_at`) VALUES
(1, 'PS-1', 'One Credit Course Portal', 'Develope a Once credit Course Portal for automating the course registration for college students.', 'Sathish P', '2025-01-30 09:21:30', '2025-01-30 09:21:30'),
(2, 'PS-2', 'Placement Student Performance Management System', 'Develop a Portal to manage the placement and performance of a student under various categories and technical parameters.', 'Santhosh S', '2025-01-30 09:24:31', '2025-01-30 09:24:31');

-- --------------------------------------------------------

--
-- Table structure for table `full_stack_review`
--

CREATE TABLE `full_stack_review` (
  `id` int(10) NOT NULL,
  `roll_no` varchar(15) NOT NULL,
  `reviewer_name` varchar(50) NOT NULL,
  `reviewer_email` varchar(100) NOT NULL,
  `rev_department` varchar(100) NOT NULL,
  `rev_venue` varchar(50) NOT NULL,
  `date` date NOT NULL,
  `time` time NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `full_stack_review`
--

INSERT INTO `full_stack_review` (`id`, `roll_no`, `reviewer_name`, `reviewer_email`, `rev_department`, `rev_venue`, `date`, `time`, `created_at`, `updated_at`) VALUES
(1, '211CS239', 'Praveen P', 'praveen@gmail.com', 'Computer Science And Engineering ', 'WW110', '2025-02-12', '11:00:00', '2025-01-30 09:24:49', '2025-01-30 09:24:49'),
(2, '211CS169', 'Sutha M', 'sutha@gmail.com', 'Traning And Placement', 'WW111', '2025-02-13', '10:30:00', '2025-01-30 09:26:56', '2025-01-30 09:26:56');

-- --------------------------------------------------------

--
-- Table structure for table `placement`
--

CREATE TABLE `placement` (
  `id` int(10) NOT NULL,
  `roll_no` varchar(25) NOT NULL,
  `placement_type` varchar(25) NOT NULL,
  `placement_status` varchar(25) NOT NULL,
  `placement_offer_1` varchar(25) NOT NULL,
  `offer1_lpa` int(25) NOT NULL,
  `placement_offer_2` varchar(25) NOT NULL,
  `offer2_lpa` varchar(25) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `placement`
--

INSERT INTO `placement` (`id`, `roll_no`, `placement_type`, `placement_status`, `placement_offer_1`, `offer1_lpa`, `placement_offer_2`, `offer2_lpa`, `created_at`, `updated_at`) VALUES
(1, '211CS239', 'IT ', 'Placed', 'Presidio', 10, 'TCS', '4', '2025-01-30 09:28:48', '2025-01-30 09:28:48'),
(2, '211CS169', 'IT', 'Not Placed', 'NA', 0, 'NA', '0', '2025-01-30 09:29:33', '2025-01-30 09:29:33'),
(3, '212IT146', 'NIP', 'NA', 'NA', 0, 'NA', '0 ', '2025-01-30 09:29:52', '2025-01-30 09:29:52'),
(4, '212IT196', 'IT', 'Placed', 'Avasoft', 0, 'NA ', '0', '2025-01-30 09:30:06', '2025-01-30 09:30:06');

-- --------------------------------------------------------

--
-- Table structure for table `ps_data`
--

CREATE TABLE `ps_data` (
  `id` int(10) NOT NULL,
  `roll_no` varchar(15) NOT NULL,
  `ps_category` varchar(20) NOT NULL,
  `category_level` varchar(20) NOT NULL,
  `slot_date` datetime NOT NULL,
  `level_status` varchar(25) NOT NULL,
  `slot_missed` varchar(10) NOT NULL,
  `negative_marks` varchar(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ps_data`
--

INSERT INTO `ps_data` (`id`, `roll_no`, `ps_category`, `category_level`, `slot_date`, `level_status`, `slot_missed`, `negative_marks`, `created_at`, `updated_at`) VALUES
(1, '211CS239', 'Technical', 'Level 1', '2025-01-22 15:01:28', 'Completed', 'No', 'no', '2025-02-03 10:14:54', '2025-02-03 10:14:54'),
(2, '211CS239', 'Aptitude', 'Level 1', '2025-01-03 15:02:11', 'Completed', 'No', 'no', '2025-02-03 10:14:54', '2025-02-03 10:14:54'),
(3, '211CS239', 'Verbal', 'Level 1', '2025-01-11 10:02:32', 'Completed', 'No', 'no', '2025-02-03 10:14:54', '2025-02-03 10:14:54'),
(4, '211CS239', 'Technical', 'Level 2', '2025-01-23 11:02:57', 'Completed', 'No', 'yes', '2025-02-03 10:14:54', '2025-02-03 10:14:54'),
(5, '211CS239', 'Technical', 'Level 3', '2025-01-24 15:03:22', 'Not Completed', 'Yes', 'no', '2025-02-03 10:14:54', '2025-02-03 10:14:54'),
(6, '211CS169', 'Technical', 'Level 1', '2025-01-02 11:00:00', 'Completed', 'No', 'yes', '2025-02-03 10:14:54', '2025-02-03 10:14:54'),
(7, '211CS169', 'Technical', 'Level 2', '2025-01-05 12:00:00', 'Completed', 'No', 'no', '2025-02-03 10:14:54', '2025-02-03 10:14:54'),
(8, '211CS169', 'Technical', 'Level 3', '2025-01-07 12:00:00', 'Not Completed', 'Yes', 'no', '2025-02-03 10:14:54', '2025-02-03 10:14:54'),
(9, '212IT146', 'Technical', 'Level 1', '2025-01-14 13:00:00', 'Completed', 'No', 'yes', '2025-02-03 10:14:54', '2025-02-03 10:14:54'),
(10, '212IT146', 'Technical', 'Level 2', '1970-01-16 10:30:00', 'Completed', 'No', 'no', '2025-02-03 10:14:54', '2025-02-03 10:14:54');

-- --------------------------------------------------------

--
-- Table structure for table `ps_ovr_data`
--

CREATE TABLE `ps_ovr_data` (
  `id` int(10) NOT NULL,
  `category` varchar(50) NOT NULL,
  `total_levels` int(40) DEFAULT NULL,
  `monthly_level` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `Updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ps_ovr_data`
--

INSERT INTO `ps_ovr_data` (`id`, `category`, `total_levels`, `monthly_level`, `created_at`, `Updated_at`) VALUES
(1, 'Technical', 15, '6', '2025-01-30 09:40:23', '2025-01-30 09:40:23'),
(2, 'Aptitude', 15, '5', '2025-01-30 09:41:00', '2025-01-30 09:40:43'),
(3, 'Verbal', 15, '4', '2025-01-30 09:41:17', '2025-01-30 09:41:17');

-- --------------------------------------------------------

--
-- Table structure for table `students_fa_data`
--

CREATE TABLE `students_fa_data` (
  `id` int(15) NOT NULL,
  `roll_no` varchar(100) NOT NULL,
  `placement_fa` int(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `students_fa_data`
--

INSERT INTO `students_fa_data` (`id`, `roll_no`, `placement_fa`, `created_at`, `updated_at`) VALUES
(1, '211CS169 ', 80, '2025-02-03 06:25:23', '2025-02-03 06:25:23'),
(2, '211CS248', 0, '2025-02-03 06:25:31', '2025-02-03 06:25:31');

-- --------------------------------------------------------

--
-- Table structure for table `student_academic_data`
--

CREATE TABLE `student_academic_data` (
  `id` int(10) NOT NULL,
  `roll_no` varchar(25) NOT NULL,
  `placement_fa` double NOT NULL,
  `academic_fa` double NOT NULL,
  `cgpa` double NOT NULL,
  `sem1` double NOT NULL,
  `sem2` double NOT NULL,
  `sem3` double NOT NULL,
  `sem4` double NOT NULL,
  `sem5` double NOT NULL,
  `sem6` double NOT NULL,
  `sem7` double NOT NULL,
  `sem8` double NOT NULL,
  `arrears` int(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `student_academic_data`
--

INSERT INTO `student_academic_data` (`id`, `roll_no`, `placement_fa`, `academic_fa`, `cgpa`, `sem1`, `sem2`, `sem3`, `sem4`, `sem5`, `sem6`, `sem7`, `sem8`, `arrears`, `created_at`, `updated_at`) VALUES
(1, '211CS239', 85.2, 45.2, 8.45, 8.12, 8.21, 8.24, 8.28, 8.58, 8.64, 8.9, 0, 0, '2025-02-03 06:28:00', '2025-02-03 06:28:00'),
(2, '211CS169', 62.8, 88.5, 8.69, 8.12, 8.45, 8.6, 8.1, 8.42, 8.4, 9.12, 0, 0, '2025-02-03 06:28:00', '2025-02-03 06:28:00'),
(3, '212IT146', 68, 90, 8.5, 8.1, 8.3, 8.4, 8.6, 8.7, 8.5, 8.2, 0, 0, '2025-02-03 06:28:00', '2025-02-03 06:28:00'),
(4, '212IT196', 12.5, 55.6, 7.89, 8.1, 8.3, 8.4, 8.6, 8.7, 8.5, 8.2, 0, 1, '2025-02-03 06:28:00', '2025-02-03 06:28:00');

-- --------------------------------------------------------

--
-- Table structure for table `student_education`
--

CREATE TABLE `student_education` (
  `id` int(10) NOT NULL,
  `roll_no` varchar(10) NOT NULL,
  `sslc_school_name` varchar(100) NOT NULL,
  `sslc_board` varchar(100) NOT NULL,
  `sslc_year_completed` varchar(10) NOT NULL,
  `sslc_percentage` varchar(10) NOT NULL,
  `hsc_school_name` varchar(100) NOT NULL,
  `hsc_board` varchar(100) NOT NULL,
  `hsc_year_completed` varchar(10) NOT NULL,
  `hsc_percentage` varchar(10) NOT NULL,
  `degree_name` varchar(10) NOT NULL,
  `department` varchar(100) NOT NULL,
  `clg_name` varchar(100) NOT NULL,
  `year_graduation` varchar(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `student_education`
--

INSERT INTO `student_education` (`id`, `roll_no`, `sslc_school_name`, `sslc_board`, `sslc_year_completed`, `sslc_percentage`, `hsc_school_name`, `hsc_board`, `hsc_year_completed`, `hsc_percentage`, `degree_name`, `department`, `clg_name`, `year_graduation`, `created_at`, `updated_at`) VALUES
(1, '211CS169', 'Government Higher Secondary School, Edappadi', 'State Board Of Tamil Nadu', '2019', '92', 'Government Higher Secondary School, Edappadi', 'State Board Of Tamil Nadu', '2021', '95', 'B.E', 'Computer Science And Engineering', 'Bannari Amman Institute Of Technology', '2025', '2025-01-30 14:10:43', '2025-01-30 14:10:43'),
(2, '211CS239', 'SNS Higher Secondary School, Dharmapuri', 'State Board Of Tamil Nadu', '2019', '92', 'SNS Higher Secondary School, Dharmapuri', 'State Board Of Tamil Nadu', '2021', '95', 'B.E', 'Computer Science And Engineering', 'Bannari Amman Institute Of Technology', '2025', '2025-01-30 14:10:43', '2025-01-30 14:10:43'),
(3, '212IT146', 'Government Higher Secondary School, Palacode', 'State Board Of Tamil Nadu', '2019', '92', 'Government Higher Secondary School, Palacode', 'State Board Of Tamil Nadu', '2021', '95', 'B.Tech', 'Information Technology', 'Bannari Amman Institute Of Technology', '2025', '2025-01-30 14:10:43', '2025-01-30 14:10:43'),
(4, '212IT196', 'Blue Bird Higher Secondary School, Hosur', 'State Board Of Tamil Nadu', '2019', '92', 'Blue Bird Higher Secondary School, Hosur', 'State Board Of Tamil Nadu', '2021', '95', 'B.Tech', 'Information Technology', 'Bannari Amman Institute Of Technology', '2025', '2025-01-30 14:10:43', '2025-01-30 14:10:43');

-- --------------------------------------------------------

--
-- Table structure for table `student_full_stack_map`
--

CREATE TABLE `student_full_stack_map` (
  `id` int(10) NOT NULL,
  `roll_no` varchar(15) NOT NULL,
  `full_stack_point` int(10) NOT NULL,
  `full_stack_rank` int(10) NOT NULL,
  `current_level` int(2) NOT NULL,
  `assigned_stack` varchar(50) NOT NULL,
  `project_id` varchar(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `student_full_stack_map`
--

INSERT INTO `student_full_stack_map` (`id`, `roll_no`, `full_stack_point`, `full_stack_rank`, `current_level`, `assigned_stack`, `project_id`, `created_at`, `updated_at`) VALUES
(1, '211CS239', 2000, 20, 3, 'MERN Stack', 'PS-1', '2025-01-30 14:16:57', '2025-01-30 14:16:57'),
(2, '211CS169', 3100, 12, 4, 'MEAN Stack', 'PS-2', '2025-01-30 14:16:57', '2025-01-30 14:16:57'),
(3, '212IT146', 2800, 40, 2, 'LAMP Stack', 'PS-1', '2025-01-30 14:16:57', '2025-01-30 14:16:57'),
(4, '212IT196', 2000, 50, 2, 'MERN Stack', 'PS-2', '2025-01-30 14:16:57', '2025-01-30 14:16:57');

-- --------------------------------------------------------

--
-- Table structure for table `suggestion`
--

CREATE TABLE `suggestion` (
  `id` int(10) NOT NULL,
  `roll_no` varchar(15) NOT NULL,
  `comments` text NOT NULL,
  `reading_status` enum('0','1','2') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `suggestion`
--

INSERT INTO `suggestion` (`id`, `roll_no`, `comments`, `reading_status`, `created_at`, `updated_at`) VALUES
(1, '211CS239', 'Try to complete the PS Technical Level 5 on or before feb 6', '1', '2025-02-07 14:01:44', '2025-01-30 14:34:41'),
(2, '211CS169', 'Try to complete aptitude level 5 before first week of feb', '0', '2025-01-30 14:35:25', '2025-01-30 14:35:25'),
(3, '212IT146', 'Try to complete full stack review properly on time', '0', '2025-01-30 14:35:37', '2025-01-30 14:35:37'),
(4, '212IT196', 'Complete the PS levels to attend placements', '0', '2025-01-30 14:36:06', '2025-01-30 14:36:06'),
(5, '211CS169', 'twv', '1', '2025-02-05 15:19:03', '2025-02-05 15:19:03'),
(6, '211CS239', 'Testing ah iruku', '1', '2025-02-06 05:56:05', '2025-02-06 05:56:05'),
(7, '211CS239', 'Proper ah coding practice panitu daily ps level complete pannu', '1', '2025-02-07 14:14:02', '2025-02-07 14:14:02');

-- --------------------------------------------------------

--
-- Table structure for table `users_profile`
--

CREATE TABLE `users_profile` (
  `id` int(10) NOT NULL,
  `roll_no` varchar(20) NOT NULL,
  `name` varchar(100) NOT NULL,
  `dob` date NOT NULL,
  `gender` varchar(10) NOT NULL,
  `phone_no` varchar(100) NOT NULL,
  `regulation` varchar(10) NOT NULL,
  `batch` varchar(12) NOT NULL,
  `department` varchar(150) NOT NULL,
  `year` varchar(10) NOT NULL,
  `mentor_id` varchar(50) NOT NULL,
  `mentor_name` varchar(50) NOT NULL,
  `aadhaar_no` varchar(100) NOT NULL,
  `pan_no` varchar(20) NOT NULL,
  `college_email` varchar(50) NOT NULL,
  `personal_email` varchar(50) NOT NULL,
  `parent_mobile` varchar(100) NOT NULL,
  `leetcode_username` varchar(50) NOT NULL,
  `github_username` varchar(50) NOT NULL,
  `role` enum('0','1','2','3') NOT NULL DEFAULT '2',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users_profile`
--

INSERT INTO `users_profile` (`id`, `roll_no`, `name`, `dob`, `gender`, `phone_no`, `regulation`, `batch`, `department`, `year`, `mentor_id`, `mentor_name`, `aadhaar_no`, `pan_no`, `college_email`, `personal_email`, `parent_mobile`, `leetcode_username`, `github_username`, `role`, `created_at`, `updated_at`) VALUES
(1, '211CS239', 'Poovarasan S', '2004-03-08', 'Male', '8667536844', '2018', '2025', 'Computer Science And Engineering', 'IV', 'CS1582', 'Santhosh P', '224382428218', 'OINPS2700C', 'poovarasan.cs21@bitsathy.ac.in', 'poovarasansivan3@gmail.com', '8098235380', 'poovarasan12', 'poovarasansivan', '1', '2025-01-30 05:20:55', '2025-01-30 05:20:55'),
(2, '211CS169', 'Jayaprakash P', '2003-10-31', 'Male', '7010760359', '2018', '2025', 'Computer Science And Engineering', 'IV', 'CS1538', 'Praveen P', '224538648299', 'OINSP4685C', 'jayaprakash.cs21@bitsathy.ac.in', 'jayaprakash3110@gmail.com', '9663883884', 'pjayaprakash3110', 'pjayaprakash3110', '3', '2025-02-03 06:03:32', '2025-02-03 06:03:32'),
(3, '212IT146', 'Hari Prasad S', '2004-08-25', 'Male', '8667536850', '2018', '2025', 'Computer Science And Engineering', 'IV', 'CS1568', 'Sathish S', '224568792043', 'OINPS2341D', 'hariprasad.it21@bitsathy.ac.in', 'hariprasad@gmail.com', '4562314562', 'pjayaprakash3110', 'hariprasad-it21', '3', '2025-02-03 06:09:00', '2025-02-03 06:09:00'),
(4, '212IT196', 'Nivetha M', '2004-11-21', 'Female', '6453515963', '2018', '2025', 'Computer Science And Engineering', 'IV', 'IT1583', 'Naveena M', '321456878572', 'OINPS3245F', 'nivetha.it21@bitsathy.ac.in', 'nivetham@gmail.com', '8526934725', 'pjayaprakash3110', 'poovarasansivan', '3', '2025-02-03 06:12:52', '2025-02-03 06:12:52');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `full_stack_projects`
--
ALTER TABLE `full_stack_projects`
  ADD KEY `id` (`id`);

--
-- Indexes for table `full_stack_review`
--
ALTER TABLE `full_stack_review`
  ADD KEY `id` (`id`);

--
-- Indexes for table `placement`
--
ALTER TABLE `placement`
  ADD KEY `id` (`id`);

--
-- Indexes for table `ps_data`
--
ALTER TABLE `ps_data`
  ADD KEY `id` (`id`);

--
-- Indexes for table `ps_ovr_data`
--
ALTER TABLE `ps_ovr_data`
  ADD KEY `id` (`id`);

--
-- Indexes for table `students_fa_data`
--
ALTER TABLE `students_fa_data`
  ADD KEY `id` (`id`);

--
-- Indexes for table `student_academic_data`
--
ALTER TABLE `student_academic_data`
  ADD KEY `id` (`id`);

--
-- Indexes for table `student_education`
--
ALTER TABLE `student_education`
  ADD KEY `id` (`id`);

--
-- Indexes for table `student_full_stack_map`
--
ALTER TABLE `student_full_stack_map`
  ADD KEY `id` (`id`);

--
-- Indexes for table `suggestion`
--
ALTER TABLE `suggestion`
  ADD KEY `id` (`id`);

--
-- Indexes for table `users_profile`
--
ALTER TABLE `users_profile`
  ADD KEY `id` (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `full_stack_projects`
--
ALTER TABLE `full_stack_projects`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `full_stack_review`
--
ALTER TABLE `full_stack_review`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `placement`
--
ALTER TABLE `placement`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `ps_data`
--
ALTER TABLE `ps_data`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `ps_ovr_data`
--
ALTER TABLE `ps_ovr_data`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `students_fa_data`
--
ALTER TABLE `students_fa_data`
  MODIFY `id` int(15) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `student_academic_data`
--
ALTER TABLE `student_academic_data`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `student_education`
--
ALTER TABLE `student_education`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `student_full_stack_map`
--
ALTER TABLE `student_full_stack_map`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `suggestion`
--
ALTER TABLE `suggestion`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `users_profile`
--
ALTER TABLE `users_profile`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
