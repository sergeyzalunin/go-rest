CREATE TABlE IF NOT EXISTS users (
    id bigserial,
    email varchar not null,
    encrypted_password varchar not null
);