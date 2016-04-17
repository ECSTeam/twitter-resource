FROM concourse/busyboxplus:base

COPY assets/out /opt/resource/out
COPY assets/check /opt/resource/check
COPY assets/in /opt/resource/in
