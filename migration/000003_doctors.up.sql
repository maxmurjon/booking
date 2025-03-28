CREATE TABLE IF NOT EXISTS doctors (
    id UUID PRIMARY KEY,
    user_id UUID UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    specialty VARCHAR(100) NOT NULL,  
    work_start TIME NOT NULL,         
    work_end TIME NOT NULL,          
    created_at TIMESTAMP DEFAULT NOW()
);