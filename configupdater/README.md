# k8s
# configupdater
This is a small application to configmap using go kuberntes client

Build Application
    
    make clean binaries build-image
    
kubectl apply -f ./test

    deployment.apps/configupdater created
    serviceaccount/developer created
    role.rbac.authorization.k8s.io/cm-mgr created
    rolebinding.rbac.authorization.k8s.io/cm-mgr-rb created
    configmap/test-cm created

kubectl get po,sa,role,rolebindings,cm 

    NAME                                 READY   STATUS    RESTARTS   AGE
    pod/configupdater-64db7ff476-lmbh8   1/1     Running   0          4m23s
    
    NAME                       SECRETS   AGE
    serviceaccount/default     1         9m7s
    serviceaccount/developer   1         4m23s
    
    NAME                                    CREATED AT
    role.rbac.authorization.k8s.io/cm-mgr   2021-12-29T19:40:58Z
    
    NAME                                              ROLE          AGE
    rolebinding.rbac.authorization.k8s.io/cm-mgr-rb   Role/cm-mgr   4m23s
    
    NAME                         DATA   AGE
    configmap/kube-root-ca.crt   1      9m7s
    configmap/test-cm            3      4m23s

kubectl describe cm test-cm
    Name:         test-cm
    Namespace:    default
    Labels:       <none>
    Annotations:  <none>
    
    Data
    ====
    dbservice:
    ----
    enterprise
    dbhost:
    ----
    localhost
    dbport:
    ----
    3306
    
    BinaryData
    ====
    
    Events:  <none>
kubectl get po

    NAME                             READY   STATUS    RESTARTS   AGE
    configupdater-67685bc774-ktzjp   1/1     Running   0          96s

kubectl exec configupdater-67685bc774-ktzjp -- /app/bin/configupdater test-cm -a dbuser=user,dbpwd=pwd -u dbhost=enterprise_db -d dbservice

    2021/12/29 20:03:57 Updating configmap:  test-cm
    2021/12/29 20:03:57 Adding key: dbuser, value: user
    2021/12/29 20:03:57 Adding key: dbpwd, value: pwd
    2021/12/29 20:03:57 Key: dbhost already exists, updating
    2021/12/29 20:03:57 Deleting key: dbservice

kubectl describe cm test-cm 
    Name:         test-cm
    Namespace:    default
    Labels:       <none>
    Annotations:  <none>
    
    Data
    ====
    dbuser:
    ----
    user
    dbhost:
    ----
    enterprise_db
    dbport:
    ----
    3306
    dbpwd:
    ----
    pwd
    
    BinaryData
    ====
    
    Events:  <none>

