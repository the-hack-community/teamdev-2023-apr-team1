-- +goose Up
CREATE TABLE users (
    user_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE
);

CREATE TABLE locations (
    location_id SERIAL PRIMARY KEY,
    lat FLOAT(24) NOT NULL,
    long FLOAT(24) NOT NULL
);

CREATE TABLE stray_cats (
    cat_id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL REFERENCES users(user_id),
    photo_data TEXT NOT NULL,
    capture_date_time TIMESTAMP NOT NULL,
    location_id INT NOT NULL REFERENCES locations(location_id),
    name VARCHAR(255) NOT NULL,
    features VARCHAR(255) NOT NULL,
    condition VARCHAR(255) NOT NULL
);

CREATE TABLE reactions (
    reaction_id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL REFERENCES users(user_id),
    cat_id VARCHAR(255) NOT NULL REFERENCES stray_cats(cat_id),
    date_time TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE reactions;
DROP TABLE stray_cats;
DROP TABLE locations;
DROP TABLE users;
