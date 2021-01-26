FROM maven:3.6.3-openjdk-14-slim AS build

COPY settings.xml /usr/share/maven/conf/

COPY pom.xml pom.xml
COPY ncs-api/pom.xml ncs-api/pom.xml
COPY ncs-model/pom.xml ncs-model/pom.xml
COPY ncs-base/pom.xml ncs-base/pom.xml
COPY ncs-script/pom.xml ncs-script/pom.xml

RUN mvn dependency:go-offline package -B

COPY ncs-api/src ncs-api/src
COPY ncs-model/src ncs-model/src
COPY ncs-base/src ncs-base/src
COPY ncs-script/src ncs-script/src

RUN mvn install -Prunnable

FROM groovy:3.0.5-jdk14
USER root

WORKDIR /

RUN mkdir service

COPY --from=build /ncs-base/target/ /service/
COPY /ncs-script/src/main/groovy/com/atlas/ncs/script /service/script/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait

RUN chmod +x /wait

ENV JAVA_TOOL_OPTIONS -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:5005

EXPOSE 5005

CMD /wait && java --enable-preview -jar /service/ncs-base-1.0-SNAPSHOT.jar -Xdebug