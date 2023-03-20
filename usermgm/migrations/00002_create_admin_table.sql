-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO users(
        first_name,
        last_name,
        email,
        is_admin,
        username,
        role,
        password
    )
VALUES(
        'admin1',
        'admin2',
        'admin@gmail.com',
        true,
        'admin',
        'admin',
        '12345'
    );
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.