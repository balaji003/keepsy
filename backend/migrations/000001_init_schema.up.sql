CREATE TABLE IF NOT EXISTS keepsy_users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(50),
    uuid VARCHAR(36) NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS keepsy_categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    parent_id INT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES keepsy_categories(id)
);

CREATE TABLE IF NOT EXISTS keepsy_products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    category_id INT,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255),
    model VARCHAR(255),
    location VARCHAR(255), -- "Guest Bedroom", "Kitchen"
    purchase_date DATE,
    warranty_end_date DATE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES keepsy_users(id),
    FOREIGN KEY (category_id) REFERENCES keepsy_categories(id)
);

CREATE TABLE IF NOT EXISTS keepsy_product_purchase_details (
    product_id INT PRIMARY KEY,
    shop_name VARCHAR(255),
    shop_address TEXT,
    contact_person VARCHAR(255),
    contact_number VARCHAR(50),
    order_id VARCHAR(255),
    delivery_status VARCHAR(50),
    FOREIGN KEY (product_id) REFERENCES keepsy_products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS keepsy_bills (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    file_url TEXT NOT NULL,
    ocr_text TEXT,
    uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES keepsy_products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS keepsy_reminders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    due_date DATETIME NOT NULL,
    is_completed BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES keepsy_users(id),
    FOREIGN KEY (product_id) REFERENCES keepsy_products(id)
);
