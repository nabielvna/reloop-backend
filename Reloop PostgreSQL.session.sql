-- Check users count
SELECT COUNT(*) as user_count FROM users;

-- Check categories
SELECT id, name, is_active FROM categories ORDER BY id;

-- Check items with prices
SELECT id, name, price, status FROM items ORDER BY price DESC LIMIT 5;

-- Check sellers
SELECT id, business_name, verification_status FROM sellers;