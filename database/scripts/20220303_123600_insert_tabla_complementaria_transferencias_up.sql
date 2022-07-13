-- Resetear secuencia dañada por registros con IDs quemados
-- Lo siguiente debería ser una forma segura de resetear el serial
-- y permitir que de ahora en más no se admita quemar los IDs
-- (o por lo menos para terceros.grupo_info_complementaria)
-- (Referencia: https://stackoverflow.com/a/244265/3180052)
-- (Otra ref: https://hcmc.uvic.ca/blogs/index.php/how_to_fix_postgresql_error_duplicate_ke?blog=22)
-- Equivale a (también funciona pero no es tan seguro):
-- ALTER SEQUENCE terceros.grupo_info_complementaria RESTART WITH 78
-- Otras formas de alterar secuencias:
-- https://stackoverflow.com/questions/8745051/postgres-manually-alter-sequence
BEGIN;
LOCK TABLE terceros.grupo_info_complementaria IN EXCLUSIVE MODE;
SELECT setval(
    'terceros.grupo_info_complementaria_id_seq',
    COALESCE((SELECT MAX(id)+1 FROM terceros.grupo_info_complementaria), 78),
    false);
COMMIT;

-- Crear el registro INFO_PERSONAL en caso que no exista
INSERT INTO terceros.grupo_info_complementaria
  (nombre, descripcion, codigo_abreviacion,
  activo, fecha_creacion, fecha_modificacion)
SELECT
  'Informacion Personal', 'Informacion personal', 'INFO_PERSONAL',
  true, LOCALTIMESTAMP, LOCALTIMESTAMP
WHERE
  NOT EXISTS (
	  SELECT codigo_abreviacion FROM terceros.grupo_info_complementaria
	  WHERE codigo_abreviacion = 'INFO_PERSONAL'
  )
;

-- Crear el registro HERMANOS_EN_LA_UNIVERSIDAD en caso que no exista
INSERT INTO terceros.info_complementaria
  (nombre, codigo_abreviacion, activo,
   grupo_info_complementaria_id,
   fecha_creacion, fecha_modificacion)
(SELECT
  'HERMANOS_EN_LA_UNIVERSIDAD', 'HERM_UD', true,
  (SELECT id FROM terceros.grupo_info_complementaria
  WHERE codigo_abreviacion = 'INFO_PERSONAL'),
  LOCALTIMESTAMP, LOCALTIMESTAMP
WHERE
  NOT EXISTS (
	  SELECT codigo_abreviacion FROM terceros.info_complementaria
	  WHERE codigo_abreviacion = 'HERM_UD'
  ))
;
