-- Campo relacion info complementaria tercero padre
ALTER TABLE terceros.info_complementaria_tercero ADD COLUMN info_complementaria_tercero_padre_id integer;


-- object: fk_info_complementaria_tercero_padre_info_complementaria_tercero | type: CONSTRAINT --
-- ALTER TABLE erceros.info_complementaria_tercero DROP CONSTRAINT IF EXISTS fk_info_complementaria_tercero_padre_info_complementaria_tercero CASCADE;
ALTER TABLE terceros.info_complementaria_tercero ADD CONSTRAINT fk_info_complementaria_tercero_padre_info_complementaria_tercero FOREIGN KEY (info_complementaria_tercero_padre_id)
REFERENCES terceros.info_complementaria_tercero (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;