
INSERT INTO users (id, name, email, phone, address) VALUES
(gen_random_uuid(),'Koushik Kumar','koushik@gmail.com','9876543210','Bangalore'),
(gen_random_uuid(),'Vinay Rastogi','vinay@gmail.com','9123456789','Noida'),
(gen_random_uuid(),'Rahul Sharma','rahul@gmail.com','9988776655','Delhi'),
(gen_random_uuid(),'Ankit Verma','ankit@gmail.com','9090909090','Pune'),
(gen_random_uuid(),'Neha Singh','neha@gmail.com','9887766554','Jaipur'),
(gen_random_uuid(),'Aman Gupta','aman@gmail.com','9776655443','Indore'),
(gen_random_uuid(),'Rohit Jain','rohit@gmail.com','9665544332','Mumbai'),
(gen_random_uuid(),'Pooja Mehta','pooja@gmail.com','9554433221','Ahmedabad'),
(gen_random_uuid(),'Saurabh Mishra','saurabh@gmail.com','9443322110','Lucknow'),
(gen_random_uuid(),'Nikita Agarwal','nikita@gmail.com','9332211009','Kolkata');

INSERT INTO stores (id, name, status) VALUES
('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa','Apple World','active'),
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb','Samsung Plaza','active'),
('cccccccc-cccc-cccc-cccc-cccccccccccc','Laptop Hub','active'),
('dddddddd-dddd-dddd-dddd-dddddddddddd','Accessory Store','active'),
('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee','TV SuperMart','active'),
('ffffffff-ffff-ffff-ffff-ffffffffffff','Budget Electronics','active');

-- =========================================
-- PRODUCTS (SEARCH-READY DATASET)
-- =========================================

INSERT INTO products
(id, store_id, name, description, brand, category, price, rating, availability, created_at)
VALUES

-- ================== APPLE ==================
(gen_random_uuid(), 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
 'iPhone 14',
 'Apple smartphone with A15 Bionic chip and OLED display',
 'Apple', 'Mobile', 69999, 4.6, true, NOW()),

(gen_random_uuid(), 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
 'iPhone 13',
 'Apple smartphone with powerful camera and long battery life',
 'Apple', 'Mobile', 59999, 4.5, false, NOW()),

(gen_random_uuid(), 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
 'MacBook Air M1',
 'Lightweight Apple laptop with M1 processor',
 'Apple', 'Laptop', 89999, 4.8, true, NOW()),

(gen_random_uuid(), 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
 'MacBook Pro M2',
 'Professional Apple laptop with M2 chip',
 'Apple', 'Laptop', 149999, 4.9, false, NOW()),

-- ================== SAMSUNG ==================
(gen_random_uuid(), 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
 'Samsung Galaxy S23',
 'Samsung flagship Android smartphone',
 'Samsung', 'Mobile', 74999, 4.6, true, NOW()),

(gen_random_uuid(), 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
 'Samsung Galaxy A54',
 'Mid-range Samsung phone with AMOLED display',
 'Samsung', 'Mobile', 34999, 4.3, true, NOW()),

(gen_random_uuid(), 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
 'Samsung Smart TV 55',
 '55 inch 4K UHD Smart TV',
 'Samsung', 'TV', 79999, 4.5, false, NOW()),

-- ================== DELL ==================
(gen_random_uuid(), 'cccccccc-cccc-cccc-cccc-cccccccccccc',
 'Dell XPS 13',
 'Premium ultrabook with InfinityEdge display',
 'Dell', 'Laptop', 99999, 4.7, true, NOW()),

(gen_random_uuid(), 'cccccccc-cccc-cccc-cccc-cccccccccccc',
 'Dell Inspiron 15',
 'Affordable laptop for everyday use',
 'Dell', 'Laptop', 55999, 4.1, true, NOW()),

(gen_random_uuid(), 'cccccccc-cccc-cccc-cccc-cccccccccccc',
 'Dell 27 Inch Monitor',
 'IPS display monitor for professionals',
 'Dell', 'Accessories', 22999, 4.4, false, NOW()),

-- ================== SONY ==================
(gen_random_uuid(), 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
 'Sony WH-1000XM5',
 'Industry leading noise cancelling headphones',
 'Sony', 'Accessories', 29999, 4.8, true, NOW()),

(gen_random_uuid(), 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
 'Sony Bravia 55',
 '4K Ultra HD Smart LED TV',
 'Sony', 'TV', 82999, 4.6, false, NOW()),

-- ================== HP ==================
(gen_random_uuid(), 'cccccccc-cccc-cccc-cccc-cccccccccccc',
 'HP Pavilion 14',
 'Laptop for students and professionals',
 'HP', 'Laptop', 62999, 4.2, true, NOW()),

(gen_random_uuid(), 'cccccccc-cccc-cccc-cccc-cccccccccccc',
 'HP DeskJet Printer',
 'Wireless inkjet printer',
 'HP', 'Accessories', 8999, 4.0, false, NOW()),

-- ================== BOAT ==================
(gen_random_uuid(), 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
 'Boat Rockerz 550',
 'Wireless over-ear headphones',
 'Boat', 'Accessories', 1999, 4.1, true, NOW()),

(gen_random_uuid(), 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
 'Boat Smart Watch',
 'Fitness tracking smartwatch',
 'Boat', 'Accessories', 3499, 4.0, false, NOW());

INSERT INTO orders (id, user_id, status)
SELECT gen_random_uuid(), id, 'created'
FROM users
LIMIT 5;

INSERT INTO order_items (order_id, product_id, quantity)
SELECT
 o.id,
 p.id,
 1
FROM orders o
JOIN products p ON p.availability = true
LIMIT 15;

INSERT INTO payments (id, order_id, amount, status)
SELECT
 gen_random_uuid(),
 o.id,
 (SELECT price FROM products ORDER BY RANDOM() LIMIT 1),
 'success'
FROM orders o
LIMIT 3;


INSERT INTO subscriptions (id, user_id, product_id, start_date, status)
SELECT
 gen_random_uuid(),
 u.id,
 p.id,
 CURRENT_DATE,
 'active'
FROM users u
JOIN products p ON p.category = 'Accessories'
LIMIT 5;
