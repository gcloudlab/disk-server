FROM mysql:5.7

# 设置环境变量
ENV MYSQL_DATABASE=gcloud \
  MYSQL_USER=song \
  MYSQL_PASSWORD=123456 \
  MYSQL_ROOT_PASSWORD=123456

COPY ./init.sql /docker-entrypoint-initdb.d/
COPY mysql.cnf /etc/mysql/mysql.conf.d/mysql.cnf

EXPOSE 3306