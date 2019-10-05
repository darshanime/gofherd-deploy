FROM grafana/grafana:6.4.1
USER root
RUN apk update && apk add curl
COPY ./datasources /etc/grafana/datasources
COPY ./dashboards /etc/grafana/dashboards
COPY ./setup.sh /setup.sh
ENTRYPOINT ["/bin/bash", "/setup.sh"]
