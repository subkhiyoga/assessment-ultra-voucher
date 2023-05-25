CREATE TABLE test (
	id INTEGER PRIMARY KEY,
	name VARCHAR,
	parent_id INTEGER,
);

INSERT INTO test (id, name, parent_id)
VALUES
(1, 'Zaki', 2),
(2, 'Ilham', NULL),
(3, 'Irwan', 2),
(4, 'Arka', 3);

SELECT t1.id, t1.name, t2.name AS parent_name
FROM test AS t1
LEFT JOIN test AS t2 ON t1.parent_id = t2.id;
