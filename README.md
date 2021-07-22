# Gitlab-POC-Template
Starter of minimal microservice api based application for gitlab team using 
- Gin-gonic/gin as the REST framework
- Logrus for logging
- Viper and  for configs
- Fsnotify for live config updates
- Testcases using testify (https://github.com/stretchr/testify)

### Directory discription
-- table here-- 

### Running go code on local machine 
- Install go (go1.16)
- Navigate to project root directory (which contains the Makefile)
- Start the test database container `docker run -d -e MYSQL_ROOT_PASSWORD=test --name mysql8 -p 127.0.0.1:3307:3306 mysql:8`
- Create a test database schema and updated the value of config.yml -> database.name (laer this can be achived with migrations)
- Cross check configuration values in `./config.yml`
- Run `make go-test`
- Run `make go-run`

### Running go code in docker container
- Make sure docker daemon is up and running
- If same machine is used for minikube setup, make sure env variable listed in output of `echo $(minikube docker-env)` are not already set
- Navigate to project root directory (which contains the Makefile)
- 
- Build the go code to generate the binary `make go-build`
- Generate the docker image using `make docker-build`
- Run the container `make docker-run`
- Access the web portal on `http://localhost:8080`

### Setup on Minikube 
- Start the minikube `minikube start`
- Set docker env `eval $(minikube docker-env)`
- Build the docker image `make docker-build`
- -- run sql container and service --
- Cross check configuration values in `./config.yml`
- Create the pod `kubectl run ginapp --image=ginapp:1 --image-pull-policy=Never`
- Create service pointing to minikube pod `kubectl expose pod ginapp --port 8080`
- Expose the service on local machine `kubectl port-forward ginapp  8080`
- Access the web portal on `http://localhost:8080`
- Check the server logs using `Kubectl logs ginapp -f`

### References
- Docker setup, fsnotify, viper, logrus reference: 
    - https://github.com/udaya2899/go-gin-starter
- Detailed information around the file structure and naming conventions: 
    - https://www.youtube.com/watch?v=LOn1GUsjOF4
    - https://github.com/vikramaroskar/Go_rest_gin
    - https://github.com/shohiebsense/Go-Gin-REST-API

### To do next
- Database migrations
- Add orm for database
- 