UPDATE terceros.info_complementaria SET activo=FALSE WHERE grupo_info_complementaria_id=6 AND nombre='NO APLICA';
INSERT INTO terceros.info_complementaria(nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('INTERSEXUAL', 'INTERSEXUAL', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.grupo_info_complementaria(id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1636, 'Orientación Sexual', 'Se refiere a los gustos propios de cada persona', 'ORIENTACION_SEXUAL', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria(nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Bisexual', 'BISEXUAL', TRUE, 1636, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria(nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Heterosexual', 'HETEROSEXUAL', TRUE, 1636, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria(nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Homosexual', 'HOMOSEXUAL', TRUE, 1636, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria(nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('No deseo contestar', 'NO_CONTESTA', TRUE, 1636, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.grupo_info_complementaria(id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1637, 'Identidad de Género', 'Se refiere a como quiere ser reconocido socialmente', 'IDENTIDAD_GENERO', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria(nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Cisgénero', 'CISGENERO', TRUE, 1637, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria(nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Trans Femenino', 'TRANSFEMENINO', TRUE, 1637, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria(nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Trans Masculino', 'TRANSMASCULINO', TRUE, 1637, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria(nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('No deseo contestar', 'NO_CONTESTA', TRUE, 1637, LOCALTIMESTAMP, LOCALTIMESTAMP);