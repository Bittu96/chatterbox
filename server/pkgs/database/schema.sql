-- Active: 1711619722854@@ec2-3-84-86-24.compute-1.amazonaws.com@5432@postgres
CREATE SCHEMA chatterbox;

CREATE TABLE chatterbox.user (
    user_id serial PRIMARY KEY,
    username varchar(64) UNIQUE,
    email varchar(64) UNIQUE NOT NULL,
    password varchar(64) NOT NULL,
    role varchar(8) NOT NULL DEFAULT 'user',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL
);

CREATE TABLE chatterbox.follower (
    user_id int NOT NULL,
    follower_id int NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(user_id, follower_id),
    CONSTRAINT fk_user_follower_1 FOREIGN KEY (user_id) REFERENCES chatterbox.user (user_id) ON DELETE CASCADE,
    CONSTRAINT fk_user_follower_2 FOREIGN KEY (follower_id) REFERENCES chatterbox.user (user_id) ON DELETE CASCADE
);

CREATE TABLE chatterbox.session (
    session_id varchar(64) PRIMARY KEY,
    user_id int NOT NULL,
    status bool NOT NULL,
    started_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ended_at timestamp,
    CONSTRAINT fk_user_session FOREIGN KEY (user_id) REFERENCES chatterbox.user (user_id) ON DELETE CASCADE
);

CREATE TABLE chatterbox.chat (
    sender_id int NOT NULL,
    receiver_id int NOT NULL,
    status bool NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(sender_id, receiver_id),
    CONSTRAINT fk_user_chat_1 FOREIGN KEY (sender_id) REFERENCES chatterbox.user (user_id) ON DELETE CASCADE,
    CONSTRAINT fk_user_chat_2 FOREIGN KEY (receiver_id) REFERENCES chatterbox.user (user_id) ON DELETE CASCADE
);

CREATE TABLE chatterbox.post (
    post_id varchar(64) PRIMARY KEY,
    user_id int NOT NULL,
    media_path varchar(64) NOT NULL,
    media_description varchar(128) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL,
    CONSTRAINT fk_user_5 FOREIGN KEY (user_id) REFERENCES chatterbox.user (user_id) ON DELETE CASCADE
);
