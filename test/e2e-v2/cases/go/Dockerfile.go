# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.13@sha256:8ebb6d5a48deef738381b56b1d4cd33d99a5d608e0d03c5fe8dfa3f68d41a1f8 AS builder

ARG SW_AGENT_GO_COMMIT
ARG GO2SKY_CODE=${SW_AGENT_GO_COMMIT}.tar.gz
ARG GO2SKY_CODE_URL=https://github.com/SkyAPM/go2sky/archive/${GO2SKY_CODE}

ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /go2sky

ADD ${GO2SKY_CODE_URL} .
RUN tar -xf ${GO2SKY_CODE} --strip 1
RUN rm ${GO2SKY_CODE}

WORKDIR /go2sky/test/e2e/example-server
RUN go build -o main

FROM alpine:3.10@sha256:451eee8bedcb2f029756dc3e9d73bab0e7943c1ac55cff3a4861c52a0fdd3e98

COPY --from=builder /go2sky/test/e2e/example-server/main .
ENTRYPOINT ["/main"]
