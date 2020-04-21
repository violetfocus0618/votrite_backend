FROM golang

ENV VR_DB_HOST=$VR_DB_HOST
ENV VR_DB_PORT=$VR_DB_PORT
ENV VR_DB_USR=$VR_DB_USR
ENV VR_DB_PWD=$VR_DB_PWD
ENV VR_DB_NAME=$VR_DB_NAME
ENV VR_API_HOST=$VR_API_HOST
ENV VR_API_PORT=$VR_API_PORT
ENV AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
ENV AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY
ENV AWS_REGION=$AWS_REGION
ENV AWS_SES_SENDER=$AWS_SES_SENDER
ENV AWS_SNS_ARN=$AWS_SNS_ARN
ENV GCP_API_KEY=$GCP_API_KEY

WORKDIR /usr/app

COPY . .

RUN go get github.com/lib/pq
RUN go get github.com/google/uuid
RUN go get rsc.io/qr
RUN go get github.com/aws/aws-sdk-go/...
RUN go get cloud.google.com/go/translate

RUN mkdir bin
RUN go build -o bin/vrapid

EXPOSE 9090

CMD ["bin/vrapid"]