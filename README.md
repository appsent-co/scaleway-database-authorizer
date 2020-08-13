# scaleway-database-authorizer

This is a simple micro-service that runs as a DaemonSet and authorizes every new node in a kubernetes cluster in the scaleway elements managed database ACL rules.

## Getting started

You can see an example deamonset in the file "database-authorizer-deamonset.yaml"

You need to configure the following env variables:

    - SCALEWAY_ACCESS_KEY => Your scaleway access key
    - SCALEWAY_SECRET_KEY => Your scaleway secret key
    - SCALEWAY_ORG_ID => Your scaleway org id
    - SCALEWAY_INSTANCE_ID => Your scaleway elements manages database instance id
    - SCALEWAY_DATABASE_REGION => Your scaleway elements manages database region (fr-par/nl-ams)
    - POD_NAME => Example in the yaml file. Name of the pod
    - POD_IP => Example in the yaml file. IP of the pod
