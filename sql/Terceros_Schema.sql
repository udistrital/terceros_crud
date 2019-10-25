--Schema Terceros :)
CREATE SCHEMA terceros;

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



--Creacion tabla grupo_info_complementaria
CREATE SEQUENCE terceros.grupo_info_complementaria_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.grupo_info_complementaria(
	id integer NOT NULL DEFAULT nextval('terceros.grupo_info_complementaria_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_grupo_info_complementaria PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.grupo_info_complementaria IS 'Tabla que parametriza la información complementaria de un tercero que pueda ser relacionada a un mismo grupo o tipo como pueden ser: discapacidades, estados civiles, grupos sanguíneo, grupos étnicos, géneros entre otros. También se requiere un grupo para valores cuyos datos internos no van a ser parametrizados.';
COMMENT ON COLUMN terceros.grupo_info_complementaria.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.grupo_info_complementaria.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del grupo o tipo.';
COMMENT ON COLUMN terceros.grupo_info_complementaria.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.grupo_info_complementaria.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.grupo_info_complementaria.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.grupo_info_complementaria.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';



--Creacion tabla info_complementaria
CREATE SEQUENCE terceros.info_complementaria_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.info_complementaria(
	id integer NOT NULL DEFAULT nextval('terceros.info_complementaria_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
    tipo_de_dato character varying(20), --Evaluar uso de UNIQUE KEY
	grupo_info_complementaria_id integer NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_info_complementaria PRIMARY KEY (id),
	CONSTRAINT fk_grupo_info_complementaria_info_complementaria FOREIGN KEY (grupo_info_complementaria_id) REFERENCES terceros.grupo_info_complementaria(id)
);
COMMENT ON TABLE terceros.info_complementaria IS 'Tabla que parametriza los tipos concretos de grupos sanguineo, grupos etnicos, discapacidades, estados civiles, generos y cualquier informacion complementaria relacionada a un tercero.';
COMMENT ON COLUMN terceros.info_complementaria.id IS 'Identificador unico de la tabla info_complementaria.';
COMMENT ON COLUMN terceros.info_complementaria.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.info_complementaria.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.info_complementaria.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.info_complementaria.tipo_de_dato IS 'Valor que indica el tipo de dato almacenado en la tabla info_complementaria_tercero en la columna dato.';
COMMENT ON COLUMN terceros.info_complementaria.grupo_info_complementaria_id IS 'Identificador unico de la tabla tipo_informacion_basica.';
COMMENT ON COLUMN terceros.info_complementaria.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.info_complementaria.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';



--Creacion tabla tipo_tercero
CREATE SEQUENCE terceros.tipo_tercero_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tipo_tercero(
	id integer NOT NULL DEFAULT nextval('terceros.tipo_tercero_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_tercero PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.tipo_tercero IS 'Tabla que parametriza a los tipos de terceros, sean: Personas/Empresas/Eps/Caja Compensacion/Arl, Entidades de caracter Publico/Privado/Mixto entre otros.';
COMMENT ON COLUMN terceros.tipo_tercero.id IS 'Identificador unico de la tabla tipo_tercero.';
COMMENT ON COLUMN terceros.tipo_tercero.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.tipo_tercero.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del tipo tercero.';
COMMENT ON COLUMN terceros.tipo_tercero.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.tipo_tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tipo_tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tipo_tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';



--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||



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
	lugar_origen integer,
	fecha_nacimiento TIMESTAMP,
    activo boolean NOT NULL,
    tipo_contribuyente_id integer NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,    
	CONSTRAINT pk_tercero PRIMARY KEY (id),
    CONSTRAINT fk_tipo_contribuyente_tercero FOREIGN KEY (tipo_contribuyente_id) REFERENCES terceros.tipo_contribuyente(id)
);
COMMENT ON TABLE terceros.tercero IS 'Tabla que parametriza los Tipos de contribuyentes: PersonaNatural - Persona Juridica';
COMMENT ON COLUMN terceros.tercero.id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.tercero.nombre_completo IS 'Campo obligatorio de la tabla que indica el nombre de la persona natural o juridica';
COMMENT ON COLUMN terceros.tercero.primer_nombre IS 'Primer nombre persona natural.';
COMMENT ON COLUMN terceros.tercero.segundo_nombre IS 'Segundo nombre persona natural.';
COMMENT ON COLUMN terceros.tercero.primer_apellido IS 'Primer apellido persona natural.';
COMMENT ON COLUMN terceros.tercero.segundo_apellido IS 'Segundo apellido persona natural.';
COMMENT ON COLUMN terceros.tercero.lugar_origen IS 'Id del pais de origen de la persona / empresa / entidad.';
COMMENT ON COLUMN terceros.tercero.fecha_nacimiento IS 'Campo que indica la fecha de nacimiento de la persona natural / indica la fecha de creacion de la empresa.';
COMMENT ON COLUMN terceros.tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tercero.tipo_contribuyente_id IS 'Identificador de la tabla tipo_contribuyente, diferencia entre persona_natural, persona_juridica.';
COMMENT ON COLUMN terceros.tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
--|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||


--Creacion tabla tercero_tipo_tercero
CREATE SEQUENCE terceros.tercero_tipo_tercero_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tercero_tipo_tercero(
	id integer NOT NULL DEFAULT nextval('terceros.tercero_tipo_tercero_id_seq'::regclass),
    tercero_id integer NOT NULL,
	tipo_tercero_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tercero_tipo_tercero PRIMARY KEY (id),
	CONSTRAINT fk_tercero_tercero_tipo_tercero FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id),
	CONSTRAINT fk_tipo_tercero_tercero_tipo_tercero FOREIGN KEY (tipo_tercero_id) REFERENCES terceros.tipo_tercero(id)
);
COMMENT ON TABLE terceros.tercero_tipo_tercero IS 'Tabla que relaciona un tercero con uno o mas tipos de tercero Ej: Un colegio es de tipo colegio / organizacion  y a su vez es de carácter publico-privado-mixto.';
COMMENT ON COLUMN terceros.tercero_tipo_tercero.id IS 'Identificador unico de la tabla tercero_tipo_tercero.';
COMMENT ON COLUMN terceros.tercero_tipo_tercero.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.tercero_tipo_tercero.tipo_tercero_id IS 'Identificador unico de la tabla tipo_tercero.';
COMMENT ON COLUMN terceros.tercero_tipo_tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tercero_tipo_tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tercero_tipo_tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';



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
	ciudad_expedicion integer,
	fecha_expedicion TIMESTAMP,
	activo boolean NOT NULL,
	documento_soporte integer,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_datos_identificacion PRIMARY KEY (id),
    CONSTRAINT fk_tipo_documento_datos_identificacion FOREIGN KEY (tipo_documento_id) REFERENCES terceros.tipo_documento(id),
    CONSTRAINT fk_tercero_datos_identificacion FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.datos_identificacion IS 'Tabla que almacena el numero de identificación de un tercero y lo relaciona con su tipo pudiendo ser este CC, CE, NIT entre otros.';
COMMENT ON COLUMN terceros.datos_identificacion.id IS 'Identificador unico de la tabla datos_identificacion.';
COMMENT ON COLUMN terceros.datos_identificacion.tipo_documento_id IS 'Identificador unico de la tabla tipo_documento';
COMMENT ON COLUMN terceros.datos_identificacion.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.datos_identificacion.numero IS 'Numero de Identificacion del Tercero.';
COMMENT ON COLUMN terceros.datos_identificacion.digito_verificacion IS 'Dígito de verificación de la DIAN. Toda persona natural o jurídica que tenga un RUT debe tener un DV (NOT NULL para personas_juridicas).';
COMMENT ON COLUMN terceros.datos_identificacion.ciudad_expedicion IS 'Valor que almacena el ID de la ciudad en que fue expedido el documento.';
COMMENT ON COLUMN terceros.datos_identificacion.fecha_expedicion IS 'Fecha en la que fue expedido el documento de identificación.';
COMMENT ON COLUMN terceros.datos_identificacion.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.datos_identificacion.documento_soporte IS 'Identificador asociado al soporte del documento, correspondiente para el tercero.';
COMMENT ON COLUMN terceros.datos_identificacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.datos_identificacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';



--Creacion tabla info_complementaria_tercero
CREATE SEQUENCE terceros.info_complementaria_tercero_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.info_complementaria_tercero(
	id integer NOT NULL DEFAULT nextval('terceros.info_complementaria_tercero_id_seq'::regclass),
	tercero_id integer NOT NULL,
	info_complementaria_id integer NOT NULL,
    dato json,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_info_complementaria_tercero PRIMARY KEY (id),
	CONSTRAINT fk_tercero_info_complementaria_tercero FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id),
	CONSTRAINT fk_info_complementaria_info_complementaria_tercero FOREIGN KEY (info_complementaria_id) REFERENCES terceros.info_complementaria(id)
);
COMMENT ON TABLE terceros.info_complementaria_tercero IS 'Tabla que relaciona a un Tercero con su información complementaria, el valor concreto con el que se relaciona puede ser encontrado en la columna dato o en la tabla info_complementaria dependiendo de si es parametrisable o no.';
COMMENT ON COLUMN terceros.info_complementaria_tercero.id IS 'Identificador unico de la tabla info_complementaria_tercero.';
COMMENT ON COLUMN terceros.info_complementaria_tercero.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.info_complementaria_tercero.info_complementaria_id IS 'Identificador unico de la tabla info_complementaria';
COMMENT ON COLUMN terceros.info_complementaria_tercero.dato IS 'Campo para guardar informacion que no se pueda parametrizar, el tipo de dato que se guarda en esta estructura se relacionara en la tabla info_complementaria.';
COMMENT ON COLUMN terceros.info_complementaria_tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.info_complementaria_tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.info_complementaria_tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';



--Creacion tabla seguridad_social_tercero
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
	activo boolean NOT NULL,
	fecha_inicio_vinculacion TIMESTAMP NOT NULL,
	fecha_fin_vinculacion TIMESTAMP,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_seguridad_social_tercero PRIMARY KEY (id),
	CONSTRAINT fk_tercero_seguridad_social_tercero FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id),
	CONSTRAINT fk_tercero_entidad_seguridad_social_tercero FOREIGN KEY (tercero_entidad_id) REFERENCES terceros.tercero(id)
);
COMMENT ON TABLE terceros.seguridad_social_tercero IS 'Tabla que relaciona a un Tercero con su:  Eps / Caja Compensacion / Arl / fondo de Pension entre otros.';
COMMENT ON COLUMN terceros.seguridad_social_tercero.id IS 'Identificador unico de la tabla info_complementaria_tercero.';
COMMENT ON COLUMN terceros.seguridad_social_tercero.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.seguridad_social_tercero.tercero_entidad_id IS 'Identificador unico de la tabla tercero que corresponde a la entidad.';
COMMENT ON COLUMN terceros.seguridad_social_tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.seguridad_social_tercero.fecha_inicio_vinculacion IS 'Fecha en la que el tercero se vincula a la entidad.';
COMMENT ON COLUMN terceros.seguridad_social_tercero.fecha_fin_vinculacion IS 'Fecha en la que el tercero se retira/desvincula de la entidad, cambio de activo a FALSE.';
COMMENT ON COLUMN terceros.seguridad_social_tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.seguridad_social_tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--Convertir en Inserts y grupos de datos.
/**
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
);**/


--Creacion tabla tipo_parentesco
CREATE SEQUENCE terceros.tipo_parentesco_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tipo_parentesco(
	id integer NOT NULL DEFAULT nextval('terceros.tipo_parentesco_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_parentesco PRIMARY KEY (id)
);
COMMENT ON TABLE terceros.tipo_parentesco IS 'Tabla que parametriza a los tipos de parentescos que puede tener un tercero en relacion a otro, sean estos: Madre, Padre, Conyuge, Hijos, entre otros.';
COMMENT ON COLUMN terceros.tipo_parentesco.id IS 'Identificador unico de la tabla tipo_parentesco.';
COMMENT ON COLUMN terceros.tipo_parentesco.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN terceros.tipo_parentesco.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del tipo de parentesco.';
COMMENT ON COLUMN terceros.tipo_parentesco.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.tipo_parentesco.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tipo_parentesco.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tipo_parentesco.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--Creacion tabla tercero_familiar
CREATE SEQUENCE terceros.tercero_familiar_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE terceros.tercero_familiar(
	id integer NOT NULL DEFAULT nextval('terceros.tercero_familiar_id_seq'::regclass),
    tercero_id integer NOT NULL,
    tercero_familiar_id integer NOT NULL,
    tipo_parentesco_id integer NOT NULL,
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tercero_familiar PRIMARY KEY (id),
    CONSTRAINT fk_tercero_tercero_familiar FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id),
    CONSTRAINT fk_tercero_familiar_tercero_familiar FOREIGN KEY (tercero_familiar_id) REFERENCES terceros.tercero(id),
    CONSTRAINT fk_tipo_parentesco_tercero_familiar FOREIGN KEY (tipo_parentesco_id) REFERENCES terceros.tipo_parentesco(id)
);
COMMENT ON TABLE terceros.tercero_familiar IS 'Tabla que parametriza a los tipos de parentescos que puede tener un tercero en relacion a otro, sean estos: Madre, Padre, Conyuge, Hijos, entre otros.';
COMMENT ON COLUMN terceros.tercero_familiar.id IS 'Identificador unico de la tabla tercero_familiar.';
COMMENT ON COLUMN terceros.tercero_familiar.tercero_id IS 'Identificador unico de la tabla tercero.';
COMMENT ON COLUMN terceros.tercero_familiar.tercero_familiar_id IS 'Identificador unico de la tabla tercero correspondiente al familiar.';
COMMENT ON COLUMN terceros.tercero_familiar.tipo_parentesco_id IS 'Identificador unico de la relacion de parentesco que existe entre tercero_id y tercero_familiar_id.';
COMMENT ON COLUMN terceros.tercero_familiar.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN terceros.tercero_familiar.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN terceros.tercero_familiar.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.tercero_familiar.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
