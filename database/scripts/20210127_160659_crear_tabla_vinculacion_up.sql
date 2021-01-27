CREATE TABLE terceros.vinculacion(
	id serial NOT NULL,
	tercero_principal_id integer NOT NULL,
    tercero_relacionado_id integer,
    tipo_vinculacion_id integer NOT NULL,
    cargo_id integer,
	dependencia_id integer,
	soporte integer,
	periodo_id integer,	
    fecha_inicio_vinculacion TIMESTAMP,
    fecha_fin_vinculacion TIMESTAMP,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_vinculacion PRIMARY KEY (id)    
);

COMMENT ON TABLE terceros.vinculacion IS 'Tabla que muestra el histórico de vinculaciones que tiene un tercero, ya sea con otro tercero o con un cargo dentro de la universidad.';
COMMENT ON COLUMN terceros.vinculacion.id IS 'Identificador unico de la tabla vinculacion';
COMMENT ON COLUMN terceros.vinculacion.tercero_principal_id IS 'Campo obligatorio que referencia a la tabla tercero. Indica qué tercero tiene una vinculación.';
COMMENT ON COLUMN terceros.vinculacion.tercero_relacionado_id IS 'Campo no obligatorio que referencia a la tabla tercero. Es el tercero relacionado al tercero principal.';
COMMENT ON COLUMN terceros.vinculacion.tipo_vinculacion_id  IS 'Campo obligatorio que referencia a un registro de la tabla parametro del esquema parametros. Indica el tipo de vinculación que posee el tercero principal.';
COMMENT ON COLUMN terceros.vinculacion.cargo_id IS 'Campo no obligatorio que referencia a un registro de la tabla parametro del esquema parametros. Indica, de aplicarse, qué cargo específico desempeña el tercero principal.';
COMMENT ON COLUMN terceros.vinculacion.dependencia_id IS 'Campo no obligatorio que referencia a esquema oikos. Indica la dependencia en la que el tercero principal desempeña la vinculación.';
COMMENT ON COLUMN terceros.vinculacion.soporte IS 'Campo no obligatorio que referencia a esquema documentos. Relaciona un documento que sirva de soporte para la vinculación indicada en el registro.';
COMMENT ON COLUMN terceros.vinculacion.periodo_id IS 'Campo no obligatorio que referencia a esquema parametros_estandar (tabla unidad). Indica, de ser válido para la vinculación, el periodo al que corresponde.';
COMMENT ON COLUMN terceros.vinculacion.fecha_inicio_vinculacion IS 'Campo no obligatorio que indica la fecha exacta del inicio de la vinculación del tercero. ';
COMMENT ON COLUMN terceros.vinculacion.fecha_fin_vinculacion IS 'Campo no obligatorio que indica la fecha exacta en la que termina la vinculación del tercero. ';
COMMENT ON COLUMN terceros.vinculacion.activo IS 'Campo obligatorio que indica el estado de la vinculación.';
COMMENT ON COLUMN terceros.vinculacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN terceros.vinculacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

ALTER TABLE terceros.vinculacion ADD CONSTRAINT fk_vinculacion_tercero_principal FOREIGN KEY (tercero_principal_id)
REFERENCES terceros.tercero (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE terceros.vinculacion ADD CONSTRAINT fk_vinculacion_tercero_relacionado FOREIGN KEY (tercero_relacionado_id)
REFERENCES terceros.tercero (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;