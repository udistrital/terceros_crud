--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--Schema Terceros
CREATE SCHEMA terceros;
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

--Creacion tabla tipo_documento

CREATE SEQUENCE terceros.tipo_documento_id_seq
	INCREMENT BY 1 
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tipo_documento(
	id integer NOT NULL DEFAULT nextval('terceros.tipo_documento_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_documento PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.tipo_documento IS 'Tabla que parametriza los diferentes tipos de documento que puede tener una persona juridica o natural.';
COMMENT ON COLUMN terceros.tipo_documento.id IS 'Identificador unico de la tabla tipo_documento.';
COMMENT ON COLUMN terceros.tipo_documento.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.tipo_documento.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del tipo_documento.';
COMMENT ON COLUMN terceros.tipo_documento.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.tipo_documento.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tipo_documento.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN terceros.tipo_documento.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tipo_documento.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--Creacion tabla tipo_contribuyente

CREATE SEQUENCE terceros.tipo_contribuyente_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tipo_contribuyente(
	id integer NOT NULL DEFAULT nextval('terceros.tipo_contribuyente_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_contribuyente PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.tipo_contribuyente IS 'Tabla que parametriza los Tipos de contribuyentes: PersonaNatural - Persona Juridica';
COMMENT ON COLUMN terceros.tipo_contribuyente.id IS 'Identificador unico de la tabla tipo_contribuyente.';
COMMENT ON COLUMN terceros.tipo_contribuyente.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.tipo_contribuyente.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_contribuyente.';
COMMENT ON COLUMN terceros.tipo_contribuyente.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.tipo_contribuyente.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tipo_contribuyente.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tipo_contribuyente.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

CREATE SEQUENCE terceros.tipo_info_basica_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tipo_info_basica(
	id integer NOT NULL DEFAULT nextval('terceros.tipo_info_basica_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_info_basica PRIMARY KEY (id)
);

-- Pendiente !!! !!! 	!!!
COMMENT ON TABLE terceros.tipo_info_basica IS 'Tabla que parametriza los grupos de grupos sanguineo, grupos etnicos, discapacidades, estados civiles, generos entre otros.';


CREATE SEQUENCE terceros.tipo_seguridad_social_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tipo_seguridad_social(
	id integer NOT NULL DEFAULT nextval('terceros.tipo_seguridad_social_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_seguridad_social PRIMARY KEY (id)
);

-- Pendiente !!! !!! 	!!!
COMMENT ON TABLE terceros.tipo_seguridad_social IS 'Tabla que parametriza a que tipo de entidad pertenece el tercero sea:  Eps / Caja Compensacion / Arl / fondo de Pension entre otros.';



--||||||||||||||||||||||||||||||||||||
--Creacion tabla informacion_basica para:
--               tipo_grupo_sanguineo
--               tipo_grupo_etnico
--               tipo_discapacidad
--               estado_civil
--               genero

CREATE SEQUENCE terceros.info_basica_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.info_basica(
	id integer NOT NULL DEFAULT nextval('terceros.info_basica_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	tipo_info_basica_id integer NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_info_basica PRIMARY KEY (id),
	CONSTRAINT fk_tipo_info_basica_info_basica FOREIGN KEY (tipo_info_basica_id) REFERENCES terceros.tipo_info_basica(id)
);

-- Pendiente !!! !!! 	!!!
COMMENT ON TABLE terceros.info_basica IS 'Tabla que parametriza los tipos de grupos sanguineo, grupos etnicos, discapacidades, estados civiles, generos entre otros.';
COMMENT ON COLUMN terceros.info_basica.id IS 'Identificador unico de la tabla WWWWWW.';
COMMENT ON COLUMN terceros.info_basica.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.info_basica.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.info_basica.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.info_basica.tipo_info_basica_id IS 'Identificador unico de la tabla tipo_informacion_basica.';
COMMENT ON COLUMN terceros.info_basica.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.info_basica.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--°°Tercero°°


--Creacion tabla tercero

CREATE SEQUENCE terceros.tercero_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tercero(
	id integer NOT NULL DEFAULT nextval('terceros.tercero_id_seq'::regclass),
	nombre_completo character varying(255) NOT NULL,
    primer_nombre character varying(100),
    segundo_nombre character varying(100),
    primer_apellido character varying(100),
    segundo_apellido character varying(100),
    tipo_contribuyente_id integer NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
    CONSTRAINT fk_tipo_contribuyente_tercero FOREIGN KEY (tipo_contribuyente_id) REFERENCES terceros.tipo_contribuyente(id),
	CONSTRAINT pk_tercero PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.tercero IS 'Tabla que parametriza los Tipos de contribuyentes: PersonaNatural - Persona Juridica';



--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||


--Creacion tabla datos_identificacion

CREATE SEQUENCE terceros.datos_identificacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.datos_identificacion(
	id integer NOT NULL DEFAULT nextval('terceros.datos_identificacion_id_seq'::regclass),
    tipo_documento_id integer NOT NULL,
    tercero_id integer NOT NULL,
    numero character varying(50) NOT NULL,
    digito_verificacion numeric(2,0),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_datos_identificacion PRIMARY KEY (id),
    CONSTRAINT fk_tipo_documento_datos_identificacion FOREIGN KEY (tipo_documento_id) REFERENCES terceros.tipo_documento(id),
    CONSTRAINT fk_tercero_datos_identificacion FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.datos_identificacion IS 'Tabla que almacena el numero de identificación de un tercero y lo relaciona con su tipo.';
COMMENT ON COLUMN terceros.datos_identificacion.id IS 'Identificador unico de la tabla datos_identificacion.';
COMMENT ON COLUMN terceros.datos_identificacion.tipo_documento_id IS 'Identificador unico de la tabla tipo_documento';
COMMENT ON COLUMN terceros.datos_identificacion.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.datos_identificacion.numero IS 'Numero de Identificacion del Tercero.';
COMMENT ON COLUMN terceros.datos_identificacion.digito_verificacion IS 'Dígito de verificación de la DIAN. Toda persona natural o jurídica que tenga un RUT debe tener un DV (NOT NULL para personas_juridicas).';
COMMENT ON COLUMN terceros.datos_identificacion.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.datos_identificacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.datos_identificacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--||||||||||||||||||||||||||||||||||||
--Creacion tabla info_basica_tercero para relacionar:
--               tipo_grupo_sanguineo
--               tipo_grupo_etnico
--               tipo_discapacidad
--               estado_civil
--               genero
--Con tabla terceros.

CREATE SEQUENCE terceros.info_basica_tercero_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.info_basica_tercero(
	id integer NOT NULL DEFAULT nextval('terceros.info_basica_tercero_id_seq'::regclass),
	tercero_id integer NOT NULL,
	info_basica_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_info_basica_tercero PRIMARY KEY (id),
	CONSTRAINT fk_tercero_info_basica_tercero FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id),
	CONSTRAINT fk_info_basica_info_basica_tercero FOREIGN KEY (info_basica_id) REFERENCES terceros.info_basica(id)
);
COMMENT ON TABLE terceros.info_basica_tercero IS 'Tabla que relaciona a un Tercero con su informacion basica sea esta: ';


CREATE SEQUENCE terceros.seguridad_social_tercero_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.seguridad_social_tercero(
	id integer NOT NULL DEFAULT nextval('terceros.seguridad_social_tercero_id_seq'::regclass),
	tercero_id integer NOT NULL,
	tercero_entidad_id integer NOT NULL,
	tipo_seguridad_social_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_inicio_vinculacion TIMESTAMP NOT NULL,
	fecha_fin_vinculacion TIMESTAMP NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_seguridad_social_tercero PRIMARY KEY (id),
	CONSTRAINT fk_tercero_seguridad_social_tercero FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id),
	CONSTRAINT fk_tercero_entidad_seguridad_social_tercero FOREIGN KEY (tercero_entidad_id) REFERENCES terceros.tercero(id),
	CONSTRAINT fk_tipo_seguridad_social_seguridad_social_tercero FOREIGN KEY (tipo_seguridad_social_id) REFERENCES terceros.tipo_seguridad_social(id)
);
COMMENT ON TABLE terceros.seguridad_social_tercero IS 'Tabla que relaciona a un Tercero con su:  Eps / Caja Compensacion / Arl / fondo de Pension entre otros.';



--Creacion tabla informacion_socioeconomica 

CREATE SEQUENCE terceros.informacion_socioeconomica_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.informacion_socioeconomica(
	id integer NOT NULL DEFAULT nextval('terceros.informacion_socioeconomica_id_seq'::regclass),
    tercero_id integer NOT NULL,
    estrato numeric (1,0),
    puntaje_sisben numeric (3,2),
    declarante_renta boolean, 
    comunidad_lgbt boolean,
    estado_civil boolean,
    cabeza_familia boolean,
    personas_a_cargo boolean,
    numero_personas_a_cargo numeric(2,0),
    hijos boolean,
    numero_hijos numeric(2,0),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_informacion_socioeconomica PRIMARY KEY (id),
    CONSTRAINT fk_tercero_informacion_socieconomica FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);




/**

--Creacion tabla informacion_socioeconomica 

CREATE SEQUENCE terceros.informacion_socioeconomica_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;



comunidad_lgbt
cabeza_familia
personas_a_cargo
número de personas a cargo
declarante renta
Hijos
numero hijos
Pensionado

estrato
puntaje sisben



CREATE TABLE terceros.informacion_socioeconomica(
	id integer NOT NULL DEFAULT nextval('terceros.informacion_socioeconomica_id_seq'::regclass),
    tercero_id integer NOT NULL,

    --NOT NULL DEFAULT false,
    --estado_civil character varying(20) NOT NULL DEFAULT 'SOLTERO'::character varying

    estrato numeric (1,0),
    puntaje_sisben numeric (3,2),
    declarante_renta boolean,
    

    comunidad_lgbt boolean,    

    estado_civil boolean,
    cabeza_familia boolean,
    personas_a_cargo boolean,
    numero_personas_a_cargo numeric(2,0),
    hijos boolean,
    numero_hijos numeric(2,0),
    

	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_informacion_socioeconomica PRIMARY KEY (id),
    CONSTRAINT fk_tercero_informacion_socieconomica FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.informacion_socioeconomica IS 'Tabla que relaciona toda la información socioeconomica de un tercero.';
COMMENT ON COLUMN terceros.informacion_socioeconomica.id IS 'Identificador unico de la tabla informacion_socioeconomica.';
COMMENT ON COLUMN terceros.informacion_socioeconomica.tercero_id IS 'Identificador unico de la tabla tercero.';


COMMENT ON COLUMN terceros.informacion_socioeconomica.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.informacion_socioeconomica.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.informacion_socioeconomica.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
**/