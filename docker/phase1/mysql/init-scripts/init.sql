-- create_database.sql
    CREATE DATABASE IF NOT EXISTS mydb;
    USE mydb;
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL UNIQUE
    );
    INSERT INTO users(name) VALUES ("thuy");