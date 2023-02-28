# Base image
FROM ubuntu:20.04

# Install MySQL and Redis
RUN apt-get update && \
  apt-get install -y mysql-server redis-server && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

# Copy MySQL and Redis configuration files
COPY mysql.cnf /etc/mysql/mysql.conf.d/mysql.cnf
COPY redis.conf /etc/redis/redis.conf

# Enable remote access for MySQL
RUN sed -i 's/^bind-address.*/bind-address = 0.0.0.0/' /etc/mysql/mysql.conf.d/mysql.cnf
RUN echo 'skip-networking=0' >> /etc/mysql/mysql.conf.d/mysql.cnf

# Enable remote access for Redis
RUN sed -i 's/^bind 127.0.0.1/bind 0.0.0.0/' /etc/redis/redis.conf

# Expose MySQL and Redis ports
EXPOSE 3306 6379

# Start MySQL and Redis services
CMD service mysql start && service redis-server start && bash
