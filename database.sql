CREATE SCHEMA supercars;

USE supercars;

CREATE TABLE cars(
    car_id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    car_name VARCHAR(255) NOT NULL,
    car_year YEAR NOT NULL,
    car_price DECIMAL(15,2) NOT NULL,
    car_status TINYINT DEFAULT 1
);

CREATE TABLE orders(
    order_id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    order_passenger_name VARCHAR(255) NOT NULL,
    order_passenger_phone VARCHAR(30) NOT NULL,
    car_id INT UNSIGNED NOT NULL,
    order_pickup_address VARCHAR(500),
    order_destination_address VARCHAR(500),
    order_price DECIMAL(15,2),
    FOREIGN KEY (car_id) REFERENCES cars(car_id) ON DELETE CASCADE
);
