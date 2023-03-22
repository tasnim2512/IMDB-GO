-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movie_genre (
    id BIGSERIAL,
    movie_id BIGINT,
    genre_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    PRIMARY KEY(id),
    CONSTRAINT fk_id
    FOREIGN KEY( movie_id)
	REFERENCES movies (id) on delete SET NULL,
    FOREIGN KEY(genre_id)
	REFERENCES genres (id) on delete SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movie_genre;
-- +goose StatementEnd
