FROM python:3.8-slim-buster

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .
RUN apt-get update -y && apt-get install -y --no-install-recommends build-essential gcc libsndfile1


CMD [ "python3", "main.py" ]