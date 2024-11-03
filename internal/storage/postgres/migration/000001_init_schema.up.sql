CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(32) UNIQUE NOT NULL,
	email VARCHAR(128) UNIQUE NOT NULL,
	password VARCHAR(128) NOT NULL,
	bio VARCHAR(255),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE INDEX idx_user_id ON users(id);
CREATE INDEX idx_user_username ON users(username);

CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);
CREATE INDEX idx_refresh_token_id ON refresh_tokens(id);

CREATE TABLE follows (
	follower_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	following_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	PRIMARY KEY(follower_id, following_id), 
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE INDEX idx_follows_id ON follows(follower_id, following_id);

CREATE TABLE tweet (
	id SERIAL PRIMARY KEY, 
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	content TEXT NOT NULL,
	media JSONB,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE INDEX idx_tweet_id ON tweet(id);
CREATE INDEX idx_tweet_created_at ON tweet(created_at);

-- check if retweet with parent tweet already exists, if so delete old retweet and create a new one, retweet id changes, parent tweet id stays the same
CREATE TABLE retweet (
	id SERIAL PRIMARY KEY,
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	parent_tweet_id INT REFERENCES tweet(id) ON DELETE CASCADE,
	parent_retweet_id INT REFERENCES retweet(id) ON DELETE CASCADE,
	is_quote BOOLEAN DEFAULT FALSE NOT NULL,
	quote TEXT,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE INDEX idx_retweet_id ON retweet(id);
CREATE INDEX idx_retweet_created_at ON retweet(created_at);

CREATE TYPE tweet_type AS ENUM ('tweet', 'retweet');

CREATE TABLE comment (
	id SERIAL PRIMARY KEY,
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	tweet_id INT REFERENCES tweet(id) ON DELETE CASCADE,
	retweet_id INT REFERENCES retweet(id) ON DELETE CASCADE,
	parent_comment_id INT REFERENCES comment(id) ON DELETE CASCADE,
	post_type tweet_type NOT NULL,
	content TEXT NOT NULL,
	media JSONB,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE like_tweet (
	id SERIAL PRIMARY KEY,
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	tweet_id INT REFERENCES tweet(id) ON DELETE CASCADE,
	retweet_id INT REFERENCES retweet(id) ON DELETE CASCADE,
	post_type tweet_type NOT NULL,
	UNIQUE (user_id, tweet_id),
	UNIQUE (user_id, retweet_id),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE like_comment (
	id SERIAL PRIMARY KEY,
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	comment_id INT REFERENCES comment(id) ON DELETE CASCADE NOT NULL,
	UNIQUE (user_id, comment_id),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE OR REPLACE FUNCTION f_update_timestamp()
RETURNS TRIGGER AS 
$BODY$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$BODY$ 
language 'plpgsql';
--
CREATE TRIGGER tg_user_update_timestamp
BEFORE UPDATE ON users
FOR EACH ROW 
WHEN (OLD.* IS DISTINCT FROM NEW.*)
EXECUTE PROCEDURE f_update_timestamp();

CREATE TRIGGER tg_follows_update_timestamp
BEFORE UPDATE ON follows
FOR EACH ROW 
WHEN (OLD.* IS DISTINCT FROM NEW.*)
EXECUTE PROCEDURE f_update_timestamp();

CREATE TRIGGER tg_tweet_update_timestamp
BEFORE UPDATE ON tweet
FOR EACH ROW 
WHEN (OLD.* IS DISTINCT FROM NEW.*)
EXECUTE PROCEDURE f_update_timestamp();

CREATE TRIGGER tg_retweet_update_timestamp
BEFORE UPDATE ON retweet
FOR EACH ROW 
WHEN (OLD.* IS DISTINCT FROM NEW.*)
EXECUTE PROCEDURE f_update_timestamp();

CREATE TRIGGER tg_comment_update_timestamp
BEFORE UPDATE ON comment
FOR EACH ROW 
WHEN (OLD.* IS DISTINCT FROM NEW.*)
EXECUTE PROCEDURE f_update_timestamp();

CREATE TRIGGER tg_like_tweet_update_timestamp
BEFORE UPDATE ON like_tweet
FOR EACH ROW 
WHEN (OLD.* IS DISTINCT FROM NEW.*)
EXECUTE PROCEDURE f_update_timestamp();

CREATE TRIGGER tg_like_comment_update_timestamp
BEFORE UPDATE ON like_comment
FOR EACH ROW 
WHEN (OLD.* IS DISTINCT FROM NEW.*)
EXECUTE PROCEDURE f_update_timestamp();

