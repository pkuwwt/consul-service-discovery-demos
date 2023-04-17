FROM maven

WORKDIR /dist/

ADD pom.xml /dist/
ADD src /dist/src

RUN mvn package


FROM openjdk:11
WORKDIR /dist/
COPY --from=0 /dist/target/spring-boot-0.0.1-SNAPSHOT.jar /dist/spring-boot-0.0.1.jar

EXPOSE 8080
ENTRYPOINT ["java", "-jar", "spring-boot-0.0.1.jar"]
