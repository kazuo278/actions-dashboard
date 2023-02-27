FROM postgres:13

RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
COPY ./sql/01_create_table.sql /docker-entrypoint-initdb.d
#COPY ./sql/02_create_testdata.sql /docker-entrypoint-initdb.d