helmBuild :
	helm install api-server ../Helm/api-server/
	sleep 10
run : helmBuild
	kubectl port-forward services/api-server-service 8088:8081
