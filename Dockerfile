FROM golang
MAINTAINER mathoperations
RUN mkdir /app
ADD mathoperations /app/
ADD pkcs8.key /app/
ADD cacert.pem /app/
ADD nuvem-intera.local.pem /app/
WORKDIR /app
EXPOSE 10000
CMD ["/app/mathoperations"]
