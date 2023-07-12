
-- +migrate Up
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    mail_address VARCHAR(255) NOT NULL,
    signup_flag CHAR(1) NOT NULL,
    verify_number INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    active_task VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +migrate Down
DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS users;