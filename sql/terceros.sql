--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--Schema Terceros
CREATE SCHEMA terceros;
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

--Creacion tabla tipo_grupo_sanguineo

CREATE SEQUENCE terceros.tipo_grupo_sanguineo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tipo_grupo_sanguineo(
	id integer NOT NULL DEFAULT nextval('terceros.tipo_grupo_sanguineo_id_seq'::regclass),
    tipo_sangre character varying(2) NOT NULL,
    factor_rh character varying(1) NOT NULL,
	codigo_abreviacion character varying(3) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_grupo_sanguineo PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.tipo_grupo_sanguineo IS 'Tabla que parametriza los grupo sanguíneos de acuerdo al sistema ABO.';
COMMENT ON COLUMN terceros.tipo_grupo_sanguineo.id IS 'Identificador unico de la tabla tipo_grupo_sanguineo.';
COMMENT ON COLUMN terceros.tipo_grupo_sanguineo.tipo_sangre IS 'Tipo de grupo que se identifica en el sistema ABO: A, B, O, AB.';
COMMENT ON COLUMN terceros.tipo_grupo_sanguineo.factor_rh IS 'Valor positivo o negativo (+, -) del factor RH.';
COMMENT ON COLUMN terceros.tipo_grupo_sanguineo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.tipo_grupo_sanguineo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tipo_grupo_sanguineo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tipo_grupo_sanguineo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';



--Creacion tabla tipo_grupo_etnico

CREATE SEQUENCE terceros.tipo_grupo_etnico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tipo_grupo_etnico(
	id integer NOT NULL DEFAULT nextval('terceros.tipo_grupo_etnico_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_grupo_etnico PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.tipo_grupo_etnico IS 'Tabla que parametriza los grupos Étnicos.';
COMMENT ON COLUMN terceros.tipo_grupo_etnico.id IS 'Identificador unico de la tabla tipo_grupo_etnico.';
COMMENT ON COLUMN terceros.tipo_grupo_etnico.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.tipo_grupo_etnico.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del grupo étnico.';
COMMENT ON COLUMN terceros.tipo_grupo_etnico.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.tipo_grupo_etnico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tipo_grupo_etnico.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN terceros.tipo_grupo_etnico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tipo_grupo_etnico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

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

--Creacion tabla tipo_discapacidad

CREATE SEQUENCE terceros.tipo_discapacidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tipo_discapacidad(
	id integer NOT NULL DEFAULT nextval('terceros.tipo_discapacidad_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_discapacidad PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.tipo_discapacidad IS 'Tabla que parametriza los tipos de discapacidades.';
COMMENT ON COLUMN terceros.tipo_discapacidad.id IS 'Identificador unico de la tabla tipo_discapacidad.';
COMMENT ON COLUMN terceros.tipo_discapacidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.tipo_discapacidad.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de la discapacidad.';
COMMENT ON COLUMN terceros.tipo_discapacidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.tipo_discapacidad.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tipo_discapacidad.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tipo_discapacidad.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--Creacion tabla estado_civil

CREATE SEQUENCE terceros.estado_civil_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.estado_civil(
	id integer NOT NULL DEFAULT nextval('terceros.estado_civil_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_estado_civil PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.estado_civil IS 'Tabla que parametriza los estados civiles.';
COMMENT ON COLUMN terceros.estado_civil.id IS 'Identificador unico de la tabla estado_civil.';
COMMENT ON COLUMN terceros.estado_civil.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.estado_civil.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion referente al estado civil.';
COMMENT ON COLUMN terceros.estado_civil.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.estado_civil.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.estado_civil.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.estado_civil.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--Creacion tabla genero

CREATE SEQUENCE terceros.genero_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.genero(
	id integer NOT NULL DEFAULT nextval('terceros.genero_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_genero PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.genero IS 'Tabla que parametriza los generos entendidos como Sexo biológico: características biológicas y físicas usadas típicamente para asignar el género al nacer.';
COMMENT ON COLUMN terceros.genero.id IS 'Identificador unico de la tabla genero.';
COMMENT ON COLUMN terceros.genero.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.genero.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion referente al genero.';
COMMENT ON COLUMN terceros.genero.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.genero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.genero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.genero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--°°Tercero°°

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

--Creacion tabla grupo_sanguineo

CREATE SEQUENCE terceros.grupo_sanguineo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.grupo_sanguineo(
	id integer NOT NULL DEFAULT nextval('terceros.grupo_sanguineo_id_seq'::regclass),
    tipo_grupo_sanguineo_id integer NOT NULL,
    tercero_id integer NOT NULL,	
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_grupo_sanguineo PRIMARY KEY (id),
    CONSTRAINT fk_tipo_grupo_sanguineo_grupo_sanguineo FOREIGN KEY (tipo_grupo_sanguineo_id) REFERENCES terceros.tipo_grupo_sanguineo(id),
    CONSTRAINT fk_tercero_grupo_sanguineo FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.grupo_sanguineo IS 'Tabla que relaciona a un tercero con un tipo_grupo_sanguineo.';
COMMENT ON COLUMN terceros.grupo_sanguineo.id IS 'Identificador unico de la tabla grupo_sanguineo.';
COMMENT ON COLUMN terceros.grupo_sanguineo.tipo_grupo_sanguineo_id IS 'Identificador unico de la tabla tipo_grupo_sanguineo.';
COMMENT ON COLUMN terceros.grupo_sanguineo.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.grupo_sanguineo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.grupo_sanguineo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--Creacion tabla grupo_etnico

CREATE SEQUENCE terceros.grupo_etnico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.grupo_etnico(
	id integer NOT NULL DEFAULT nextval('terceros.grupo_etnico_id_seq'::regclass),
    tipo_grupo_etnico_id integer NOT NULL,
    tercero_id integer NOT NULL,	
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_grupo_etnico PRIMARY KEY (id),
    CONSTRAINT fk_tipo_grupo_etnico_grupo_etnico FOREIGN KEY (tipo_grupo_etnico_id) REFERENCES terceros.tipo_grupo_etnico(id),
    CONSTRAINT fk_tercero_grupo_etnico FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.grupo_etnico IS 'Tabla que relaciona a un tercero con un tipo_grupo_etnico.';
COMMENT ON COLUMN terceros.grupo_etnico.id IS 'Identificador unico de la tabla grupo_etnico.';
COMMENT ON COLUMN terceros.grupo_etnico.tipo_grupo_etnico_id IS 'Identificador unico de la tabla tipo_grupo_etnico.';
COMMENT ON COLUMN terceros.grupo_etnico.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.grupo_etnico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.grupo_etnico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--Creacion tabla numero_identificacion

CREATE SEQUENCE terceros.numero_identificacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.numero_identificacion(
	id integer NOT NULL DEFAULT nextval('terceros.numero_identificacion_id_seq'::regclass),
    tipo_documento_id integer NOT NULL,
    tercero_id integer NOT NULL,
    numero character varying(50) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_numero_identificacion PRIMARY KEY (id),
    CONSTRAINT fk_tipo_documento_numero_identificacion FOREIGN KEY (tipo_documento_id) REFERENCES terceros.tipo_documento(id),
    CONSTRAINT fk_tercero_numero_identificacion FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.numero_identificacion IS 'Tabla que almacena el numero de identificación de un tercero y lo relaciona con su tipo.';
COMMENT ON COLUMN terceros.numero_identificacion.id IS 'Identificador unico de la tabla tipo_documento.';
COMMENT ON COLUMN terceros.numero_identificacion.tipo_documento_id IS 'Identificador unico de la tabla tipo_documento';
COMMENT ON COLUMN terceros.numero_identificacion.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.numero_identificacion.numero IS 'Numero de Identificacion del Tercero.';
COMMENT ON COLUMN terceros.numero_identificacion.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.numero_identificacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.numero_identificacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--Creacion tabla discapacidad

CREATE SEQUENCE terceros.discapacidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.discapacidad(
	id integer NOT NULL DEFAULT nextval('terceros.discapacidad_id_seq'::regclass),
    tipo_discapacidad_id integer NOT NULL,
    tercero_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_discapacidad PRIMARY KEY (id),
    CONSTRAINT fk_tipo_discapacidad_discapacidad FOREIGN KEY (tipo_discapacidad_id) REFERENCES terceros.tipo_discapacidad(id),
    CONSTRAINT fk_tercero_discapacidad FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.discapacidad IS 'Tabla que relaciona a un tercero con un tipo de discapacidad.';
COMMENT ON COLUMN terceros.discapacidad.id IS 'Identificador unico de la tabla discapacidad.';
COMMENT ON COLUMN terceros.discapacidad.tipo_discapacidad_id IS 'Identificador unico de la tabla tipo_discapacidad';
COMMENT ON COLUMN terceros.discapacidad.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.discapacidad.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.discapacidad.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.discapacidad.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--Creacion tabla estado_civil_persona

CREATE SEQUENCE terceros.estado_civil_persona_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.estado_civil_persona(
	id integer NOT NULL DEFAULT nextval('terceros.estado_civil_persona_id_seq'::regclass),
    estado_civil_id integer NOT NULL,
    tercero_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_estado_civil_persona PRIMARY KEY (id),
    CONSTRAINT fk_estado_civil_estado_civil_persona FOREIGN KEY (estado_civil_id) REFERENCES terceros.estado_civil(id),
    CONSTRAINT fk_tercero_estado_civil FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.estado_civil_persona IS 'Tabla que relaciona a un tercero con un tipo de estado civil.';
COMMENT ON COLUMN terceros.estado_civil_persona.id IS 'Identificador unico de la tabla estado_civil_id.';
COMMENT ON COLUMN terceros.estado_civil_persona.estado_civil_id IS 'Identificador unico de la tabla estado_civil';
COMMENT ON COLUMN terceros.estado_civil_persona.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.estado_civil_persona.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.estado_civil_persona.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.estado_civil_persona.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--Creacion tabla genero_persona

CREATE SEQUENCE terceros.genero_persona_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.genero_persona(
	id integer NOT NULL DEFAULT nextval('terceros.genero_persona_id_seq'::regclass),
    genero_id integer NOT NULL,
    tercero_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_genero_persona PRIMARY KEY (id),
    CONSTRAINT fk_genero_genero_persona FOREIGN KEY (genero_id) REFERENCES terceros.genero(id),
    CONSTRAINT fk_tercero_genero_persona FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.genero_persona IS 'Tabla que relaciona a un tercero con un genero.';
COMMENT ON COLUMN terceros.genero_persona.id IS 'Identificador unico de la tabla genero_persona.';
COMMENT ON COLUMN terceros.genero_persona.genero_id IS 'Identificador unico de la tabla genero';
COMMENT ON COLUMN terceros.genero_persona.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.genero_persona.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.genero_persona.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.genero_persona.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||


--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||


