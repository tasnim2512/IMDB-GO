-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movie_rating (
    id BIGSERIAL,
    movie_id BIGINT,
    user_id BIGINT,
    rating INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    PRIMARY KEY(movie_id),
    CONSTRAINT fk_id
    FOREIGN KEY( movie_id)
	REFERENCES movies (id) on delete SET NULL,
    FOREIGN KEY(user_id)
	REFERENCES users (id) on delete SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movie_rating;
-- +goose StatementEnd
