FROM openjdk:8-jdk-alpine
ARG JAR_FILE=target/attack-provider-server.jar
ARG FILES_DIR=files
COPY ${JAR_FILE} app.jar
COPY ${FILES_DIR} ${FILES_DIR}
ENTRYPOINT ["java","-jar","/app.jar"]
