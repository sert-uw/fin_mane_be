FROM alpine

COPY /workspace/github-sert-uw-fin_mane_be /fin_mane_be

CMD [ "/fin_mane_be" ]