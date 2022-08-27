# Sleeping Barber
Barber saloon has n chairs in the waiting room and one barber chair in the barber room. Barber sleeps in the barber chair if there are no customers. If a customer enters the barber saloon and there is no one but the sleeping barber, customer wakes up the barber and sits in the barber chair. If a customer enters and barber is busy with another customer, he takes a seat in the waiting room. If a customer enters and there is no place in the waiting room, he leaves.

## Solution
One channel is used for storing clients, and ’blocking’ them inside until barber is ready to service them. Clients are serviced in the order of coming to the saloon without additional mechanisms. There is no necessity for semaphore and mutex. If the saloon channel is full, clients will not enter. State of the saloon is known to the clients without additional variable, that would need its own critical section.

## Wikipedia
https://en.wikipedia.org/wiki/Sleeping_barber_problem