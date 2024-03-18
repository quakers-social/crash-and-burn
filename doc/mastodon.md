MASTODON
========


MISSION
-------

Installation of Mastodon on Kubernetes.

Finde the last version and check the file [`Chart.yaml`](https://github.com/bitnami/charts/blob/main/bitnami/mastodon/Chart.yaml)
in the Chart Repo.

You finde the information about the Helm Chart parmaeters in the file [`README.md`](https://github.com/bitnami/charts/blob/main/bitnami/mastodon/README.md)


CREATE VAPID KEYs
-----------------

```bash
openssl ecparam -name prime256v1 -genkey -noout -out vapid_private_key.pem
openssl ec -in vapid_private_key.pem -pubout -out vapid_public_key.pem
echo -n VAPID_PRIVATE_KEY=;cat vapid_private_key.pem | sed -e "1 d" -e "$ d" | tr -d "\n"; echo
echo -n VAPID_PUBLIC_KEY=;cat vapid_public_key.pem | sed -e "1 d" -e "$ d" | tr -d "\n"; echo

```


EXTERNAL LINKS
--------------

- [Bitnami package for Mastodon](https://github.com/bitnami/charts/tree/main/bitnami/mastodon/)
- [Mastodon on DigitalOcean Kubernetes](https://docs.digitalocean.com/developer-center/mastodon-on-digitalocean-kubernetes/)
- [Running Mastodon server on Kubernetes](https://softwaremill.com/running-mastodon-server-on-kubernetes/)