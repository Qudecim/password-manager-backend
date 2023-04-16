CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;


CREATE TABLE secrets
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    `data` TEXT NOT NULL,
    CONSTRAINT secrets_user_fk
    FOREIGN KEY (user_id) REFERENCES users (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;

CREATE TABLE devices
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    uuid VARCHAR(255) NOT NULL,
    public_key VARCHAR(255) NOT NULL,
    CONSTRAINT devices_user_fk
    FOREIGN KEY (user_id) REFERENCES users (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;

CREATE TABLE tokens
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    token VARCHAR(255) NOT NULL,
    expired_at DATETIME NOT NULL,
    is_deleted TINYINT DEFAULT 0,
    CONSTRAINT tokens_user_fk
    FOREIGN KEY (user_id) REFERENCES users (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;