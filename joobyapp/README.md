# joobyapp

Welcome to joobyapp!!

## running

    mvn clean jooby:run

## building

    mvn clean package

    java -jar target/joobyapp-1.0.0.jar
    java -server -Xms2g -Xmx2g -XX:+UseNUMA -XX:+UseParallelGC -Dio.netty.disableHttpHeadersValidation=true -Dio.netty.buffer.checkBounds=false -Dio.netty.buffer.checkAccessible=false -jar target/joobyapp-1.0.0.jar