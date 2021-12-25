# podipfinder
This is a small application to find pod ip address by service name using go

How to build this application

    make clean binaries build-image

Testing 

    kubectl apply -f ./test/
    
    deployment.apps/nginx-depl created
    service/nginx-service created
    pod/podipfinder created
    
    kubectl get po
    NAME                          READY   STATUS    RESTARTS   AGE
    nginx-depl-5c8bf76b5b-cf985   1/1     Running   0          8s
    nginx-depl-5c8bf76b5b-ml5gr   1/1     Running   0          8s
    podipfinder                   1/1     Running   0          8s
    
    kubectl exec podipfinder -- /podipfinder nginx-service
    2021/12/25 20:25:39 IP Addresses for service: nginx-service
    2021/12/25 20:25:39 10.244.2.8
    2021/12/25 20:25:39 10.244.1.1
