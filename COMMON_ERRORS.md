| Error | Possible solution |
| ----- | ------ |
| ERROR: for mongo  Cannot start service mongo: driver failed programming external connectivity on endpoint mongo (xxxxxx): Error starting userland proxy: listen tcp 0.0.0.0:27017: bind: address already in use. ERROR: Encountered errors while bringing up the project. | Stop the local instance of mongoDB running on port 27017 (`sudo service mongodb stop`) |
| 2020/05/29 07:58:03 open *.txt: no such file or directory gochanges_1 exit status 1 | file *.txt doesn't exist. Try with an existing file |
| If everything start but you are monitoring the websites you don't want, try to delete the old image and build new one | `docker rmi gochanges_gochanges:latest`,successivamente modificare il Dockerfile ed il file di configurazione come scritto nel README.md e far ripartire tutto con `docker-compose up` |
