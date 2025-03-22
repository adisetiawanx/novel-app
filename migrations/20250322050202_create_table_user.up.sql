CREATE TABLE users (
    id           CHAR(36) PRIMARY KEY NOT NULL,
    name         VARCHAR(255) NOT NULL,
    email        VARCHAR(255) NOT NULL UNIQUE,
    password     VARCHAR(255) NOT NULL,
    phone        VARCHAR(255) NOT NULL,
    profile      VARCHAR(255),
    role      ENUM('admin', 'visitor') NOT NULL DEFAULT 'visitor',
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;