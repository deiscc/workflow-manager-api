# API Persistent Data Architecture

The Workflow Manager Service API stores persistent data via a postgres database.

## Prerequisites

* Existing DB using the `PostgreSQL 9.4.5` engine

## Bootstrapping

The `data` package is designed to bootstrap from scratch a working (empty) table schema if one does not exist, or to use an existing implementation. This is currently a one-time operation that happens at application launch time.

## Table Schemas

Data is organized into three tables, with fields outlined below:

* `clusters`, a table that stores "Cluster" records (each cluster record maps to a unique deis cluster seen in the wild)
  * `cluster_id uuid PRIMARY KEY`
  * `data json`
* `clusters_checkins`, a table that stores records of deis clusters checking in with the API
  * `checkins_id bigserial PRIMARY KEY`
  * `cluster_id uuid`
  * `created_at timestamp`
  * `data json`
* `versions`, a table that stores authoritative deis component version information
  * `version_id bigserial PRIMARY KEY`
  * `component_name varchar(32)`
  * `train varchar(24)`
  * `version varchar(32)`
  * `release_timestamp timestamp`
  * `data json`
  * with a uniqueness constraint `unique (component_name, train, version)`

## License

Copyright 2016 Engine Yard, Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
