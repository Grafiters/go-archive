CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    role varchar(50) NOT NULL,
    action varchar(50) NOT NULL,
    verb varchar(50) NOT NULL,
    path varchar(150) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);