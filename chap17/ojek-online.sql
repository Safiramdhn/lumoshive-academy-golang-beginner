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
  "city" varchar(255),
  "district" varchar(255),
  "neighborhood" varchar(255),
  "street_name" varchar(255),
  "order_time" time,
  "order_date" date,
  "order_status" order_status_enum default "on_the_way",
  "status" status_enum DEFAULT 'active',
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW())
);

ALTER TABLE "customer" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "driver" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("driver_id") REFERENCES "driver" ("id");

INSERT INTO users (email, password, login_time, logout_time) VALUES
    ('john.doe@example.com', 'password123', '2024-10-01 08:00', NULL),
    ('jane.smith@example.com', 'password123', '2024-10-01 09:30', NULL),
    ('alice.johnson@example.com', 'password123', NULL, '2024-10-02 17:30'),
    ('bob.williams@example.com', 'password123', NULL, '2024-10-03 21:00'),
    ('sarah.brown@example.com', 'password123', '2024-10-04 06:45', NULL),
    ('tom.hanks@example.com', 'password123', '2024-10-01 08:15', NULL),
    ('emma.watson@example.com', 'password123', NULL, '2024-10-01 17:45'),
    ('robert.downey@example.com', 'password123', NULL, '2024-10-02 18:00'),
    ('chris.evans@example.com', 'password123', '2024-10-03 12:15', NULL),
    ('scarlett.j@example.com', 'password123', '2024-10-04 07:00', NULL);
    
 insert into customer (user_id, first_name, last_name) values
 	(1, 'John', 'Doe'),
 	(2, 'Jane', 'Smith'),
 	(3, 'Alice', 'Johnson'),
 	(4, 'Bob', 'Williams'),
 	(5, 'Sarah', 'Brown');
 	
insert into driver (user_id, first_name, last_name) values
 	(6, 'Tom', 'Hanks'),
 	(7, 'Emma', 'Watson'),
 	(8, 'Robert', 'Downey'),
 	(9, 'Chris', 'Evans'),
 	(10, 'Scarlett', 'Johansson');
 	
INSERT INTO orders 
    (customer_id, driver_id, order_date, order_time, city, district, neighborhood, street_name) 
VALUES
    (1, 2, '2024-08-01', '08:30', 'New York', 'Manhattan', 'Midtown', '5th Avenue'),
    (2, 1, '2024-08-01', '09:15', 'New York', 'Manhattan', 'Midtown', '5th Avenue'),
    (3, 3, '2024-08-02', '10:45', 'Los Angeles', 'Hollywood', 'Central LA', 'Sunset Boulevard'),
    (1, 4, '2024-09-03', '15:20', 'Chicago', 'Lincoln Park', 'North Side', 'Clark Street'),
    (4, 5, '2024-09-04', '18:00', 'Chicago', 'Lincoln Park', 'North Side', 'Clark Street'),
    (5, 2, '2024-09-05', '20:00', 'New York', 'Manhattan', 'Midtown', '5th Avenue'),
    (2, 1, '2024-10-06', '11:30', 'Chicago', 'Hyde Park', 'South Side', 'University Avenue'),
    (3, 3, '2024-10-06', '13:15', 'Los Angeles', 'Venice', 'West LA', 'Venice Boulevard'),
    (1, 4, '2024-10-07', '16:45', 'San Francisco', 'SoMa', 'South of Market', 'Howard Street'),
    (5, 5, '2024-10-08', '22:30', 'Los Angeles', 'Venice', 'West LA', 'Venice Boulevard');
   
--dapat melihat total order setiap bulan
SELECT DATE_TRUNC('month', order_date) AS month,
       COUNT(id) AS total_orders
FROM orders
GROUP BY DATE_TRUNC('month', order_date)
ORDER BY month desc;

--dapat melihat customer yang sering order tiap bulan (tampilkan namanya)  
SELECT DATE_TRUNC('month', order_date) AS month,
       concat(c.first_name, ' ', c.last_name) AS customer_name,
       COUNT(o.id) AS total_orders
FROM orders o
JOIN customer c ON o.customer_id = c.id
GROUP BY month,
         customer_name
HAVING COUNT(o.id) > 0 -- replace with the desired threshold
ORDER BY month,
         total_orders DESC;

--dapat melihat daerah mana saja yang banyak ordernya
SELECT 
city,
district,
neighborhood,
street_name,
COUNT(id) AS total_orders
FROM orders
GROUP BY city,
         district,
         neighborhood,
         street_name
ORDER BY total_orders DESC;


--dapat melihat jumlah customer yang masih login dan logout
SELECT COUNT(c.id) AS total_customer_login
FROM customer c
JOIN users u ON c.user_id = u.id
WHERE u.login_time IS NOT NULL;


SELECT COUNT(d.id) AS total_driver_logout
FROM driver d
JOIN users u ON d.user_id = u.id
where u.logout_time is not null

--dapat melihat driver yang rajin mengambil order setiap bulan

SELECT
    o.driver_id,
    CONCAT(d.first_name, ' ', d.last_name) AS driver_full_name,
    COUNT(o.driver_id) AS total_order,
    DATE_TRUNC('month', o.order_time) AS month
FROM
    orders o
JOIN
    driver d ON o.driver_id = d.id
GROUP BY
    month, o.driver_id, d.first_name, d.last_name
ORDER BY
    total_order DESC;

SELECT D.ID,
       CONCAT(D.FIRST_NAME, ' ', D.LAST_NAME) AS DRIVER_FULL_NAME,
  ( SELECT COUNT(O.ID)
   FROM ORDERS O
   WHERE O.DRIVER_ID = D.ID
     AND O.ORDER_TIME BETWEEN '2023-10-01 00:00:00' AND '2023-10-31 23:59:59' )
FROM DRIVER D;

-- dapat melihat pukul berapa saja order yang ramai dan sepi
SELECT 
    EXTRACT(HOUR FROM order_time) AS hour,
    COUNT(id) AS total_orders
FROM orders
GROUP BY hour
ORDER BY total_orders DESC;

SELECT 
    TO_CHAR(order_time, 'HH24:00') || ' - ' || TO_CHAR(order_time + INTERVAL '1 hour', 'HH24:00') AS hour_range,
    COUNT(id) AS total_orders
FROM orders
GROUP BY hour_range
ORDER BY total_orders DESC;

SELECT 
    TO_CHAR(order_time, 'HH24:00') AS hour,
    COUNT(id) AS total_orders
FROM orders
GROUP BY hour
ORDER BY total_orders DESC;

select * from customer
select * from driver