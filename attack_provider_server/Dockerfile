FROM openjdk:8-jdk-alpine
ARG JAR_FILE=target/attack-provider-server.jar
COPY ${JAR_FILE} app.jar
ENTRYPOINT ["java","-jar","/app.jar"]
