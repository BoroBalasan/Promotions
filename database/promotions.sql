CREATE TABLE `promotions`
(
    id varchar(255) NOT NULL PRIMARY KEY,
    price DECIMAL(8,6),
    expiration_date DATETIME
);