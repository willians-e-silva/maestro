CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    docker_image VARCHAR(255),
    docker_tag VARCHAR(255),
    port VARCHAR(50),
    env_vars JSONB,
    volumes TEXT[],
    network_mode VARCHAR(50),
    restart_policy VARCHAR(50),
    cpu_limit INT,
    memory_limit INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);