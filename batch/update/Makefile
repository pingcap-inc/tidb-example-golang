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

.PHONY: all prepare build run

all:
	make prepare build run

prepare:
	~/.tiup/bin/tiup demo bookshop prepare --drop-tables
	mycli --host 127.0.0.1 --port 4000 -u root --no-warn < add_attr_ten_point.sql

build:
	go build -o bin/batch-update

run:
	./bin/batch-update