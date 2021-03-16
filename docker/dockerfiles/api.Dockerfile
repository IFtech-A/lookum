FROM ubuntu:20.04
RUN DEBIAN_FRONTEND=noninteractive
RUN sed -Ei "s:/(([a-z]{2}\.)?archive|security)\.ubuntu\.com:/mirror.kakao.com:" /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y locales ca-certificates tzdata && \
    apt-get clean && \
    locale-gen en_US.UTF-8
WORKDIR /lookum
ADD . ./
ENTRYPOINT [ "/lookum/bin/lookum" ]