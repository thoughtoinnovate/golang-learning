#Two Types of Channels
** UnBuffered Channel
** Buffered Channel

#UnBuffered Channel - In this channel sender can put as many as messages in the buffer , condition: if there is no receiver for the message , sender gets blocked and cant add any more messages.
Syntax :  myChannel:=make(chan int)

#Buffered Channnel - this is fixed buffer size channel , if capacity is e.g. 5 then sender go  routine can put max five messages in the buffer without getting blocked even if receiver doesnt pick it up , but if buffer os full say has max 5 messages in this case , then sender go routine gets blocked here , until any receiver go routine picks message from buffer.
Syntax: MyRoutine:=make(chant int,5)

#Note:
In the both the buffers if receiver go routine tries to get message from empty buffer , it gets blocked until a new message is added to the buffer.