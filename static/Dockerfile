FROM busybox

VOLUME  /www/horodata/static
WORKDIR /www/horodata/static

ADD roboto                  roboto
ADD material_icons          material_icons
ADD bower_components        bower_components
ADD CGU.pdf                 CGU.pdf
ADD CGV.pdf                 CGV.pdf
ADD images                  images

ADD front/front.css         front/front.css

ADD app/horodata.css        app/horodata.css
ADD app/horodata.min.js     app/horodata.js



CMD ["true"]
