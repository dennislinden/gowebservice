#kind create cluster

#kubectl create namespace argocd
#kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

#while ! [ $(kubectl get pods -n argocd | grep Running -c) -eq 5 ];
#	do
#	echo wait for pods
#	done
kubectl port-forward svc/argocd-server -n argocd 8081:443
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d && echo
kubectl apply -f appofapps/appofapps.yaml 
