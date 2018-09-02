FROM alpine

COPY ./fin_mane_be /fin_mane_be

CMD [ "/fin_mane_be" ]