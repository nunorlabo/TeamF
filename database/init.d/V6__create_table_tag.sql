CREATE DATABASE IF NOT EXISTS nexus_db;
USE nexus_db;
CREATE TABLE tag (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(255),
  type VARCHAR(255),
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;