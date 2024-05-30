CREATE TABLE storage_drinks (
    id SERIAL PRIMARY KEY,
    public_id VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL,
    stock INT NOT NULL,
    created_at timestamp NOT NULL,
    created_by VARCHAR(100) NOT NULL,
    updated_at timestamp NOT NULL,
    updated_by VARCHAR(100) NOT NULL,
    deleted_at timestamp NULL,
    deleted_by VARCHAR(100) NULL
)