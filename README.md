# **PLEO SRE CHALLENGE**

This is my submission for the Pleo SRE Challenge. I have created a payment microservice written in Go and wired it up with the Antaeus API. Both Services are running in kubernetes.


# **How To Test**
To test this solution, simply run the script file by executing ./k8s-script.sh in your terminal. This script will deploy the microservices into the cluster. The containers within the k8s cluster are created are based on the image created from the "Dockerfile" & "Dockerfile-paymentsvc" and are pushed to the dockerhub.
On calling the /rest/v1/invoices/pay endpoint, the payment service is invoked and pending invoices are paid.

## Discussion Points
In a production environment, Logging/monitoring  is important. Tools like ELK & prometheus could help log & monitor the cluster. Also, Helm charts could be used to manage deployments.
RBAC(Role based access control) can be used to ensure a developer is given permission to a particular service.