POSTGRESQL
==========

- [POSTGRESQL](#postgresql)
  - [MISSION](#mission)
  - [INSTALLATION OF OPERATOR](#installation-of-operator)
  - [BOOTSTRAP OF POSTGRESQL INSTANCE / CLUSTER](#bootstrap-of-postgresql-instance--cluster)
  - [POSTGRESQL CONNECTION](#postgresql-connection)
  - [BACKUP \& RESTORE](#backup--restore)
  - [COMPARING SOLUTIONS AND ALTERNATIVES](#comparing-solutions-and-alternatives)


MISSION
-------

We use a separate PostgreSQL installation for more control. And to do this
we use the operator from [Crunchy Data](https://www.crunchydata.com/).

INSTALLATION OF OPERATOR
------------------------

When installing the operator, the official instructions were used as a
guide. The installation and configuration has been automated with Pulumi.

See:
 - [Docu](https://access.crunchydata.com/documentation/postgres-operator/latest/installation/helm)
 - [git repo with Install example for Helm](https://github.com/CrunchyData/postgres-operator-examples/tree/main/helm/install)


BOOTSTRAP OF POSTGRESQL INSTANCE / CLUSTER
------------------------------------------

The official documentation also served as a template when setting up the
PostgreSQL cluster itself. The installation and configuration was
automated with Pulumi.

See: [Create a Postgres Cluster](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/basic-setup/create-cluster)


POSTGRESQL CONNECTION
---------------------

See: [Connect to a Postgres Cluster](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/basic-setup/connect-cluster)


BACKUP & RESTORE
----------------

An S3 serves as backup storage. The PostgreSQL operator is responsible
for the backup. This configuration is also managed with Pulumi and
stored in Git and S3.

To create adhoc backups in Mastodon upgrade, the following command is
used:

```bash
export PG_CLUSTER_NAME="mastodon-db"
kubectl annotate \
  postgrescluster ${PG_CLUSTER_NAME} \
  --overwrite \
  postgres-operator.crunchydata.com/pgbackrest-backup="$(date)" \
  -n pulumi-apps
```

See:
- [Backup Configuration](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/backups-disaster-recovery/backups)
- [Using S3 Credentials](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/backups-disaster-recovery/backups#using-s3-credentials)
- [Taking a One-Off Backup](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/backups-disaster-recovery/backup-management#taking-a-one-off-backup)


COMPARING SOLUTIONS AND ALTERNATIVES
------------------------------------

When deciding on the PostgreSQL operator, the following options were evaluated

Node: (Watcher/Forks/Stars/Issues)

- [Postgres Operator from Crunchy Data (67/566/3,7k/112)](https://github.com/CrunchyData/postgres-operator)
  - [Install](https://access.crunchydata.com/documentation/postgres-operator/latest/installation)
  - [Backup](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/basic-setup)
- [CloudNativePG (23/223/3k/218)](https://github.com/cloudnative-pg/cloudnative-pg/tree/main)
- [Zalando Postgres Operator (66/935/3,9k/532)](https://github.com/zalando/postgres-operator)

**Ultimately, the solution that was most popular and had the best documentation was chosen.**