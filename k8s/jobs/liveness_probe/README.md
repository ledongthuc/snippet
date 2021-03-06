# Liveness probe

We use a testing service `ledongthuc/liveness-probe-api` to simulate the case liveness route's down or return error status.
 - 1st and 2th requests to `/liveness_probe_status`, they return status code 200
 - 3rd and 4nd request, `/liveness_probe_status` returns status code 500, it makes the pod's restarted
 - Request counter will be reset after restarting pod
 - After 5 times restarting, we will got error CrashLoopBackOff

Docker Image: https://hub.docker.com/r/ledongthuc/liveness-probe-api

1. Apply deployment and service

```
kubectl apply  -f./
```

2. Watch pod status

```
$ kubectl get pod --watch

http-app-xxx   0/1     ContainerCreating   0          5s
http-app-xxx   1/1     Running             0          5s
http-app-xxx   1/1     Running             1          30s
http-app-xxx   1/1     Running             2          54s
http-app-xxx   1/1     Running             3          80s
http-app-xxx   1/1     Running             4          104s
http-app-xxx   0/1     CrashLoopBackOff    4          2m7s
```

3. To check detail, use describe

```
$ kubectl describe pod [pod-name]
... 
Events:
  Type     Reason     Age                     From               Message
  ----     ------     ----                    ----               -------
  Normal   Scheduled  <unknown>               default-scheduler  Successfully assigned xxx to minikube
  Normal   Pulling    15m (x3 over 19m)       kubelet, minikube  Pulling image "ledongthuc/liveness-probe-api:latest"
  Normal   Killing    15m (x2 over 17m)       kubelet, minikube  Container http-server failed liveness probe, will be restarted
  Normal   Pulled     15m (x3 over 19m)       kubelet, minikube  Successfully pulled image "ledongthuc/liveness-probe-api:latest"
  Normal   Created    15m (x3 over 19m)       kubelet, minikube  Created container http-server
  Normal   Started    15m (x3 over 19m)       kubelet, minikube  Started container http-server
  Warning  Unhealthy  8m57s (x26 over 18m)    kubelet, minikube  Liveness probe failed: HTTP probe failed with statuscode: 500
  Warning  BackOff    4m10s (x13 over 6m27s)  kubelet, minikube  Back-off restarting failed container
```

4. Check application logs:

```
$ kubectl logs http-app-xxx

Env NUMBER_OF_LIVENESS_SUCCESS:  2
Env NUMBER_OF_SKIP_STARTUP:
Start with port 8080
/liveness_probe_status, Counter:  0
/liveness_probe_status, Counter:  1
/liveness_probe_status, Counter:  2
/liveness_probe_status, Counter:  3
/liveness_probe_status, Counter:  4
/liveness_probe_status, Counter:  5
/liveness_probe_status, Error result
/liveness_probe_status, Counter:  5
/liveness_probe_status, Error result
