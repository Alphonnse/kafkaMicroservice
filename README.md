#High available micro service with kafka
In this project i have 3 main files. One of them is the producer and two of them is the consumers.  
Meaning of this code is that if any of consumers are down that the second can keep consume messages. If the one that was down gets back in line it can again consume messages with no downtime.
