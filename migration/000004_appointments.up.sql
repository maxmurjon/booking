CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    patient_id UUID REFERENCES users(id) ON DELETE CASCADE, 
    doctor_id INT REFERENCES doctors(id) ON DELETE CASCADE,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    status VARCHAR(20) CHECK (status IN ('pending', 'confirmed', 'cancelled')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (doctor_id, appointment_date, appointment_time) -- Bir vaqtga bir nechta bemor yozilishining oldini olish
);
