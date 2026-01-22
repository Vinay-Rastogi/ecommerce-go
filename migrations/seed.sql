-- ============================
-- USERS
-- ============================
INSERT INTO users (id, name, email, phone, address) VALUES
('11111111-1111-1111-1111-111111111111', 'Koushik', 'koushik@gmail.com', '9876543210', 'Bangalore'),
('22222222-2222-2222-2222-222222222222', 'Vinay', 'vinay@gmail.com', '9123456789', 'Noida');

-- ============================
-- STORES
-- ============================
INSERT INTO stores (id, name, status) VALUES
('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Tech Store', 'active'),
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Grocery Store', 'active');

-- ============================
-- PRODUCTS
-- ============================
INSERT INTO products (id, store_id, name, price, availability) VALUES
('aaaaaaaa-1111-1111-1111-111111111111', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Laptop', 75000.00, true),
('bbbbbbbb-2222-2222-2222-222222222222', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Wireless Mouse', 1500.00, true),
('cccccccc-3333-3333-3333-333333333333', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Rice Bag (5kg)', 1200.00, true),
('dddddddd-4444-4444-4444-444444444444', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Cooking Oil (1L)', 180.00, true);

-- ============================
-- ORDERS
-- ============================
INSERT INTO orders (id, user_id, status) VALUES
('eeeeeeee-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 'created'),
('ffffffff-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', 'paid');

-- ============================
-- ORDER ITEMS
-- ============================
INSERT INTO order_items (order_id, product_id, quantity) VALUES
('eeeeeeee-1111-1111-1111-111111111111', 'aaaaaaaa-1111-1111-1111-111111111111', 1),
('eeeeeeee-1111-1111-1111-111111111111', 'bbbbbbbb-2222-2222-2222-222222222222', 2),
('ffffffff-2222-2222-2222-222222222222', 'cccccccc-3333-3333-3333-333333333333', 1);

-- ============================
-- PAYMENTS
-- ============================
INSERT INTO payments (id, order_id, amount, status) VALUES
('99999999-9999-9999-9999-999999999999', 'ffffffff-2222-2222-2222-222222222222', 1200.00, 'success');

-- ============================
-- SUBSCRIPTIONS
-- ============================
INSERT INTO subscriptions (id, user_id, product_id, start_date, status) VALUES
('88888888-8888-8888-8888-888888888888', '11111111-1111-1111-1111-111111111111', 'dddddddd-4444-4444-4444-444444444444', CURRENT_DATE, 'active'),
('77777777-7777-7777-7777-777777777777', '22222222-2222-2222-2222-222222222222', 'cccccccc-3333-3333-3333-333333333333', CURRENT_DATE, 'paused');
