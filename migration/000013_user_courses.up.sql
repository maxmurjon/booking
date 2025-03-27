CREATE TABLE IF NOT EXISTS user_courses (
    id SERIAL PRIMARY KEY,              -- Foydalanuvchi-kurs ID
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- Foydalanuvchi ID
    course_id INT REFERENCES courses(id) ON DELETE CASCADE,  -- Kurs ID
    access_granted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Kirish huquqi berilgan sana
);