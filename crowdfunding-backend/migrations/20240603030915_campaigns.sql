-- +goose Up
-- +goose StatementBegin

CREATE TABLE campaigns (
    id int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT(11) DEFAULT NULL,
    name VARCHAR(255) DEFAULT NULL,
    short_description VARCHAR(255) DEFAULT NULL,
    description TEXT,
    perks TEXT,
    backer_count INT(11) DEFAULT NULL,
    goal_amount INT(11) DEFAULT NULL,
    current_amount INT(11) DEFAULT NULL,
    slug VARCHAR(255) DEFAULT NULL,
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE campaigns;
-- +goose StatementEnd