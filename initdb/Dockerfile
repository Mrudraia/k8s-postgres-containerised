FROM postgres:13.1
ENV POSTGRES_USER postgres
ENV POSTGRES_DB go-postgres-app
ENV POSTGRES_PASSWORD password 
VOLUME /var/lib/posgresql/data
COPY init.sql /docker-entrypoint-initdb.d/