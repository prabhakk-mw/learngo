# LearnGo

This repository lists the various evolutions of learning go and grpc.


## Thoughts

* What is the point of gRPC communication, when you can simply do a function call to make the communication possible?


Preference of communication patterns:
1. Simple function call
2. Async (threads) communication
3. Inter process communication
     When to use IPC? 
     - IPC should be used when it is beneficial to break your project up into separable microservices, or executables.
     - If everything is within a single executable, then why do you even care about starting gRPC comms, you could simply do a function call!
