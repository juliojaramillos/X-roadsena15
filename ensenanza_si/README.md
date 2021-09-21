# X-Road Adapter Example

Esta implementación está completamente basada en el adaptador ejemplo provisto por la librería [XRd4J](https://github.com/nordic-institute/xrd4j)

### Ejecución

En orden de construír y ejecutar ese proyecto, use los siguientes comandos

```
mvn clean install
docker build -t example-adapter .
docker run -p 8080:8080 example-adapter 
```

### Requisitos de software

* Java 8
* Tomcat >= 6
