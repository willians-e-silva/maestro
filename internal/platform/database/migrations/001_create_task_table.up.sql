CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    docker_image VARCHAR(255),
    docker_tag VARCHAR(255),
    port VARCHAR(50),
    network_mode VARCHAR(50),
    restart_policy VARCHAR(50),
    cpu_limit INT,
    memory_limit INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
