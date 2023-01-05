FROM adoptopenjdk/openjdk8:x86_64-centos-jre8u312-b07

ENV TZ=Asia/Shanghai
#设置时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo '$TZ' > /etc/timezone
#设置编码
RUN localedef -c -f UTF-8 -i zh_CN zh_CN.utf8
#设置环境变量
ENV LC_ALL zh_CN.utf8

RUN mkdir -p /home/admin/cloudrun/

ADD ./target/cloudrun-0.0.1.jar /home/admin/cloudrun/

RUN chown -R 777 /home/admin/cloudrun/

WORKDIR /home/admin/cloudrun/

ENTRYPOINT ["java","-jar","cloudrun-0.0.1.jar"]
