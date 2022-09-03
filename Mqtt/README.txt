docker-compose up
kafka-topics
kafka-topics --bootstrap-server=localhost:9092 --list 
kafka-topics --bootstrap-server=localhost:9092 --topic=my-queue-1 --create

display message 
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=my-queue-1

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=my-queue-1 --group=myconsumer

send message
kafka-console-producer --bootstrap-server=localhost:9092 --topic=my-queue-1