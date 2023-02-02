# golang-amqp-consumer

A Golang consumer for RabbitMQ queues running in Docker.

Forwards the message data to a handler based on the messageEvent header.

![image](https://user-images.githubusercontent.com/9871509/216378330-4ad66f93-678a-4c3e-b72e-0d4df60956c1.png)
> Example: Working as a message proxy


#### Run locally

`$ docker compose up`


#### Add new event treatments
In `cmd/main.go` add the line below according with the handle function created in the `handlers` directory:
```golang
  for message := range messages {
    switch message.Headers["messageEvent"] {

    // YOUR CODE HERE
    case "yourEvent":
      yourHandler.Handle(message.Headers, string(message.Body))
      continue
    // YOUR CODE HERE
      
    default:
      fmt.Println("Message received with unknown event")
    }
  }

```
