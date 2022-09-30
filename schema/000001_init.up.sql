CREATE TABLE IF NOT EXISTS users
(
    id            serial       PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_lists
(
    id          serial       PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users_lists
(
    id      serial                                          PRIMARY KEY,
    user_id int references users(id) on delete cascade      NOT NULL,
    list_id int references todo_lists(id) on delete cascade NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_items
(
    id          serial       PRIMARY KEY,
    title       VARCHAR(255) NOT NULL, 
    description VARCHAR(255),
    done        boolean      NOT NULL default false
);

CREATE TABLE IF NOT EXISTS lists_items
(
    id      serial                                          PRIMARY KEY,
    item_id int references todo_items(id) on delete cascade NOT NULL,
    list_id int references todo_lists(id) on delete cascade NOT NULL
);

