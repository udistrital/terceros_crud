-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.0-alpha1
-- PostgreSQL version: 9.6
-- Project Site: pgmodeler.com.br
-- Model Author: ---


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: "Prupuesta_persona_con_relacion_cargo_universidad" | type: DATABASE --
-- -- DROP DATABASE IF EXISTS "Prupuesta_persona_con_relacion_cargo_universidad";
-- CREATE DATABASE "Prupuesta_persona_con_relacion_cargo_universidad"
-- ;
-- -- ddl-end --
-- 

-- object: core_new | type: SCHEMA --
-- DROP SCHEMA IF EXISTS core_new CASCADE;
CREATE SCHEMA core_new;
-- ddl-end --

SET search_path TO pg_catalog,public,core_new;
-- ddl-end --

-- object: core_new.persona | type: TABLE --
-- DROP TABLE IF EXISTS core_new.persona CASCADE;
CREATE TABLE core_new.persona(
	id bigserial NOT NULL,
	primer_nombre character varying(50) NOT NULL,
	segundo_nombre character varying(50),
	primer_apellido character varying(50) NOT NULL,
	segundo_apellido character varying(50),
	fecha_nacimiento date,
	usuario varchar(50),
	ente integer,
	foto integer,
	CONSTRAINT pk_persona PRIMARY KEY (id),
	CONSTRAINT uq_usuario UNIQUE (usuario),
	CONSTRAINT uq_ente UNIQUE (ente)

);
-- ddl-end --
COMMENT ON TABLE core_new.persona IS 'tabla que almacena la informacion basica geral relacionada a las personas';
-- ddl-end --
COMMENT ON COLUMN core_new.persona.id IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON COLUMN core_new.persona.primer_nombre IS 'Corresponde al primer nombre de la persona.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona.segundo_nombre IS 'Corresponde al segundo nombre de la persona.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona.primer_apellido IS 'Corresponde al primer apellidode la persona.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona.segundo_apellido IS 'Corresponde al segundo apellido de la persona.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona.fecha_nacimiento IS 'Fecha de nacimiento de la persona.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona.usuario IS 'identificador del usuario relacionado al sistema de autenticación';
-- ddl-end --
COMMENT ON COLUMN core_new.persona.foto IS 'Foto de la perosna';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona ON core_new.persona  IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --

-- object: core_new.genero | type: TABLE --
-- DROP TABLE IF EXISTS core_new.genero CASCADE;
CREATE TABLE core_new.genero(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_genero PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.genero IS 'Almacena los tipos de genero ';
-- ddl-end --
COMMENT ON COLUMN core_new.genero.id IS 'identificador de la tabla genero';
-- ddl-end --
COMMENT ON COLUMN core_new.genero.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro de tipo genero';
-- ddl-end --
COMMENT ON COLUMN core_new.genero.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN core_new.genero.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN core_new.genero.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN core_new.genero.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_genero ON core_new.genero  IS 'llave primaria de la tabla tipo genero';
-- ddl-end --

-- object: core_new.persona_genero | type: TABLE --
-- DROP TABLE IF EXISTS core_new.persona_genero CASCADE;
CREATE TABLE core_new.persona_genero(
	id serial NOT NULL,
	genero integer NOT NULL,
	persona integer NOT NULL,
	CONSTRAINT pk_persona_genero PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.persona_genero IS 'Tabla que relaciona la persona con el correspondiente genero';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_genero.id IS 'Identificador de la tabla persona_genero';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_genero.genero IS 'Identificador de la tabla genero.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_genero.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_genero ON core_new.persona_genero  IS 'Restriccion de llave primaria de la tabla persona genero.';
-- ddl-end --

-- object: core_new.grupo_etnico | type: TABLE --
-- DROP TABLE IF EXISTS core_new.grupo_etnico CASCADE;
CREATE TABLE core_new.grupo_etnico(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_grupo_etnico PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.grupo_etnico IS 'Almacena los grupos etnicos en los que se catalogan las personas.';
-- ddl-end --
COMMENT ON COLUMN core_new.grupo_etnico.id IS 'identificador de la tabla grupo_etnico';
-- ddl-end --
COMMENT ON COLUMN core_new.grupo_etnico.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro.';
-- ddl-end --
COMMENT ON COLUMN core_new.grupo_etnico.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN core_new.grupo_etnico.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN core_new.grupo_etnico.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN core_new.grupo_etnico.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_grupo_etnico ON core_new.grupo_etnico  IS 'llave primaria de la tabla';
-- ddl-end --

-- object: core_new.persona_grupo_etnico | type: TABLE --
-- DROP TABLE IF EXISTS core_new.persona_grupo_etnico CASCADE;
CREATE TABLE core_new.persona_grupo_etnico(
	id serial NOT NULL,
	grupo_etnico integer NOT NULL,
	persona integer NOT NULL,
	CONSTRAINT pk_persona_grupo_etnico PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.persona_grupo_etnico IS 'Tabla que relaciona la persona con el grupo étnico.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_grupo_etnico.id IS 'Identificador de la tabla persona_grupo_etnico';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_grupo_etnico.grupo_etnico IS 'Identificador de la tbala grupo_etnico';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_grupo_etnico.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_grupo_etnico ON core_new.persona_grupo_etnico  IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --

-- object: core_new.estado_civil | type: TABLE --
-- DROP TABLE IF EXISTS core_new.estado_civil CASCADE;
CREATE TABLE core_new.estado_civil(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_estado_civil PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.estado_civil IS 'Almacena los parametros de estado civil.';
-- ddl-end --
COMMENT ON COLUMN core_new.estado_civil.id IS 'identificador de la tabla estado_civil';
-- ddl-end --
COMMENT ON COLUMN core_new.estado_civil.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro.';
-- ddl-end --
COMMENT ON COLUMN core_new.estado_civil.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN core_new.estado_civil.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN core_new.estado_civil.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN core_new.estado_civil.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_civil ON core_new.estado_civil  IS 'llave primaria de la tabla';
-- ddl-end --

-- object: core_new.persona_estado_civil | type: TABLE --
-- DROP TABLE IF EXISTS core_new.persona_estado_civil CASCADE;
CREATE TABLE core_new.persona_estado_civil(
	id serial NOT NULL,
	estado_civil integer NOT NULL,
	persona integer NOT NULL,
	CONSTRAINT pk_persona_estado_civil PRIMARY KEY (id),
	CONSTRAINT uq_estado_civil_persona UNIQUE (estado_civil,persona)

);
-- ddl-end --
COMMENT ON TABLE core_new.persona_estado_civil IS 'Tabla que relaciona la persona con el estado civil.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_estado_civil.id IS 'Identificador de la tabla persona_estado_civil';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_estado_civil.estado_civil IS 'Identificador foraneo a la tabla parametrica estado_civil.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_estado_civil.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_estado_civil ON core_new.persona_estado_civil  IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --

-- object: core_new.tipo_discapacidad | type: TABLE --
-- DROP TABLE IF EXISTS core_new.tipo_discapacidad CASCADE;
CREATE TABLE core_new.tipo_discapacidad(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_tipo_discapacidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.tipo_discapacidad IS 'Almacena los parametros de tipo discapacidad.';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_discapacidad.id IS 'identificador unico de la tabla tipo_discapacidad';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_discapacidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro.';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_discapacidad.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_discapacidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_discapacidad.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_discapacidad.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_discapacidad ON core_new.tipo_discapacidad  IS 'llave primaria de la tabla';
-- ddl-end --

-- object: core_new.persona_tipo_discapacidad | type: TABLE --
-- DROP TABLE IF EXISTS core_new.persona_tipo_discapacidad CASCADE;
CREATE TABLE core_new.persona_tipo_discapacidad(
	id serial NOT NULL,
	tipo_discapacidad integer NOT NULL,
	persona integer NOT NULL,
	CONSTRAINT pk_persona_tipo_discapacidad PRIMARY KEY (id),
	CONSTRAINT uq_discapacidad_persona UNIQUE (tipo_discapacidad,persona)

);
-- ddl-end --
COMMENT ON TABLE core_new.persona_tipo_discapacidad IS 'Tabla que relaciona la persona con el tipo de discapacidad';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_tipo_discapacidad.id IS 'Identificador  de la tabla persona_tipo_discapacidad';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_tipo_discapacidad.tipo_discapacidad IS 'Identificador de la tabla tipo_discapacidad';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_tipo_discapacidad.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_tipo_discapacidad ON core_new.persona_tipo_discapacidad  IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --

-- object: core_new.perfil_profesional | type: TABLE --
-- DROP TABLE IF EXISTS core_new.perfil_profesional CASCADE;
CREATE TABLE core_new.perfil_profesional(
	id integer NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_perfil PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.perfil_profesional IS 'tabla en la que se registran los diferntes perfiles profesionales para las personas.';
-- ddl-end --
COMMENT ON COLUMN core_new.perfil_profesional.id IS 'identificador de la tabla perfil_profesional';
-- ddl-end --
COMMENT ON COLUMN core_new.perfil_profesional.nombre IS 'Campo que indica el nombre del perfil.';
-- ddl-end --
COMMENT ON COLUMN core_new.perfil_profesional.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN core_new.perfil_profesional.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN core_new.perfil_profesional.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN core_new.perfil_profesional.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_perfil ON core_new.perfil_profesional  IS 'llave primaria de la tabla';
-- ddl-end --

-- object: core_new.persona_perfil_profesional | type: TABLE --
-- DROP TABLE IF EXISTS core_new.persona_perfil_profesional CASCADE;
CREATE TABLE core_new.persona_perfil_profesional(
	id serial NOT NULL,
	perfil_profesional integer NOT NULL,
	persona integer NOT NULL,
	CONSTRAINT pk_persona_perfil_profesional PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.persona_perfil_profesional IS 'Tabla que relaciona la persona con el perfil.';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_perfil_profesional.id IS 'Identificador primario de la tabla';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_perfil_profesional.perfil_profesional IS 'Identificador de la tabla perfil_profesional';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_perfil_profesional.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_perfil_profesional ON core_new.persona_perfil_profesional  IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --

-- object: core_new.tipo_relacion_personas | type: TABLE --
-- DROP TABLE IF EXISTS core_new.tipo_relacion_personas CASCADE;
CREATE TABLE core_new.tipo_relacion_personas(
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_tipo_relacion_personas PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.tipo_relacion_personas IS 'Tabla parametrica que indica el tipo de relacion que pueden tener las personas. Ejemplo: Madre, Padre, Conyugue';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_relacion_personas.id IS 'identificador de la tabla tipo_relacion_personas';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_relacion_personas.nombre IS 'Campo que indica el nombre del tipo de relación';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_relacion_personas.descripcion IS 'Campo para agregar una descripcion acerca del tipo de relacion';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_relacion_personas.codigo_abreviacion IS 'Codigo de abreviacion opcional para referirse al registro';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_relacion_personas.activo IS 'Campo que indica si el tipo de relación esta activo';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_relacion_personas.numero_orden IS 'Campo para organizar el orden en el que se presentan los registros';
-- ddl-end --

-- object: core_new.relacion_personas | type: TABLE --
-- DROP TABLE IF EXISTS core_new.relacion_personas CASCADE;
CREATE TABLE core_new.relacion_personas(
	id serial NOT NULL,
	persona_principal integer NOT NULL,
	persona_relacionada integer NOT NULL,
	tipo_relacion_personas integer NOT NULL,
	CONSTRAINT pk_relacion_personas PRIMARY KEY (id),
	CONSTRAINT uq_relacion_personas UNIQUE (persona_principal,persona_relacionada,tipo_relacion_personas)

);
-- ddl-end --
COMMENT ON TABLE core_new.relacion_personas IS 'Describe la relación de parentesco que tiene una persona principal con una persona relacionada';
-- ddl-end --
COMMENT ON COLUMN core_new.relacion_personas.id IS 'identificador de la tabla relacion_personas';
-- ddl-end --
COMMENT ON COLUMN core_new.relacion_personas.persona_principal IS 'identificador de la tabla persona, que indica la persona a la que se le indica la relacion';
-- ddl-end --
COMMENT ON COLUMN core_new.relacion_personas.persona_relacionada IS 'identificador de la tabla persona, que indica la persona que es relacionada';
-- ddl-end --

-- object: core_new.grupo_sanguineo_persona | type: TABLE --
-- DROP TABLE IF EXISTS core_new.grupo_sanguineo_persona CASCADE;
CREATE TABLE core_new.grupo_sanguineo_persona(
	id integer NOT NULL,
	factor_rh character varying(1) NOT NULL,
	grupo_sanguineo character varying(1) NOT NULL,
	persona bigint NOT NULL,
	CONSTRAINT pk_grupo_sanguineo_persona PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.grupo_sanguineo_persona IS 'Tabla que contiene informacion del grupo sanguineo de la persona';
-- ddl-end --
COMMENT ON COLUMN core_new.grupo_sanguineo_persona.id IS 'Campo de id de la tabla grupo sanguineo person';
-- ddl-end --
COMMENT ON COLUMN core_new.grupo_sanguineo_persona.factor_rh IS 'Campo para el facto del rh (Ejemplo: + -)';
-- ddl-end --
COMMENT ON COLUMN core_new.grupo_sanguineo_persona.grupo_sanguineo IS 'campo para el grupo sanguineo (Ejemplo : A,B,O,AB)';
-- ddl-end --
ALTER TABLE core_new.grupo_sanguineo_persona OWNER TO postgres;
-- ddl-end --

-- object: grupo_sanguineo_persona_persona | type: CONSTRAINT --
-- ALTER TABLE core_new.grupo_sanguineo_persona DROP CONSTRAINT IF EXISTS grupo_sanguineo_persona_persona CASCADE;
ALTER TABLE core_new.grupo_sanguineo_persona ADD CONSTRAINT grupo_sanguineo_persona_persona FOREIGN KEY (persona)
REFERENCES core_new.persona (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_grupo_sanguineo_persona | type: CONSTRAINT --
-- ALTER TABLE core_new.grupo_sanguineo_persona DROP CONSTRAINT IF EXISTS uq_grupo_sanguineo_persona CASCADE;
ALTER TABLE core_new.grupo_sanguineo_persona ADD CONSTRAINT uq_grupo_sanguineo_persona UNIQUE (persona);
-- ddl-end --

-- object: core_new.persona_vinculo_persona | type: TABLE --
-- DROP TABLE IF EXISTS core_new.persona_vinculo_persona CASCADE;
CREATE TABLE core_new.persona_vinculo_persona(
	id integer NOT NULL,
	tipo_vinculo_persona integer NOT NULL,
	dependencia integer NOT NULL,
	persona bigint NOT NULL,
	CONSTRAINT persona_vinculo_persona_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.persona_vinculo_persona IS 'Tabla de rompimiento entre persona y tipo de vinculo con la organizacion';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_vinculo_persona.id IS 'Campo para el idectificador de la tabla persona_vinculo_persona';
-- ddl-end --
COMMENT ON COLUMN core_new.persona_vinculo_persona.dependencia IS 'Campo para relacionar con la tabla dependencia en el esquema espacio_logico, relaciona la depencia (proyecto curricular) con la vinculación de la persona.';
-- ddl-end --
ALTER TABLE core_new.persona_vinculo_persona OWNER TO postgres;
-- ddl-end --

-- object: core_new.tipo_vinculo_persona | type: TABLE --
-- DROP TABLE IF EXISTS core_new.tipo_vinculo_persona CASCADE;
CREATE TABLE core_new.tipo_vinculo_persona(
	id integer NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_tipo_vinculo_persona PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE core_new.tipo_vinculo_persona IS 'Tabla que contiene el tipo de vinculacion de la persona con la organizacion, y tabla parametrica
';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_vinculo_persona.id IS 'Campo de identificacion de la tabla tipo_vinculo_persona';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_vinculo_persona.nombre IS 'Campo que indica el nombre de vinculacion (Ejemplo: Estudiante, contratista ,docente.etc)';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_vinculo_persona.descripcion IS 'Descripción opcional del parámetro';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_vinculo_persona.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_vinculo_persona.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN core_new.tipo_vinculo_persona.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión';
-- ddl-end --
ALTER TABLE core_new.tipo_vinculo_persona OWNER TO postgres;
-- ddl-end --

-- object: persona_vinculo_persona_tipo_vinculo_persona | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_vinculo_persona DROP CONSTRAINT IF EXISTS persona_vinculo_persona_tipo_vinculo_persona CASCADE;
ALTER TABLE core_new.persona_vinculo_persona ADD CONSTRAINT persona_vinculo_persona_tipo_vinculo_persona FOREIGN KEY (tipo_vinculo_persona)
REFERENCES core_new.tipo_vinculo_persona (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: persona_vinculo_persona_persona | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_vinculo_persona DROP CONSTRAINT IF EXISTS persona_vinculo_persona_persona CASCADE;
ALTER TABLE core_new.persona_vinculo_persona ADD CONSTRAINT persona_vinculo_persona_persona FOREIGN KEY (persona)
REFERENCES core_new.persona (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona_genero_genero | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_genero DROP CONSTRAINT IF EXISTS fk_persona_genero_genero CASCADE;
ALTER TABLE core_new.persona_genero ADD CONSTRAINT fk_persona_genero_genero FOREIGN KEY (genero)
REFERENCES core_new.genero (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_genero DROP CONSTRAINT IF EXISTS fk_persona CASCADE;
ALTER TABLE core_new.persona_genero ADD CONSTRAINT fk_persona FOREIGN KEY (persona)
REFERENCES core_new.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_grupo_etnico | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_grupo_etnico DROP CONSTRAINT IF EXISTS fk_grupo_etnico CASCADE;
ALTER TABLE core_new.persona_grupo_etnico ADD CONSTRAINT fk_grupo_etnico FOREIGN KEY (grupo_etnico)
REFERENCES core_new.grupo_etnico (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_grupo_etnico DROP CONSTRAINT IF EXISTS fk_persona CASCADE;
ALTER TABLE core_new.persona_grupo_etnico ADD CONSTRAINT fk_persona FOREIGN KEY (persona)
REFERENCES core_new.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_persona_estado_civil_estado_civil | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_estado_civil DROP CONSTRAINT IF EXISTS fk_persona_estado_civil_estado_civil CASCADE;
ALTER TABLE core_new.persona_estado_civil ADD CONSTRAINT fk_persona_estado_civil_estado_civil FOREIGN KEY (estado_civil)
REFERENCES core_new.estado_civil (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_estado_civil DROP CONSTRAINT IF EXISTS fk_persona CASCADE;
ALTER TABLE core_new.persona_estado_civil ADD CONSTRAINT fk_persona FOREIGN KEY (persona)
REFERENCES core_new.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_persona_tipo_discapacidad | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_tipo_discapacidad DROP CONSTRAINT IF EXISTS fk_persona_tipo_discapacidad CASCADE;
ALTER TABLE core_new.persona_tipo_discapacidad ADD CONSTRAINT fk_persona_tipo_discapacidad FOREIGN KEY (tipo_discapacidad)
REFERENCES core_new.tipo_discapacidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_tipo_discapacidad DROP CONSTRAINT IF EXISTS fk_persona CASCADE;
ALTER TABLE core_new.persona_tipo_discapacidad ADD CONSTRAINT fk_persona FOREIGN KEY (persona)
REFERENCES core_new.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_perfil_profesional | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_perfil_profesional DROP CONSTRAINT IF EXISTS fk_perfil_profesional CASCADE;
ALTER TABLE core_new.persona_perfil_profesional ADD CONSTRAINT fk_perfil_profesional FOREIGN KEY (perfil_profesional)
REFERENCES core_new.perfil_profesional (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona | type: CONSTRAINT --
-- ALTER TABLE core_new.persona_perfil_profesional DROP CONSTRAINT IF EXISTS fk_persona CASCADE;
ALTER TABLE core_new.persona_perfil_profesional ADD CONSTRAINT fk_persona FOREIGN KEY (persona)
REFERENCES core_new.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_tipo_relacion_personas | type: CONSTRAINT --
-- ALTER TABLE core_new.relacion_personas DROP CONSTRAINT IF EXISTS fk_tipo_relacion_personas CASCADE;
ALTER TABLE core_new.relacion_personas ADD CONSTRAINT fk_tipo_relacion_personas FOREIGN KEY (tipo_relacion_personas)
REFERENCES core_new.tipo_relacion_personas (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_persona_principal | type: CONSTRAINT --
-- ALTER TABLE core_new.relacion_personas DROP CONSTRAINT IF EXISTS fk_persona_principal CASCADE;
ALTER TABLE core_new.relacion_personas ADD CONSTRAINT fk_persona_principal FOREIGN KEY (persona_principal)
REFERENCES core_new.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_persona_relacionada | type: CONSTRAINT --
-- ALTER TABLE core_new.relacion_personas DROP CONSTRAINT IF EXISTS fk_persona_relacionada CASCADE;
ALTER TABLE core_new.relacion_personas ADD CONSTRAINT fk_persona_relacionada FOREIGN KEY (persona_relacionada)
REFERENCES core_new.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


