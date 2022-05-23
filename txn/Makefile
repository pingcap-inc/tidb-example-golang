# Copyright 2022 PingCAP, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: all prepare re-prepare build pessimistic optimistic pessimistic-not-oversell pessimistic-oversell optimistic-not-oversell optimistic-oversell

pessimistic:
	make build prepare pessimistic-not-oversell re-prepare pessimistic-oversell

optimistic:
	make build prepare optimistic-not-oversell re-prepare optimistic-oversell

prepare:
	tiup demo bookshop prepare --drop-tables --books 0 --authors 0 --orders 0 --ratings 0 --users 0

re-prepare:
	tiup demo bookshop prepare --drop-tables --books 0 --authors 0 --orders 0 --ratings 0 --users 0

build:
	go build -o bin/txn

pessimistic-not-oversell:
	./bin/txn -a 4 -b 6

pessimistic-oversell:
	./bin/txn -a 4 -b 7

optimistic-not-oversell:
	./bin/txn -a 4 -b 6 -o true

optimistic-oversell:
	./bin/txn -a 4 -b 7 -o true