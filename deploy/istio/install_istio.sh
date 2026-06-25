kubectl create namespace istio-system
helm install istio-base istio/base -n istio-system
helm install istiod istio/istiod -n istio-system -f deploy/istio/istiod-values.yaml

kubectl create namespace istio-ingress
helm install istio-ingress istio/gateway -n istio-ingress -f deploy/istio/istiod-values.yaml

