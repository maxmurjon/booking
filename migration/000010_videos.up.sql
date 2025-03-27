CREATE TABLE IF NOT EXISTS videos (
    id SERIAL PRIMARY KEY,
    section_id INT REFERENCES sections(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    url TEXT NOT NULL,                        -- MinIO URL
    duration INT NOT NULL,                    -- Video davomiyligi (sekundda)
    "order" INT NOT NULL,                     -- Tartib
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);