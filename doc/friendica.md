FRIENDICA
=========



BASICS
------

* **FQDN:** friendica.the-independent-friend.de
* **IP:** 2a01:4f9:c010:8e06::1

DATABSE / MARIADB
-----------------

* **URL:** mariadb.argocd.svc


PORT FORWART
------------

```bash
$ kubectl port-forward service/friendica-app -n argocd-aoa 8081:8080
```


LINKS
-----

* [Website](https://friendi.ca/)
* [Installation](https://friendi.ca/resources/installation/)
* [On Github](https://github.com/friendica/friendica)
* [Public server](https://dir.friendica.social/servers)
* [On openhub.net](https://openhub.net/p/friendica)
* [Docker container](https://hub.docker.com/_/friendica)
* [Docker image code](https://github.com/friendica/docker)
* [Docker compose](https://github.com/friendica/docker/blob/stable/.examples/docker-compose/opensocial.at/docker-compose.yml)
* [How to Import People You Follow on Mastodon to Friendica?](https://sanguok.com/en/blog/how-to-import-the-people-you-follow-on-mastodon-to-friendica/)



