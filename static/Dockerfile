FROM busybox

VOLUME /www/horodata/static
WORKDIR /www/horodata/static

ADD bower_components        bower_components
ADD roboto                  roboto
ADD app/horodata.css        app/horodata.css
ADD app/horodata.min.js     app/horodata.js

CMD ["true"]
