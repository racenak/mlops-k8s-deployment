kubectl create namespace istio-system
helm install istio-base istio/base -n istio-system
helm install istiod istio/istiod -n istio-system -f deploy/istio/istiod-values.yaml

kubectl apply -k deploy/istio/addons