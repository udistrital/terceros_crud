-- Eliminar relación
ALTER TABLE terceros.info_complementaria_tercero DROP CONSTRAINT IF EXISTS fk_info_complementaria_tercero_padre_info_complementaria_tercero CASCADE;

-- Eliminar campos
ALTER TABLE terceros.info_complementaria_tercero
DROP COLUMN IF EXISTS info_complementaria_tercero_padre_id;
