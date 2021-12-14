create table if not exists todos(
    id serial primary key,
    assignee varchar(64),
    title varchar(64),
    summary varchar(256),
    deadline timestamp,
    status varchar(32)
);