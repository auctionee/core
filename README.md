# core
auctionee core service

core service communicates with other services in order to process auction

list of services:
- goods service
    - goods service contains info about all goods in organisation
- auth service
    - auth service contains users login/passwords and checks for permissions
- balance service
    - balance service return users balance and can modify user balance
- front service
    - service that server templates and styles, allows user to interact with core service
- core service
    - is a communication point between services. Do all auction-related work

DB-Service list:
- goods   <-->  key-value user/goods DB
- auth    <-->  key-value login/password DB
- balance <-->  key-value login/balance DB
- front   <-->  key-value resources DB
- core    <-->  key-value auctions/info DB

each service communicates with others using JSON-RPC

deployment - GCloud
