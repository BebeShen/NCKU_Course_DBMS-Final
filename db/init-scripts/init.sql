CREATE TABLE `Course` (
  `id` varchar(10) NOT NULL,
  `date` DATE NOT NULL DEFAULT (CURRENT_DATE),
  `name` varchar(100) DEFAULT '',
  `classroom` varchar(100) DEFAULT '',
  `seats` int unsigned DEFAULT '0',
  `cols` int unsigned DEFAULT '0',
  PRIMARY KEY (`id`, `date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `Student` (
  `id` varchar(100) NOT NULL,
  `name` varchar(100) DEFAULT '',
  `password` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `Seat` (
  `course_id` varchar(10) NOT NULL,
  `course_date` DATE NOT NULL,
  `seat_id` int unsigned NOT NULL,
  `reserved_by` varchar(100) NOT NULL,
  PRIMARY KEY (`course_id`, `course_date`, `seat_id`),
  FOREIGN KEY (`course_id`,`course_date`) REFERENCES `Course`(`id`,`date`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`reserved_by`) REFERENCES `Student`(`id`)ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
