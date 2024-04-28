FROM golang
#COPY FILES
ADD ./backend /app
WORKDIR /app
#BUILD BIN
#RUN ls
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#RUN go build
#DEBUG
#RUN ls
#RUN pwd
#RUN stat /app FÜR FOLDER
#RUN chmod 700 /app
#RUN stat /app
ENTRYPOINT /app/app
#Build Final Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
#ADD ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
#COPY EXECUTEABLE FILE
COPY --from=0 /app/main .
#COPY --from=0 /app/app.env .
#DEBUG
#RUN pwd
#RUN ls -l app.env
#RUN ls
EXPOSE 3000
ENTRYPOINT [ "./main" ]
