| Error | Possible solution |
| ----- | ------ |
| ERROR: for mongo  Cannot start service mongo: driver failed programming external connectivity on endpoint mongo (xxxxxx): Error starting userland proxy: listen tcp 0.0.0.0:27017: bind: address already in use. ERROR: Encountered errors while bringing up the project. | Stop the local instance of mongoDB running on port 27017 (`sudo service mongodb stop`) |
