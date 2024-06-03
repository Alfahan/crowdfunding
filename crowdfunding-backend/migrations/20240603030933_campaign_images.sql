-- +goose Up
-- +goose StatementBegin
CREATE TABLE campaign_images (
    id int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    campaign_id INT(11) DEFAULT NULL,
    file_name VARCHAR(255) DEFAULT NULL,
    is_primary TINYINT(4) DEFAULT NULL,
    created_at datetime DEFAULT NULL,
    updated datetime DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE campaign_images;
-- +goose StatementEnd