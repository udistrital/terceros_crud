INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (21, 'Derechos pecuniarios', 'Información de los derechos pecunarios que se le asignan a una solicitud de un estudiante', 'Grupo_21', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('RECIBO_DERECHO_PECUNARIO', 'RECIBO_D_PECUNARIO', TRUE, 21, LOCALTIMESTAMP, LOCALTIMESTAMP);