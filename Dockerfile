FROM concourse/busyboxplus:base

COPY assets/check /opt/resource/check
COPY assets/in    /opt/resource/in
COPY assets/out   /opt/resource/out
