# Sleeping Barber

- Client function uses select statement to make decision based on the saloon channel state.  
- Barber function reads clients id's from the saloon channel, work is simulated with sleep.  
- When all new clients are either inside or didn't get into the saloon, channel gets closed.  
- Clients that are inside are serviced, range over channel is done when saloon channel is empty and closed.  
- Barber goroutines are done -> end.  
