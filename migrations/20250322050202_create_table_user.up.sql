CREATE TABLE auth.users (
   id           UUID PRIMARY KEY NOT NULL,
   name         VARCHAR(150) NOT NULL,
   email        VARCHAR(255) NOT NULL UNIQUE,
   password     VARCHAR(150),
   profile      VARCHAR(150),
   role         VARCHAR(20) NOT NULL DEFAULT 'visitor' CHECK (role IN ('admin', 'visitor')),
   provider     VARCHAR(50) NOT NULL CHECK (provider IN ('google', 'twitter')),
   provider_id  VARCHAR(100),
   created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
