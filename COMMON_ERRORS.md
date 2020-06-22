| Error | Possible solution |
| ----- | ------ |
| ERROR: for mongo  Cannot start service mongo: driver failed programming external connectivity on endpoint mongo (xxxxxx): Error starting userland proxy: listen tcp 0.0.0.0:27017: bind: address already in use. ERROR: Encountered errors while bringing up the project. | Stop the local instance of mongoDB running on port 27017 (`sudo service mongodb stop`) |
| Everything's working but I can't see the changes I did to the code | Try to restart the service with `make restart` (This is because the old container's still up |
