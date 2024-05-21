-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
insert into features
select id, ('ftr №' || id) as name  from generate_series(1, 1000, 1) as id;

insert into tags
select id, ('tag №' || id) as name from generate_series(1, 1000, 1) as id;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
truncate features cascade ;
truncate tags cascade ;
-- +goose StatementEnd
