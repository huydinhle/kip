# The flows

# Implementing
- VirtualService struct inside of IstioCanaryDeployment --> make sure it build
- Operator can create IstioCanaryDeployment inside of the cluster --> install istio inside of minikube, make sure operator can create virtual service
- Include Prometheus-Client inside of the operator-code to query stuff from it
- Install prometheus operator with only prometheus-operator and prometheus
- Do some sample Query

### Isitio-client is currently not completed because of the schema lived inside the repo don't have all json tags in it, we will have to switch it off and use the name of virtualservice instead

### Reconcile method
1. **Receive** Reconcile Request which has IstioCanaryDeployment instance Whenever there is an operations add/update/delete happens to our istiocanarydeployment
2. **Validate** the IstioCanaryDeployment instance
4. If Deployment **DOES NOT** exist
  - **Generate** dep,svc,istioVS from CR instance
  - **DELETE** service, and istio stuff if they existed ---> not done
  - **Set references** for dep, svc, istioSVC to the CR instance( we are doing this so that whenever we delete CR, these got destroyed )
  - **Create**
    - Deployment
    - Service
    - Istio Virtual Service 
  - **Done** and return 
4. If Deployment **DOES** exist
  - **Canary Process** the instance

### Canary Process
1. **Delete** Canary dev,svc if there is
2. **Generate** dep,svc,istioVS from CR instance
3. **Set references** for dep, svc, istioSVC to the CR instance( we are doing this so that whenever we delete CR, these got destroyed )
4. **Create**
  - Deployment
  - Service
  - Istio Virtual Service 
5. **Analyze** Canary

#### Generate Stuffs From our CR
1. **DELETE** canary if it exist
2. **Generate** dep,svc,istioVS from CR instance
  - if this is a canary one, give the CR's name instance in there, and defer the action of reseting the name of the instance
3. **Set references** for dep, svc, istioSVC to the CR instance( we are doing this so that whenever we delete CR, these got destroyed )
4. **Create**
  - Deployment
  - Service
  - if this is not a canary
    - Istio Virtual Service 


#### Delete Canary
1. **LIST/GET** dep,svc for canary
2. **NOT** exist, return 
3. **DOES** exist, **DELETE** it

### Validate a CanaryInstance
1. How would you do that 

### Analyze Canary
1. To be continued

### Notes to come back to look at
- Istio Virtual Service object isn't in the CRD yet
  - **Complicated**
  - https://github.com/istio/istio/issues/8772
    - https://github.com/istio/tools/pull/37
    - https://github.com/istio/api/pull/764/files
- We should have our Reconciler doing action on CREATE/UPDATE only
  - We don't do shit on Delete because delete the instance will in turns trigger delete all the resources that are created by that instance already
  - Sending DELETE into the reconciler would cause confusion to our operator because it will
    - Try to create regular deployment and canary deployment
- When a Canary Action should be performed? ---> Translate this into code because we only want CanaryProcess to run whenever this happen
  - Only when deploymentSpecs changes ---> which in turns telling me only when the podSpec actually changes, we trigger our Canary
    - Except for replicas, if replicas is the only changes , then you don't really 
  - We can have a flag to trigger CanaryProcess --> this has limitations, and shitty because people have to remember to set this stupid flag on and off all the time. 
    - TRUE: do a canaryProcess
    - FALSE: do not do a canaryProcess
  - **Our best bet right now is to do a comparison between the new PodSpecTemplate and the old one using reflect"**
- Have a flag to override the CanaryProcess forcing the deployment to go to the next version immediately 
  - This flag needs to be reset to false after the deployment got updated



### Questions that we need to answer later on
1. What happens when there are changes that are done to our deployment, service, our istio virtual service stuff?
  - This is the watch functionality 
2. How to use `Status IstioCanaryDeploymentStatus` to make our operator become more efficient
