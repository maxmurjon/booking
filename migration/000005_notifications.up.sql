CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    appointment_id INT REFERENCES appointments(id) ON DELETE CASCADE,
    message TEXT NOT NULL, -- Xabar matni
    sent_at TIMESTAMP DEFAULT NOW() -- Qachon yuborilgan
);