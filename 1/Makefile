
REPO?=runyonsolutions
IMAGE?=litecoind
TAG?=0.17.1

docker:
	docker build -t ${REPO}/${IMAGE}:${TAG} .

push:
	docker push ${REPO}/${IMAGE}:${TAG}

scan: 
	trivy -severity HIGH,CRITICAL ${REPO}/${IMAGE}:${TAG}

# To install Trivy on a Mac:
# brew install aquasecurity/trivy/trivy
# Other installs on https://github.com/aquasecurity/trivy#installation