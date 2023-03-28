-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movie_watched (
    id BIGSERIAL,
    movie_id BIGINT,
    user_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(movie_id,user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movie_watched;
-- +goose StatementEnd
