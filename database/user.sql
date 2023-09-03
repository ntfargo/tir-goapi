CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
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
    CONSTRAINT chk_full_name_length CHECK (LENGTH(full_name) <= 100),
    CONSTRAINT chk_email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,4}$'),
    CONSTRAINT chk_username_format CHECK (username ~* '^[A-Za-z0-9_]+$')
);

CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_api_key ON users (api_key);