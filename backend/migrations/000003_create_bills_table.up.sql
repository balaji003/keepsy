CREATE TABLE IF NOT EXISTS bills (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    category_id INT, -- Optional: link to category
    name VARCHAR(255) NOT NULL,
    file_url VARCHAR(2048) NOT NULL,
    file_type VARCHAR(50), -- pdf, image/jpeg, etc.
    amount DECIMAL(10, 2), -- Optional: amount on the bill
    due_date DATE,         -- Optional: due date
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);
