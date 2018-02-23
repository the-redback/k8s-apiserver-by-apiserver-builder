# k8s-apiserver-by-apiserver-builder

To build vendor

```console
$ dep ensure -v
```

To run project

```console
$ apiserver-boot run local-minikube
```

To test create CRD

```console
$ kc create -f sample/mykind.yaml
mykind "mykind-example" created
```

For more details, see [Apiserver Builder](https://github.com/kubernetes-incubator/apiserver-builder).
