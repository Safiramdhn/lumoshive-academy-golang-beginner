CREATE TYPE "status_enum" AS ENUM (
  'active',
  'deleted'
);

CREATE TYPE "order_status_enum" AS ENUM (
  'on_the_way',
  'with_customer',
  'done'
);

CREATE TABLE "users" (
  "id" serial UNIQUE PRIMARY KEY NOT NULL,
  "email" varchar(255) NOT NULL,
  "password" varchar NOT NULL,
  "login_time" timestamp,
  "logout_time" timestamp DEFAULT null,
  "status" status_enum DEFAULT 'active',
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW())
);

CREATE TABLE "customer" (
  "id" serial UNIQUE PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "first_name" varchar(255),
  "last_name" varchar(255),
  "status" status_enum DEFAULT 'active',
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW())
);

CREATE TABLE "driver" (
  "id" serial UNIQUE PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "first_name" varchar(255),
  "last_name" varchar(255),
  "status" status_enum DEFAULT 'active',
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW())
);

CREATE TABLE "orders" (
  "id" serial UNIQUE PRIMARY KEY NOT NULL,
  "customer_id" int NOT NULL,
  "driver_id" int NOT NULL,
  "pick_up_location" varchar NOT NULL,
  "destination" varchar NOT NULL,
  "order_time" timestamp,
  "order_status" order_status_enum,
  "status" status_enum DEFAULT 'active',
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW())
);

ALTER TABLE "customer" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "driver" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("driver_id") REFERENCES "driver" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("id");


INSERT INTO users (email, password, login_time, logout_time, status, created_at, updated_at) VALUES
('user1@example.com', 'password1', NOW(), NULL, 'active', NOW(), NOW()),
('user2@example.com', 'password2', NOW(), NULL, 'active', NOW(), NOW()),
('user3@example.com', 'password3', NOW(), NULL, 'active', NOW(), NOW()),
('user4@example.com', 'password4', NOW(), NULL, 'active', NOW(), NOW()),
('user5@example.com', 'password5', NOW(), NULL, 'active', NOW(), NOW()),
('user6@example.com', 'password6', NOW(), NULL, 'active', NOW(), NOW()),
('user7@example.com', 'password7', NOW(), NULL, 'active', NOW(), NOW()),
('user8@example.com', 'password8', NOW(), NULL, 'active', NOW(), NOW()),
('user9@example.com', 'password9', NOW(), NULL, 'active', NOW(), NOW()),
('user10@example.com', 'password10', NOW(), NULL, 'active', NOW(), NOW()),
('user11@example.com', 'password11', NOW(), NULL, 'active', NOW(), NOW()),
('user12@example.com', 'password12', NOW(), NULL, 'active', NOW(), NOW()),
('user13@example.com', 'password13', NOW(), NULL, 'active', NOW(), NOW()),
('user14@example.com', 'password14', NOW(), NULL, 'active', NOW(), NOW()),
('user15@example.com', 'password15', NOW(), NULL, 'active', NOW(), NOW()),
('user16@example.com', 'password16', NOW(), NULL, 'active', NOW(), NOW()),
('user17@example.com', 'password17', NOW(), NULL, 'active', NOW(), NOW()),
('user18@example.com', 'password18', NOW(), NULL, 'active', NOW(), NOW()),
('user19@example.com', 'password19', NOW(), NULL, 'active', NOW(), NOW()),
('user20@example.com', 'password20', NOW(), NULL, 'active', NOW(), NOW());

INSERT INTO customer (user_id, first_name, last_name, last_order, status, created_at, updated_at) VALUES
(1, 'Alice', 'Smith', NULL, 'active', NOW(), NOW()),
(2, 'Bob', 'Johnson', NULL, 'active', NOW(), NOW()),
(3, 'Charlie', 'Williams', NULL, 'active', NOW(), NOW()),
(4, 'David', 'Brown', NULL, 'active', NOW(), NOW()),
(5, 'Eve', 'Jones', NULL, 'active', NOW(), NOW()),
(6, 'Frank', 'Garcia', NULL, 'active', NOW(), NOW()),
(7, 'Grace', 'Martinez', NULL, 'active', NOW(), NOW()),
(8, 'Hank', 'Davis', NULL, 'active', NOW(), NOW()),
(9, 'Ivy', 'Rodriguez', NULL, 'active', NOW(), NOW()),
(10, 'Jack', 'Wilson', NULL, 'active', NOW(), NOW()),
(11, 'Kathy', 'Anderson', NULL, 'active', NOW(), NOW()),
(12, 'Leo', 'Thomas', NULL, 'active', NOW(), NOW()),
(13, 'Mona', 'Taylor', NULL, 'active', NOW(), NOW()),
(14, 'Nina', 'Moore', NULL, 'active', NOW(), NOW()),
(15, 'Oscar', 'Jackson', NULL, 'active', NOW(), NOW()),
(16, 'Paul', 'Martin', NULL, 'active', NOW(), NOW()),
(17, 'Quinn', 'Lee', NULL, 'active', NOW(), NOW()),
(18, 'Rita', 'Harris', NULL, 'active', NOW(), NOW()),
(19, 'Sam', 'Clark', NULL, 'active', NOW(), NOW()),
(20, 'Tina', 'Lewis', NULL, 'active', NOW(), NOW());

INSERT INTO driver (user_id, first_name, last_name, last_order, status, created_at, updated_at) VALUES
(11, 'Andy', 'Walker', NULL, 'active', NOW(), NOW()),
(12, 'Bella', 'Hall', NULL, 'active', NOW(), NOW()),
(13, 'Cody', 'Allen', NULL, 'active', NOW(), NOW()),
(14, 'Diana', 'Young',NULL, 'active', NOW(), NOW()),
(15, 'Ethan', 'King', NULL, 'active', NOW(), NOW()),
(16, 'Florence', 'Ward', NULL, 'active', NOW(), NOW()),
(17, 'Gabriel', 'Ross', NULL, 'active', NOW(), NOW()),
(18, 'Hannah', 'Parker', NULL, 'active', NOW(), NOW()),
(19, 'Isaac', 'Morris', NULL, 'active', NOW(), NOW()),
(20, 'Julia', 'Gomez', NULL, 'active', NOW(), NOW()),
(1, 'Kevin', 'White', NULL, 'active', NOW(), NOW()),
(2, 'Lily', 'Hernandez', NULL, 'active', NOW(), NOW()),
(3, 'Matthew', 'Sanchez', NULL, 'active', NOW(), NOW()),
(4, 'Natalie', 'Patel', NULL, 'active', NOW(), NOW()),
(5, 'Oliver', 'Brooks', NULL, 'active', NOW(), NOW()),
(6, 'Penelope', 'Price', NULL, 'active', NOW(), NOW()),
(7, 'Quincy', 'Gonzalez', NULL, 'active', NOW(), NOW()),
(8, 'Rachel', 'Foster', NULL, 'active', NOW(), NOW()),
(9, 'Sophia', 'Reed', NULL, 'active', NOW(), NOW()),
(10, 'Toby', 'Russell', NULL, 'active', NOW(), NOW());

select * from customer
select * from driver

INSERT INTO orders (customer_id, driver_id, pick_up_location, destination, order_time, order_status, status, created_at, updated_at) VALUES
(20, 18, 'Downtown', 'Office Complex', '2023-08-01 08:00:00', 'on_the_way', 'active', NOW(), NOW()),
(19, 1, 'Airport', 'Hotel', '2023-04-01 09:30:00', 'with_customer', 'active', NOW(), NOW()),
(1, 12, 'Train Station', 'Hotel', '2023-03-01 10:15:00', 'done', 'active', NOW(), NOW()),
(16, 14, 'Shopping Mall', 'Hotel', '2023-08-01 11:45:00', 'on_the_way', 'active', NOW(), NOW()),
(4, 19, 'University', 'Residental Neighborhood', '2023-02-01 12:30:00', 'with_customer', 'active', NOW(), NOW()),
(5, 10, 'Hospital', 'Train Station', '2023-10-01 13:00:00', 'done', 'active', NOW(), NOW()),
(12, 6, 'City Park', 'Residental Neighborhood', '2023-02-01 14:00:00', 'on_the_way', 'active', NOW(), NOW()),
(2, 3, 'Office Complex', 'Residental Neighborhood', '2023-03-01 15:30:00', 'with_customer', 'active', NOW(), NOW()),
(3, 16, 'Suburban Area', 'Residental Neighborhood', '2023-05-01 16:45:00', 'done', 'active', NOW(), NOW()),
(19, 7, 'Residential Neighborhood', 'Hospital', '2023-04-01 17:00:00', 'on_the_way', 'active', NOW(), NOW()),
(16, 13, 'Residential Neighborhood', 'Office Complex', '2023-08-02 08:00:00', 'with_customer', 'active', NOW(), NOW()),
(12, 3, 'Airport', 'Train Station', '2023-10-02 09:30:00', 'done', 'active', NOW(), NOW()),
(2, 5, 'Shopping Mall', 'Office Complex', '2023-09-02 10:15:00', 'on_the_way', 'active', NOW(), NOW()),
(5, 17, 'Shopping Mall', 'University', '2023-04-02 11:45:00', 'with_customer', 'active', NOW(), NOW()),
(10, 8, 'Hospital', 'University','2023-10-02 12:30:00', 'done', 'active', NOW(), NOW()),
(3, 8, 'City Park', 'Hotel', '2023-05-02 13:00:00', 'on_the_way', 'active', NOW(), NOW()),
(18, 11, 'University', 'Airport', '2023-12-02 14:00:00', 'with_customer', 'active', NOW(), NOW()),
(10, 14, 'University', 'Shopping Mall', '2023-06-02 15:30:00', 'done', 'active', NOW(), NOW()),
(18, 5, 'University', 'Train Station', '2023-11-02 16:45:00', 'on_the_way', 'active', NOW(), NOW()),
(1, 2, 'Train Station', 'Airport', '2023-5-02 17:00:00', 'with_customer', 'active', NOW(), NOW());

select * from orders


--dapat melihat total order setiap bulan
SELECT
	DATE_TRUNC('month', O.ORDER_TIME) AS MONTH,
	COUNT(O.ID)
FROM
	ORDERS O
GROUP BY
	MONTH
ORDER BY
	MONTH ASC;

--dapat melihat customer yang sering order tiap bulan (tampilkan namanya)  
SELECT
    c.id,
    CONCAT(c.first_name, ' ', c.last_name) AS full_name,
    order_counts.total_orders
FROM
    customer c
JOIN (
    SELECT
        o.customer_id,
        COUNT(o.customer_id) AS total_orders
    FROM
        orders o
    GROUP BY
        o.customer_id
    HAVING
        COUNT(o.customer_id) > 1
) AS order_counts ON c.id = order_counts.customer_id;

--dapat melihat daerah mana saja yang banyak ordernya
SELECT
	O.PICK_UP_LOCATION,
	COUNT(O.PICK_UP_LOCATION) AS TOTAL_ORDER
FROM
	ORDERS O
GROUP BY
	O.PICK_UP_LOCATION
HAVING
	COUNT(O.ID) > 0
ORDER BY
	TOTAL_ORDER DESC

--dapat melihat pukul berapa saja order yang ramai dan sepi
SELECT 
    DATE_TRUNC('hour', o.order_time) AS order_hour, 
    COUNT(o.id) AS order_count
FROM 
    orders o
GROUP BY 
    order_hour
ORDER BY 
    order_hour ASC;


--dapat melihat jumlah customer yang masih login dan logout
SELECT COUNT(c.id) AS cus_login
FROM customer c
JOIN users u ON c.user_id = u.id
WHERE u.login_time IS NOT NULL;

update users set login_time = null, logout_time = now() where id = 1

SELECT COUNT(c.id) AS cus_login
FROM customer c
JOIN users u ON c.user_id = u.id
WHERE u.logout_time IS NOT NULL;


--dapat melihat driver yang rajin mengambil order setiap bulan
-- SELECT 
--     o.driver_id,
--     CONCAT(d.first_name, ' ', d.last_name) AS driver_full_name,
--     COUNT(o.driver_id) AS total_order,
--     DATE_TRUNC('month', o.order_time) AS month
-- FROM 
--     orders o 
-- JOIN 
--     driver d ON o.driver_id = d.id 
-- GROUP BY 
--     month, o.driver_id, d.first_name, d.last_name
-- ORDER BY 
--     total_order DESC;

SELECT
	D.ID,
	CONCAT(D.FIRST_NAME, ' ', D.LAST_NAME) AS DRIVER_FULL_NAME,
	(
		SELECT
			COUNT(O.ID)
		FROM
			ORDERS O
		WHERE
			O.DRIVER_ID = D.ID 
			AND O.ORDER_TIME BETWEEN '2023-10-01 00:00:00' AND '2023-10-31 23:59:59'
	)
FROM
	DRIVER D;