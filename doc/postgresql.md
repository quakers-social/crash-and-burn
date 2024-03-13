POSTGRESQL
==========


MISSION
-------

We use a separate PostgreSQL installation for more control. And to do this we use the operator from [Crunchy Data](https://www.crunchydata.com/).

INSTALLATION OF OPERATOR
------------------------

Over HELM
See:
 - [Docu](https://access.crunchydata.com/documentation/postgres-operator/latest/installation/helm)
 - [git repo with Install example for Helm](https://github.com/CrunchyData/postgres-operator-examples/tree/main/helm/install)


BOOTSTRAP OF POSTGRESQL INSTANCE / CLUSTER
------------------------------------------

See: [Create a Postgres Cluster](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/basic-setup/create-cluster)


POSTGRESQL CONNECTION
---------------------

See: [Connect to a Postgres Cluster](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/basic-setup/connect-cluster)


BACKUP & RESTORE
----------------

See:
- [Backup Configuration](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/backups-disaster-recovery/backups)
- [Using S3 Credentials](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/backups-disaster-recovery/backups#using-s3-credentials)


COMPARING SOLUTIONS AND ALTERNATIVES
------------------------------------

Comparing: (Watcher/Forks/Stars/Issues)

- [Postgres Operator from Crunchy Data (67/566/3,7k/112)](https://github.com/CrunchyData/postgres-operator)
  - [Install](https://access.crunchydata.com/documentation/postgres-operator/latest/installation)
  - [Backup](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorials/basic-setup)
- [CloudNativePG (23/223/3k/218)](https://github.com/cloudnative-pg/cloudnative-pg/tree/main)
- [Zalando Postgres Operator (66/935/3,9k/532)](https://github.com/zalando/postgres-operator)