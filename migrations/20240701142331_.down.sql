TRUNCATE items;

TRUNCATE warehouses;

INSERT INTO warehouses(id, name, is_available) VALUES (0, 'unstored', true);

ALTER SEQUENCE warehouses_id_seq restart;