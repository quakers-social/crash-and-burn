ARGOCD
======




PORT FORWART
------------


```bash
kubectl port-forward service/argocd-server -n argocd 8080:80
```

after there open http://localhost:8080


GET PASSWORD
------------

Get password (default user admin):

```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```