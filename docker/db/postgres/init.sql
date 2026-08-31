-- 3. Connect to the newly created database
\c test_db
-- 4. Create a new schema inside the new database
CREATE SCHEMA test_schema AUTHORIZATION test_user;
-- 5. Grant permissions to the user on that schema
GRANT ALL PRIVILEGES ON SCHEMA test_schema TO test_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA test_schema TO test_user;

CREATE TABLE IF NOT EXISTS test_schema.customers (
    customer_id SERIAL Primary Key,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS test_schema.products (
    product_id SERIAL Primary Key,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS test_schema.orders (
    order_id	  SERIAL Primary key,
    customer_id	  INT REFERENCES test_schema.customers(customer_id),
    product_id	  INT REFERENCES test_schema.products(product_id),
    quantity	  INT DEFAULT 1,
    total_amount  DECIMAL(10, 2),
    order_date	  TIMESTAMP DEFAULT NOW()
);

INSERT INTO test_schema.customers (first_name, last_name, email)
VALUES ('Le', 'Loan', 'leloan@gmail.com'),
       ('Vuong', 'Kieu', 'vkieu@gmail.com'),
       ('Vo', 'Trong', 'vtrong@gmail.com');

INSERT INTO test_schema.products (name, description, price)
VALUES ('Laptop', 'Laptop HP core i5, i3, i7', 999.99),
       ('Mouse', 'An external wireless mouse improves laptop productivity and precision over a built-in trackpad.', 25.00),
       ('Keyboard', 'An external keyboard for a laptop improves typing comfort, increases speed, and protects the built-in keys from wear', 75.00);

INSERT INTO test_schema.orders (customer_id, product_id, quantity, total_amount, order_date)
VALUES (1, 1, 1, 999.99, NOW()),
       (2, 2, 1, 25.00, NOW()),
       (3, 3, 1, 75.00, '2026-08-28 14:30:00');

CREATE TABLE IF NOT EXISTS test_schema.attendee (
    attendee_id SERIAL Primary Key,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS test_schema.curriculum (
    curriculum_id SERIAL Primary Key,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS test_schema.enroll (
    enroll_id	  SERIAL Primary key,
    attendee_id	  INT REFERENCES test_schema.attendee(attendee_id),
    curriculum_id	  INT REFERENCES test_schema.curriculum(curriculum_id),
    enroll_date	  TIMESTAMP DEFAULT NOW()
);

INSERT INTO test_schema.attendee (first_name, last_name, email)
VALUES ('Le', 'Loan', 'leloan@gmail.com'),
       ('Vuong', 'Kieu', 'vkieu@gmail.com'),
       ('Vo', 'Trong', 'vtrong@gmail.com');

INSERT INTO test_schema.curriculum (name, description, price)
VALUES ('Foundational Knowledge', 'Introduction to basic concepts', 999.99),
       ('AI Techniques and Applications', 'Hands-on experience with generative AI tools', 25.00),
       ('Ethics and Human-Centered Thinking', 'Emphasizing responsible AI use, personal data protection, copyright awareness', 75.00);

INSERT INTO test_schema.enroll (attendee_id, curriculum_id, enroll_date)
VALUES (1, 1, NOW()),
       (2, 2, NOW()),
       (3, 3, '2026-08-28 14:30:00');