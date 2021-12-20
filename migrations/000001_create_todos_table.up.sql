create table todos(
    id serial primary key,
    assignee varchar(50),
    title varchar(50),
    summary varchar(100),
    deadline timestamp null,
    status varchar(20),
    created_at timestamp null,
    updated_at timestamp null,
    deleted_at timestamp null
);