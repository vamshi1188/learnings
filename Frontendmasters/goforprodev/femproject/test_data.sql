-- Insert test user if it doesn't exist
INSERT INTO users (username, email, password_hash)
SELECT 'testuser', 'test@example.com', 'password123'
WHERE NOT EXISTS (
    SELECT 1 FROM users WHERE username = 'testuser'
);
