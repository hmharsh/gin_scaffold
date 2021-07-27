# Microservice API Docker Image
Starter of minimal microservice api based application
- Gin-gonic/gin as the REST framework
- Logrus for logging
- Viper and  for configs
- Fsnotify for live config updates
- Testcases using testify (https://github.com/stretchr/testify)

### Directory discription 
| Path                   |  Purpose                                                                      |
| ---------------------- |  ---------------------------------------------------------------------------- |
| ./bin/                 |  After code build it will contains the executable binary                      |
| ./configuration/       |  Go package to load the config and keep it updated on live changes            |
| ./connection           |  Code for database connectivity                                               | 
| ./httpd/               |  Contains the entrypoint (main.go) and handler logic                          | 
| ./httpd/handler/ (*)   |  Contains REST endpoint(s) information with corresponding logic reference     | 
| `./platform/<endpoint>`|  Contains logic and unit tests corresponds to different API endpoints         | 
| ./storage              |  Contains instance of storage/database, that can be commonly imported         | 
 
* File name in this directory will follow a standard which is `<http_endpoint_name>_<http_endpoint_method>.go` for example `ping_get.go`

### Dependencies 
- Go (1.16.6 darwin/amd64)
- Docker (20.10.6)
- Minikube (v1.22.0)

 
### Running go code on local machine 
- Git Clone and Navigate to project root directory (which contains the Makefile)
- Start the test database container `docker run -d -e MYSQL_ROOT_PASSWORD=test -e MYSQL_DATABASE=ginapp --name mysql8 -p 127.0.0.1:3306:3306 mysql:8`
- Cross check configuration values in `./config.yml`, make sure database.host point to `localhost`
- Test the code by `make go-test`
- Run the application locally `make go-run`

### Generate the docker image 


### Running go code in docker container
- Make sure docker daemon is up and running
- If same machine is used for minikube setup, make sure env variable listed in output of `echo $(minikube docker-env)` are not already set
- Git Clone and Navigate to project root directory (which contains the Makefile)
- In config.yaml make sure database.host is pointing to name of db service from docker-compose file ie `mysql8`
- Test the code by `make go-test`
- Build the go code to generate the binary `make go-build`
- Run `docker compose up`
- Access `http://localhost:8080`

### Setup on Minikube 
- Git Clone and Navigate to project root directory (which contains the Makefile)
- Start the minikube `minikube start`
- Set docker env `eval $(minikube docker-env)` to make sure local images can be directly used by minikube pods
- Test the code by `make go-test`
- Run test sql container and service `kubectl run mysql8  --image=mysql:8 --expose --port 3306 --env="MYSQL_ROOT_PASSWORD=test" --env="MYSQL_DATABASE=ginapp" --labels="app=ginapp"`
- Update the ./config.yml file under database, set host to `mysql8` which is the name of mysql service
- Build the go code to generate the binary `make go-build`
- Build the docker image `make docker-build`
- Create the application pod `kubectl run ginapp --image=ginapp:latest --image-pull-policy=Never --labels="app=ginapp"`
- Port-forward for local access `kubectl port-forward pod/ginapp 8080:8080`
- Access the web portal on `http://localhost:8080`
- Check the server logs using `Kubectl logs ginapp -f`
- For cleanup run `kubectl delete pods,services -l="app=ginapp"`

Note: listing down the saparate steps below, later based on development/deployment usecase, we will able to combine instructions/command into single Make command supported for multiple environment (like local, prod etc)


### References
- Detailed information around the file structure and naming conventions: 
    - https://www.youtube.com/watch?v=LOn1GUsjOF4
    - https://github.com/shohiebsense/Go-Gin-REST-API
- Minimal Usage of Fsnotify, Viper, Logrus
    - https://github.com/udaya2899/go-gin-starter    
