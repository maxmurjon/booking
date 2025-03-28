CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    patient_id UUID REFERENCES users(id) ON DELETE CASCADE, 
    doctor_id UUID REFERENCES doctors(id) ON DELETE CASCADE,
    appointment_date DATE NOT NULL,
    appointment_time TIMESTAMP NOT NULL,  -- TIME o‘rniga TIMESTAMP ishlatildi
    status VARCHAR(20) CHECK (status IN ('pending', 'confirmed', 'cancelled')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (doctor_id, appointment_time) -- appointment_date ni olib tashladik
);
w