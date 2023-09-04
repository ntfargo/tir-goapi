CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    bio TEXT,
    full_name VARCHAR(255),
    api_key varchar(50) NOT NULL,
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    last_login TIMESTAMPTZ,
    CONSTRAINT chk_role CHECK (role IN ('ADMIN', 'MODERATOR', 'MEMBER')),
    CONSTRAINT chk_password_length CHECK (LENGTH(password) >= 10),
    CONSTRAINT chk_full_name_length CHECK (LENGTH(full_name) <= 100)
);

CREATE INDEX idx_api_key ON users (api_key);