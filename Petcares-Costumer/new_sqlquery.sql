CREATE DATABASE petcares;

USE petcares;

CREATE TABLE animal (
  Animal_id VARCHAR(10) NOT NULL,
  Customer_id VARCHAR(10) NOT NULL,
  Animal_name VARCHAR(50) NOT NULL,
  Species VARCHAR(50) NOT NULL,
  Age INT(11) NOT NULL,
  Gender CHAR(1) NOT NULL,
  PRIMARY KEY (Animal_id)
);

CREATE TABLE customer (
  Customer_id VARCHAR(10) NOT NULL,
  Firstname VARCHAR(50) NOT NULL,
  Lastname VARCHAR(50) NOT NULL,
  Email VARCHAR(100) NOT NULL,
  Phone VARCHAR(15) NOT NULL,
  PRIMARY KEY (Customer_id)
);

CREATE TABLE hotel (
  Hotel_id VARCHAR(10) NOT NULL,
  Name VARCHAR(50) NOT NULL,
  Location VARCHAR(100) NOT NULL,
  Capacity INT(11) NOT NULL,
  PRIMARY KEY (Hotel_id)
);

CREATE TABLE doctor (
  Doctor_id VARCHAR(10) NOT NULL,
  Firstname VARCHAR(50) NOT NULL,
  Lastname VARCHAR(50) NOT NULL,
  Specialty VARCHAR(50) NOT NULL,
  Phone VARCHAR(15) NOT NULL,
  PRIMARY KEY (Doctor_id)
);

CREATE TABLE reservation (
  Reservation_id VARCHAR(10) NOT NULL,
  Animal_id VARCHAR(10) NOT NULL,
  Hotel_id VARCHAR(10) NOT NULL,
  Doctor_id VARCHAR(10) NOT NULL,
  Start_date DATE NOT NULL,
  End_date DATE NOT NULL,
  PRIMARY KEY (Reservation_id),
  FOREIGN KEY (Animal_id) REFERENCES animal(Animal_id),
  FOREIGN KEY (Hotel_id) REFERENCES hotel(Hotel_id),
  FOREIGN KEY (Doctor_id) REFERENCES doctor(Doctor_id)
);

CREATE TABLE doctor_history (
  History_id INT AUTO_INCREMENT PRIMARY KEY,
  Doctor_id VARCHAR(10) NOT NULL,
  Doctor_name VARCHAR(100) NOT NULL,
  Animal_name VARCHAR(50) NOT NULL,
  Reservation_id VARCHAR(10) NOT NULL,
  Start_date DATE NOT NULL,
  End_date DATE NOT NULL,
  FOREIGN KEY (Doctor_id) REFERENCES doctor(Doctor_id)
);

CREATE TABLE hotel_history (
  History_id INT AUTO_INCREMENT PRIMARY KEY,
  Hotel_id VARCHAR(10) NOT NULL,
  Hotel_name VARCHAR(100) NOT NULL,
  Animal_name VARCHAR(50) NOT NULL,
  Reservation_id VARCHAR(10) NOT NULL,
  Start_date DATE NOT NULL,
  End_date DATE NOT NULL,
  FOREIGN KEY (Hotel_id) REFERENCES hotel(Hotel_id)
);
