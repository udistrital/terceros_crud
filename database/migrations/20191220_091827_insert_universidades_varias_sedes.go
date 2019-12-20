package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertUniversidadesVariasSedes_20191220_091827 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertUniversidadesVariasSedes_20191220_091827{}
	m.Created = "20191220_091827"

	migration.Register("InsertUniversidadesVariasSedes_20191220_091827", m)
}

// Run the migrations
func (m *InsertUniversidadesVariasSedes_20191220_091827) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('SERVICIO NACIONAL DE APRENDIZAJE-SENA-', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_OFICIAL'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, digito_verificacion, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 899999034,1, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BOGOTA D.C.\",\"SNIES\":\"9110\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"CALI\",\"SNIES\":\"9111\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins6 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"MEDELLIN\",\"SNIES\":\"9112\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins7 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"MOSQUERA\",\"SNIES\":\"9113\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins8 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"ARMENIA.\",\"SNIES\":\"9114\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"GUADALAJARA DE BUGA\",\"SNIES\":\"9115\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('UNIVERSIDAD DE ANTIOQUIA', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_OFICIAL'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, digito_verificacion, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 890980040,8, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"MEDELLIN\",\"SNIES\":\"1201\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"CARMEN DE VIBORAL\",\"SNIES\":\"1219\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins6 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"ANDES\",\"SNIES\":\"1220\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins7 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"CAUCASIA\",\"SNIES\":\"1221\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins8 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"PUERTO BERRIO\",\"SNIES\":\"1222\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins9 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"TURBO\",\"SNIES\":\"1223\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"SANTA FE DE ANTIOQUIA\",\"SNIES\":\"9125\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('UNIVERSIDAD NACIONAL DE COLOMBIA', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_OFICIAL'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, digito_verificacion, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 899999063,3, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BOGOTA D.C.\",\"SNIES\":\"1101\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"MEDELLIN\",\"SNIES\":\"1102\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins6 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"MANIZALES\",\"SNIES\":\"1103\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins7 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"PALMIRA\",\"SNIES\":\"1104\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins8 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"LETICIA\",\"SNIES\":\"1125\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins9 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"SAN ANDRES\",\"SNIES\":\"1126\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins10 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"TUMACO\",\"SNIES\":\"9920\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"LA PAZ\",\"SNIES\":\"9933\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('UNIVERSIDAD DEL VALLE', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_OFICIAL'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, digito_verificacion, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 890399010,6, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"CALI\",\"SNIES\":\"1203\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"GUADALAJARA DE BUGA\",\"SNIES\":\"9908\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins6 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"ZARZAL\",\"SNIES\":\"9909\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins7 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BUENAVENTURA\",\"SNIES\":\"9911\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"PALMIRA\",\"SNIES\":\"9912\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('UNIVERSIDAD COOPERATIVA DE COLOMBIA', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_PRIVADA'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, digito_verificacion, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 860029924,7, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"MEDELLIN\",\"SNIES\":\"1816\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BUCARAMANGA\",\"SNIES\":\"1817\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins6 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BOGOTA D.C.\",\"SNIES\":\"1818\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins7 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BARRANCABERMEJA\",\"SNIES\":\"1819\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"SANTA MARTA\",\"SNIES\":\"1820\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('UNIVERSIDAD DE SAN BUENAVENTURA', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_PRIVADA'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, digito_verificacion, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 890307400, 1, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"CALI\",\"SNIES\":\"1716\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"MEDELLIN\",\"SNIES\":\"1717\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins6 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BOGOTA D.C.\",\"SNIES\":\"1718\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"CARTAGENA\",\"SNIES\":\"1724\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('UNIVERSIDAD PEDAGOGICA Y TECNOLOGICA DE COLOMBIA - UPTC', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_OFICIAL'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, digito_verificacion, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 891800330, 1, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"TUNJA\",\"SNIES\":\"1106\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"DUITAMA\",\"SNIES\":\"1107\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins6 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"SOGAMOSO\",\"SNIES\":\"1108\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"CHIQUINQUIRA\",\"SNIES\":\"1109\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('UNIVERSIDAD PONTIFICIA BOLIVARIANA', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_PRIVADA'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, digito_verificacion, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 890902922, 6, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"MEDELLIN\",\"SNIES\":\"1710\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BUCARAMANGA\",\"SNIES\":\"1723\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins6 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"MONTERIA\",\"SNIES\":\"1727\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"PALMIRA\",\"SNIES\":\"1730\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('UNIVERSIDAD SANTO TOMAS', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_PRIVADA'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, digito_verificacion, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 860012357, 6, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BOGOTA D.C.\",\"SNIES\":\"1704\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"BUCARAMANGA\",\"SNIES\":\"1705\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"TUNJA\",\"SNIES\":\"1732\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
	m.SQL("WITH ins1 AS ( INSERT INTO terceros.tercero(nombre_completo, activo, tipo_contribuyente_id, fecha_creacion, fecha_modificacion) VALUES ('UNIVERSIDAD DE CUNDINAMARCA-UDEC', TRUE, (SELECT id FROM terceros.tipo_contribuyente WHERE codigo_abreviacion = 'P_JURIDICA'), LOCALTIMESTAMP, LOCALTIMESTAMP) RETURNING id AS user_id ) , ins2 AS ( INSERT INTO terceros.tercero_tipo_tercero(tercero_id, tipo_tercero_id, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_tercero WHERE codigo_abreviacion = 'UNIVERSIDAD_OFICIAL'), TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins3 AS ( INSERT INTO terceros.datos_identificacion(tercero_id, tipo_documento_id, numero, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.tipo_documento WHERE codigo_abreviacion = 'NIT'), 890680062, TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins4 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"FUSAGASUGA.\",\"SNIES\":\"1214\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) , ins5 AS ( INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"GIRARDOT\",\"SNIES\":\"1215\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1 ) INSERT INTO terceros.info_complementaria_tercero (tercero_id, info_complementaria_id, dato, activo, fecha_creacion, fecha_modificacion) SELECT user_id, (SELECT id FROM terceros.info_complementaria WHERE codigo_abreviacion = 'DIRECCIÓN'), '{\"SEDE\":\"UBATE\",\"SNIES\":\"1216\"}', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP FROM ins1;")
}

// Reverse the migrations
func (m *InsertUniversidadesVariasSedes_20191220_091827) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tercero WHERE id = 10000;")
}