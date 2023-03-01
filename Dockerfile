FROM mysql:5.7

# 设置环境变量
ENV MYSQL_DATABASE=gcloud \
  MYSQL_USER=song \
  MYSQL_PASSWORD=123456 \
  MYSQL_ROOT_PASSWORD=123456

# 复制初始化脚本到容器中
COPY ./init.sql /docker-entrypoint-initdb.d/

EXPOSE 3306

# 容器启动时运行初始化脚本
CMD ["mysqld", "--user=mysql", "--console", "--init-file", "/docker-entrypoint-initdb.d/init.sql"]