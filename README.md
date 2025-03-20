# gator
rss feed agregator

#requisitos:
- necesita tener postgres instalado, y una base de datos creada
- hay un archivo de configuración: .gatorconfig en el home del usuario, el mismo es un archivo json con el string de conexión a  la BD, y el usuario de gator activo:
contenido de .gatorconfig.json de ejemplo:
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":"fermar"}



