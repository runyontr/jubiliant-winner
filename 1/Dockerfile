FROM ubuntu:18.04

ARG LITECOIN_VERSION=0.17.1
# install dependencies
# wget and gpg for downloading and validating binaries
RUN apt-get update &&  apt-get install -y wget gpg 

WORKDIR /opt/litecoin

RUN wget -q https://download.litecoin.org/litecoin-${LITECOIN_VERSION}/linux/litecoin-${LITECOIN_VERSION}-linux-signatures.asc
RUN wget -q https://download.litecoin.org/litecoin-${LITECOIN_VERSION}/linux/litecoin-${LITECOIN_VERSION}-x86_64-linux-gnu.tar.gz

# https://download.litecoin.org/README-HOWTO-GPG-VERIFY-TEAM-MEMBERS-KEY.txt
# Download the key.  Server in readme above was not working, but was able to find correct key in Ubuntu's server
RUN wget -q -O gpg.key 'https://keyserver.ubuntu.com/pks/lookup?op=get&search=0xFE3348877809386C'
RUN gpg --import gpg.key
# Make sure the checksum file is valid
RUN gpg --verify litecoin-${LITECOIN_VERSION}-linux-signatures.asc  
# Check the sha of the tarball is correct

# Reference: https://forums.linuxmint.com/viewtopic.php?t=225855&start=20
RUN sha256sum --check --ignore-missing litecoin-${LITECOIN_VERSION}-linux-signatures.asc
RUN tar -xvf litecoin-${LITECOIN_VERSION}-x86_64-linux-gnu.tar.gz

# Create data directory
RUN mkdir -p /data/litecoin
RUN chmod -R 777 /data/litecoin

# New user setup
RUN useradd -ms /bin/bash litecoin
# https://docs.docker.com/engine/reference/builder/#add add --chown so fould isn't owned by root
ADD --chown=litecoin:litecoin litecoin.conf /home/litecoin/.litecoin/
USER litecoin
WORKDIR /home/litecoin

RUN mkdir bin && cp /opt/litecoin/litecoin-*/bin/litecoind bin/
CMD ["./bin/litecoind"]
