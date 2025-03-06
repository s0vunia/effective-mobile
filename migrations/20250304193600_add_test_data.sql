-- +goose Up
-- +goose StatementBegin
INSERT INTO groups (name, created_at, updated_at) VALUES
    ('Muse', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Radiohead', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Pink Floyd', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Queen', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('The Beatles', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO songs (group_id, title, release_date, link, created_at, updated_at) VALUES
    -- Muse
    ((SELECT id FROM groups WHERE name = 'Muse'), 'Supermassive Black Hole', '2006-07-19', 'https://www.youtube.com/watch?v=pta-gf6JaHQ', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ((SELECT id FROM groups WHERE name = 'Muse'), 'Starlight', '2006-09-04', 'https://www.youtube.com/watch?v=Pgum6OT_VH8', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    
    -- Radiohead
    ((SELECT id FROM groups WHERE name = 'Radiohead'), 'Karma Police', '1997-08-25', 'https://www.youtube.com/watch?v=1uYWYWPc9HU', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ((SELECT id FROM groups WHERE name = 'Radiohead'), 'Creep', '1992-09-21', 'https://www.youtube.com/watch?v=XFkzRNyygfk', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    
    -- Pink Floyd
    ((SELECT id FROM groups WHERE name = 'Pink Floyd'), 'Wish You Were Here', '1975-09-12', 'https://www.youtube.com/watch?v=IXdNnw99-Ic', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ((SELECT id FROM groups WHERE name = 'Pink Floyd'), 'Another Brick in the Wall', '1979-11-23', 'https://www.youtube.com/watch?v=YR5ApYxkU-U', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    
    -- Queen
    ((SELECT id FROM groups WHERE name = 'Queen'), 'Bohemian Rhapsody', '1975-10-31', 'https://www.youtube.com/watch?v=fJ9rUzIMcZQ', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ((SELECT id FROM groups WHERE name = 'Queen'), 'We Will Rock You', '1977-10-07', 'https://www.youtube.com/watch?v=-tJYN-eG1zk', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    
    -- The Beatles
    ((SELECT id FROM groups WHERE name = 'The Beatles'), 'Hey Jude', '1968-08-26', 'https://www.youtube.com/watch?v=A_MjCqQoLLA', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ((SELECT id FROM groups WHERE name = 'The Beatles'), 'Let It Be', '1970-03-06', 'https://www.youtube.com/watch?v=QDYfEBY9NM4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO verses (song_id, verse_number, text, created_at, updated_at) VALUES
    -- Supermassive Black Hole
    ((SELECT id FROM songs WHERE title = 'Supermassive Black Hole'), 1, 'Ooh baby, don''t you know I suffer?\nOh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ((SELECT id FROM songs WHERE title = 'Supermassive Black Hole'), 2, 'You set my soul alight\nYou set my soul alight\nGlaciers melting in the dead of night\nAnd the superstars sucked into the supermassive', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    
    -- Bohemian Rhapsody
    ((SELECT id FROM songs WHERE title = 'Bohemian Rhapsody'), 1, 'Is this the real life?\nIs this just fantasy?\nCaught in a landslide\nNo escape from reality', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ((SELECT id FROM songs WHERE title = 'Bohemian Rhapsody'), 2, 'Open your eyes\nLook up to the skies and see\nI''m just a poor boy, I need no sympathy\nBecause I''m easy come, easy go\nLittle high, little low', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    
    -- Wish You Were Here
    ((SELECT id FROM songs WHERE title = 'Wish You Were Here'), 1, 'So, so you think you can tell\nHeaven from hell?\nBlue skies from pain?\nCan you tell a green field\nFrom a cold steel rail?', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ((SELECT id FROM songs WHERE title = 'Wish You Were Here'), 2, 'A smile from a veil?\nDo you think you can tell?\nDid they get you to trade\nYour heroes for ghosts?', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM verses;
DELETE FROM songs;
DELETE FROM groups;
-- +goose StatementEnd 