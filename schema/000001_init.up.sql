CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;


CREATE TABLE secrets
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    `encryption` TEXT NOT NULL,
    CONSTRAINT secrets_user_fk
    FOREIGN KEY (user_id) REFERENCES users (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;

CREATE TABLE devices
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `uid` VARCHAR(255) NOT NULL UNIQUE,
    public_key VARCHAR(255) NOT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;

CREATE TABLE user_device
(
    user_id INT NOT NULL,
    device_id INT NOT NULL,
    CONSTRAINT userdevice_user_fk
    FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT userdevice_device_fk
    FOREIGN KEY (device_id) REFERENCES devices (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;

CREATE TABLE device_confirmations
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    device_id INT NOT NULL,
    confirmation VARCHAR(255) NOT NULL,
    CONSTRAINT device_confirm_device_fk
    FOREIGN KEY (device_id) REFERENCES devices (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;