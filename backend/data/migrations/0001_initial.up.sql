-- our user
create table if not exists users (
    id integer primary key autoincrement,
    uuid text not null,
    name text not null not null,
    email text unique,
    folder_uuid text not null,
    is_draft bool not null,
    created_at timestamp default current_timestamp
);
-- auth sessions
create table if not exists sessions (
    id integer primary key autoincrement,
    user_id integer not null,
    email text not null,
    expiry timestamp not null,
    token text not null,
    activate_code text not null,
    refresh_token text not null,
    user_ip text not null,
    type text not null,
    post_suspend_expiry timestamp,
    is_expired bool default false,
    foreign key (user_id) references users(id)
);

create table if not exists pdfs (
    id integer primary key autoincrement,
    uuid text not null,
    user_id integer not null,
    name text not null,
    created_at timestamp default current_timestamp,
    last_open_at timestamp default current_timestamp,
    foreign key (user_id) references users(id)
)