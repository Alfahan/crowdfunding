-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions (
    id int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    campaign_id INT(11) DEFAULT NULL,
    user_id INT(11) DEFAULT NULL,
    amount INT(11) DEFAULT NULL,
    status VARCHAR(255) DEFAULT NULL,
    code VARCHAR(255) DEFAULT NULL,
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE transactions;
-- +goose StatementEnd
