INSERT INTO users (name, username, email, password)
VALUES
('user 1', 'user1', 'user1@email.com', '$2a$10$kuvphcaD23yvWP6sJ/sNGO/JukG7QNh8HDLaOpzKd8nsTZwmHAsam'),
('user 2', 'user2', 'user2@email.com', '$2a$10$kuvphcaD23yvWP6sJ/sNGO/JukG7QNh8HDLaOpzKd8nsTZwmHAsam'),
('user 3', 'user3', 'user3@email.com', '$2a$10$kuvphcaD23yvWP6sJ/sNGO/JukG7QNh8HDLaOpzKd8nsTZwmHAsam'),
('user 4', 'user4', 'user4@email.com', '$2a$10$kuvphcaD23yvWP6sJ/sNGO/JukG7QNh8HDLaOpzKd8nsTZwmHAsam'),
('user 5', 'user5', 'user5@email.com', '$2a$10$kuvphcaD23yvWP6sJ/sNGO/JukG7QNh8HDLaOpzKd8nsTZwmHAsam'),
('user 6', 'user6', 'user6@email.com', '$2a$10$kuvphcaD23yvWP6sJ/sNGO/JukG7QNh8HDLaOpzKd8nsTZwmHAsam'),
('user 7', 'user7', 'user7@email.com', '$2a$10$kuvphcaD23yvWP6sJ/sNGO/JukG7QNh8HDLaOpzKd8nsTZwmHAsam'),
('user 8', 'user8', 'user8@email.com', '$2a$10$kuvphcaD23yvWP6sJ/sNGO/JukG7QNh8HDLaOpzKd8nsTZwmHAsam');

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(1, 3),
(2, 4),
(5, 8),
(5, 2),
(5, 3),
(5, 4),
(2, 5),
(6, 2);