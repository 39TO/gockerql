CREATE TABLE todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    done BOOLEAN NOT NULL,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);