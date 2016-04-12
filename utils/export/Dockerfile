FROM python:2-alpine

ENV PYTHONUNBUFFERED 1
ENV PRODUCTION_MODE=yes
EXPOSE 5000

RUN mkdir /code
WORKDIR /code
ADD requirements.txt requirements.txt
RUN pip install -r requirements.txt

ADD export.py export.py

CMD ["gunicorn", "--bind", "0.0.0.0:5000", "export:app"]
