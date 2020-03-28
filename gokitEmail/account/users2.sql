create table users2 (
    id VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(50) NOT NULL,
    body VARCHAR(1024) NOT NULL,
    toEmail VARCHAR(100) NOT NUll
);

