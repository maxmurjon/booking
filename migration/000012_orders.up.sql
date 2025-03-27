CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,               -- Buyurtma ID
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- Foydalanuvchi ID
    course_id INT REFERENCES courses(id) ON DELETE CASCADE,  -- Kurs ID
    status VARCHAR(50) DEFAULT 'pending',  -- Buyurtma holati: 'pending', 'approved', 'rejected'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Yaratilgan sana
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Yangilangan sana
);