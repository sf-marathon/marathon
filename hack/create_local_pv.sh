#!/bin/sh
kubectl create -f - <<EOF
apiVersion: v1
kind: PersistentVolume
metadata:
  name: hostpv1
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
    - ReadWriteMany
    - ReadOnlyMany
  hostPath:
    path: /var/k8s-volume
  storageClassName: default
EOF

kubectl create -f - <<EOF
apiVersion: v1
kind: PersistentVolume
metadata:
  name: hostpv2
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
    - ReadWriteMany
    - ReadOnlyMany
  hostPath:
    path: /var/k8s-volume
  storageClassName: default
EOF
