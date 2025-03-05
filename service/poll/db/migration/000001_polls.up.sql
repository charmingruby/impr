CREATE TABLE IF NOT EXISTS polls
(
    id varchar PRIMARY KEY NOT NULL,
    name varchar NOT NULL,
    description text,
	status varchar NOT NULL,
    expirationTime integer NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp
);

CREATE TABLE IF NOT EXISTS poll_options 
(
    id varchar PRIMARY KEY NOT NULL,
    content varchar NOT NULL,
    poll_id varchar NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
    
	FOREIGN KEY (poll_id) REFERENCES polls(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS votes
(
    id varchar PRIMARY KEY NOT NULL,
    poll_option_id varchar NOT NULL,
    user_id varchar NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    
	FOREIGN KEY (poll_option_id) REFERENCES poll_options(id) ON DELETE CASCADE
);