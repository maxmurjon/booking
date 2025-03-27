CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price_tutor NUMERIC(10, 2) NOT NULL,      -- Ustoz bilan narxi
    price_no_tutor NUMERIC(10, 2) NOT NULL,  -- Ustozsiz narxi
    image_url TEXT,
    video_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);