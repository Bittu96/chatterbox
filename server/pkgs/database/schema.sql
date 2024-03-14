CREATE SCHEMA chatterbox_schema;

CREATE TABLE chatterbox_schema.users (
    id serial PRIMARY KEY,
    username varchar(64) UNIQUE NOT NULL,
    email varchar(64) UNIQUE NOT NULL,
    password varchar(64) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL
);

CREATE TABLE chatterbox_schema.followers (
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    follower_id int NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_followers_1 FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_followers_2 FOREIGN KEY (follower_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT uk_followers UNIQUE (user_id, follower_id)
);

CREATE TABLE chatterbox_schema.sessions (
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    session_id varchar(64) NOT NULL,
    status bool NOT NULL,
    started_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ended_at timestamp,
    CONSTRAINT fk_user_sessions FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
);

CREATE TABLE chatterbox_schema.chats (
    id serial PRIMARY KEY,
    sender_id int NOT NULL,
    receiver_id int NOT NULL,
    chat_id varchar(64) NOT NULL,
    status bool NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_chats_1 FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_chats_2 FOREIGN KEY (receiver_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE chatterbox_schema.posts (
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    media_path varchar(64) NOT NULL,
    media_description varchar(128) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL,
    CONSTRAINT fk_user_5 FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
);