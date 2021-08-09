DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  username TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE followers(
  user_id INT NOT NULL,
  CONSTRAINT fk_userid
    FOREIGN KEY(user_id)
      REFERENCES users(id)
        ON DELETE CASCADE,
  follower_id INT NOT NULL,
  CONSTRAINT fk_followerid
    FOREIGN KEY(follower_id)
      REFERENCES users(id)
        ON DELETE CASCADE,

  PRIMARY KEY (user_id, follower_id)
)

CREATE TABLE posts(
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  author_id INT NOT NULL,
  likes INT DEFAULT 0,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_authorid
    FOREIGN KEY(author_id)
      REFERENCES users(id)
        ON DELETE CASCADE
)