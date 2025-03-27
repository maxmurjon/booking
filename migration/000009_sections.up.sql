CREATE TABLE IF NOT EXISTS sections (
    id SERIAL PRIMARY KEY,
    course_id INT REFERENCES courses(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    "order" INT NOT NULL,                     -- Tartib
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);