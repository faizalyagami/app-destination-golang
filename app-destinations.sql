CREATE TABLE destinations (
    id SERIAL PRIMARY KEY,
    country VARCHAR(100),
    destination VARCHAR(100),
    description TEXT,
    price NUMERIC,
    tour_duration VARCHAR(50),
    rating INT,
    images TEXT,
    people VARCHAR(50),
    flag_image TEXT,
    price_discount NUMERIC
);

INSERT INTO destinations (country, destination, description, price, tour_duration, rating, images, people, flag_image, price_discount)
VALUES ('Brazil', 'Amazon', 'lorem ipsum', 100000, '8 days', 5, 'https://localhost/images.png', '30 people going', 'https://localhost/brazil.jpeg', 10000);
