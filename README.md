# Emitter

Creating an ecosystem of two microservices using event-driven architecture

Create an ecosystem of two microservices (an event producer destroyer and consumer deathstar ) that communicate via an open source pub-sub messaging system. The destroyer service should have a gRPC method called acquireTargets which produces an
event targets.acquired (see schema below) on the event stream. The deathstar service consumes these events from the pub-sub system and stores the contained targets in a database. The destroyer service should also have a second gRPC method called listTargets that returns the list of targets acquired from the database of targets directly and returns results as a json array. Given the destroyer service can acquire one or more targets, the deathstar has to extract the list of targets from the event payload and store them in the database. Finally prepare or simulate a few invoca.tions of the destroyer RPC methods that acuires i) an indiviual target, ii) multiple targets and iii) list all targets from the database.

# Required Tech Stack:

- Database: Any (mongodb, mysql, postgres)
- Apache Pulsar as the Eventstore / pub-sub system or alternatively use NATs
- Can run / demo stack locally - using a containerised ( docker / docker compose) workflow Define service interfaces as protocol buffers .proto files
- Code is tested

# Bonus Points [20pts]

Use of Apache Pulsar as event store / pub-sub system
Services connect to the apache pulsar pub-sub system securley (TLS, Certs)
Both destroyer and deathstar services can also implement two additional gRPC methods:
service health check
service readiness response (connects to all the service dependencies and indicates a readiness/status for each dependency of the service ).

# Bonus++ [30pts]

Use of secret managment store like Hashicorp vault to manage TLS secrets and database credentials
Implementation of schema registry to ensure data in event streams conforms to pre- determined schema
The following article could be of use with relation to hasicorp vault: Docker compose - Hashicorp's Vault and Consul - Installing vault, unsealking, static secrets and policies. Again bear in mind this is bonus content and hence why the above link has been provided to assist.

# PROGRESS

- TEST FOR DATABASE FUNCTIONALITY ADDED
- TEST FOR GRPC FUNCTIONALITY ADDED
- DESTROYER SERVICE CREATED
- DEATHSTAR SERVICE CREATED
- POSTGRES DB USED
- APACHE PULSAR USED AS MESSENGER
- DESTROYER TEST ADDED TO TEST APP FUNCTIONALITY AND CONSUMER RESPONSE(RUN BOTH APPS FIRST)

- SIMULATION

  - DESTROYER_TEST.GO
  - CAN BE USED TO SIMULATE ACQUIRETARGET
  - LIST SINGLE TARGET
  - LIST MULTIBLE TARGET

- TODO
  - CONSUL HEALTH CHECK (STARTED BUT STILL IN REVIEW)
  - VAULT (MOVE CONFIG TO VAULT AND RUN VAULT LOCAL)
