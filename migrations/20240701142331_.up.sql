INSERT INTO warehouses (id, name, is_available)
VALUES (1, 'A', true),
    (2, 'B', true),
    (3, 'C', true),
    (4, 'D', true),
    (5, 'E', false),
    (6, 'F', false);

INSERT INTO items (unique_code, name, size, quantity, warehouse_id)
VALUES ('A0001', 'TestItem1', 'S', 10, 1),
       ('A0002', 'TestItem1', 'M', 1, 1),
       ('A0003', 'TestItem2', 'M', 30, 1),
       ('A0004', 'TestItem2', 'L', 5, 1),
       ('A0005', 'TestItem3', 'XL', 2, 1),
       ('A0006', 'TestItem3', 'M', 7, 1),
       ('B0001', 'TestShoe1', '41', 11, 2),
       ('B0002', 'TestShoe1', '39', 5, 2),
       ('B0003', 'TestShoe2', '42', 1, 2),
       ('B0004', 'TestShoe3', '40', 2, 2),
       ('C0001', 'TestSweeter', 'S', 15, 3),
       ('C0002', 'TestSweeter2', 'M', 200, 3),
       ('C0003', 'TestSweeter2', 'S', 150, 3),
       ('D0001', 'TestAnotherItem1', 'test_size1', 15, 4),
       ('D0002', 'TestAnotherItem1', 'test_size2', 7, 4),
       ('D0003', 'TestAnotherItem1', 'test_size1', 2, 4),
       ('D0004', 'TestAnotherItem1', 'test_size2', 1, 4),
       ('E0001', 'TestUnavailableItem', 'test_size1', 5, 5),
       ('E0004', 'TestUnavailableItem', 'test_size2', 1, 5),
       ('F0001', 'TestUnavailableItem2', 'test_sizeS', 1, 6);

ALTER SEQUENCE warehouses_id_seq START WITH 7;
ALTER SEQUENCE warehouses_id_seq RESTART;