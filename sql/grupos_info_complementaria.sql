- ._.
-CURRENT_TIME and CURRENT_TIMESTAMP deliver values with time zone; LOCALTIME and LOCALTIMESTAMP deliver values without time zone.

-REVISAR agora.parametro_estandar

-Tabla: grupo_info_complementaria 
-Tabla que parametriza la información complementaria de un tercero que pueda ser relacionada a un mismo grupo o tipo como pueden ser: discapacidades, estados civiles, grupos sanguíneo, grupos étnicos, géneros entre otros. También se requiere un grupo para valores cuyos datos internos no van a ser parametrizados

INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1, 'Tipo Discapacidad', 'Tipos de Discapacidades que puede tener un tercero / persona natural.', 'Grupo_1', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (2, 'Estado Civil', 'Tipos de Estado Civil que puede tener un tercero / persona_natural.', 'Grupo_2', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (3, 'Tipo Etnia', 'Tipos de Etnias a la que puede pertenecer un tercero / persona_natural.', 'Grupo_3', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (4, 'Estados', 'Tipos de Estados.', 'Grupo_4', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (5, 'Tipo Perfil', 'Tipos de Perfiles.', 'Grupo_5', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);

-Tabla: info_complementaria
-Tabla que parametriza los tipos concretos de grupos sanguineo, grupos etnicos, discapacidades, estados civiles, generos y cualquier informacion complementaria relacionada a un tercero.

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('FISICA', 'FISICA', TRUE, 1, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('SENSORIAL', 'SENSORIAL', TRUE, 1, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('AUDITIVA', 'AUDITIVA', TRUE, 1, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('VISUAL', 'VISUAL', TRUE, 1, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('PSIQUICA', 'PSIQUICA', TRUE, 1, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('MENTAL', 'MENTAL', TRUE, 1, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('NO APLICA', 'NO APLICA', TRUE, 1, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('SOLTERO', 'SOLTERO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CASADO', 'CASADO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('UNION LIBRE', 'UNION LIBRE', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('VIUDO', 'VIUDO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('DIVORCIADO', 'DIVORCIADO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('AFRODESCENDIENTES', 'AFRODESCENDIENTES', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('INDIGENAS', 'INDIGENAS', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('RAIZALES', 'RAIZALES', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('ROM', 'ROM', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('NO APLICA', 'NO APLICA', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('ACTIVO', 'ACTIVO', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('INACTIVO', 'INACTIVO', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('INHABILITADO', 'INHABILITADO', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('RECHAZADO', 'RECHAZADO', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('ASISTENCIAL', 'ASISTENCIAL', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('TÉCNICO', 'TÉCNICO', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('PROFESIONAL', 'PROFESIONAL', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('PROFESIONAL ESPECIALIZADO', 'PROF_ESPECIALIZADO', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('ASESOR 1', 'ASESOR 1', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('ASESOR 2', 'ASESOR 2', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('NO APLICA', 'NO APLICA', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);

-Tabla: tipo_documento 
-Tabla que parametriza los diferentes tipos de documento que puede tener una persona juridica o natural.

INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('REGISTRO CIVIL DE NACIMIENTO', 'REGISTRO CIVIL DE NACIMIENTO', 'RC', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('TARJETA DE IDENTIDAD', 'TARJETA DE IDENTIDAD', 'TI', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('CÉDULA DE CIUDADANÍA', 'CÉDULA DE CIUDADANÍA', 'CC', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('CERTIFICADO REGISTRADURÍA SIN IDENTIFICACIÓN', 'CERTIFICADO REGISTRADURÍA SIN IDENTIFICACIÓN', 'CRSI', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('TARJETA DE EXTRANJERÍA', 'TARJETA DE EXTRANJERÍA', 'TE', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('CÉDULA DE EXTRANJERÍA', 'CÉDULA DE EXTRANJERÍA', 'CE', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('NIT', 'NIT', 'NIT', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('IDENTIFICACIÓN DE EXTRANJEROS DIFERENTE AL NIT ASIGNADO DIAN', 'IDENTIFICACIÓN DE EXTRANJEROS DIFERENTE AL NIT ASIGNADO DIAN', 'IEDD', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('PASAPORTE', 'PASAPORTE', 'PAS', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('DOCUMENTO DE IDENTIFICACIÓN EXTRANJERO', 'DOCUMENTO DE IDENTIFICACIÓN EXTRANJERO', 'DIE', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('SIN IDENTIFICACIÓN DEL EXTERIOR O PARA USO DEFINIDO POR DIAN', 'SIN IDENTIFICACIÓN DEL EXTERIOR O PARA USO DEFINIDO POR DIAN', 'SIED', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('DOCUMENTO DE IDENTIFICACIÓN EXTRANJERO PERSONA JURÍDICA', 'DOCUMENTO DE IDENTIFICACIÓN EXTRANJERO PERSONA JURÍDICA', 'DIEPJ', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('CARNÉ DIPLOMÁTICO', 'CARNÉ DIPLOMÁTICO', 'CD', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);






-Notas 

-Revisar agora.parametro_tipo_conformacion VS agora.tipo_entidad
-Unidades tabla que se pidio revisar agora.unidad





-Tabla: tipo_contribuyente 
-Tabla que parametriza los Tipos de contribuyentes: PersonaNatural - Persona Juridica

INSERT INTO terceros.tipo_contribuyente (id, nombre, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1, 'PERSONA NATURAL', 'P_NATURAL', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_contribuyente (id, nombre, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (2, 'PERSONA JURIDICA', 'P_JURIDICA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);

-Tabla: tipo_tercero
Tabla que parametriza a los tipos de terceros, sean: Personas/Empresas/Eps/Caja Compensacion/Arl, Entidades de caracter Publico/Privado/Mixto entre otros.

INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1, 'BANCO', 'Entidad Bancaria', 'BANCO', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (2, 'E_ASEGURADORA', 'Entidad Aseguradora', 'E_ASEGURADORA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (3, 'EPS', 'Entidad Prestadora de Salud ', 'EPS', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (4, 'ARL', 'Administradora de Riesgos Laborales', 'ARL', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (5, 'F_PENSION', 'Fondo de Pensión', 'F_PENSION', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (6, 'CAJA_COMPENSACIÓN', 'Caja de Compensación', 'CAJA_COMPENSACIÓN', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (7, 'COLEGIO', null, 'COLEGIO', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (8, 'ENTIDAD_PUBLICA', null, 'ENTIDAD_PUBLICA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (9, 'ENTIDAD_PRIVADA', null, 'ENTIDAD_PRIVADA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (10, 'ENTIDAD_MIXTA', 'Eentidad de caracter tanto publica como privada', 'ENTIDAD_MIXTA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (11, 'FAMILIAR', 'Persona natural que tiene un parentesco con otra persona natural de la organización', 'FAMILIAR', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);


-Tabla: grupo_info_complementaria // info_complementaria

INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (6, 'Genero', 'Generos.', 'Grupo_6', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (7, 'Grupo sanguineo', 'Clasificacion de los grupos sanguineos establecida por Karl Landsteiner.', 'Grupo_7', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (8, 'Factor RH', 'El factor RH es una proteína de los glóbulos rojos y permite detectar el tipo de sangre sea + o -.', 'Grupo_8', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (9, 'Información Socioeconómica', 'Información Socioeconómica de una persona natural.', 'Grupo_9', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (10, 'Información Contacto', 'Información de Contacto de una persona natural o juridica.', 'Grupo_10', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);


INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('FEMENINO', 'FEMENINO', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('MASCULINO', 'MASCULINO', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('NO APLICA', 'NO APLICA', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('A', 'A', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('B', 'B', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('AB', 'AB', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('O', 'O', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('+', '+', TRUE, 8, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('-', '-', TRUE, 8, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('ESTRATO', 'ESTRATO', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('PUNTAJE_SISBEN', 'PUNTAJE_SISBEN', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('DECLARANTE_RENTA', 'DECLARANTE_RENTA', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('COMUNIDAD_LGBT', 'COMUNIDAD_LGBT', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CABEZA_FAMILIA', 'CABEZA_FAMILIA', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('PERSONAS_A_CARGO', 'PERSONAS_A_CARGO', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('NUMERO_PERSONAS_A_CARGO', 'NUM_PERSONAS_A_CARGO', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('HIJOS', 'HIJOS', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('NUMERO_HIJOS', 'NUMERO_HIJOS', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('NUMERO_HERMANOS', 'NUMERO_HERMANOS', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);

INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('TELEFONO', 'TELEFONO', TRUE, 10, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CELULAR', 'CELULAR', TRUE, 10, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CORREO', 'CORREO', TRUE, 10, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('DIRECCIÓN', 'DIRECCIÓN', TRUE, 10, LOCALTIMESTAMP, LOCALTIMESTAMP);
INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CODIGO_POSTAL', 'CODIGO_POSTAL', TRUE, 10, LOCALTIMESTAMP, LOCALTIMESTAMP);
