FROM ubuntu:latest
LABEL authors="JiangH"

# 安装mysql
RUN apt-get update \
    && apt-get install -y mysql-server

# 创建mysql用户
RUN usermod -d /var/lib/mysql/ mysql \
    && chown -R mysql:mysql /var/lib/mysql \
    && chown -R mysql:mysql /var/run/mysqld

# 配置环境变量
ENV MYSQL_ROOT_PASSWORD=123456
ENV TZ Asia/Shanghai

EXPOSE 3306

# 将library.sql复制到容器
COPY library.sql /docker-entrypoint-initdb.d/

# 启动mysql服务并执行初始化脚本
CMD ["mysqld", "--init-file", "/docker-entrypoint-initdb.d/library.sql"]
