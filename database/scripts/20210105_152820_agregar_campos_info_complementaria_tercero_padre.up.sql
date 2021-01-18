-- Campo relacion nivel formacion padre
ALTER TABLE tercero_crud.info_complementaria_tercero ADD COLUMN info_complementaria_tercero_padre_id integer;


-- object: fk_nivel_formacion_padre_nivel_formacion | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.nivel_formacion DROP CONSTRAINT IF EXISTS fk_nivel_formacion_padre_nivel_formacion CASCADE;
ALTER TABLE tercero_crud.info_complementaria_tercero ADD CONSTRAINT fk_info_complementaria_tercero_padre_info_complementaria_tercero FOREIGN KEY (info_complementaria_tercero_padre_id)
REFERENCES tercero_crud.info_complementaria_tercero (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;