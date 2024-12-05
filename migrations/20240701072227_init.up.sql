create table warehouses
(
    id           serial
        primary key,
    name         varchar(255)         not null,
    is_available boolean default true not null
);

create table items
(
    unique_code  varchar(100)      not null
        primary key,
    name         varchar(255),
    size         varchar(50),
    quantity     integer default 0 not null,
    warehouse_id integer default 0
        constraint warehouse_id_fk
            references warehouses,
    reserved     integer default 0 not null
);

INSERT INTO warehouses(id, name, is_available) VALUES (0, 'unstored', true)