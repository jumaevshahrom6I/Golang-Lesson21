-- Homework 1: SQL CRUD basics (DB: coursedb)

-- 1. Create table
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    price INT NOT NULL,
    pages INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 2. Insert data
INSERT INTO books (title, author, price, pages) VALUES
('Clean Code', 'Robert Martin', 120, 464),
('The Go Programming Language', 'Alan Donovan', 150, 380),
('Go in Action', 'William Kennedy', 90, 264),
('Learning Go', 'Jon Bodner', 110, 375),
('Concurrency in Go', 'Katherine Cox-Buday', 130, 238),
('Designing Data-Intensive Applications', 'Martin Kleppmann', 200, 616),
('The Pragmatic Programmer', 'Andrew Hunt', 95, 352),
('Refactoring', 'Martin Fowler', 140, 448),
('Code Complete', 'Steve McConnell', 180, 960),
('Head First SQL', 'Lynn Beighley', 70, 592),
('SQL Cheat Sheet', 'John Smith', 30, 80),
('PostgreSQL Up and Running', 'Regina Obe', 85, 328),
('Introduction to Algorithms', 'Thomas Cormen', 250, 1312),
('Go Pocket Guide', 'Mike Wilson', 45, 90),
('The C Programming Language', 'Brian Kernighan', 60, 272);

-- 3. Select data
SELECT * FROM books;
SELECT title, author FROM books;
SELECT title, price FROM books;

-- 4. WHERE
SELECT * FROM books WHERE author = 'Robert Martin';
SELECT * FROM books WHERE price > 100;
SELECT * FROM books WHERE price < 50;
SELECT * FROM books WHERE pages > 500;
SELECT * FROM books WHERE price BETWEEN 100 AND 200;

-- 5. AND / OR
SELECT * FROM books WHERE price > 100 AND pages > 300;
SELECT * FROM books WHERE price < 100 OR pages < 200;
SELECT * FROM books WHERE author = 'Martin Fowler' AND price > 150;

-- 6. ORDER BY
SELECT * FROM books ORDER BY price ASC;
SELECT * FROM books ORDER BY price DESC;
SELECT * FROM books ORDER BY pages DESC;
SELECT * FROM books ORDER BY title ASC;

-- 7. LIMIT
SELECT * FROM books ORDER BY price DESC LIMIT 5;
SELECT * FROM books ORDER BY pages DESC LIMIT 3;

-- 8. UPDATE + verify
UPDATE books SET price = 125 WHERE id = 1;
SELECT * FROM books WHERE id = 1;

UPDATE books SET author = 'Alan A. A. Donovan' WHERE id = 2;
SELECT * FROM books WHERE id = 2;

UPDATE books SET title = 'Go in Action 2nd Ed', price = 99, pages = 300 WHERE id = 3;
SELECT * FROM books WHERE id = 3;

-- 9. UPDATE multiple rows
UPDATE books SET price = 100 WHERE price < 100;
SELECT * FROM books ORDER BY price;

-- 10. DELETE + verify
DELETE FROM books WHERE id = 15;
SELECT * FROM books;

DELETE FROM books WHERE price < 50;
SELECT * FROM books;

DELETE FROM books WHERE pages < 100;
SELECT * FROM books;

-- 11. COUNT
SELECT COUNT(*) FROM books;
SELECT COUNT(*) FROM books WHERE price > 100;
SELECT COUNT(*) FROM books WHERE author = 'Robert Martin';

-- 12. LIKE
SELECT * FROM books WHERE title LIKE '%Go%';
SELECT * FROM books WHERE author LIKE '%Martin%';

-- 13. Custom queries
SELECT * FROM books ORDER BY price ASC LIMIT 3;
SELECT * FROM books WHERE price BETWEEN 90 AND 150 ORDER BY price;
SELECT title FROM books WHERE pages > 400 ORDER BY pages DESC;
SELECT AVG(price) FROM books;
SELECT MAX(pages), MIN(pages) FROM books;
SELECT * FROM books WHERE title ILIKE '%sql%';
SELECT * FROM books WHERE author LIKE 'M%';
UPDATE books SET price = price + 10 WHERE pages > 500;
UPDATE books SET author = 'Unknown' WHERE author = 'John Smith';
DELETE FROM books WHERE author = 'Unknown';
SELECT * FROM books ORDER BY created_at DESC LIMIT 5;
