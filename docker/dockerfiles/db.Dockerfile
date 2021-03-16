FROM postgresql:13
ENV POSTGRES_PASSWORD=1234qwer!
ENV POSTGRES_USER=dbuser
ENV POSTGRES_DB=lookum
ENV PGPASSWORD=1234qwer!
ENV PGUSER=postgres
ENV PGPORT=5432

COPY --chown=postgres ./db/sql /docker-entrypoint-initdb.d