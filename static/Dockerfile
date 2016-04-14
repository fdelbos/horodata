FROM busybox

VOLUME  /www/horodata/static
WORKDIR /www/horodata/static

ADD profile-default.png     profile-default.png
ADD roboto                  roboto
ADD material_icons          material_icons
ADD bower_components        bower_components
ADD app/horodata.css        app/horodata.css
ADD app/horodata.min.js     app/horodata.js

CMD ["true"]
