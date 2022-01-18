UPDATE terceros.info_complementaria SET Nombre = 'DOCUMENTO_SOPORTE', CodigoAbreviacion = 'DOC_SOPORTE' WHERE grupo_info_complementaria_id = 1 AND codigo_abreviacion = 'DOC_DISCAPACIDAD'; 

DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 3 AND codigo_abreviacion = 'DOC_POBLACION';