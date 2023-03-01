# FROM mysql:5.7

# # 设置环境变量
# ENV MYSQL_DATABASE=gcloud \
#   MYSQL_USER=song \
#   MYSQL_PASSWORD=123456 \
#   MYSQL_ROOT_PASSWORD=123456



# Base image
FROM ubuntu:20.04

# 修改apt-get镜像源为阿里云
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

# Install MySQL
RUN apt-get update && \
  apt-get install -y mysql-server=5.7* && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

RUN service mysql start && \
  mysql -u root -e "CREATE USER 'root'@'%' IDENTIFIED BY '123456';" && \
  mysql -u root -e "GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;" && \
  mysql -u root -e "FLUSH PRIVILEGES;"

# Copy MySQL configuration file
COPY ./init.sql /docker-entrypoint-initdb.d/
COPY mysql.cnf /etc/mysql/mysql.conf.d/mysql.cnf

# Enable remote access for MySQL
RUN sed -i 's/^bind-address.*/bind-address = 0.0.0.0/' /etc/mysql/mysql.conf.d/mysql.cnf
RUN echo 'skip-networking=0' >> /etc/mysql/mysql.conf.d/mysql.cnf

# Expose MySQL port
EXPOSE 3306

# Start MySQL service

CMD service mysql stop && usermod -d /var/lib/mysql/ mysql && service mysql start && bash
