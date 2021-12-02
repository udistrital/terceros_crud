# terceros_crud

API para la gestión de terceros y tipos de terceros vinculados a la Universidad Distrital

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

- [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
- [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno

```shell
# parametros de api
TERCEROS_CRUD_HTTP_PORT=[Puerto de exposición del API]
TERCEROS_CRUD_RUN_MODE=[Modo de ejecución del API]
# paramametros de bd
TERCEROS_CRUD_PGUSER=[Usuario de BD]
TERCEROS_CRUD_PGPASS=[Contraseña del usaurio de BD]
TERCEROS_CRUD_PGHOST=[URL, Dominio o EndPoint de la BD]
TERCEROS_CRUD_PGPORT=[Puerto de la BD]
TERCEROS_CRUD_PGDB=[Nombre de Base de Datos]
TERCEROS_CRUD_PGSCHEMA=[Nombre del Esquema de Base de Datos]
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con `TERCEROS_CRUD_...`

### Ejecución del Proyecto

```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/terceros_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/terceros_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
TERCEROS_CRUD_PORT=8080 TERCEROS_CRUD_SOME_VARIABLE bee run
```

### Ejecución Dockerfile

```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose

```shell
# 1. Obtener repositorio
git clone https://github.com/udistrital/terceros_crud.git
# 2. Ir a la carpeta del repositorio
cd $GOPATH/src/github.com/terceros_crud
# 3. Cambiar a la rama develop
git checkout develop
# 4. Crear red back_end
docker network create back_end
# 5. Ejecutar docker compose
docker-compose up --build
```

### Ejecución Pruebas

Pruebas unitarias

```shell
# En Proceso
```

## Modelo de Datos

<!-- TODO: Limpiar versiones anteriores (y el siguiente link), para eso es Git -->

[Modelo de Datos Relacional terceros_crud](./sql/Terceros_Schema.png) -
[SVG](database/terceros.svg) -
[PGModeler](database/terceros.dbm)

## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/terceros_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/terceros_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/terceros_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/terceros_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/terceros_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/terceros_crud) |

## Licencia

This file is part of terceros_crud.

terceros_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

terceros_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with terceros_crud. If not, see https://www.gnu.org/licenses/.
