FROM        maven:3.6.3-openjdk-15 as compiler
COPY        ./ciudadania_si /root/ciudadania
WORKDIR     /root/ciudadania/
RUN         mvn clean package -DskipTests

FROM        openjdk:15-jdk-oracle
RUN         mkdir /root/app
COPY        --from=compiler /root/ciudadania/target/ciudadania_si.jar /root/app/ciudadania_si.jar
ENTRYPOINT  ["java", "-jar", "/root/app/ciudadania_si.jar"]
EXPOSE      8080
