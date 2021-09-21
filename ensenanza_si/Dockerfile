FROM maven:3.6.0-jdk-8-slim AS build
RUN mkdir /ensenanza_si
WORKDIR /ensenanza_si
COPY ./ensenanza_si/src /ensenanza_si/src
COPY ./ensenanza_si/pom.xml /ensenanza_si
RUN mvn -f /ensenanza_si/pom.xml clean install

FROM java:8
COPY --from=build /ensenanza_si/target/ensenanza-si-0.0.1.jar /usr/local/lib/deploy.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/usr/local/lib/deploy.jar"]