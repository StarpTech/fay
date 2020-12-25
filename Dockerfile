FROM mcr.microsoft.com/playwright:bionic as base

# |--------------------------------------------------------------------------
# | Fonts
# |--------------------------------------------------------------------------
# |
# | Installs a handful of fonts.
# | Note: ttf-mscorefonts-installer are installed on top of this Dockerfile.
# |

# Credits: 
# https://github.com/arachnys/athenapdf/blob/master/cli/Dockerfile
# https://help.accusoft.com/PrizmDoc/v12.1/HTML/Installing_Asian_Fonts_on_Ubuntu_and_Debian.html
RUN apt-get install -y \
    culmus \
    fonts-beng \
    fonts-hosny-amiri \
    fonts-lklug-sinhala \
    fonts-lohit-guru \
    fonts-lohit-knda \
    fonts-samyak-gujr \
    fonts-samyak-mlym \
    fonts-samyak-taml \
    fonts-sarai \
    fonts-sil-abyssinica \
    fonts-sil-padauk \
    fonts-telu \
    fonts-thai-tlwg \
    fonts-liberation \
    ttf-wqy-zenhei \
    fonts-arphic-uming \
    fonts-ipafont-mincho \
    fonts-ipafont-gothic \
    fonts-unfonts-core

COPY .docker/fonts/* /usr/local/share/fonts/
COPY .docker/fonts.conf /etc/fonts/conf.d/100-gotenberg.conf

FROM golang:1.15 as builder
WORKDIR /go/src/app
COPY . .
RUN ./scripts/build-installer.sh && ./scripts/build.sh

FROM base as production
COPY --from=builder /go/src/app/fay .
COPY --from=builder /go/src/app/fayinstaller .
# Ensure that client and browser are baked into image
RUN ./fayinstaller
EXPOSE 3000
CMD ["./fay"]