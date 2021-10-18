### About the task ###

fxporxy code challenge is designed to demonstrate [sidecar pattern](https://docs.microsoft.com/en-us/azure/architecture/patterns/sidecar) implementation. fxproxy is sidecar service which is responsible to govern incoming traffic to the downstream application services based on allowed path list `allowedList` in `main.go`.


### Task objective ###

* implement `ValidatePath` and pass failing tests in `main_test.go`, however we don't want to limit you so feel free to change the function name or even project structure/packages in way you believe can suit a production grade Go project. Code and tests are added to give you an idea about the business requirement of the task.

* It won't be a sidecar service if it doesn't handle incoming Http request and send it to downstream application service, so please extend project functionality to handle http proxy responsibilities well, you can use any 3rd party packages or stick with standard libraries.

* Provide docker-compose to run the sidecar service along side with a dummy application service 

#### Good to have ####

* Implementing the task in TDD way
* Effective use of source control
* Follow Idiomatic Go
* Design a loosely coupled and highly maintainable code
* Since this service is expected to handle all traffic to pod/application service, code is expected to be fast and resilient


##### To Run the test

```

cd sidecar/pkg/middleware/ && go test

```

### Solution

* I create two project inside repository, sidecar and dummy app project. Http request are come to sidecar project and after validations sidecar redirect the request to the dummy app project. 

##### Steps to start

-  Clone the repo
  ```
  git clone https://github.com/salihkemaloglu/fxproxy.git
   ```

##### Steps to run

1. Go to main directory:
   ```
   cd fxproxy/
   ```
2. Grant permission: 
   ```
   sudo chmod +x ./app/shell/deploy_app.sh
   ```   
   ```
   sudo chmod +x ./sidecar/shell/deploy_sidecar.sh
   ```  
3. Run :
   ```
   cd app && ./shell/deploy_app.sh
   ```
   ```
   cd sidecar && ./shell/deploy_sidecar.sh
   ```

 ##### Http request
  - Validate success request sample:
   ```
     curl -X GET \
    http://localhost:8081/hello
   ```

   - Success Response:

  ```
   [
      "sidecar is working ...",
      "hello hello"
   ]

   ```
  - Validate fault request sample:
   ```
     curl -X GET \
    http://localhost:8081/hello1
   ```
  
   - Fault Response:
 
    406 not acceptable
    
## To improve
![architecture](https://freepngimg.com/thumb/street_fighter/35134-8-street-fighter-ii-image.png)


