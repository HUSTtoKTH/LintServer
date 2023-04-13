CREATE TABLE IF NOT EXISTS rules(
    id serial PRIMARY KEY,
    project_id bigint UNIQUE,
    organization_id bigint,
    rule text
);