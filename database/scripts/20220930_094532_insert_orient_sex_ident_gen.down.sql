DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id=1637;
DELETE FROM terceros.grupo_info_complementaria WHERE id=1637

DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id=1636;
DELETE FROM terceros.grupo_info_complementaria WHERE id=1636

DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id=6 AND nombre='INTERSEXUAL';
UPDATE terceros.info_complementaria SET activo=TRUE WHERE grupo_info_complementaria_id=6 AND nombre='NO APLICA';